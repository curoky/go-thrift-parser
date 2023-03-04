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

package ast

import "fmt"

func recursiveResolveIdentifierType(document *Document, typ *Type) {
	for _, field := range typ.Fields {
		if field.Type.Category == CategoryIdentifier {
			if originalType, exists := document.Types[field.Type.Name]; exists {
				field.Type = originalType
			}
		} else {
			recursiveResolveIdentifierType(document, field.Type)
		}

	}
	if typ.KeyType != nil {
		if typ.KeyType.Category == CategoryIdentifier {
			if originalType, exists := document.Types[typ.KeyType.Name]; exists {
				typ.KeyType = originalType
			}
		} else {
			recursiveResolveIdentifierType(document, typ.KeyType)
		}
	}
	if typ.ValueType != nil {
		if typ.ValueType.Category == CategoryIdentifier {
			if originalType, exists := document.Types[typ.ValueType.Name]; exists {
				typ.ValueType = originalType
			}
		} else {
			recursiveResolveIdentifierType(document, typ.ValueType)
		}
	}
	if typ.PreRefType != nil {
		if typ.PreRefType.Category == CategoryIdentifier {
			if originalType, exists := document.Types[typ.PreRefType.Name]; exists {
				typ.PreRefType = originalType
			}
		} else {
			recursiveResolveIdentifierType(document, typ.PreRefType)
		}
	}
}

func resolveIdentifierType(document *Document) {
	for _, typ := range document.Types {
		recursiveResolveIdentifierType(document, typ)
	}

	for _, service := range document.Services {
		for _, function := range service.Functions {
			for _, argument := range function.Arguments {
				if argument.Type.Category == CategoryIdentifier {
					if originalType, exists := document.Types[argument.Type.Name]; exists {
						argument.Type = originalType
					}
				} else {
					recursiveResolveIdentifierType(document, argument.Type)
				}
			}
			if function.ReturnType != nil {
				if function.ReturnType.Category == CategoryIdentifier {
					if originalType, exists := document.Types[function.ReturnType.Name]; exists {
						function.ReturnType = originalType
					}
				} else {
					recursiveResolveIdentifierType(document, function.ReturnType)
				}
			}
		}
	}
}

func (document *Document) Resolve(thrift *Thrift) error {

	// Fill Document's `Includes/Namespaces/Constants/Structs/...`,
	// just for easy to use.
	for _, st := range document.Body {
		switch v := st.([]interface{})[0].(type) {
		case *Include:
			document.Includes = append(document.Includes, v)
		case *CppInclude:
			document.CppIncludes = append(document.CppIncludes, v)
		case *Namespace:
			document.Namespaces[v.Language] = v
		case *Constant:
			document.Constants[v.Name] = v
		case *Type:
			switch v.Category {
			case CategoryEnum:
				document.Enums[v.Name] = v
			case CategoryTypedef:
				document.Typedefs[v.Name] = v
			case CategoryStruct:
				document.Structs[v.Name] = v
			case CategoryUnion:
				document.Unions[v.Name] = v
			case CategoryException:
				document.Exceptions[v.Name] = v
			}
			if v.Category != CategoryIdentifier {
				document.Types[v.Name] = v
			}
		case *Service:
			document.Services[v.Name] = v
		default:
			return fmt.Errorf("parser: unknown value %#v", v)
		}
	}

	resolveIdentifierType(document)
	return nil
}
