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

package services

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Cache_ClearAfter_Args represents the arguments for the Cache.clearAfter function.
//
// The arguments for clearAfter are sent and received over the wire as this struct.
type Cache_ClearAfter_Args struct {
	DurationMS *int64 `json:"durationMS,omitempty"`
}

// ToWire translates a Cache_ClearAfter_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Cache_ClearAfter_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.DurationMS != nil {
		w, err = wire.NewValueI64(*(v.DurationMS)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Cache_ClearAfter_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Cache_ClearAfter_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Cache_ClearAfter_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Cache_ClearAfter_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.DurationMS = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Cache_ClearAfter_Args
// struct.
func (v *Cache_ClearAfter_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.DurationMS != nil {
		fields[i] = fmt.Sprintf("DurationMS: %v", *(v.DurationMS))
		i++
	}

	return fmt.Sprintf("Cache_ClearAfter_Args{%v}", strings.Join(fields[:i], ", "))
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Cache_ClearAfter_Args match the
// provided Cache_ClearAfter_Args.
//
// This function performs a deep comparison.
func (v *Cache_ClearAfter_Args) Equals(rhs *Cache_ClearAfter_Args) bool {
	if !_I64_EqualsPtr(v.DurationMS, rhs.DurationMS) {
		return false
	}

	return true
}

// GetDurationMS returns the value of DurationMS if it is set or its
// zero value if it is unset.
func (v *Cache_ClearAfter_Args) GetDurationMS() (o int64) {
	if v.DurationMS != nil {
		return *v.DurationMS
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "clearAfter" for this struct.
func (v *Cache_ClearAfter_Args) MethodName() string {
	return "clearAfter"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *Cache_ClearAfter_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// Cache_ClearAfter_Helper provides functions that aid in handling the
// parameters and return values of the Cache.clearAfter
// function.
var Cache_ClearAfter_Helper = struct {
	// Args accepts the parameters of clearAfter in-order and returns
	// the arguments struct for the function.
	Args func(
		durationMS *int64,
	) *Cache_ClearAfter_Args
}{}

func init() {
	Cache_ClearAfter_Helper.Args = func(
		durationMS *int64,
	) *Cache_ClearAfter_Args {
		return &Cache_ClearAfter_Args{
			DurationMS: durationMS,
		}
	}

}
