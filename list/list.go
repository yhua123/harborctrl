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

package list

import (
	//"fmt"
	"errors"
	"os"
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

	log.Infof("### Listing Running Harbor container ####")

	myCmdDocker := exec.Command("docker", "ps")
	myCmdGrep := exec.Command("grep", "deploy_")

	myCmdGrep.Stdin, _ = myCmdDocker.StdoutPipe()
	myCmdGrep.Stdout = os.Stdout
	myCmdGrep.Stderr = os.Stderr
	err := myCmdGrep.Start()
	if err != nil {
		//fmt.Println("myCmdGrep.Output: ", err)
		log.Errorf("List cannot continue... %s", err)
		return err
	}
	err = myCmdDocker.Run()
	if err != nil {
		log.Errorf("List cannot continue... %s", err)
		return err
	}
	myCmdGrep.Wait()

	return nil
}
