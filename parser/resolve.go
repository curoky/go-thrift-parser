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
	"strings"

	"github.com/curoky/go-thrift-parser/parser/ast"
	log "github.com/sirupsen/logrus"
)

func getTypeFromDocument(name string, thrift *ast.Thrift) *ast.Type {
	log.Infof("getTypeFromDocument %s", name)
	if typ, ok := thrift.Enums[name]; ok {
		return typ
	}
	if typ, ok := thrift.Typedefs[name]; ok {
		return typ
	}
	if typ, ok := thrift.Structs[name]; ok {
		return typ
	}
	if typ, ok := thrift.Unions[name]; ok {
		return typ
	}
	if typ, ok := thrift.Exceptions[name]; ok {
		return typ
	}
	return nil
}

func resolveType(doc *ast.Document, typ *ast.Type) *ast.Type {
	if typ.Category != ast.CategoryIdentifier && typ.Category != ast.CategoryTypedef {
		return typ
	}
	// log.Infof("resolveType %s %v", typ.Name, typ.Category)
	// TODO(curoky): remove duplicated code
	if strings.Contains(typ.Name, ".") {
		seg := strings.Split(typ.Name, ".")
		for _, inc := range typ.Belong.Includes {
			if inc.Name == seg[0] {
				if typ.PreRefType == nil {
					typ.PreRefType = getTypeFromDocument(seg[1], inc.Reference)
				}
				typ.FinalRefType = resolveType(doc, typ.PreRefType)
			}
		}
	} else {
		if typ.PreRefType == nil {
			typ.PreRefType = getTypeFromDocument(typ.Name, typ.Belong)
		}
		typ.FinalRefType = resolveType(doc, typ.PreRefType)
	}
	return typ.FinalRefType
}

func resolve(doc *ast.Document) {
	for _, thrift := range doc.Thrifts {
		for _, typ := range thrift.AllTypes {
			resolveType(doc, typ)
		}
	}
}
