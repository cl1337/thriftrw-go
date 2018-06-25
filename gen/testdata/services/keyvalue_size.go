// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// KeyValue_Size_Args represents the arguments for the KeyValue.size function.
//
// The arguments for size are sent and received over the wire as this struct.
type KeyValue_Size_Args struct {
}

// ToWire translates a KeyValue_Size_Args struct into a Thrift-level intermediate
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
func (v *KeyValue_Size_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_Size_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_Size_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_Size_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_Size_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_Size_Args
// struct.
func (v *KeyValue_Size_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("KeyValue_Size_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_Size_Args match the
// provided KeyValue_Size_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_Size_Args) Equals(rhs *KeyValue_Size_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "size" for this struct.
func (v *KeyValue_Size_Args) MethodName() string {
	return "size"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_Size_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_Size_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.size
// function.
var KeyValue_Size_Helper = struct {
	// Args accepts the parameters of size in-order and returns
	// the arguments struct for the function.
	Args func() *KeyValue_Size_Args

	// IsException returns true if the given error can be thrown
	// by size.
	//
	// An error can be thrown by size only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for size
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// size into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by size
	//
	//   value, err := size(args)
	//   result, err := KeyValue_Size_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from size: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(int64, error) (*KeyValue_Size_Result, error)

	// UnwrapResponse takes the result struct for size
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if size threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := KeyValue_Size_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_Size_Result) (int64, error)
}{}

func init() {
	KeyValue_Size_Helper.Args = func() *KeyValue_Size_Args {
		return &KeyValue_Size_Args{}
	}

	KeyValue_Size_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	KeyValue_Size_Helper.WrapResponse = func(success int64, err error) (*KeyValue_Size_Result, error) {
		if err == nil {
			return &KeyValue_Size_Result{Success: &success}, nil
		}

		return nil, err
	}
	KeyValue_Size_Helper.UnwrapResponse = func(result *KeyValue_Size_Result) (success int64, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// KeyValue_Size_Result represents the result of a KeyValue.size function call.
//
// The result of a size execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type KeyValue_Size_Result struct {
	// Value returned by size after a successful execution.
	Success *int64 `json:"success,omitempty"`
}

// ToWire translates a KeyValue_Size_Result struct into a Thrift-level intermediate
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
func (v *KeyValue_Size_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueI64(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_Size_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_Size_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_Size_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_Size_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_Size_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_Size_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a KeyValue_Size_Result
// struct.
func (v *KeyValue_Size_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("KeyValue_Size_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_Size_Result match the
// provided KeyValue_Size_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_Size_Result) Equals(rhs *KeyValue_Size_Result) bool {
	if !_I64_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *KeyValue_Size_Result) GetSuccess() (o int64) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "size" for this struct.
func (v *KeyValue_Size_Result) MethodName() string {
	return "size"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_Size_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
