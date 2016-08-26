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

package check

import (
	"errors"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"
)

var packageName string

var checkPoint = map[string]string{
	"config":         "harbor.cfg",
	"prepare":        "prepare",
	//"docker":         "/usr/bin/docker",
	//"docker-compose": "/usr/local/bin/docker-compose",
}

// Flags ...
func Flags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "pkg, p",
			Usage:       "Checking Harbor specific package(docker/docker-compose, config, prepare).",
			Value:       "all",
			Destination: &packageName,
		},
	}

	return flags
}

//Run ...
func Run(cli *cli.Context) error {

	//fmt.Printf("PackageName=%s\n", PackageName)
	err := errors.New("")

	_, exists := checkPoint[packageName]

	if !exists && (packageName != "all") {
		err = errors.New("no package")
		return err
	}

	if packageName == "all" {
		for key, value := range checkPoint {
			err = checkFile(key, value)
			if err != nil {
				return err
			}
		}
	} else {
		err = checkFile(packageName, checkPoint[packageName])
		if err != nil {
			return err
		}
	}

	//fmt.Printf("Harbor %s package check successful.\n", packageName)

	return nil
}

func checkFile(checkPkgName string, checkFileName string) error {

	_, err := os.Stat(checkFileName)
	if err != nil {
		log.Errorf("Harbor %s package check fail.", checkPkgName)
		return err
	}

	fmt.Printf("%s package check successful.\n", checkPkgName)
	return nil
}
