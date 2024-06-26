/*
Dynamicdns.pro

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynamicdnsapi

import (
	"encoding/json"
	"fmt"
)

// Deleteip200Response struct for Deleteip200Response
type Deleteip200Response struct {
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Deleteip200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.map[string]interface{});
	if err == nil {
		jsonmap[string]interface{}, _ := json.Marshal(dst.map[string]interface{})
		if string(jsonmap[string]interface{}) == "{}" { // empty struct
			dst.map[string]interface{} = nil
		} else {
			return nil // data stored in dst.map[string]interface{}, return on the first match
		}
	} else {
		dst.map[string]interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Deleteip200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Deleteip200Response) MarshalJSON() ([]byte, error) {
	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeleteip200Response struct {
	value *Deleteip200Response
	isSet bool
}

func (v NullableDeleteip200Response) Get() *Deleteip200Response {
	return v.value
}

func (v *NullableDeleteip200Response) Set(val *Deleteip200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteip200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteip200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteip200Response(val *Deleteip200Response) *NullableDeleteip200Response {
	return &NullableDeleteip200Response{value: val, isSet: true}
}

func (v NullableDeleteip200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteip200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


