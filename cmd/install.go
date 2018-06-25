package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/fresh8/goreplay-installer/tmpl/installer"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

type upstartConfig struct {
	Port   string
	Host   string
	Filter string
}

var (
	debug bool

	goFileName      = "gor_0.16.1_x64.tar.gz"
	goReplayVersion = "https://github.com/buger/goreplay/releases/download/v0.16.1/" + goFileName
	workingDir      = `/tmp`
	destDir         = `/usr/local/bin`
	upstartDir      = `/etc/init`
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "if goreplay-installer should run in debug mode (default is false)")
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install listening_port destination_url",
	Short: "Install goreplay and install upstart config",
	Long:  `Install goreplay and install upstart config and specifiy which port to listen to and the destination address for the incomming requests`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.MinimumNArgs(2)(cmd, args) // port & host
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			logMessage(arg)
		}

		config := createConfig(args)

		logMessage(fmt.Sprintf("Downloading goreplay file to %q", workingDir))
		err := downloadFile(workingDir+"/"+goFileName, goReplayVersion)
		if err != nil {
			log.Fatal(err)
		}

		logMessage(fmt.Sprintf("Extracting goreplay file to %q", destDir))
		err = extractFile(workingDir+"/"+goFileName, destDir)
		if err != nil {
			log.Fatal(err)
		}

		logMessage("Creating Upstart script and deploying")
		upstartConfBuf, err := createUpstartConf(config)
		if err != nil {
			log.Fatal(err)
		}

		logMessage("Writing Upstart script template to file")
		err = writeToFile(upstartDir+"/goreplay.conf", upstartConfBuf.String())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf(`Success! Everything should now be in place for your environment.`))
	},
}

func logMessage(message string) {
	if debug {
		log.Println(message)
	}
}

func createUpstartConf(config upstartConfig) (buf *bytes.Buffer, err error) {
	template, err := installer.GetTemplate("/etc/init/goreplay-listen.conf")
	if err != nil {
		return buf, err
	}

	return createTemplateOutput(config, template)
}

func createTemplateOutput(config upstartConfig, templateContent string) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.New("upstart-config").Parse(templateContent)
	if err != nil {
		return buf, err
	}

	err = tmpl.Execute(buf, config)
	return buf, err
}

func downloadFile(filepath, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func extractFile(filename, destDir string) error {
	err := archiver.TarGz.Open(filename, destDir)
	return err
}

func writeToFile(filePath string, contents string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}

	return nil
}

func createConfig(args []string) upstartConfig {
	config := upstartConfig{
		Port:   args[0],
		Host:   args[1],
		Filter: "",
	}
	for i := 3; i <= len(args); i++ {
		config.Filter = strings.Trim(fmt.Sprintf("%s --http-allow-url %s", config.Filter, args[i-1]), " ")
	}
	if config.Filter == "" {
		config.Filter = "--http-disallow-url /_health --http-disallow-url /_metrics"
	}

	return config
}
