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
	"github.com/curoky/blink/blink/compiler/ast"
)

var baseTypeNameToCategory = map[string]ast.Category{
	"bool":   ast.Category_Bool,
	"byte":   ast.Category_Byte,
	"i16":    ast.Category_I16,
	"i32":    ast.Category_I32,
	"i64":    ast.Category_I64,
	"double": ast.Category_Double,
	"string": ast.Category_String,
	"binary": ast.Category_Binary,
	"list":   ast.Category_List,
	"set":    ast.Category_Set,
	"map":    ast.Category_Map,
}
