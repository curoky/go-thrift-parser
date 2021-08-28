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

//go:generate pigeon -o parser/thrift.peg.go parser/thrift.peg
//go:generate thrift --gen go --out . ast/ast.thrift

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/curoky/blink/blink"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func app() *cli.App {
	app := &cli.App{
		Name:    "blink compiler",
		Usage:   "compiler thrift IDL",
		Version: fmt.Sprintf("%s (%s)", blink.Version, runtime.Version()),
	}
	app.Flags = []cli.Flag{
		&cli.PathFlag{
			Name:      "file",
			Usage:     "input file path",
			TakesFile: true,
			Required:  true,
		},
		&cli.PathFlag{
			Name:     "out",
			Usage:    "output directory",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "lang",
			Usage:    "language need to generate",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "dump",
			Value: false,
			Usage: "dump ast",
		},
		&cli.BoolFlag{
			Name:  "readonly",
			Value: true,
			Usage: "make output file readony",
		},
		&cli.BoolFlag{
			Name:   "fmtcode",
			Value:  true,
			Usage:  "format code",
			Hidden: true,
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Value: false,
		},
	}
	return app
}

func main() {
	err := app().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}