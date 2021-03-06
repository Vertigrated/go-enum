// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2020-08-02T15:36:10-03:00
// by go-enum

package enum

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

type fooEnumValue struct {
	EnumValue stringEnumValue `json:"enum_value"`
}

type fooEnumValueOmitEmpty struct {
	EnumValue stringEnumValue `json:"enum_value,omitempty"`
}

type fooEnumValuePtr struct {
	EnumValue *stringEnumValue `json:"enum_value"`
}

type fooEnumValuePtrOmitEmpty struct {
	EnumValue *stringEnumValue `json:"enum_value,omitempty"`
}

func TestEnumValue_MarshalJSON(t *testing.T) {
	t.Run("Marshal_AnnonStructField", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		v := struct {
			A stringEnumValue `json:"enum_value"`
		}{c}
		data, err := json.Marshal(&v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"enum_value":"%s"}`, v.A.String()), string(data))
	})
	t.Run("Marshal_StructField", func(t *testing.T) {
		v := fooEnumValue{EnumValue: stringEnumValue{key: countriesIso31661Key}}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"enum_value":"%s"}`, v.EnumValue.String()), string(data))
	})
	t.Run("Marshal_OmitEmptyStruct", func(t *testing.T) {
		// encoding/json ignores omitempty on struct zero values
		// https://github.com/golang/go/issues/11939
		v := fooEnumValueOmitEmpty{}
		data, err := json.Marshal(v)
		require.Nil(t, err)
		assert.EqualValues(t, `{"enum_value":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtr{EnumValue: &stringEnumValue{key: countriesIso31661Key}}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"enum_value":""}`, string(data))
	})
	t.Run("Marshal_StructFieldPtr", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		vPtr := fooEnumValuePtr{EnumValue: &c}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, fmt.Sprintf(`{"enum_value":"%s"}`, vPtr.EnumValue.String()), string(data))
	})
	t.Run("Marshal_StructFieldNilPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtr{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{"enum_value":null}`, string(data))
	})
	t.Run("Marshal_OmitEmptyStructPtr", func(t *testing.T) {
		vPtr := fooEnumValuePtrOmitEmpty{}
		data, err := json.Marshal(vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, `{}`, string(data))
	})
}

func TestEnumValue_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal_InvalidJSON", func(t *testing.T) {
		data := `{"enum_value":"PA`
		rawData := []byte(data)

		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		v := struct {
			A stringEnumValue `json:"enum_value"`
		}{c}
		// Invalid JSON structures fail on json.Unmarshal
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_InvalidValue", func(t *testing.T) {
		data := `{"enum_value":"invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"}`
		rawData := []byte(data)

		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		v := struct {
			A stringEnumValue `json:"enum_value"`
		}{c}
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_InvalidValueJSON", func(t *testing.T) {
		data := `{"enum_value":123}`
		rawData := []byte(data)

		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		v := struct {
			A stringEnumValue `json:"enum_value"`
		}{c}
		// Invalid field values but whithin a valid JSON are caught on UnmarshalJSON
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
	t.Run("Unmarshal_AnnonStructField", func(t *testing.T) {
		data := `{"enum_value":"Ca"}`
		rawData := []byte(data)

		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		v := struct {
			A stringEnumValue `json:"enum_value"`
		}{c}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v.A.String())
	})
	t.Run("Unmarshal_StructField", func(t *testing.T) {
		data := `{"enum_value":"Ca"}`
		rawData := []byte(data)

		v := fooEnumValue{EnumValue: stringEnumValue{key: countriesIso31661Key}}
		err := json.Unmarshal(rawData, &v)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v.EnumValue.String())
	})
	t.Run("Unmarshal_OmitEmptyStruct", func(t *testing.T) {
		data := `{"enum_value":null}`
		rawData := []byte(data)

		v := fooEnumValueOmitEmpty{}
		err := json.Unmarshal(rawData, &v)
		require.NoError(t, err)
	})
	t.Run("Unmarshal_StructFieldPtr", func(t *testing.T) {
		data := `{"enum_value":"Ca"}`
		rawData := []byte(data)

		vPtr := fooEnumValuePtr{EnumValue: &stringEnumValue{key: countriesIso31661Key}}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", vPtr.EnumValue.String())
	})
	t.Run("Unmarshal_StructFieldNilPtr", func(t *testing.T) {
		data := `{"enum_value":null}`
		rawData := []byte(data)

		vPtr := fooEnumValuePtr{}
		err := json.Unmarshal(rawData, &vPtr)
		require.Nil(t, err)
		assert.Nil(t, vPtr.EnumValue)
	})
	t.Run("Unmarshal_TableUnmatch", func(t *testing.T) {
		data := `{"enum_value":"PPP"}`
		rawData := []byte(data)

		v := fooEnumValue{}
		err := json.Unmarshal(rawData, &v)
		require.NotNil(t, err)
	})
}

