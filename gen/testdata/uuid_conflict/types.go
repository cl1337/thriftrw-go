// Code generated by thriftrw v1.2.0
// @generated

package uuid_conflict

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/gen/testdata/typedefs"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type UUID string

func (v UUID) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

func (v UUID) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

func (v *UUID) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (UUID)(x)
	return err
}

func (lhs UUID) Equals(rhs UUID) bool {
	return (lhs == rhs)
}

type UUIDConflict struct {
	LocalUUID    UUID           `json:"localUUID"`
	ImportedUUID *typedefs.UUID `json:"importedUUID"`
}

func (v *UUIDConflict) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.LocalUUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.ImportedUUID == nil {
		return w, errors.New("field ImportedUUID of UUIDConflict is required")
	}
	w, err = v.ImportedUUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UUID_Read(w wire.Value) (UUID, error) {
	var x UUID
	err := x.FromWire(w)
	return x, err
}

func _UUID_1_Read(w wire.Value) (*typedefs.UUID, error) {
	var x typedefs.UUID
	err := x.FromWire(w)
	return &x, err
}

func (v *UUIDConflict) FromWire(w wire.Value) error {
	var err error
	localUUIDIsSet := false
	importedUUIDIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.LocalUUID, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				localUUIDIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ImportedUUID, err = _UUID_1_Read(field.Value)
				if err != nil {
					return err
				}
				importedUUIDIsSet = true
			}
		}
	}
	if !localUUIDIsSet {
		return errors.New("field LocalUUID of UUIDConflict is required")
	}
	if !importedUUIDIsSet {
		return errors.New("field ImportedUUID of UUIDConflict is required")
	}
	return nil
}

func (v *UUIDConflict) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("LocalUUID: %v", v.LocalUUID)
	i++
	fields[i] = fmt.Sprintf("ImportedUUID: %v", v.ImportedUUID)
	i++
	return fmt.Sprintf("UUIDConflict{%v}", strings.Join(fields[:i], ", "))
}

func (v *UUIDConflict) Equals(rhs *UUIDConflict) bool {
	if !(v.LocalUUID == rhs.LocalUUID) {
		return false
	}
	if !v.ImportedUUID.Equals(rhs.ImportedUUID) {
		return false
	}
	return true
}
