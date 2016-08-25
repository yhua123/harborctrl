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

package version

import (
	"errors"
	"fmt"
	"os/exec"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"
)

//Run ...
func Run(cli *cli.Context) error {

	if len(cli.Args()) > 0 {
		log.Errorf("Unknown argument: %s", cli.Args()[0])
		return errors.New("")
	}

	gitversionCmd := exec.Command("git", "describe", "--tags")
	gitversionOutput, _ := gitversionCmd.Output()
	ReleaseTag := fmt.Sprintf(string(gitversionOutput))

	if ReleaseTag != "" {
		fmt.Fprintf(cli.App.Writer, "Harbor version: %s\n", ReleaseTag)
	} else {

		log.Errorf("Fail get the Harbor version from .git.")
	}

	return nil
}
