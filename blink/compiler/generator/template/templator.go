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

package template

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/curoky/blink/blink"
	"github.com/curoky/blink/blink/compiler/utils"

	"github.com/flosch/pongo2/v4"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

type Templator struct {
	resource *packr.Box
	tpl_set  *pongo2.TemplateSet

	OutputList []string
}

func Init() {}

func CreateTemplator(lang string) *Templator {
	box := packr.New("Template Box", ".")
	return &Templator{
		resource: box,
		tpl_set:  pongo2.NewSet("PackrBox", NewPongo2PackrLoader(box)),
	}
}

var trimSpace, emptyLine = regexp.MustCompile(`[ ]*\n`), regexp.MustCompile(`[\r\n]+`)

func removeEmptyLine(str string) string {
	// return strings.ReplaceAll(str, "\n\n", "\n")
	str = strings.TrimSpace(str)
	str = trimSpace.ReplaceAllString(str, "\n")
	str = emptyLine.ReplaceAllString(str, "\n")
	return str
}

func (t *Templator) RenderTo(tpl_path string, context pongo2.Context, output_path string) {
	tpl, err := t.tpl_set.FromFile(tpl_path)
	if err != nil {
		log.Fatal(err)
	}

	utils.MakeFileWriteAble(output_path)
	writer, err := os.OpenFile(output_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	context["filename"] = filepath.Base(output_path)
	context["version"] = blink.Version

	content, err := tpl.Execute(context)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteString(removeEmptyLine(content))
	if err != nil {
		log.Fatal(err)
	}
	writer.Close()

	log.Infof("Render file %s", output_path)

	t.OutputList = append(t.OutputList, output_path)
}
