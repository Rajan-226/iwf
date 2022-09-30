/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// ActivityCommand struct for ActivityCommand
type ActivityCommand struct {
	CommandId string `json:"commandId"`
	ActivityType string `json:"activityType"`
	Input *EncodedObject `json:"input,omitempty"`
	ActivityOptions *ActivityOptions `json:"activityOptions,omitempty"`
}

// NewActivityCommand instantiates a new ActivityCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityCommand(commandId string, activityType string) *ActivityCommand {
	this := ActivityCommand{}
	this.CommandId = commandId
	this.ActivityType = activityType
	return &this
}

// NewActivityCommandWithDefaults instantiates a new ActivityCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityCommandWithDefaults() *ActivityCommand {
	this := ActivityCommand{}
	return &this
}

// GetCommandId returns the CommandId field value
func (o *ActivityCommand) GetCommandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value
// and a boolean to check if the value has been set.
func (o *ActivityCommand) GetCommandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandId, true
}

// SetCommandId sets field value
func (o *ActivityCommand) SetCommandId(v string) {
	o.CommandId = v
}

// GetActivityType returns the ActivityType field value
func (o *ActivityCommand) GetActivityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value
// and a boolean to check if the value has been set.
func (o *ActivityCommand) GetActivityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityType, true
}

// SetActivityType sets field value
func (o *ActivityCommand) SetActivityType(v string) {
	o.ActivityType = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ActivityCommand) GetInput() EncodedObject {
	if o == nil || o.Input == nil {
		var ret EncodedObject
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityCommand) GetInputOk() (*EncodedObject, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ActivityCommand) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given EncodedObject and assigns it to the Input field.
func (o *ActivityCommand) SetInput(v EncodedObject) {
	o.Input = &v
}

// GetActivityOptions returns the ActivityOptions field value if set, zero value otherwise.
func (o *ActivityCommand) GetActivityOptions() ActivityOptions {
	if o == nil || o.ActivityOptions == nil {
		var ret ActivityOptions
		return ret
	}
	return *o.ActivityOptions
}

// GetActivityOptionsOk returns a tuple with the ActivityOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityCommand) GetActivityOptionsOk() (*ActivityOptions, bool) {
	if o == nil || o.ActivityOptions == nil {
		return nil, false
	}
	return o.ActivityOptions, true
}

// HasActivityOptions returns a boolean if a field has been set.
func (o *ActivityCommand) HasActivityOptions() bool {
	if o != nil && o.ActivityOptions != nil {
		return true
	}

	return false
}

// SetActivityOptions gets a reference to the given ActivityOptions and assigns it to the ActivityOptions field.
func (o *ActivityCommand) SetActivityOptions(v ActivityOptions) {
	o.ActivityOptions = &v
}

func (o ActivityCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commandId"] = o.CommandId
	}
	if true {
		toSerialize["activityType"] = o.ActivityType
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.ActivityOptions != nil {
		toSerialize["activityOptions"] = o.ActivityOptions
	}
	return json.Marshal(toSerialize)
}

type NullableActivityCommand struct {
	value *ActivityCommand
	isSet bool
}

func (v NullableActivityCommand) Get() *ActivityCommand {
	return v.value
}

func (v *NullableActivityCommand) Set(val *ActivityCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityCommand(val *ActivityCommand) *NullableActivityCommand {
	return &NullableActivityCommand{value: val, isSet: true}
}

func (v NullableActivityCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


