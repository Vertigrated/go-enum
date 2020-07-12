// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-07-12T00:16:54-03:00
// by go-enum
package enum

import (
	"fmt"
)

/**
 *
 * init: register enum to the internal index
 *
 **/
//nolint:gochecknoinits // enum has to register itself for unmarshaling at runtime
func init() {
	if _, ok := enumIndex[ghostKey]; ok {
		panic(fmt.Sprintf("enum: enumeration %s is already registered", ghostKey))
	}
	enumIndex[ghostKey] = ghostValueIndex
}

/**
 *
 * Type aliases and declarations
 *
 **/

type Ghost struct {
	stringEnumValue
}

func GhostFromValue(value string, ignoreCase bool) (Ghost, bool) {
	result, found := fromValue(value, ignoreCase, ghostKey)
	return Ghost{result}, found
}

type ghostList []Ghost

// ghostEnum is a type and memory safe iterable enumeration of Ghost values
type ghostEnum struct {
	ghostList
}

func (e ghostEnum) ForEach(f func(int, Ghost)) {
	for i, e := range e.ghostList {
		f(i, e)
	}
}

func (e ghostEnum) Len() int {
	return len(e.ghostList)
}

/**
 *
 * Private value index, key
 *
 **/

var (
	ghostValueIndex = valueIndex{
		"Blinky": {},
		"Clyde":  {},
		"Inky":   {},
		"Pinky":  {},
	}
	ghostKey = "Ghost"
)

/**
 *
 * Public enumeration
 *
 **/

var (
	GhostBlinky = Ghost{stringEnumValue{"Blinky", ghostKey}}
	GhostClyde  = Ghost{stringEnumValue{"Clyde", ghostKey}}
	GhostInky   = Ghost{stringEnumValue{"Inky", ghostKey}}
	GhostPinky  = Ghost{stringEnumValue{"Pinky", ghostKey}}

	EnumGhost = ghostEnum{ghostList{
		GhostBlinky,
		GhostClyde,
		GhostInky,
		GhostPinky,
	}}
)

/**
 *
 * Proxy methods for enum unmarshaling
 *
 **/

func (e *Ghost) UnmarshalJSON(data []byte) error {
	e.key = ghostKey
	return e.stringEnumValue.UnmarshalJSON(data)
}

func (e *Ghost) UnmarshalText(text []byte) error {
	e.key = ghostKey
	return e.stringEnumValue.UnmarshalText(text)
}

func (e *Ghost) UnmarshalBSON(data []byte) error {
	e.key = ghostKey
	return e.stringEnumValue.UnmarshalBSON(data)
}

func (e *Ghost) UnmarshalBinary(data []byte) error {
	e.key = ghostKey
	return e.stringEnumValue.UnmarshalBinary(data)
}

func (e *Ghost) GobDecode(data []byte) error {
	e.key = ghostKey
	return e.stringEnumValue.GobDecode(data)
}

func (e *Ghost) Scan(raw interface{}) error {
	e.key = ghostKey
	return e.stringEnumValue.Scan(raw)
}
