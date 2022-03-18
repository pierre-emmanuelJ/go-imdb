# NameAwardEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTitle** | Pointer to **NullableString** |  | [optional] 
**OutcomeItems** | Pointer to [**[]NameAwardOutcome**](NameAwardOutcome.md) |  | [optional] 

## Methods

### NewNameAwardEvent

`func NewNameAwardEvent() *NameAwardEvent`

NewNameAwardEvent instantiates a new NameAwardEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameAwardEventWithDefaults

`func NewNameAwardEventWithDefaults() *NameAwardEvent`

NewNameAwardEventWithDefaults instantiates a new NameAwardEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTitle

`func (o *NameAwardEvent) GetEventTitle() string`

GetEventTitle returns the EventTitle field if non-nil, zero value otherwise.

### GetEventTitleOk

`func (o *NameAwardEvent) GetEventTitleOk() (*string, bool)`

GetEventTitleOk returns a tuple with the EventTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTitle

`func (o *NameAwardEvent) SetEventTitle(v string)`

SetEventTitle sets EventTitle field to given value.

### HasEventTitle

`func (o *NameAwardEvent) HasEventTitle() bool`

HasEventTitle returns a boolean if a field has been set.

### SetEventTitleNil

`func (o *NameAwardEvent) SetEventTitleNil(b bool)`

 SetEventTitleNil sets the value for EventTitle to be an explicit nil

### UnsetEventTitle
`func (o *NameAwardEvent) UnsetEventTitle()`

UnsetEventTitle ensures that no value is present for EventTitle, not even an explicit nil
### GetOutcomeItems

`func (o *NameAwardEvent) GetOutcomeItems() []NameAwardOutcome`

GetOutcomeItems returns the OutcomeItems field if non-nil, zero value otherwise.

### GetOutcomeItemsOk

`func (o *NameAwardEvent) GetOutcomeItemsOk() (*[]NameAwardOutcome, bool)`

GetOutcomeItemsOk returns a tuple with the OutcomeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeItems

`func (o *NameAwardEvent) SetOutcomeItems(v []NameAwardOutcome)`

SetOutcomeItems sets OutcomeItems field to given value.

### HasOutcomeItems

`func (o *NameAwardEvent) HasOutcomeItems() bool`

HasOutcomeItems returns a boolean if a field has been set.

### SetOutcomeItemsNil

`func (o *NameAwardEvent) SetOutcomeItemsNil(b bool)`

 SetOutcomeItemsNil sets the value for OutcomeItems to be an explicit nil

### UnsetOutcomeItems
`func (o *NameAwardEvent) UnsetOutcomeItems()`

UnsetOutcomeItems ensures that no value is present for OutcomeItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


