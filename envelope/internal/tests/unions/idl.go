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

package unions

import (
	"go.uber.org/thriftrw/envelope/internal/tests/typedefs"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "unions",
	Package:  "go.uber.org/thriftrw/envelope/internal/tests/unions",
	FilePath: "unions.thrift",
	SHA1:     "ae222d8ff3a8efe55b3d2ebb0c70b4565255623c",
	Includes: []*thriftreflect.ThriftModule{
		typedefs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./typedefs.thrift\"\n\nunion EmptyUnion {}\n\nunion Document {\n    1: typedefs.PDF pdf\n    2: string plainText\n}\n\n/**\n * ArbitraryValue allows constructing complex values without a schema.\n *\n * A value is one of,\n *\n * * Boolean\n * * Integer\n * * String\n * * A list of other values\n * * A dictionary of other values\n */\nunion ArbitraryValue {\n    1: bool boolValue\n    2: i64 int64Value\n    3: string stringValue\n    4: list<ArbitraryValue> listValue\n    5: map<string, ArbitraryValue> mapValue\n}\n"
