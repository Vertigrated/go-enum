package templates

var EnumTemplate = `{{- if not .OmitGeneratedNotice}}
// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
// by go-enum

{{- end }}
package {{ .Package }}

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

// Equals does a case sensitive comparison against a string value
func (e stringEnumValue) Equals(s string) bool {
	return e.value == s
}

// Equals does a case insensitive comparison against a string value
func (e stringEnumValue) EqualsIgnoreCase(s string) bool {
	return strings.ToUpper(e.value) == strings.ToUpper(s)
}

// IsUndefined returns whether the value is an fully initialized enum value or an undefined one (MyEnumType{}).
// Such values are permitted in order to support empty values on compile and (un)marshaling time without pointer fields 
func (e stringEnumValue) IsUndefined() bool { return e.value == "" && e.key == "" }

`
