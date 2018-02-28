package installer

import (
	"errors"
	"fmt"
)

var (
	templates = map[string]string{
		"/etc/init/goreplay-listen.conf": fmt.Sprintf(`
description "Goreplay"
author      "Adam Reynolds <adam@connected-ventures.com>"

env START=/usr/local/bin/goreplay
env PORT=:3000
env HOST=http://10.133.0.75:80

chdir /etc/

limit nofile 65536 65536

setuid root
env HOME=/tmp

start on (local-filesystems and net-device-up IFACE=eth0)
stop  on shutdown

respawn                # restart when job dies
respawn limit 5 60     # give up restart after 5 respawns in 60 seconds

exec $START --input-raw $PORT --output-http $HOST   --http-disallow-url /_health --http-disallow-url /_metrics
`),
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