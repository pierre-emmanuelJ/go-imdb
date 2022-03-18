# AwardEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTitle** | Pointer to **NullableString** |  | [optional] 
**EventYear** | Pointer to **NullableString** |  | [optional] 
**OutcomeItems** | Pointer to [**[]AwardOutcome**](AwardOutcome.md) |  | [optional] 

## Methods

### NewAwardEvent

`func NewAwardEvent() *AwardEvent`

NewAwardEvent instantiates a new AwardEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardEventWithDefaults

`func NewAwardEventWithDefaults() *AwardEvent`

NewAwardEventWithDefaults instantiates a new AwardEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTitle

`func (o *AwardEvent) GetEventTitle() string`

GetEventTitle returns the EventTitle field if non-nil, zero value otherwise.

### GetEventTitleOk

`func (o *AwardEvent) GetEventTitleOk() (*string, bool)`

GetEventTitleOk returns a tuple with the EventTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTitle

`func (o *AwardEvent) SetEventTitle(v string)`

SetEventTitle sets EventTitle field to given value.

### HasEventTitle

`func (o *AwardEvent) HasEventTitle() bool`

HasEventTitle returns a boolean if a field has been set.

### SetEventTitleNil

`func (o *AwardEvent) SetEventTitleNil(b bool)`

 SetEventTitleNil sets the value for EventTitle to be an explicit nil

### UnsetEventTitle
`func (o *AwardEvent) UnsetEventTitle()`

UnsetEventTitle ensures that no value is present for EventTitle, not even an explicit nil
### GetEventYear

`func (o *AwardEvent) GetEventYear() string`

GetEventYear returns the EventYear field if non-nil, zero value otherwise.

### GetEventYearOk

`func (o *AwardEvent) GetEventYearOk() (*string, bool)`

GetEventYearOk returns a tuple with the EventYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventYear

`func (o *AwardEvent) SetEventYear(v string)`

SetEventYear sets EventYear field to given value.

### HasEventYear

`func (o *AwardEvent) HasEventYear() bool`

HasEventYear returns a boolean if a field has been set.

### SetEventYearNil

`func (o *AwardEvent) SetEventYearNil(b bool)`

 SetEventYearNil sets the value for EventYear to be an explicit nil

### UnsetEventYear
`func (o *AwardEvent) UnsetEventYear()`

UnsetEventYear ensures that no value is present for EventYear, not even an explicit nil
### GetOutcomeItems

`func (o *AwardEvent) GetOutcomeItems() []AwardOutcome`

GetOutcomeItems returns the OutcomeItems field if non-nil, zero value otherwise.

### GetOutcomeItemsOk

`func (o *AwardEvent) GetOutcomeItemsOk() (*[]AwardOutcome, bool)`

GetOutcomeItemsOk returns a tuple with the OutcomeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeItems

`func (o *AwardEvent) SetOutcomeItems(v []AwardOutcome)`

SetOutcomeItems sets OutcomeItems field to given value.

### HasOutcomeItems

`func (o *AwardEvent) HasOutcomeItems() bool`

HasOutcomeItems returns a boolean if a field has been set.

### SetOutcomeItemsNil

`func (o *AwardEvent) SetOutcomeItemsNil(b bool)`

 SetOutcomeItemsNil sets the value for OutcomeItems to be an explicit nil

### UnsetOutcomeItems
`func (o *AwardEvent) UnsetOutcomeItems()`

UnsetOutcomeItems ensures that no value is present for OutcomeItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


