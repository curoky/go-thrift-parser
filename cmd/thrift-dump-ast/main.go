/*
 * Copyright 2021 curoky(cccuroky@gmail.com).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"
	"path/filepath"

	"github.com/curoky/go-thrift-parser/parser"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func app() *cli.App {
	app := &cli.App{}
	app.Flags = []cli.Flag{
		&cli.PathFlag{
			Name:      "file",
			Usage:     "Set the input thrift file path",
			TakesFile: true,
			Required:  true,
		},
		&cli.PathFlag{
			Name:     "out",
			Usage:    "Set the output file path",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose mode",
			Value: false,
		},
		&cli.StringSliceFlag{
			Name:  "include",
			Usage: "Add a directory to the list of directories searched for include directives",
			Value: cli.NewStringSlice(),
		},
	}

	app.Action = func(c *cli.Context) error {
		input_file, _ := filepath.Abs(c.Path("file"))
		output_dir := c.Path("out")
		log.Infof("Input file: %s", input_file)

		p := parser.CreateParser(c.Bool("verbose"), c.StringSlice("include"))
		if err := p.RecursiveParse(input_file); err != nil {
			log.Fatal(err)
		}

		log.Infof("Dump to: %s", output_dir)
		err := p.Dump(output_dir)
		if err != nil {
			return err
		}

		log.Infof("Success!")
		return nil
	}
	return app
}

func main() {
	err := app().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
