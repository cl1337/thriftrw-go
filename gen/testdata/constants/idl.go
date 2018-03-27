// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package constants

import (
	"go.uber.org/thriftrw/gen/testdata/containers"
	"go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/gen/testdata/exceptions"
	"go.uber.org/thriftrw/gen/testdata/other_constants"
	"go.uber.org/thriftrw/gen/testdata/structs"
	"go.uber.org/thriftrw/gen/testdata/typedefs"
	"go.uber.org/thriftrw/gen/testdata/unions"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "constants",
	Package:  "go.uber.org/thriftrw/gen/testdata/constants",
	FilePath: "constants.thrift",
	SHA1:     "c174d3c52157a0b78e14ad7677b2128174517eb6",
	Includes: []*thriftreflect.ThriftModule{
		containers.ThriftModule,
		enums.ThriftModule,
		exceptions.ThriftModule,
		other_constants.ThriftModule,
		structs.ThriftModule,
		typedefs.ThriftModule,
		unions.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./other_constants.thrift\"\ninclude \"./containers.thrift\"\ninclude \"./enums.thrift\"\ninclude \"./exceptions.thrift\"\ninclude \"./structs.thrift\"\ninclude \"./unions.thrift\"\ninclude \"./typedefs.thrift\"\n\nconst containers.PrimitiveContainers primitiveContainers = {\n    \"listOfInts\": other_constants.listOfInts, // imported constant\n    \"setOfStrings\": [\"foo\", \"bar\"],\n    \"setOfBytes\": other_constants.listOfInts, // imported constant with type casting\n    \"mapOfIntToString\": {\n        1: \"1\",\n        2: \"2\",\n        3: \"3\",\n    },\n    \"mapOfStringToBool\": {\n        \"1\": 0,\n        \"2\": 1,\n        \"3\": 1,\n    }\n}\n\nconst containers.EnumContainers enumContainers = {\n    \"listOfEnums\": [1, enums.EnumDefault.Foo],\n    \"setOfEnums\": [123, enums.EnumWithValues.Y],\n    \"mapOfEnums\": {\n        0: 1,\n        enums.EnumWithDuplicateValues.Q: 2,\n    },\n}\n\nconst containers.ContainersOfContainers containersOfContainers = {\n    \"listOfLists\": [[1, 2, 3], [4, 5, 6]],\n    \"listOfSets\": [[1, 2, 3], [4, 5, 6]],\n    \"listOfMaps\": [{1: 2, 3: 4, 5: 6}, {7: 8, 9: 10, 11: 12}],\n    \"setOfSets\": [[\"1\", \"2\", \"3\"], [\"4\", \"5\", \"6\"]],\n    \"setOfLists\": [[\"1\", \"2\", \"3\"], [\"4\", \"5\", \"6\"]],\n    \"setOfMaps\": [\n        {\"1\": \"2\", \"3\": \"4\", \"5\": \"6\"},\n        {\"7\": \"8\", \"9\": \"10\", \"11\": \"12\"},\n    ],\n    \"mapOfMapToInt\": {\n        {\"1\": 1, \"2\": 2, \"3\": 3}: 100,\n        {\"4\": 4, \"5\": 5, \"6\": 6}: 200,\n    },\n    \"mapOfListToSet\": {\n        // more type casting\n        other_constants.listOfInts: other_constants.listOfInts,\n        [4, 5, 6]: [4, 5, 6],\n    },\n    \"mapOfSetToListOfDouble\": {\n        [1, 2, 3]: [1.2, 3.4],\n        [4, 5, 6]: [5.6, 7.8],\n    },\n}\n\nconst enums.StructWithOptionalEnum structWithOptionalEnum = {\n    \"e\": enums.EnumDefault.Baz\n}\n\nconst exceptions.EmptyException emptyException = {}\n\nconst structs.Graph graph = {\n    \"edges\": [\n        {\"startPoint\": other_constants.some_point, \"endPoint\": {\"x\": 3, \"y\": 4}},\n        {\"startPoint\": {\"x\": 5, \"y\": 6}, \"endPoint\": {\"x\": 7, \"y\": 8}},\n    ]\n}\n\nconst structs.Node lastNode = {\"value\": 3}\nconst structs.Node node = {\n    \"value\": 1,\n    \"tail\": {\"value\": 2, \"tail\": lastNode},\n}\n\nconst unions.ArbitraryValue arbitraryValue = {\n    \"listValue\": [\n        {\"boolValue\": 1},\n        {\"int64Value\": 2},\n        {\"stringValue\": \"hello\"},\n        {\"mapValue\": {\"foo\": {\"stringValue\": \"bar\"}}},\n    ],\n}\n// TODO: union validation for constants?\n\nconst typedefs.i128 i128 = uuid\nconst typedefs.UUID uuid = {\"high\": 1234, \"low\": 5678}\n\n/** Timestamp at which time began. */\nconst typedefs.Timestamp beginningOfTime = 0\n\n/**\n * An example frame group.\n *\n * Contains two frames.\n */\nconst typedefs.FrameGroup frameGroup = [\n    {\n        \"topLeft\": {\"x\": 1, \"y\": 2},\n        \"size\": {\"width\": 100, \"height\": 200},\n    }\n    {\n        \"topLeft\": {\"x\": 3, \"y\": 4},\n        \"size\": {\"width\": 300, \"height\": 400},\n    },\n]\n\nconst typedefs.MyEnum myEnum = enums.EnumWithValues.Y\n\nconst enums.RecordType NAME = enums.RecordType.NAME\nconst enums.RecordType HOME = enums.RecordType.HOME_ADDRESS\nconst enums.RecordType WORK_ADDRESS = enums.RecordType.WORK_ADDRESS\n\nconst enums.lowerCaseEnum lower = enums.lowerCaseEnum.items\n"
