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

// Updateip200Response struct for Updateip200Response
type Updateip200Response struct {
	Updateip200ResponseAnyOf *Updateip200ResponseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Updateip200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Updateip200ResponseAnyOf
	err = json.Unmarshal(data, &dst.Updateip200ResponseAnyOf);
	if err == nil {
		jsonUpdateip200ResponseAnyOf, _ := json.Marshal(dst.Updateip200ResponseAnyOf)
		if string(jsonUpdateip200ResponseAnyOf) == "{}" { // empty struct
			dst.Updateip200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.Updateip200ResponseAnyOf, return on the first match
		}
	} else {
		dst.Updateip200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Updateip200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Updateip200Response) MarshalJSON() ([]byte, error) {
	if src.Updateip200ResponseAnyOf != nil {
		return json.Marshal(&src.Updateip200ResponseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUpdateip200Response struct {
	value *Updateip200Response
	isSet bool
}

func (v NullableUpdateip200Response) Get() *Updateip200Response {
	return v.value
}

func (v *NullableUpdateip200Response) Set(val *Updateip200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateip200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateip200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateip200Response(val *Updateip200Response) *NullableUpdateip200Response {
	return &NullableUpdateip200Response{value: val, isSet: true}
}

func (v NullableUpdateip200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateip200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


