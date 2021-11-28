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
package generator

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/curoky/blink/blink/compiler/ast"
	"github.com/curoky/blink/blink/compiler/generator/template"
	"github.com/curoky/blink/blink/compiler/utils"
	"github.com/flosch/pongo2/v4"
)

type CppGenerator struct {
	tpl *template.Templator
}

func (g *CppGenerator) Generate(thrift *ast.Thrift, conf Config) {
	ctx := pongo2.Context{"thrift": thrift}

	g.tpl.RenderTo("cpp/types.h.j2", ctx, fmt.Sprintf("%s/%s.h", conf.OutputPrefix, filepath.Base(thrift.Filename)))

	if conf.FormatCode {
		if _, err := exec.LookPath("clang-format"); err == nil {
			utils.ClangFormat(g.tpl.OutputList)
		}
	}

	if conf.MakeReadOnly {
		for _, f := range g.tpl.OutputList {
			utils.MakeFileReadOnly(f)
		}
	}
}
