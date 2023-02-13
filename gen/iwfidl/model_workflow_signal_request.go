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

// WorkflowSignalRequest struct for WorkflowSignalRequest
type WorkflowSignalRequest struct {
	WorkflowId        string         `json:"workflowId"`
	WorkflowRunId     *string        `json:"workflowRunId,omitempty"`
	SignalChannelName string         `json:"signalChannelName"`
	SignalValue       *EncodedObject `json:"signalValue,omitempty"`
}

// NewWorkflowSignalRequest instantiates a new WorkflowSignalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSignalRequest(workflowId string, signalChannelName string) *WorkflowSignalRequest {
	this := WorkflowSignalRequest{}
	this.WorkflowId = workflowId
	this.SignalChannelName = signalChannelName
	return &this
}

// NewWorkflowSignalRequestWithDefaults instantiates a new WorkflowSignalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSignalRequestWithDefaults() *WorkflowSignalRequest {
	this := WorkflowSignalRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowSignalRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSignalRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowSignalRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowSignalRequest) GetWorkflowRunId() string {
	if o == nil || isNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSignalRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowSignalRequest) HasWorkflowRunId() bool {
	if o != nil && !isNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowSignalRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetSignalChannelName returns the SignalChannelName field value
func (o *WorkflowSignalRequest) GetSignalChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalChannelName
}

// GetSignalChannelNameOk returns a tuple with the SignalChannelName field value
// and a boolean to check if the value has been set.
func (o *WorkflowSignalRequest) GetSignalChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalChannelName, true
}

// SetSignalChannelName sets field value
func (o *WorkflowSignalRequest) SetSignalChannelName(v string) {
	o.SignalChannelName = v
}

// GetSignalValue returns the SignalValue field value if set, zero value otherwise.
func (o *WorkflowSignalRequest) GetSignalValue() EncodedObject {
	if o == nil || isNil(o.SignalValue) {
		var ret EncodedObject
		return ret
	}
	return *o.SignalValue
}

// GetSignalValueOk returns a tuple with the SignalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSignalRequest) GetSignalValueOk() (*EncodedObject, bool) {
	if o == nil || isNil(o.SignalValue) {
		return nil, false
	}
	return o.SignalValue, true
}

// HasSignalValue returns a boolean if a field has been set.
func (o *WorkflowSignalRequest) HasSignalValue() bool {
	if o != nil && !isNil(o.SignalValue) {
		return true
	}

	return false
}

// SetSignalValue gets a reference to the given EncodedObject and assigns it to the SignalValue field.
func (o *WorkflowSignalRequest) SetSignalValue(v EncodedObject) {
	o.SignalValue = &v
}

func (o WorkflowSignalRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if !isNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if true {
		toSerialize["signalChannelName"] = o.SignalChannelName
	}
	if !isNil(o.SignalValue) {
		toSerialize["signalValue"] = o.SignalValue
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSignalRequest struct {
	value *WorkflowSignalRequest
	isSet bool
}

func (v NullableWorkflowSignalRequest) Get() *WorkflowSignalRequest {
	return v.value
}

func (v *NullableWorkflowSignalRequest) Set(val *WorkflowSignalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSignalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSignalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSignalRequest(val *WorkflowSignalRequest) *NullableWorkflowSignalRequest {
	return &NullableWorkflowSignalRequest{value: val, isSet: true}
}

func (v NullableWorkflowSignalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSignalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