func TestEnumValue_TextCodec(t *testing.T) {
	t.Run("MarshalText_Valid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		data, err := c.MarshalText()
		require.Nil(t, err)
		require.True(t, len(data) > 0)
	})
	t.Run("UnmarshalText_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.UnmarshalText([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, CountriesIso31661CA.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.UnmarshalText([]byte(""))
		require.NoError(t, err)
		err = c.UnmarshalText([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_Stringer(t *testing.T) {
	c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
	require.EqualValues(t, c.value, c.String())
}

func TestEnumValue_DriverValues(t *testing.T) {
	t.Run("Scan_String", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		require.Nil(t, c.Scan(CountriesIso31661CA.String()))
		require.EqualValues(t, CountriesIso31661CA.String(), c.value)
	})
	t.Run("Scan_Bytes", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		require.Nil(t, c.Scan([]byte(CountriesIso31661CA.String())))
		require.EqualValues(t, CountriesIso31661CA.String(), c.value)
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		require.NotNil(t, c.Scan(1))
	})
	t.Run("Scan_Invalid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		v, err := c.Value()
		assert.Nil(t, err)
		require.EqualValues(t, CountriesIso31661CA.String(), fmt.Sprintf("%v", v))
	})
}

func TestEnumValue_BinaryCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		data, err := c.MarshalBinary()
		require.Nil(t, err)
		assert.Len(t, data, 2)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.UnmarshalBinary([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, CountriesIso31661CA.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.UnmarshalBinary(nil)
		require.NoError(t, err)
		err = c.UnmarshalBinary([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_GobCodec(t *testing.T) {
	t.Run("MarshalBinary_Valid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		data, err := c.GobEncode()
		require.Nil(t, err)
		assert.Len(t, data, 2)
	})
	t.Run("UnmarshalBinary_Valid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.GobDecode([]byte("Ca"))
		require.Nil(t, err)
		require.EqualValues(t, CountriesIso31661CA.String(), c.value)
	})
	t.Run("UnmarshalText_Invalid", func(t *testing.T) {
		c := &stringEnumValue{key: countriesIso31661Key}
		err := c.GobDecode(nil)
		require.NoError(t, err)
		err = c.GobDecode([]byte("invalid_e55bd60f-6352-4861-a0be-048acf7c3d03"))
		require.NotNil(t, err)
	})
}

func TestEnumValue_MarshalBSON(t *testing.T) {
	t.Run("MarshalBSON_Valid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		data, err := c.MarshalBSON()
		require.Nil(t, err)
		assert.Len(t, data, 18)
	})
}

func TestEnumValue_UnmarshalBSON(t *testing.T) {
	t.Run("UnmarshalBSON_Valid", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}

		v1 := &fooEnumValue{EnumValue: c}
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooEnumValue{EnumValue: c}
		err = bson.Unmarshal(rawData, &v2)
		require.Nil(t, err)
		assert.EqualValues(t, "Ca", v2.EnumValue.String())
	})
	t.Run("UnmarshalBSON_InvalidTable", func(t *testing.T) {
		c := stringEnumValue{CountriesIso31661CA.String(), countriesIso31661Key}
		ptr := &c
		ptr.value = "PPP"
		v1 := &fooEnumValue{EnumValue: c}
		rawData, err := bson.Marshal(v1)
		require.Nil(t, err)

		v2 := &fooEnumValue{}
		err = bson.Unmarshal(rawData, &v2)
		require.NotNil(t, err)
	})
}
