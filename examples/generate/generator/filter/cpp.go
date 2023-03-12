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
	"strings"

	"github.com/curoky/go-thrift-parser/parser/ast"
	"github.com/flosch/pongo2/v6"
)

var cpp_type_map = map[string]string{
	"bool":   "bool",
	"byte":   "char",
	"i16":    "int16_t",
	"i32":    "int32_t",
	"i64":    "int64_t",
	"double": "double",
	"string": "std::string",
	"binary": "std::string",
	"list":   "std::vector",
	"set":    "std::set",
	"map":    "std::map",
}

func CppType(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	name := resolveType(in.Interface().(*ast.Type))
	for k, v := range cpp_type_map {
		name = strings.Replace(name, k, v, -1)
	}
	// TODO(curoky): remove this trick
	name = strings.Replace(name, "std::std::", "std::", -1)
	return pongo2.AsSafeValue(name), nil
}

func AnnCppType(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := in.Interface().(map[string]*ast.Annotation)["cpp.type"]
	return pongo2.AsSafeValue(t.Value), nil
}

func resolveCppValue(t *ast.ConstValue) (res string) {
	switch t.Type {
	case ast.ConstTypeDouble:
		res = fmt.Sprint(*t.TypedValue.Double)
	case ast.ConstTypeInt:
		res = fmt.Sprint(*t.TypedValue.Int)
	case ast.ConstTypeLiteral:
		res = fmt.Sprintf("\"%s\"", *t.TypedValue.Literal)
	case ast.ConstTypeList:
	case ast.ConstTypeMap:
	}
	return
}

func CppValue(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsSafeValue(resolveCppValue(in.Interface().(*ast.ConstValue))), nil
}
