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

//go:generate go-enum --marshal -f=$GOFILE -a "+:Plus,#:Sharp"

package ast

// ENUM(
// Void,
// Constant,
// Bool
// Byte
// I16
// I32
// I64
// Double
// String
// Binary
// Map
// List
// Set
// Enum
// Struct
// Union
// Exception
// Service
// Typedef
// Identifier
// Unknown
// )
type Category int32

type SourceInfo struct {
	Line, Col, Offset int
	Text              string
}

type Annotation struct {
	Name  string
	Value string

	SourceInfo *SourceInfo
}

type Type struct {
	Name        string
	Category    Category
	Annotations map[string]*Annotation
	Belong      *Document `json:"-"`

	// for list/set/map
	KeyType   *Type
	ValueType *Type

	// for typedef or reference
	PreRefType   *Type
	FinalRefType *Type

	// for struct
	Fields []*Field

	// for enum
	Values []*EnumValue

	ExpandCategoryStr string

	SourceInfo *SourceInfo
}

func (t *Type) FinalType() *Type {
	if t.FinalRefType != nil {
		return t.FinalRefType
	}
	return t
}

type Namespace struct {
	Name        string
	Language    string
	Annotations map[string]*Annotation

	SourceInfo *SourceInfo
}

type EnumValue struct {
	Name        string
	Value       int64
	Annotations map[string]*Annotation

	SourceInfo *SourceInfo
}

// ENUM(
// Double,
// Int,
// Literal,
// Identifier,
// List,
// Map
// )
type ConstType int32

type ConstValueExtra struct {
	Name   string
	IsEnum bool
	Index  int64
	Sel    string

	SourceInfo *SourceInfo
}

type ConstValue struct {
	Type       ConstType
	TypedValue *ConstTypedValue
	Extra      *ConstValueExtra

	SourceInfo *SourceInfo
}

type ConstTypedValue struct {
	Double     *float64
	Int        *int64
	Literal    *string
	Identifier *string
	List       []*ConstValue
	Map        []*MapConstValue

	SourceInfo *SourceInfo
}

type MapConstValue struct {
	Key   *ConstValue
	Value *ConstValue

	SourceInfo *SourceInfo
}

type Constant struct {
	Name        string
	Type        *Type
	Value       *ConstValue
	Annotations map[string]*Annotation

	SourceInfo *SourceInfo
}

// ENUM(
// Default,
// Required,
// Optional,
// )
type FieldType int32

type Field struct {
	ID           int64
	Name         string
	Requiredness FieldType
	Type         *Type
	Default      *ConstValue
	Annotations  map[string]*Annotation

	SourceInfo *SourceInfo
}

type Function struct {
	Name        string
	ReturnType  *Type
	Arguments   []*Field
	Exceptions  []*Field
	Annotations map[string]*Annotation

	SourceInfo *SourceInfo
}

type Service struct {
	Name        string
	Extends     string
	Functions   []*Function
	Annotations map[string]*Annotation

	SourceInfo *SourceInfo
}

type Include struct {
	Path      string
	Name      string // short name when include
	Reference *Document

	SourceInfo *SourceInfo
}

type CppInclude struct {
	Name string

	SourceInfo *SourceInfo
}

type Document struct {
	Filename string
	Body     []interface{}

	Includes    []*Include            `json:"-"`
	CppIncludes []*CppInclude         `json:"-"`
	Namespaces  map[string]*Namespace `json:"-"`
	Constants   map[string]*Constant  `json:"-"`
	Enums       map[string]*Type      `json:"-"`
	Typedefs    map[string]*Type      `json:"-"`
	Structs     map[string]*Type      `json:"-"`
	Unions      map[string]*Type      `json:"-"`
	Exceptions  map[string]*Type      `json:"-"`
	Services    map[string]*Service   `json:"-"`

	// TODO: remove this
	AllTypes []*Type `json:"-"`
}

type Thrift struct {
	Documents map[string]*Document
}
