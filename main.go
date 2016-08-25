/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"os"

	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"

	"harborctrl/check"
	"harborctrl/down"
	"harborctrl/install"
	"harborctrl/list"
	"harborctrl/version"
)

//harborctrl version, buildid and commitid
var (
	Version  string
	BuildID  string
	CommitID string
)

//harborctral log
const (
	LogFile = "harborctrl.log"
)

func main() {
	app := cli.NewApp()

	app.Name = filepath.Base(os.Args[0])
	app.Usage = `Create and Manage Harbor Host
   Please confirm harborctrl running under Harbor root directory.`
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:      "version",
			ShortName: "ver",
			Usage:     "Show Harbor version information.",
			ArgsUsage: " ",
			Action:    version.Run,
		},
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "List running Harbor continers.",
			ArgsUsage: " ",
			Action:    list.Run,
		},
		{
			Name:      "check",
			ShortName: "chk",
			Usage:     "Inspect Harbor code and install env.",
			Action:    check.Run,
			Flags:     check.Flags(),
		},
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "install Harbor continers.",
			Action:    install.Run,
			Flags:     install.Flags(),
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Shutdown Harbor continers.",
			Action:    down.Run,
			Flags:     down.Flags(),
		},
	}

	if Version != "" {
		app.Version = fmt.Sprintf("%s-%s-%s", Version, BuildID, CommitID)
	} else {
		app.Version = fmt.Sprintf("%s-%s", BuildID, CommitID)
	}

	logs := []io.Writer{app.Writer}
	// Open log file
	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening logfile %s: %v\n", LogFile, err)
	} else {
		defer f.Close()
		logs = append(logs, f)
	}

	// Initiliaze logger with default TextFormatter
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
	// SetOutput to io.MultiWriter so that we can log to stdout and a file
	log.SetOutput(io.MultiWriter(logs...))

	if err := app.Run(os.Args); err != nil {
		log.Errorf("--------------------")
		log.Errorf("%s failed: %s\n", app.Name, err)
		os.Exit(1)
	}
}
