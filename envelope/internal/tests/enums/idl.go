// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package enums

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "enums",
	Package:  "go.uber.org/thriftrw/envelope/internal/tests/enums",
	FilePath: "enums.thrift",
	SHA1:     "e5f5f95817a87808e0963a9340b3e67f9a09cdc6",
	Raw:      rawIDL,
}

const rawIDL = "enum EmptyEnum {}\n\nenum EnumDefault {\n    Foo, Bar, Baz\n}\n\nenum EnumWithValues {\n    X = 123,\n    Y = 456,\n    Z = 789,\n}\n\nenum EnumWithDuplicateValues {\n    P, // 0\n    Q = -1,\n    R, // 0\n}\n\n// enum with item names conflicting with those of another enum\nenum EnumWithDuplicateName {\n    A, B, C, P, Q, R, X, Y, Z\n}\n\n// Enum treated as optional inside a struct\nstruct StructWithOptionalEnum {\n    1: optional EnumDefault e\n}\n\n/**\n * Kinds of records stored in the database.\n */\nenum RecordType {\n  /** Name of the user. */\n  NAME,\n\n  /**\n   * Home address of the user.\n   *\n   * This record is always present.\n   */\n  HOME_ADDRESS,\n\n  /**\n   * Home address of the user.\n   *\n   * This record may not be present.\n   */\n  WORK_ADDRESS\n}\n\nenum lowerCaseEnum {\n    containing, lower_case, items\n}\n\n// collision with RecordType_Values() function.\nenum RecordType_Values { FOO, BAR }\n"
