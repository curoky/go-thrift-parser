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

// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package ast

import (
	"errors"
	"fmt"
)

const (
	// CategoryVoid is a Category of type Void.
	CategoryVoid Category = iota
	// CategoryConstant is a Category of type Constant.
	CategoryConstant
	// CategoryBool is a Category of type Bool.
	CategoryBool
	// CategoryByte is a Category of type Byte.
	CategoryByte
	// CategoryI16 is a Category of type I16.
	CategoryI16
	// CategoryI32 is a Category of type I32.
	CategoryI32
	// CategoryI64 is a Category of type I64.
	CategoryI64
	// CategoryDouble is a Category of type Double.
	CategoryDouble
	// CategoryString is a Category of type String.
	CategoryString
	// CategoryBinary is a Category of type Binary.
	CategoryBinary
	// CategoryMap is a Category of type Map.
	CategoryMap
	// CategoryList is a Category of type List.
	CategoryList
	// CategorySet is a Category of type Set.
	CategorySet
	// CategoryEnum is a Category of type Enum.
	CategoryEnum
	// CategoryStruct is a Category of type Struct.
	CategoryStruct
	// CategoryUnion is a Category of type Union.
	CategoryUnion
	// CategoryException is a Category of type Exception.
	CategoryException
	// CategoryService is a Category of type Service.
	CategoryService
	// CategoryTypedef is a Category of type Typedef.
	CategoryTypedef
	// CategoryIdentifier is a Category of type Identifier.
	CategoryIdentifier
	// CategoryUnknown is a Category of type Unknown.
	CategoryUnknown
)

var ErrInvalidCategory = errors.New("not a valid Category")

const _CategoryName = "VoidConstantBoolByteI16I32I64DoubleStringBinaryMapListSetEnumStructUnionExceptionServiceTypedefIdentifierUnknown"

var _CategoryMap = map[Category]string{
	CategoryVoid:       _CategoryName[0:4],
	CategoryConstant:   _CategoryName[4:12],
	CategoryBool:       _CategoryName[12:16],
	CategoryByte:       _CategoryName[16:20],
	CategoryI16:        _CategoryName[20:23],
	CategoryI32:        _CategoryName[23:26],
	CategoryI64:        _CategoryName[26:29],
	CategoryDouble:     _CategoryName[29:35],
	CategoryString:     _CategoryName[35:41],
	CategoryBinary:     _CategoryName[41:47],
	CategoryMap:        _CategoryName[47:50],
	CategoryList:       _CategoryName[50:54],
	CategorySet:        _CategoryName[54:57],
	CategoryEnum:       _CategoryName[57:61],
	CategoryStruct:     _CategoryName[61:67],
	CategoryUnion:      _CategoryName[67:72],
	CategoryException:  _CategoryName[72:81],
	CategoryService:    _CategoryName[81:88],
	CategoryTypedef:    _CategoryName[88:95],
	CategoryIdentifier: _CategoryName[95:105],
	CategoryUnknown:    _CategoryName[105:112],
}

// String implements the Stringer interface.
func (x Category) String() string {
	if str, ok := _CategoryMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Category(%d)", x)
}

var _CategoryValue = map[string]Category{
	_CategoryName[0:4]:     CategoryVoid,
	_CategoryName[4:12]:    CategoryConstant,
	_CategoryName[12:16]:   CategoryBool,
	_CategoryName[16:20]:   CategoryByte,
	_CategoryName[20:23]:   CategoryI16,
	_CategoryName[23:26]:   CategoryI32,
	_CategoryName[26:29]:   CategoryI64,
	_CategoryName[29:35]:   CategoryDouble,
	_CategoryName[35:41]:   CategoryString,
	_CategoryName[41:47]:   CategoryBinary,
	_CategoryName[47:50]:   CategoryMap,
	_CategoryName[50:54]:   CategoryList,
	_CategoryName[54:57]:   CategorySet,
	_CategoryName[57:61]:   CategoryEnum,
	_CategoryName[61:67]:   CategoryStruct,
	_CategoryName[67:72]:   CategoryUnion,
	_CategoryName[72:81]:   CategoryException,
	_CategoryName[81:88]:   CategoryService,
	_CategoryName[88:95]:   CategoryTypedef,
	_CategoryName[95:105]:  CategoryIdentifier,
	_CategoryName[105:112]: CategoryUnknown,
}

