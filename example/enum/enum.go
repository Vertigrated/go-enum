// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-07-12T00:01:07-03:00
// by go-enum
package enum

import (
	"strings"
)

// Base value index for internal validations
var (
	enumIndex = map[string]valueIndex{}
)

type valueIndex map[string]struct{}

// stringEnumValue is a typed designed to be embeddable in other structs so
// they can expose type safe string enumerations and all of their
// corresponding marshaling and unmarshaling operations
//
// This type is NOT meant to be consumed directly
type stringEnumValue struct {
	value string
	key   string
}

func fromValue(value string, ignoreCase bool, key string) (stringEnumValue, bool) {
	if index, ok := enumIndex[key]; ok {
		if ignoreCase {
			value = strings.ToUpper(value)
		}
		for v := range index {
			if ignoreCase {
				v = strings.ToUpper(v)
			}
			if v == value {
				return stringEnumValue{v, key}, true
			}
		}
	}

	return stringEnumValue{}, false
}

func (e stringEnumValue) Equals(s string) bool {
	return e.value == s
}

func (e stringEnumValue) EqualsIgnoreCase(s string) bool {
	return strings.ToUpper(e.value) == strings.ToUpper(s)
}

func (e stringEnumValue) IsUndefined() bool { return e.value == "" && e.key == "" }
