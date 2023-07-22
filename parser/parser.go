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
	"os"
	"path/filepath"

	"github.com/curoky/go-thrift-parser/parser/ast"
	log "github.com/sirupsen/logrus"
)

func Dump(thrift *ast.Thrift, filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	content, err := json.MarshalIndent(thrift, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(absPath, []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

func searchThriftFileInIncludePath(filename string, includePaths []string) (string, error) {
	// TODO(curoky): check if there are any files with same name in different directory.
	for _, includePath := range includePaths {
		log.Debugf("parser: searching <%s> from <%s>", filename, includePath)
		path := filepath.Join(includePath, filename)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", &os.PathError{Op: "search", Path: filename, Err: os.ErrNotExist}
}

func ParseThriftFile(filename string, includePaths []string, recursive bool, verbose bool) (*ast.Thrift, error) {
	thrift := &ast.Thrift{
		Documents: make(map[string]*ast.Document),
	}
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	includePaths = append(includePaths, ".", "/")
	err = recursiveParseThriftFile(thrift, filename, includePaths, verbose)
	if err != nil {
		return nil, err
	}
	return thrift, err
}

func parseThriftString(filename string, fileContent []byte, verbose bool) (*ast.Document, error) {
	docIf, err := Parse(filename, fileContent, Debug(verbose))
	if err != nil {
		return nil, err
	}
	doc := docIf.(*ast.Document)
	doc.Filename = filename

	err = doc.Resolve(nil)
	if err != nil {
		return nil, err
	}
	return doc, err
}

func recursiveParseThriftFile(thrift *ast.Thrift, filename string, includePaths []string, verbose bool) error {
	reslovedPath, err := searchThriftFileInIncludePath(filename, includePaths)
	log.Debugf("parser: searched %s", reslovedPath)
	if err != nil {
		log.Errorf("parser: can't find %s", filename)
		return err
	}

	if _, ok := thrift.Documents[reslovedPath]; ok {
		log.Debugf("parser: already in cached, skip %s", reslovedPath)
		return nil
	}

	log.Debugf("parser: parse %s", reslovedPath)
	content, err := os.ReadFile(reslovedPath)
	if err != nil {
		log.Errorf("parser: ReadFile failed %s, err %s", reslovedPath, err)
		return err
	}

	doc, err := parseThriftString(reslovedPath, content, verbose)
	if err != nil {
		log.Errorf("parser: parseThriftString failed %s, err %s", reslovedPath, err)
		return err
	}

	log.Debugf("parser: parse %s success", doc.Filename)
	thrift.Documents[reslovedPath] = doc

	for _, inc := range doc.Includes {
		err = recursiveParseThriftFile(thrift, inc.Path, append(includePaths, filepath.Dir(reslovedPath)), verbose)
		if err != nil {
			return err
		}
		inc.Reference = thrift.Documents[inc.Path]
	}
	return nil
}
