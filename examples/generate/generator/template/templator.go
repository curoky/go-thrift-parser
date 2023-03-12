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

package template

import (
	"embed"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flosch/pongo2/v6"
	log "github.com/sirupsen/logrus"
)

//go:embed cpp/*
var templateData embed.FS

type Templator struct {
	tpl_set *pongo2.TemplateSet

	OutputList []string
}

func Init() {}

func CreateTemplator(lang string) *Templator {
	return &Templator{
		tpl_set: pongo2.NewSet("TemplateData", pongo2.NewFSLoader(templateData)),
	}
}

var trimSpaceRE = regexp.MustCompile(`[ ]*\n`)
var emptyLineRE = regexp.MustCompile(`[\r\n]+`)

func removeEmptyLine(str string) string {
	// return strings.ReplaceAll(str, "\n\n", "\n")
	str = strings.TrimSpace(str)
	str = trimSpaceRE.ReplaceAllString(str, "\n")
	str = emptyLineRE.ReplaceAllString(str, "\n")
	return str
}

func (t *Templator) RenderTo(tpl_path string, context pongo2.Context, output_path string) error {
	tpl, err := t.tpl_set.FromFile(tpl_path)
	if err != nil {
		return err
	}

	writer, err := os.OpenFile(output_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	context["filename"] = filepath.Base(output_path)
	context["version"] = "0.1.0"

	content, err := tpl.Execute(context)
	if err != nil {
		return err
	}

	_, err = writer.WriteString(removeEmptyLine(content))
	if err != nil {
		return err
	}
	writer.Close()

	log.Infof("Render file %s", output_path)

	t.OutputList = append(t.OutputList, output_path)
	return nil
}
