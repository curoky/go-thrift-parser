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

package parser

import (
	"context"
	"os"
	"path/filepath"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/curoky/blink/blink/compiler/ast"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
	Document ast.Document
	Verbose  bool
}

func CreateParser(verbose bool) *Parser {
	p := &Parser{Verbose: verbose}
	p.Document.Thrifts = make(map[string]*ast.Thrift)
	return p
}

func (p *Parser) Dump(filename string) {
	if len(filename) == 0 {
		filename = "ast.json"
	}
	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Error(err)
	}
	trans := thrift.NewTMemoryBuffer()
	serializer := thrift.TSerializer{
		Transport: trans,
		Protocol:  thrift.NewTSimpleJSONProtocolConf(trans, nil),
	}
	content, err := serializer.WriteString(context.TODO(), &p.Document)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(absPath, []byte(content), 0600)
	log.Error(err)
}

func (p *Parser) RecursiveParse(filename string) error {
	log.Infof("RecursiveParse: start process %s", filename)

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	if _, ok := p.Document.Thrifts[absPath]; ok {
		log.Infof("RecursiveParse: already in cached, skip %s", absPath)
		return nil
	}

	log.Infof("RecursiveParse: parse %s", absPath)
	i, err := ParseFile(absPath, Debug(p.Verbose))
	if err != nil {
		log.Errorf("RecursiveParse: parse failed %s, err %s", absPath, err)
		return err
	}
	tt := i.(*ast.Thrift)
	tt.Filename = absPath
	log.Infof("RecursiveParse: parse %s success", tt.Filename)
	p.Document.Thrifts[absPath] = tt

	for _, inc := range tt.Includes {
		inc.Path = filepath.Join(filepath.Dir(absPath), inc.Path)
		err = p.RecursiveParse(inc.Path)
		if err != nil {
			return err
		}
		inc.Reference = p.Document.Thrifts[inc.Path]
	}
	return nil
}
