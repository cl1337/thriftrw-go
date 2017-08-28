// Code generated by thriftrw v1.7.0. DO NOT EDIT.
// @generated

package constants

import (
	"go.uber.org/thriftrw/gen/testdata/containers"
	"go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/gen/testdata/exceptions"
	"go.uber.org/thriftrw/gen/testdata/structs"
	"go.uber.org/thriftrw/gen/testdata/typedefs"
	"go.uber.org/thriftrw/gen/testdata/unions"
	"go.uber.org/thriftrw/ptr"
)

const Home enums.RecordType = enums.RecordTypeHomeAddress

const Name enums.RecordType = enums.RecordTypeName

const WorkAddress enums.RecordType = enums.RecordTypeWorkAddress

var ArbitraryValue *unions.ArbitraryValue = &unions.ArbitraryValue{ListValue: []*unions.ArbitraryValue{&unions.ArbitraryValue{BoolValue: ptr.Bool(true)}, &unions.ArbitraryValue{Int64Value: ptr.Int64(2)}, &unions.ArbitraryValue{StringValue: ptr.String("hello")}, &unions.ArbitraryValue{MapValue: map[string]*unions.ArbitraryValue{"foo": &unions.ArbitraryValue{StringValue: ptr.String("bar")}}}}}

const BeginningOfTime typedefs.Timestamp = typedefs.Timestamp(0)

var ContainersOfContainers *containers.ContainersOfContainers = &containers.ContainersOfContainers{ListOfLists: [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 6}}, ListOfMaps: []map[int32]int32{map[int32]int32{1: 2, 3: 4, 5: 6}, map[int32]int32{7: 8, 9: 10, 11: 12}}, ListOfSets: []map[int32]struct{}{map[int32]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}, map[int32]struct{}{4: struct{}{}, 5: struct{}{}, 6: struct{}{}}}, MapOfListToSet: []struct {
	Key   []int32
	Value map[int64]struct{}
}{{Key: []int32{1, 2, 3}, Value: map[int64]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}}, {Key: []int32{4, 5, 6}, Value: map[int64]struct{}{4: struct{}{}, 5: struct{}{}, 6: struct{}{}}}}, MapOfMapToInt: []struct {
	Key   map[string]int32
	Value int64
}{{Key: map[string]int32{"1": 1, "2": 2, "3": 3}, Value: 100}, {Key: map[string]int32{"4": 4, "5": 5, "6": 6}, Value: 200}}, MapOfSetToListOfDouble: []struct {
	Key   map[int32]struct{}
	Value []float64
}{{Key: map[int32]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}, Value: []float64{1.2, 3.4}}, {Key: map[int32]struct{}{4: struct{}{}, 5: struct{}{}, 6: struct{}{}}, Value: []float64{5.6, 7.8}}}, SetOfLists: [][]string{[]string{"1", "2", "3"}, []string{"4", "5", "6"}}, SetOfMaps: []map[string]string{map[string]string{"1": "2", "3": "4", "5": "6"}, map[string]string{"7": "8", "9": "10", "11": "12"}}, SetOfSets: []map[string]struct{}{map[string]struct{}{"1": struct{}{}, "2": struct{}{}, "3": struct{}{}}, map[string]struct{}{"4": struct{}{}, "5": struct{}{}, "6": struct{}{}}}}

var EmptyException *exceptions.EmptyException = &exceptions.EmptyException{}

var EnumContainers *containers.EnumContainers = &containers.EnumContainers{ListOfEnums: []enums.EnumDefault{enums.EnumDefaultBar, enums.EnumDefaultFoo}, MapOfEnums: map[enums.EnumWithDuplicateValues]int32{enums.EnumWithDuplicateValuesP: 1, enums.EnumWithDuplicateValuesQ: 2}, SetOfEnums: map[enums.EnumWithValues]struct{}{enums.EnumWithValuesX: struct{}{}, enums.EnumWithValuesY: struct{}{}}}

var FrameGroup typedefs.FrameGroup = typedefs.FrameGroup{&structs.Frame{Size: &structs.Size{Height: 200, Width: 100}, TopLeft: &structs.Point{X: 1, Y: 2}}, &structs.Frame{Size: &structs.Size{Height: 400, Width: 300}, TopLeft: &structs.Point{X: 3, Y: 4}}}

var Graph *structs.Graph = &structs.Graph{Edges: []*structs.Edge{&structs.Edge{EndPoint: &structs.Point{X: 3, Y: 4}, StartPoint: &structs.Point{X: 1, Y: 2}}, &structs.Edge{EndPoint: &structs.Point{X: 7, Y: 8}, StartPoint: &structs.Point{X: 5, Y: 6}}}}

var I128 *typedefs.I128 = &typedefs.I128{High: 1234, Low: 5678}

var LastNode *structs.Node = &structs.Node{Value: 3}

const Lower enums.LowerCaseEnum = enums.LowerCaseEnumItems

const MyEnum typedefs.MyEnum = typedefs.MyEnum(enums.EnumWithValuesY)

var Node *structs.Node = &structs.Node{Tail: &structs.List{Tail: &structs.List{Value: 3}, Value: 2}, Value: 1}

var PrimitiveContainers *containers.PrimitiveContainers = &containers.PrimitiveContainers{ListOfInts: []int64{1, 2, 3}, MapOfIntToString: map[int32]string{1: "1", 2: "2", 3: "3"}, MapOfStringToBool: map[string]bool{"1": false, "2": true, "3": true}, SetOfBytes: map[int8]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}, SetOfStrings: map[string]struct{}{"foo": struct{}{}, "bar": struct{}{}}}

func _EnumDefault_ptr(v enums.EnumDefault) *enums.EnumDefault {
	return &v
}

var StructWithOptionalEnum *enums.StructWithOptionalEnum = &enums.StructWithOptionalEnum{E: _EnumDefault_ptr(enums.EnumDefaultBaz)}

var UUID *typedefs.UUID = &typedefs.UUID{High: 1234, Low: 5678}
