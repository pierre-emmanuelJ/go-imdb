# AwardOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutcomeTitle** | Pointer to **NullableString** |  | [optional] 
**OutcomeCategory** | Pointer to **NullableString** |  | [optional] 
**OutcomeDetails** | Pointer to [**[]AwardOutcomeDetail**](AwardOutcomeDetail.md) |  | [optional] 

## Methods

### NewAwardOutcome

`func NewAwardOutcome() *AwardOutcome`

NewAwardOutcome instantiates a new AwardOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardOutcomeWithDefaults

`func NewAwardOutcomeWithDefaults() *AwardOutcome`

NewAwardOutcomeWithDefaults instantiates a new AwardOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcomeTitle

`func (o *AwardOutcome) GetOutcomeTitle() string`

GetOutcomeTitle returns the OutcomeTitle field if non-nil, zero value otherwise.

### GetOutcomeTitleOk

`func (o *AwardOutcome) GetOutcomeTitleOk() (*string, bool)`

GetOutcomeTitleOk returns a tuple with the OutcomeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeTitle

`func (o *AwardOutcome) SetOutcomeTitle(v string)`

SetOutcomeTitle sets OutcomeTitle field to given value.

### HasOutcomeTitle

`func (o *AwardOutcome) HasOutcomeTitle() bool`

HasOutcomeTitle returns a boolean if a field has been set.

### SetOutcomeTitleNil

`func (o *AwardOutcome) SetOutcomeTitleNil(b bool)`

 SetOutcomeTitleNil sets the value for OutcomeTitle to be an explicit nil

### UnsetOutcomeTitle
`func (o *AwardOutcome) UnsetOutcomeTitle()`

UnsetOutcomeTitle ensures that no value is present for OutcomeTitle, not even an explicit nil
### GetOutcomeCategory

`func (o *AwardOutcome) GetOutcomeCategory() string`

GetOutcomeCategory returns the OutcomeCategory field if non-nil, zero value otherwise.

### GetOutcomeCategoryOk

`func (o *AwardOutcome) GetOutcomeCategoryOk() (*string, bool)`

GetOutcomeCategoryOk returns a tuple with the OutcomeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeCategory

`func (o *AwardOutcome) SetOutcomeCategory(v string)`

SetOutcomeCategory sets OutcomeCategory field to given value.

### HasOutcomeCategory

`func (o *AwardOutcome) HasOutcomeCategory() bool`

HasOutcomeCategory returns a boolean if a field has been set.

### SetOutcomeCategoryNil

`func (o *AwardOutcome) SetOutcomeCategoryNil(b bool)`

 SetOutcomeCategoryNil sets the value for OutcomeCategory to be an explicit nil

### UnsetOutcomeCategory
`func (o *AwardOutcome) UnsetOutcomeCategory()`

UnsetOutcomeCategory ensures that no value is present for OutcomeCategory, not even an explicit nil
### GetOutcomeDetails

`func (o *AwardOutcome) GetOutcomeDetails() []AwardOutcomeDetail`

GetOutcomeDetails returns the OutcomeDetails field if non-nil, zero value otherwise.

### GetOutcomeDetailsOk

`func (o *AwardOutcome) GetOutcomeDetailsOk() (*[]AwardOutcomeDetail, bool)`

GetOutcomeDetailsOk returns a tuple with the OutcomeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeDetails

`func (o *AwardOutcome) SetOutcomeDetails(v []AwardOutcomeDetail)`

SetOutcomeDetails sets OutcomeDetails field to given value.

### HasOutcomeDetails

`func (o *AwardOutcome) HasOutcomeDetails() bool`

HasOutcomeDetails returns a boolean if a field has been set.

### SetOutcomeDetailsNil

`func (o *AwardOutcome) SetOutcomeDetailsNil(b bool)`

 SetOutcomeDetailsNil sets the value for OutcomeDetails to be an explicit nil

### UnsetOutcomeDetails
`func (o *AwardOutcome) UnsetOutcomeDetails()`

UnsetOutcomeDetails ensures that no value is present for OutcomeDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


