
[![CircleCI](https://circleci.com/gh/fresh8/goreplay-installer.svg?style=svg)](https://circleci.com/gh/fresh8/goreplay-installer)[![Go Report Card](https://goreportcard.com/badge/github.com/fresh8/goreplay-installer)](https://goreportcard.com/report/github.com/fresh8/goreplay-installer)

# goreplay-installer
A tool for deploying goreplay to autoscalled instances http://fresh8gaming.com

Current version: v0.0.8

## Getting Started
### Installing
Download a version from the [releases](https://github.com/fresh8/goreplay-installer/releases) page on Github, and place it into your local bin folder.

### Running
goreplay-installer providers a number of options, depending on the command you wish to run. For a full list of commands, just run:
```
goreplay-installer help
```

### Install
This command downloads goreplay to `/tmp` then copies this file to `/usr/local/bin`. An upstart script is installed to `/etc/init/goreplay.conf`. It expects the PORT and HOST environment variables have been set for this instance.
```
goreplay-installer install
```

Additional options can be viewed via the help command.
