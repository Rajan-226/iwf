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

// WorkflowGetDataObjectsRequest struct for WorkflowGetDataObjectsRequest
type WorkflowGetDataObjectsRequest struct {
	WorkflowId    string   `json:"workflowId"`
	WorkflowRunId *string  `json:"workflowRunId,omitempty"`
	Keys          []string `json:"keys,omitempty"`
}

// NewWorkflowGetDataObjectsRequest instantiates a new WorkflowGetDataObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowGetDataObjectsRequest(workflowId string) *WorkflowGetDataObjectsRequest {
	this := WorkflowGetDataObjectsRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowGetDataObjectsRequestWithDefaults instantiates a new WorkflowGetDataObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowGetDataObjectsRequestWithDefaults() *WorkflowGetDataObjectsRequest {
	this := WorkflowGetDataObjectsRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowGetDataObjectsRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetDataObjectsRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowGetDataObjectsRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowGetDataObjectsRequest) GetWorkflowRunId() string {
	if o == nil || isNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetDataObjectsRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowGetDataObjectsRequest) HasWorkflowRunId() bool {
	if o != nil && !isNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowGetDataObjectsRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *WorkflowGetDataObjectsRequest) GetKeys() []string {
	if o == nil || isNil(o.Keys) {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetDataObjectsRequest) GetKeysOk() ([]string, bool) {
	if o == nil || isNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *WorkflowGetDataObjectsRequest) HasKeys() bool {
	if o != nil && !isNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *WorkflowGetDataObjectsRequest) SetKeys(v []string) {
	o.Keys = v
}

func (o WorkflowGetDataObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if !isNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if !isNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowGetDataObjectsRequest struct {
	value *WorkflowGetDataObjectsRequest
	isSet bool
}

func (v NullableWorkflowGetDataObjectsRequest) Get() *WorkflowGetDataObjectsRequest {
	return v.value
}

func (v *NullableWorkflowGetDataObjectsRequest) Set(val *WorkflowGetDataObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowGetDataObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowGetDataObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowGetDataObjectsRequest(val *WorkflowGetDataObjectsRequest) *NullableWorkflowGetDataObjectsRequest {
	return &NullableWorkflowGetDataObjectsRequest{value: val, isSet: true}
}

func (v NullableWorkflowGetDataObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowGetDataObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
