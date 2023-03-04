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

//go:generate pigeon -o thrift.peg.go thrift.peg

package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/curoky/go-thrift-parser/parser/ast"
	log "github.com/sirupsen/logrus"
)

type PreProcessorType func(filename string) ([]byte, error)

type Parser struct {
	Thrift       ast.Thrift
	IncludePaths []string
	Verbose      bool
	PreProcessor PreProcessorType
}

func readOnlyProcessor(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func CreateParser(verbose bool, includePaths []string) *Parser {
	p := &Parser{Verbose: verbose}
	p.Thrift.Documents = make(map[string]*ast.Document)
	p.IncludePaths = includePaths
	p.PreProcessor = readOnlyProcessor
	return p
}

func (p *Parser) Dump(filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	content, err := json.MarshalIndent(p.Thrift, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(absPath, []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) findThriftFileInIncludePath(filename string) *string {
	if filepath.IsAbs(filename) {
		if _, err := os.Stat(filename); err == nil {
			return &filename
		}
	}
	for _, includePath := range p.IncludePaths {
		log.Debugf("findThriftFileInIncludePath: %s %s", includePath, filename)
		path := filepath.Join(includePath, filename)
		if _, err := os.Stat(path); err == nil {
			return &path
		}
	}
	return nil
}

func (p *Parser) RecursiveParse(filename string) error {
	log.Debugf("RecursiveParse: start process %s", filename)

	absPath := ""
	if path := p.findThriftFileInIncludePath(filename); path != nil {
		absPath = *path
	} else {
		return fmt.Errorf("RecursiveParse: can't find %s", filename)
	}

	if _, ok := p.Thrift.Documents[absPath]; ok {
		log.Debugf("RecursiveParse: already in cached, skip %s", absPath)
		return nil
	}

	log.Debugf("RecursiveParse: parse %s", absPath)
	content, err := p.PreProcessor(absPath)
	if err != nil {
		log.Errorf("RecursiveParse: PreProcessor failed %s, err %s", absPath, err)
		return err
	}

	docIf, err := Parse(absPath, content, Debug(p.Verbose))
	if err != nil {
		log.Errorf("RecursiveParse: parse failed %s, err %s", absPath, err)
		return err
	}
	doc := docIf.(*ast.Document)
	doc.Filename = absPath

	err = doc.Resolve(&p.Thrift)
	if err != nil {
		return err
	}

	log.Debugf("RecursiveParse: parse %s success", doc.Filename)
	p.Thrift.Documents[absPath] = doc

	for _, inc := range doc.Includes {
		err = p.RecursiveParse(inc.Path)
		if err != nil {
			return err
		}
		inc.Reference = p.Thrift.Documents[inc.Path]
	}
	return nil
}
