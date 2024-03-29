/*
 * Copyright (c) 2020-2023 curoky(cccuroky@gmail.com).
 *
 * This file is part of go-thrift-parser.
 * See https://github.com/curoky/blink for further info.
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
	"fmt"

	"github.com/curoky/go-thrift-parser/parser"
	"github.com/curoky/go-thrift-parser/parser/ast"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	thrift, err := parser.ParseThriftFile("types.thrift", []string{}, true, false)
	if err != nil {
		panic(err)
	}

	for name, doc := range thrift.Documents {
		fmt.Printf("%s\n", name)
		fmt.Printf("%s\n", doc.Filename)
		for _, body := range doc.Body {
			switch v := body.([]interface{})[0].(type) {
			case *ast.Include:
				fmt.Printf("include: %s\n", v.Name)
			case *ast.Type:
				fmt.Printf("type: %s\n", v.Name)
			case *ast.Service:
				fmt.Printf("service: %s\n", v.Name)
			}
		}
	}

}
