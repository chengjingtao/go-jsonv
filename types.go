package jsonv

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JsonV struct {
	Type         Type             `json:"type"`
	IntVal       int64            `json:"intVal"`
	StringVal    string           `json:"strVal"`
	BoolVal      bool             `json:"boolVal"`
	StringMapVal map[string]JsonV `json:"strMapVal"`
	ArrayVal     []JsonV          `json:"arrayVal"`
}

func (v *JsonV) UnmarshalJSON(value []byte) (err error) {
	v.Type, err = v.typ(value)
	if err != nil {
		return err
	}

	switch v.Type {
	case Null:
		return nil
	case Int:
		return json.Unmarshal(value, &v.IntVal)
	case String:
		return json.Unmarshal(value, &v.StringVal)
	case Bool:
		return json.Unmarshal(value, &v.BoolVal)
	case StringMap:
		return json.Unmarshal(value, &v.StringMapVal)
	case Arrary:
		return json.Unmarshal(value, &v.ArrayVal)
	default:
		return fmt.Errorf("UnKnown type when unmarshal json: %s", string(value))
	}
}

func (v *JsonV) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var nullV = &struct{}{}
	err := unmarshal(nullV)
	if err == nil && nullV == nil {
		v.Type = Null
		return nil
	}

	var intV int64
	err = unmarshal(&intV)
	if err == nil {
		v.Type = Int
		v.IntVal = intV
		return nil
	}

	var boolV bool
	err = unmarshal(&boolV)
	if err == nil {
		v.Type = Bool
		v.BoolVal = boolV
		return nil
	}

	var stringV string
	err = unmarshal(&stringV)
	if err == nil {
		v.Type = String
		v.StringVal = stringV
		return nil
	}

	var stringMapV = map[string]JsonV{}
	err = unmarshal(&stringMapV)
	if err == nil {
		v.Type = StringMap
		v.StringMapVal = stringMapV
		return nil
	}

	var arrayV = []JsonV{}
	err = unmarshal(&arrayV)
	if err == nil {
		v.Type = Arrary
		v.ArrayVal = arrayV
		return nil
	}

	return fmt.Errorf("UnKnown type when unmarshal yaml")
}

func (v JsonV) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case Null:
		return json.Marshal(nil)
	case Int:
		return json.Marshal(v.IntVal)
	case String:
		return json.Marshal(v.StringVal)
	case Bool:
		return json.Marshal(v.BoolVal)
	case StringMap:
		return json.Marshal(v.StringMapVal)
	case Arrary:
		return json.Marshal(v.ArrayVal)
	default:
		return []byte{}, fmt.Errorf("impossible V.Type: %#v", v.Type)
	}
}

func (v JsonV) MarshalYAML() (interface{}, error) {
	switch v.Type {
	case Null:
		return nil, nil
	case Int:
		return v.IntVal, nil
	case String:
		return v.StringVal, nil
	case Bool:
		return v.BoolVal, nil
	case StringMap:
		return v.StringMapVal, nil
	case Arrary:
		return v.ArrayVal, nil
	default:
		return nil, fmt.Errorf("impossible V.Type: %#v", v.Type)
	}
}

func (v *JsonV) typ(value []byte) (Type, error) {
	start := value[0]
	if start == '"' {
		return String, nil
	}
	if start == '{' {
		return StringMap, nil
	}
	if start == '[' {
		return Arrary, nil
	}

	str := string(value)
	if str == "false" || str == "true" {
		return Bool, nil
	}

	if str == "null" {
		return Null, nil
	}

	_, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return Int, nil
	}

	return Null, fmt.Errorf("UnKnown type of value: %s", string(value))
}

// Type represents the stored type of IntOrString.
type Type int

const (
	// Null indicates the type of Object is Null
	Null Type = iota
	Int
	String
	Bool
	StringMap
	Arrary
)
