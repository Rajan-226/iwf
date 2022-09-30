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

// WorkflowStartRequest struct for WorkflowStartRequest
type WorkflowStartRequest struct {
	WorkflowId string `json:"workflowId"`
	IwfWorkflowType string `json:"iwfWorkflowType"`
	WorkflowTimeoutSeconds int32 `json:"workflowTimeoutSeconds"`
	IwfWorkerUrl string `json:"iwfWorkerUrl"`
	StartStateId string `json:"startStateId"`
	StateInput *EncodedObject `json:"stateInput,omitempty"`
	StateOptions *WorkflowStateOptions `json:"stateOptions,omitempty"`
}

// NewWorkflowStartRequest instantiates a new WorkflowStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStartRequest(workflowId string, iwfWorkflowType string, workflowTimeoutSeconds int32, iwfWorkerUrl string, startStateId string) *WorkflowStartRequest {
	this := WorkflowStartRequest{}
	this.WorkflowId = workflowId
	this.IwfWorkflowType = iwfWorkflowType
	this.WorkflowTimeoutSeconds = workflowTimeoutSeconds
	this.IwfWorkerUrl = iwfWorkerUrl
	this.StartStateId = startStateId
	return &this
}

// NewWorkflowStartRequestWithDefaults instantiates a new WorkflowStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStartRequestWithDefaults() *WorkflowStartRequest {
	this := WorkflowStartRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowStartRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowStartRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetIwfWorkflowType returns the IwfWorkflowType field value
func (o *WorkflowStartRequest) GetIwfWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IwfWorkflowType
}

// GetIwfWorkflowTypeOk returns a tuple with the IwfWorkflowType field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetIwfWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IwfWorkflowType, true
}

// SetIwfWorkflowType sets field value
func (o *WorkflowStartRequest) SetIwfWorkflowType(v string) {
	o.IwfWorkflowType = v
}

// GetWorkflowTimeoutSeconds returns the WorkflowTimeoutSeconds field value
func (o *WorkflowStartRequest) GetWorkflowTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkflowTimeoutSeconds
}

// GetWorkflowTimeoutSecondsOk returns a tuple with the WorkflowTimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWorkflowTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowTimeoutSeconds, true
}

// SetWorkflowTimeoutSeconds sets field value
func (o *WorkflowStartRequest) SetWorkflowTimeoutSeconds(v int32) {
	o.WorkflowTimeoutSeconds = v
}

// GetIwfWorkerUrl returns the IwfWorkerUrl field value
func (o *WorkflowStartRequest) GetIwfWorkerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IwfWorkerUrl
}

// GetIwfWorkerUrlOk returns a tuple with the IwfWorkerUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetIwfWorkerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IwfWorkerUrl, true
}

// SetIwfWorkerUrl sets field value
func (o *WorkflowStartRequest) SetIwfWorkerUrl(v string) {
	o.IwfWorkerUrl = v
}

// GetStartStateId returns the StartStateId field value
func (o *WorkflowStartRequest) GetStartStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartStateId
}

// GetStartStateIdOk returns a tuple with the StartStateId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStartStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartStateId, true
}

// SetStartStateId sets field value
func (o *WorkflowStartRequest) SetStartStateId(v string) {
	o.StartStateId = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetStateInput() EncodedObject {
	if o == nil || o.StateInput == nil {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || o.StateInput == nil {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasStateInput() bool {
	if o != nil && o.StateInput != nil {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *WorkflowStartRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetStateOptions returns the StateOptions field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetStateOptions() WorkflowStateOptions {
	if o == nil || o.StateOptions == nil {
		var ret WorkflowStateOptions
		return ret
	}
	return *o.StateOptions
}

// GetStateOptionsOk returns a tuple with the StateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStateOptionsOk() (*WorkflowStateOptions, bool) {
	if o == nil || o.StateOptions == nil {
		return nil, false
	}
	return o.StateOptions, true
}

// HasStateOptions returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasStateOptions() bool {
	if o != nil && o.StateOptions != nil {
		return true
	}

	return false
}

// SetStateOptions gets a reference to the given WorkflowStateOptions and assigns it to the StateOptions field.
func (o *WorkflowStartRequest) SetStateOptions(v WorkflowStateOptions) {
	o.StateOptions = &v
}

func (o WorkflowStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if true {
		toSerialize["iwfWorkflowType"] = o.IwfWorkflowType
	}
	if true {
		toSerialize["workflowTimeoutSeconds"] = o.WorkflowTimeoutSeconds
	}
	if true {
		toSerialize["iwfWorkerUrl"] = o.IwfWorkerUrl
	}
	if true {
		toSerialize["startStateId"] = o.StartStateId
	}
	if o.StateInput != nil {
		toSerialize["stateInput"] = o.StateInput
	}
	if o.StateOptions != nil {
		toSerialize["stateOptions"] = o.StateOptions
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStartRequest struct {
	value *WorkflowStartRequest
	isSet bool
}

func (v NullableWorkflowStartRequest) Get() *WorkflowStartRequest {
	return v.value
}

func (v *NullableWorkflowStartRequest) Set(val *WorkflowStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStartRequest(val *WorkflowStartRequest) *NullableWorkflowStartRequest {
	return &NullableWorkflowStartRequest{value: val, isSet: true}
}

func (v NullableWorkflowStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


