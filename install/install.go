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

package install

import (
	//"fmt"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"
)

var configFile string
var composeFile string

//Flags ...
func Flags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "configfile, cf",
			Usage:       "Specify an alternate config file (default: Deploy/harbor.cfg)",
			Value:       "Deploy/harbor.cfg",
			Destination: &configFile,
		},
		cli.StringFlag{
			Name:        "composefile, df",
			Usage:       "Specify an alternate docker-compose file (default: Deploy/docker-compose.yml)",
			Value:       "Deploy/docker-compose.yml",
			Destination: &composeFile,
		},
	}

	return flags
}

//Run ...
func Run(cli *cli.Context) error {

	//fmt.Printf("configfile: %s\n", configFile)

	//prepare
	prepare := exec.Command("prepare", "-conf", configFile)

	prepare.Stdout = os.Stdout
	prepare.Stderr = os.Stderr
	err := prepare.Run()
	if err != nil {
		log.Errorf("prepare fail... %s", err)
		return err
	}

	//load images
	loadImage := exec.Command("load_image.sh")
	loadImage.Stdout = os.Stdout
	loadImage.Stderr = os.Stderr
	err := loadImage.Run()
	if err != nil {
		log.Errorf("load images fail... %s", err)
		return err
	}

	//build
	compose := exec.Command("docker-compose", "-f", composeFile, "up", "-d")
	compose.Stdout = os.Stdout
	compose.Stderr = os.Stderr
	err = compose.Run()
	if err != nil {
		log.Errorf("create fail... %s", err)
		return err
	}

	return nil
}
