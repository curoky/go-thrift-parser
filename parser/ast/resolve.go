package ast

import "fmt"

func (document *Document) Resolve(thrift *Thrift) error {

	// Fill Document's `Includes/Namespaces/Constants/Structs/...`,
	// just for easy to use.
	for _, st := range document.Body {
		switch v := st.([]interface{})[0].(type) {
		case *Include:
			document.Includes = append(document.Includes, v)
		case CppInclude:
			document.CppIncludes = append(document.CppIncludes, CppInclude(v))
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
		case *Service:
			document.Services[v.Name] = v
		default:
			return fmt.Errorf("parser: unknown value %#v", v)
		}
	}
	return nil
}
