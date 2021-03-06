/*
IMDb-API

The IMDb-API Documentation. You need a <a href='/Identity/Account/Manage' target='_blank'><code>API Key</code></a> for testing APIs.<br/><a class='link' href='/API'>Back to API Tester</a>

API version: 1.8.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NameAwardOutcome struct for NameAwardOutcome
type NameAwardOutcome struct {
	OutcomeYear NullableString `json:"outcomeYear,omitempty"`
	OutcomeTitle NullableString `json:"outcomeTitle,omitempty"`
	OutcomeCategory NullableString `json:"outcomeCategory,omitempty"`
	OutcomeDetails []NameAwardOutcomeDetail `json:"outcomeDetails,omitempty"`
}

// NewNameAwardOutcome instantiates a new NameAwardOutcome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameAwardOutcome() *NameAwardOutcome {
	this := NameAwardOutcome{}
	return &this
}

// NewNameAwardOutcomeWithDefaults instantiates a new NameAwardOutcome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameAwardOutcomeWithDefaults() *NameAwardOutcome {
	this := NameAwardOutcome{}
	return &this
}

// GetOutcomeYear returns the OutcomeYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameAwardOutcome) GetOutcomeYear() string {
	if o == nil || o.OutcomeYear.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutcomeYear.Get()
}

// GetOutcomeYearOk returns a tuple with the OutcomeYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameAwardOutcome) GetOutcomeYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutcomeYear.Get(), o.OutcomeYear.IsSet()
}

// HasOutcomeYear returns a boolean if a field has been set.
func (o *NameAwardOutcome) HasOutcomeYear() bool {
	if o != nil && o.OutcomeYear.IsSet() {
		return true
	}

	return false
}

// SetOutcomeYear gets a reference to the given NullableString and assigns it to the OutcomeYear field.
func (o *NameAwardOutcome) SetOutcomeYear(v string) {
	o.OutcomeYear.Set(&v)
}
// SetOutcomeYearNil sets the value for OutcomeYear to be an explicit nil
func (o *NameAwardOutcome) SetOutcomeYearNil() {
	o.OutcomeYear.Set(nil)
}

// UnsetOutcomeYear ensures that no value is present for OutcomeYear, not even an explicit nil
func (o *NameAwardOutcome) UnsetOutcomeYear() {
	o.OutcomeYear.Unset()
}

// GetOutcomeTitle returns the OutcomeTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameAwardOutcome) GetOutcomeTitle() string {
	if o == nil || o.OutcomeTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutcomeTitle.Get()
}

// GetOutcomeTitleOk returns a tuple with the OutcomeTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameAwardOutcome) GetOutcomeTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutcomeTitle.Get(), o.OutcomeTitle.IsSet()
}

// HasOutcomeTitle returns a boolean if a field has been set.
func (o *NameAwardOutcome) HasOutcomeTitle() bool {
	if o != nil && o.OutcomeTitle.IsSet() {
		return true
	}

	return false
}

// SetOutcomeTitle gets a reference to the given NullableString and assigns it to the OutcomeTitle field.
func (o *NameAwardOutcome) SetOutcomeTitle(v string) {
	o.OutcomeTitle.Set(&v)
}
// SetOutcomeTitleNil sets the value for OutcomeTitle to be an explicit nil
func (o *NameAwardOutcome) SetOutcomeTitleNil() {
	o.OutcomeTitle.Set(nil)
}

// UnsetOutcomeTitle ensures that no value is present for OutcomeTitle, not even an explicit nil
func (o *NameAwardOutcome) UnsetOutcomeTitle() {
	o.OutcomeTitle.Unset()
}

// GetOutcomeCategory returns the OutcomeCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameAwardOutcome) GetOutcomeCategory() string {
	if o == nil || o.OutcomeCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutcomeCategory.Get()
}

// GetOutcomeCategoryOk returns a tuple with the OutcomeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameAwardOutcome) GetOutcomeCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutcomeCategory.Get(), o.OutcomeCategory.IsSet()
}

// HasOutcomeCategory returns a boolean if a field has been set.
func (o *NameAwardOutcome) HasOutcomeCategory() bool {
	if o != nil && o.OutcomeCategory.IsSet() {
		return true
	}

	return false
}

// SetOutcomeCategory gets a reference to the given NullableString and assigns it to the OutcomeCategory field.
func (o *NameAwardOutcome) SetOutcomeCategory(v string) {
	o.OutcomeCategory.Set(&v)
}
// SetOutcomeCategoryNil sets the value for OutcomeCategory to be an explicit nil
func (o *NameAwardOutcome) SetOutcomeCategoryNil() {
	o.OutcomeCategory.Set(nil)
}

// UnsetOutcomeCategory ensures that no value is present for OutcomeCategory, not even an explicit nil
func (o *NameAwardOutcome) UnsetOutcomeCategory() {
	o.OutcomeCategory.Unset()
}

// GetOutcomeDetails returns the OutcomeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameAwardOutcome) GetOutcomeDetails() []NameAwardOutcomeDetail {
	if o == nil {
		var ret []NameAwardOutcomeDetail
		return ret
	}
	return o.OutcomeDetails
}

// GetOutcomeDetailsOk returns a tuple with the OutcomeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameAwardOutcome) GetOutcomeDetailsOk() ([]NameAwardOutcomeDetail, bool) {
	if o == nil || o.OutcomeDetails == nil {
		return nil, false
	}
	return o.OutcomeDetails, true
}

// HasOutcomeDetails returns a boolean if a field has been set.
func (o *NameAwardOutcome) HasOutcomeDetails() bool {
	if o != nil && o.OutcomeDetails != nil {
		return true
	}

	return false
}

// SetOutcomeDetails gets a reference to the given []NameAwardOutcomeDetail and assigns it to the OutcomeDetails field.
func (o *NameAwardOutcome) SetOutcomeDetails(v []NameAwardOutcomeDetail) {
	o.OutcomeDetails = v
}

func (o NameAwardOutcome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OutcomeYear.IsSet() {
		toSerialize["outcomeYear"] = o.OutcomeYear.Get()
	}
	if o.OutcomeTitle.IsSet() {
		toSerialize["outcomeTitle"] = o.OutcomeTitle.Get()
	}
	if o.OutcomeCategory.IsSet() {
		toSerialize["outcomeCategory"] = o.OutcomeCategory.Get()
	}
	if o.OutcomeDetails != nil {
		toSerialize["outcomeDetails"] = o.OutcomeDetails
	}
	return json.Marshal(toSerialize)
}

type NullableNameAwardOutcome struct {
	value *NameAwardOutcome
	isSet bool
}

func (v NullableNameAwardOutcome) Get() *NameAwardOutcome {
	return v.value
}

func (v *NullableNameAwardOutcome) Set(val *NameAwardOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableNameAwardOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableNameAwardOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameAwardOutcome(val *NameAwardOutcome) *NullableNameAwardOutcome {
	return &NullableNameAwardOutcome{value: val, isSet: true}
}

func (v NullableNameAwardOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameAwardOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


