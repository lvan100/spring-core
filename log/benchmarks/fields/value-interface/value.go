/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package value_interface

import (
	"fields/encoder"
)

// Value is an interface for types that can encode themselves using an Encoder.
type Value interface {
	Encode(enc encoder.Encoder)
}

// BoolValue represents a bool carried by Field.
type BoolValue bool

// Encode encodes the data represented by v to an Encoder.
func (v BoolValue) Encode(enc encoder.Encoder) {
	enc.AppendBool(bool(v))
}

// Int64Value represents an int64 carried by Field.
type Int64Value int64

// Encode encodes the data represented by v to an Encoder.
func (v Int64Value) Encode(enc encoder.Encoder) {
	enc.AppendInt64(int64(v))
}

// Uint64Value represents an uint64 carried by Field.
type Uint64Value uint64

// Encode encodes the data represented by v to an Encoder.
func (v Uint64Value) Encode(enc encoder.Encoder) {
	enc.AppendUint64(uint64(v))
}

// Float64Value represents a float64 carried by Field.
type Float64Value float64

// Encode encodes the data represented by v to an Encoder.
func (v Float64Value) Encode(enc encoder.Encoder) {
	enc.AppendFloat64(float64(v))
}

// StringValue represents a string carried by Field.
type StringValue string

// Encode encodes the data represented by v to an Encoder.
func (v StringValue) Encode(enc encoder.Encoder) {
	enc.AppendString(string(v))
}

// ReflectValue represents an interface{} carried by Field.
type ReflectValue struct {
	Val interface{}
}

// Encode encodes the data represented by v to an Encoder.
func (v ReflectValue) Encode(enc encoder.Encoder) {
	enc.AppendReflect(v.Val)
}

// BoolsValue represents a slice of bool carried by Field.
type BoolsValue []bool

// Encode encodes the data represented by v to an Encoder.
func (v BoolsValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendBool(val)
	}
	enc.AppendArrayEnd()
}

// IntsValue represents a slice of int carried by Field.
type IntsValue []int

// Encode encodes the data represented by v to an Encoder.
func (v IntsValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendInt64(int64(val))
	}
	enc.AppendArrayEnd()
}

// Int8sValue represents a slice of int8 carried by Field.
type Int8sValue []int8

// Encode encodes the data represented by v to an Encoder.
func (v Int8sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendInt64(int64(val))
	}
	enc.AppendArrayEnd()
}

// Int16sValue represents a slice of int16 carried by Field.
type Int16sValue []int16

// Encode encodes the data represented by v to an Encoder.
func (v Int16sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendInt64(int64(val))
	}
	enc.AppendArrayEnd()
}

// Int32sValue represents a slice of int32 carried by Field.
type Int32sValue []int32

// Encode encodes the data represented by v to an Encoder.
func (v Int32sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendInt64(int64(val))
	}
	enc.AppendArrayEnd()
}

// Int64sValue represents a slice of int64 carried by Field.
type Int64sValue []int64

// Encode encodes the data represented by v to an Encoder.
func (v Int64sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendInt64(val)
	}
	enc.AppendArrayEnd()
}

// UintsValue represents a slice of uint carried by Field.
type UintsValue []uint

// Encode encodes the data represented by v to an Encoder.
func (v UintsValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendUint64(uint64(val))
	}
	enc.AppendArrayEnd()
}

// Uint8sValue represents a slice of uint8 carried by Field.
type Uint8sValue []uint8

// Encode encodes the data represented by v to an Encoder.
func (v Uint8sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendUint64(uint64(val))
	}
	enc.AppendArrayEnd()
}

// Uint16sValue represents a slice of uint16 carried by Field.
type Uint16sValue []uint16

// Encode encodes the data represented by v to an Encoder.
func (v Uint16sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendUint64(uint64(val))
	}
	enc.AppendArrayEnd()
}

// Uint32sValue represents a slice of uint32 carried by Field.
type Uint32sValue []uint32

// Encode encodes the data represented by v to an Encoder.
func (v Uint32sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendUint64(uint64(val))
	}
	enc.AppendArrayEnd()
}

// Uint64sValue represents a slice of uint64 carried by Field.
type Uint64sValue []uint64

// Encode encodes the data represented by v to an Encoder.
func (v Uint64sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendUint64(val)
	}
	enc.AppendArrayEnd()
}

// Float32sValue represents a slice of float32 carried by Field.
type Float32sValue []float32

// Encode encodes the data represented by v to an Encoder.
func (v Float32sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendFloat64(float64(val))
	}
	enc.AppendArrayEnd()
}

// Float64sValue represents a slice of float64 carried by Field.
type Float64sValue []float64

// Encode encodes the data represented by v to an Encoder.
func (v Float64sValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendFloat64(val)
	}
	enc.AppendArrayEnd()
}

// StringsValue represents a slice of string carried by Field.
type StringsValue []string

// Encode encodes the data represented by v to an Encoder.
func (v StringsValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		enc.AppendString(val)
	}
	enc.AppendArrayEnd()
}

// ArrayValue represents a slice of Value carried by Field.
type ArrayValue []Value

// Encode encodes the data represented by v to an Encoder.
func (v ArrayValue) Encode(enc encoder.Encoder) {
	enc.AppendArrayBegin()
	for _, val := range v {
		val.Encode(enc)
	}
	enc.AppendArrayEnd()
}

// ObjectValue represents a slice of Field carried by Field.
type ObjectValue []Field

// Encode encodes the data represented by v to an Encoder.
func (v ObjectValue) Encode(enc encoder.Encoder) {
	enc.AppendObjectBegin()
	WriteFields(enc, v)
	enc.AppendObjectEnd()
}

// WriteFields writes a slice of Field objects to the encoder.
func WriteFields(enc encoder.Encoder, fields []Field) {
	for _, f := range fields {
		enc.AppendKey(f.Key)
		f.Val.Encode(enc)
	}
}
