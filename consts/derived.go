package consts

import (
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewlang/llvm/types"
)

// Vector represents a vector constant which is a vetor containing only
// constants.
//
// Examples:
//    <i32 37, i32 42>   ; type: <2 x i32>
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Vector struct {
	typ   *types.Vector
	elems []Constant
}

// NewVector returns a vector constant based on the given vector type and vector
// elements.
func NewVector(typ types.Type, elems []Constant) (*Vector, error) {
	v := new(Vector)
	var ok bool
	v.typ, ok = typ.(*types.Vector)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected vector type", typ)
	}

	// Verify vector element types.
	for _, elem := range elems {
		if !elem.Type().Equal(v.typ.Elem()) {
			return nil, errutil.Newf("invalid vector element type %q; expected %q", elem.Type(), v.typ.Elem())
		}
	}
	v.elems = elems

	return v, nil
}

// Array represents an array constant which is an array containing only
// constants.
//
// Examples:
//    [double 3.14, double 1.5]      ; type: [2 x double]
//    [<2 x i32> <i32 15, i32 20>]   ; type: [1 x <2 x i32>]
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Array struct {
	typ   *types.Array
	elems []Constant
}

// NewArray returns an array constant based on the given array type and array
// elements.
func NewArray(typ types.Type, elems []Constant) (*Array, error) {
	v := new(Array)
	var ok bool
	v.typ, ok = typ.(*types.Array)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected array type", typ)
	}

	// Verify array element types.
	for _, elem := range elems {
		if !elem.Type().Equal(v.typ.Elem()) {
			return nil, errutil.Newf("invalid array element type %q; expected %q", elem.Type(), v.typ.Elem())
		}
	}
	v.elems = elems

	return v, nil
}

// Struct represents a structure constant which is a structure containing only
// constants.
//
// Examples:
//    {i32 7, i8 3}                    ; type: {i32, i8}
//    {i32 7, {i8, i8} {i8 3, i8 5}}   ; type: {i32, {i8, i8}}
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Struct struct {
	typ    *types.Struct
	fields []Constant
}

// NewStruct returns a structure constant based on the given structure type and
// structure fields.
func NewStruct(typ types.Type, fields []Constant) (*Struct, error) {
	v := new(Struct)
	var ok bool
	v.typ, ok = typ.(*types.Struct)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected structure type", typ)
	}

	// Verify structure field types.
	fieldTypes := v.typ.Fields()
	if len(fields) != len(fieldTypes) {
		return nil, errutil.Newf("incorrect number of fields in structure constant; expected %d, got %d", len(fieldTypes), len(fields))
	}
	for i := range fields {
		if !fieldTypes[i].Equal(fields[i].Type()) {
			return nil, errutil.Newf("invalid field type %q in structure constant; expected %q", fields[i].Type(), fieldTypes[i])
		}
	}
	v.fields = fields

	return v, nil
}
