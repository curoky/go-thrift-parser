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

//go:generate pigeon -o thrift.peg.go thrift.peg

package parser

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/curoky/go-thrift-parser/parser/ast"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
	Thrift  ast.Thrift
	Verbose bool
}

func CreateParser(verbose bool) *Parser {
	p := &Parser{Verbose: verbose}
	p.Thrift.Documents = make(map[string]*ast.Document)
	return p
}

func (p *Parser) Dump(filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	content, err := json.Marshal(p.Thrift)
	if err != nil {
		return err
	}
	err = os.WriteFile(absPath, []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) RecursiveParse(filename string) error {
	log.Debugf("RecursiveParse: start process %s", filename)

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	if _, ok := p.Thrift.Documents[absPath]; ok {
		log.Debugf("RecursiveParse: already in cached, skip %s", absPath)
		return nil
	}

	log.Debugf("RecursiveParse: parse %s", absPath)
	docIf, err := ParseFile(absPath, Debug(p.Verbose))
	if err != nil {
		log.Errorf("RecursiveParse: parse failed %s, err %s", absPath, err)
		return err
	}
	doc := docIf.(*ast.Document)
	doc.Filename = absPath
	doc.Resolve(&p.Thrift)

	log.Debugf("RecursiveParse: parse %s success", doc.Filename)
	p.Thrift.Documents[absPath] = doc

	for _, inc := range doc.Includes {
		inc.Path = filepath.Join(filepath.Dir(absPath), inc.Path)
		err = p.RecursiveParse(inc.Path)
		if err != nil {
			return err
		}
		inc.Reference = p.Thrift.Documents[inc.Path]
	}
	return nil
}

func (p *Parser) Resolve() {
	resolve(&p.Thrift)
}
