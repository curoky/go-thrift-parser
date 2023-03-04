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

	for _, typ := range document.Types {
		for _, field := range typ.Fields {
			if field.Type.Category == CategoryIdentifier {
				if originalType, exists := document.Types[field.Type.Name]; exists {
					field.Type = originalType
				}
			}
		}
	}
	return nil
}
