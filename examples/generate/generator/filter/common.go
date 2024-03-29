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
package filter

import (
	"fmt"
	"path/filepath"

	"github.com/curoky/go-thrift-parser/parser/ast"
	"github.com/flosch/pongo2/v6"
)

func resolveType(t *ast.Type) (name string) {
	switch t.Name {
	case "map":
		name = fmt.Sprintf("map<%s,%s>", resolveType(t.KeyType), resolveType(t.ValueType))
	case "list":
		name = fmt.Sprintf("list<%s>", resolveType(t.ValueType))
	case "set":
		name = fmt.Sprintf("set<%s>", resolveType(t.ValueType))
	default:
		name = t.Name
	}
	return
}

func expandCategory(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := in.Interface().(*ast.Type)
	name := t.Name
	switch t.Category {
	case ast.CategoryMap:
		name = fmt.Sprintf("map|%s|%s", t.KeyType.FinalType().Category.String(), t.ValueType.FinalType().Category.String())
	case ast.CategoryList:
		name = fmt.Sprintf("list|%s", t.ValueType.FinalType().Category.String())
	case ast.CategorySet:
		name = fmt.Sprintf("set|%s", t.ValueType.FinalType().Category.String())
	}
	return pongo2.AsSafeValue(name), nil
}

func BaseName(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsSafeValue(filepath.Base(in.Interface().(string))), nil
}
