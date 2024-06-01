/*
Dynamicdns.pro

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynamicdnsapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Updateip200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Updateip200ResponseAnyOf{}

// Updateip200ResponseAnyOf struct for Updateip200ResponseAnyOf
type Updateip200ResponseAnyOf struct {
	Message string `json:"message"`
	Ip string `json:"ip"`
}

type _Updateip200ResponseAnyOf Updateip200ResponseAnyOf

// NewUpdateip200ResponseAnyOf instantiates a new Updateip200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateip200ResponseAnyOf(message string, ip string) *Updateip200ResponseAnyOf {
	this := Updateip200ResponseAnyOf{}
	this.Message = message
	this.Ip = ip
	return &this
}

// NewUpdateip200ResponseAnyOfWithDefaults instantiates a new Updateip200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateip200ResponseAnyOfWithDefaults() *Updateip200ResponseAnyOf {
	this := Updateip200ResponseAnyOf{}
	return &this
}

// GetMessage returns the Message field value
func (o *Updateip200ResponseAnyOf) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Updateip200ResponseAnyOf) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Updateip200ResponseAnyOf) SetMessage(v string) {
	o.Message = v
}

// GetIp returns the Ip field value
func (o *Updateip200ResponseAnyOf) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Updateip200ResponseAnyOf) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *Updateip200ResponseAnyOf) SetIp(v string) {
	o.Ip = v
}

func (o Updateip200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Updateip200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

func (o *Updateip200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"ip",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateip200ResponseAnyOf := _Updateip200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateip200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = Updateip200ResponseAnyOf(varUpdateip200ResponseAnyOf)

	return err
}

type NullableUpdateip200ResponseAnyOf struct {
	value *Updateip200ResponseAnyOf
	isSet bool
}

func (v NullableUpdateip200ResponseAnyOf) Get() *Updateip200ResponseAnyOf {
	return v.value
}

func (v *NullableUpdateip200ResponseAnyOf) Set(val *Updateip200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateip200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateip200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateip200ResponseAnyOf(val *Updateip200ResponseAnyOf) *NullableUpdateip200ResponseAnyOf {
	return &NullableUpdateip200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableUpdateip200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateip200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