// ParseCategory attempts to convert a string to a Category.
func ParseCategory(name string) (Category, error) {
	if x, ok := _CategoryValue[name]; ok {
		return x, nil
	}
	return Category(0), fmt.Errorf("%s is %w", name, ErrInvalidCategory)
}

// MarshalText implements the text marshaller method.
func (x Category) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Category) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseCategory(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ConstTypeDouble is a ConstType of type Double.
	ConstTypeDouble ConstType = iota
	// ConstTypeInt is a ConstType of type Int.
	ConstTypeInt
	// ConstTypeLiteral is a ConstType of type Literal.
	ConstTypeLiteral
	// ConstTypeIdentifier is a ConstType of type Identifier.
	ConstTypeIdentifier
	// ConstTypeList is a ConstType of type List.
	ConstTypeList
	// ConstTypeMap is a ConstType of type Map.
	ConstTypeMap
)

var ErrInvalidConstType = errors.New("not a valid ConstType")

const _ConstTypeName = "DoubleIntLiteralIdentifierListMap"

var _ConstTypeMap = map[ConstType]string{
	ConstTypeDouble:     _ConstTypeName[0:6],
	ConstTypeInt:        _ConstTypeName[6:9],
	ConstTypeLiteral:    _ConstTypeName[9:16],
	ConstTypeIdentifier: _ConstTypeName[16:26],
	ConstTypeList:       _ConstTypeName[26:30],
	ConstTypeMap:        _ConstTypeName[30:33],
}

// String implements the Stringer interface.
func (x ConstType) String() string {
	if str, ok := _ConstTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ConstType(%d)", x)
}

var _ConstTypeValue = map[string]ConstType{
	_ConstTypeName[0:6]:   ConstTypeDouble,
	_ConstTypeName[6:9]:   ConstTypeInt,
	_ConstTypeName[9:16]:  ConstTypeLiteral,
	_ConstTypeName[16:26]: ConstTypeIdentifier,
	_ConstTypeName[26:30]: ConstTypeList,
	_ConstTypeName[30:33]: ConstTypeMap,
}

// ParseConstType attempts to convert a string to a ConstType.
func ParseConstType(name string) (ConstType, error) {
	if x, ok := _ConstTypeValue[name]; ok {
		return x, nil
	}
	return ConstType(0), fmt.Errorf("%s is %w", name, ErrInvalidConstType)
}

// MarshalText implements the text marshaller method.
func (x ConstType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ConstType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseConstType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// FieldTypeDefault is a FieldType of type Default.
	FieldTypeDefault FieldType = iota
	// FieldTypeRequired is a FieldType of type Required.
	FieldTypeRequired
	// FieldTypeOptional is a FieldType of type Optional.
	FieldTypeOptional
)

var ErrInvalidFieldType = errors.New("not a valid FieldType")

const _FieldTypeName = "DefaultRequiredOptional"

var _FieldTypeMap = map[FieldType]string{
	FieldTypeDefault:  _FieldTypeName[0:7],
	FieldTypeRequired: _FieldTypeName[7:15],
	FieldTypeOptional: _FieldTypeName[15:23],
}

// String implements the Stringer interface.
func (x FieldType) String() string {
	if str, ok := _FieldTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("FieldType(%d)", x)
}

var _FieldTypeValue = map[string]FieldType{
	_FieldTypeName[0:7]:   FieldTypeDefault,
	_FieldTypeName[7:15]:  FieldTypeRequired,
	_FieldTypeName[15:23]: FieldTypeOptional,
}

// ParseFieldType attempts to convert a string to a FieldType.
func ParseFieldType(name string) (FieldType, error) {
	if x, ok := _FieldTypeValue[name]; ok {
		return x, nil
	}
	return FieldType(0), fmt.Errorf("%s is %w", name, ErrInvalidFieldType)
}

// MarshalText implements the text marshaller method.
func (x FieldType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *FieldType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseFieldType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
