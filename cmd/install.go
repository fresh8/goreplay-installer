package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/fresh8/goreplay-installer/tmpl/installer"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

type upstartConfig struct {
	Body string
}

var (
	debug bool

	goFileName      = "gor_0.16.1_x64.tar.gz"
	goReplayVersion = "https://github.com/buger/goreplay/releases/download/v0.16.1/" + goFileName
	workingDir      = `/tmp`
	destDir         = `/usr/local/bin`
	upstartDir      = `/etc/init/`
	outputUpstart   = ``
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "if terraplate should run in debug mode (default is false)")
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "TODO",
	Long:  `TODO`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.MinimumNArgs(0)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
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
		upstartConfBuf, err := createUpstartConf()
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

func createUpstartConf() (buf *bytes.Buffer, err error) {
	template, err := installer.GetTemplate("/etc/init/goreplay-listen.conf")
	if err != nil {
		return buf, err
	}

	return createTemplateOutput(template)
}

func createTemplateOutput(templateContent string) (*bytes.Buffer, error) {
	var config upstartConfig
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.New("thing").Parse(templateContent)
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
