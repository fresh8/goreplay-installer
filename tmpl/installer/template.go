package installer

import (
	"errors"
)

var (
	templates = map[string]string{
		"/etc/init/goreplay-listen.conf": `description "Goreplay"
author      "Adam Reynolds <adam@connected-ventures.com>"

env START=/usr/local/bin/goreplay
env PORT={{.Port}}
env HOST={{.Host}}

chdir /etc/

limit nofile 65536 65536

setuid root
env HOME=/tmp

start on (local-filesystems and net-device-up IFACE=eth0)
stop  on shutdown

respawn                # restart when job dies
respawn limit 5 60     # give up restart after 5 respawns in 60 seconds

exec $START --input-raw $PORT --output-http $HOST {{.Filter}}
`,
	}
)

// GetTemplate returns the template contents for a file key, or an error if it
// does not exist.
func GetTemplate(key string) (string, error) {
	val, ok := templates[key]
	if !ok {
		return "", ErrTemplateNotExist
	}

	return val, nil
}

var (
	// ErrTemplateNotExist returns when a template does not exist
	ErrTemplateNotExist = errors.New("template does not exist")
)
