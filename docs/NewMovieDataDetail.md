# NewMovieDataDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**ReleaseState** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**RuntimeMins** | Pointer to **NullableString** |  | [optional] 
**RuntimeStr** | Pointer to **NullableString** |  | [optional] 
**Plot** | Pointer to **NullableString** |  | [optional] 
**ContentRating** | Pointer to **NullableString** |  | [optional] 
**ImDbRating** | Pointer to **NullableString** |  | [optional] 
**ImDbRatingCount** | Pointer to **NullableString** |  | [optional] 
**MetacriticRating** | Pointer to **NullableString** |  | [optional] 
**Genres** | Pointer to **NullableString** |  | [optional] 
**GenreList** | Pointer to [**[]KeyValueItem**](KeyValueItem.md) |  | [optional] 
**Directors** | Pointer to **NullableString** |  | [optional] 
**DirectorList** | Pointer to [**[]StarShort**](StarShort.md) |  | [optional] 
**Stars** | Pointer to **NullableString** |  | [optional] 
**StarList** | Pointer to [**[]StarShort**](StarShort.md) |  | [optional] 

## Methods

### NewNewMovieDataDetail

`func NewNewMovieDataDetail() *NewMovieDataDetail`

NewNewMovieDataDetail instantiates a new NewMovieDataDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMovieDataDetailWithDefaults

`func NewNewMovieDataDetailWithDefaults() *NewMovieDataDetail`

NewNewMovieDataDetailWithDefaults instantiates a new NewMovieDataDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewMovieDataDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewMovieDataDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewMovieDataDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewMovieDataDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NewMovieDataDetail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NewMovieDataDetail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *NewMovieDataDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewMovieDataDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewMovieDataDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewMovieDataDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *NewMovieDataDetail) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NewMovieDataDetail) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *NewMovieDataDetail) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *NewMovieDataDetail) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *NewMovieDataDetail) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *NewMovieDataDetail) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *NewMovieDataDetail) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *NewMovieDataDetail) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetYear

`func (o *NewMovieDataDetail) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *NewMovieDataDetail) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *NewMovieDataDetail) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *NewMovieDataDetail) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *NewMovieDataDetail) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *NewMovieDataDetail) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetReleaseState

`func (o *NewMovieDataDetail) GetReleaseState() string`

GetReleaseState returns the ReleaseState field if non-nil, zero value otherwise.

### GetReleaseStateOk

`func (o *NewMovieDataDetail) GetReleaseStateOk() (*string, bool)`

GetReleaseStateOk returns a tuple with the ReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseState

`func (o *NewMovieDataDetail) SetReleaseState(v string)`

SetReleaseState sets ReleaseState field to given value.

### HasReleaseState

`func (o *NewMovieDataDetail) HasReleaseState() bool`

HasReleaseState returns a boolean if a field has been set.

### SetReleaseStateNil

`func (o *NewMovieDataDetail) SetReleaseStateNil(b bool)`

 SetReleaseStateNil sets the value for ReleaseState to be an explicit nil

### UnsetReleaseState
`func (o *NewMovieDataDetail) UnsetReleaseState()`

UnsetReleaseState ensures that no value is present for ReleaseState, not even an explicit nil
### GetImage

`func (o *NewMovieDataDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *NewMovieDataDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *NewMovieDataDetail) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *NewMovieDataDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *NewMovieDataDetail) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *NewMovieDataDetail) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetRuntimeMins

`func (o *NewMovieDataDetail) GetRuntimeMins() string`

GetRuntimeMins returns the RuntimeMins field if non-nil, zero value otherwise.

### GetRuntimeMinsOk

`func (o *NewMovieDataDetail) GetRuntimeMinsOk() (*string, bool)`

GetRuntimeMinsOk returns a tuple with the RuntimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMins

`func (o *NewMovieDataDetail) SetRuntimeMins(v string)`

SetRuntimeMins sets RuntimeMins field to given value.

### HasRuntimeMins

`func (o *NewMovieDataDetail) HasRuntimeMins() bool`

HasRuntimeMins returns a boolean if a field has been set.

### SetRuntimeMinsNil

`func (o *NewMovieDataDetail) SetRuntimeMinsNil(b bool)`

 SetRuntimeMinsNil sets the value for RuntimeMins to be an explicit nil

### UnsetRuntimeMins
`func (o *NewMovieDataDetail) UnsetRuntimeMins()`

UnsetRuntimeMins ensures that no value is present for RuntimeMins, not even an explicit nil
### GetRuntimeStr

`func (o *NewMovieDataDetail) GetRuntimeStr() string`

GetRuntimeStr returns the RuntimeStr field if non-nil, zero value otherwise.

### GetRuntimeStrOk

`func (o *NewMovieDataDetail) GetRuntimeStrOk() (*string, bool)`

GetRuntimeStrOk returns a tuple with the RuntimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStr

`func (o *NewMovieDataDetail) SetRuntimeStr(v string)`

SetRuntimeStr sets RuntimeStr field to given value.

### HasRuntimeStr

`func (o *NewMovieDataDetail) HasRuntimeStr() bool`

HasRuntimeStr returns a boolean if a field has been set.

### SetRuntimeStrNil

`func (o *NewMovieDataDetail) SetRuntimeStrNil(b bool)`

 SetRuntimeStrNil sets the value for RuntimeStr to be an explicit nil

### UnsetRuntimeStr
`func (o *NewMovieDataDetail) UnsetRuntimeStr()`

UnsetRuntimeStr ensures that no value is present for RuntimeStr, not even an explicit nil
### GetPlot

`func (o *NewMovieDataDetail) GetPlot() string`

GetPlot returns the Plot field if non-nil, zero value otherwise.

### GetPlotOk

`func (o *NewMovieDataDetail) GetPlotOk() (*string, bool)`

GetPlotOk returns a tuple with the Plot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlot

`func (o *NewMovieDataDetail) SetPlot(v string)`

SetPlot sets Plot field to given value.

### HasPlot

`func (o *NewMovieDataDetail) HasPlot() bool`

HasPlot returns a boolean if a field has been set.

### SetPlotNil

`func (o *NewMovieDataDetail) SetPlotNil(b bool)`

 SetPlotNil sets the value for Plot to be an explicit nil

### UnsetPlot
`func (o *NewMovieDataDetail) UnsetPlot()`

UnsetPlot ensures that no value is present for Plot, not even an explicit nil
### GetContentRating

`func (o *NewMovieDataDetail) GetContentRating() string`

GetContentRating returns the ContentRating field if non-nil, zero value otherwise.

### GetContentRatingOk

`func (o *NewMovieDataDetail) GetContentRatingOk() (*string, bool)`

GetContentRatingOk returns a tuple with the ContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRating

`func (o *NewMovieDataDetail) SetContentRating(v string)`

SetContentRating sets ContentRating field to given value.

### HasContentRating

`func (o *NewMovieDataDetail) HasContentRating() bool`

HasContentRating returns a boolean if a field has been set.

### SetContentRatingNil

`func (o *NewMovieDataDetail) SetContentRatingNil(b bool)`

 SetContentRatingNil sets the value for ContentRating to be an explicit nil

### UnsetContentRating
`func (o *NewMovieDataDetail) UnsetContentRating()`

UnsetContentRating ensures that no value is present for ContentRating, not even an explicit nil
### GetImDbRating

`func (o *NewMovieDataDetail) GetImDbRating() string`

GetImDbRating returns the ImDbRating field if non-nil, zero value otherwise.

### GetImDbRatingOk

`func (o *NewMovieDataDetail) GetImDbRatingOk() (*string, bool)`

GetImDbRatingOk returns a tuple with the ImDbRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbRating

`func (o *NewMovieDataDetail) SetImDbRating(v string)`

SetImDbRating sets ImDbRating field to given value.

### HasImDbRating

`func (o *NewMovieDataDetail) HasImDbRating() bool`

HasImDbRating returns a boolean if a field has been set.

### SetImDbRatingNil

`func (o *NewMovieDataDetail) SetImDbRatingNil(b bool)`

 SetImDbRatingNil sets the value for ImDbRating to be an explicit nil

### UnsetImDbRating
`func (o *NewMovieDataDetail) UnsetImDbRating()`

UnsetImDbRating ensures that no value is present for ImDbRating, not even an explicit nil
### GetImDbRatingCount

`func (o *NewMovieDataDetail) GetImDbRatingCount() string`

GetImDbRatingCount returns the ImDbRatingCount field if non-nil, zero value otherwise.

### GetImDbRatingCountOk

`func (o *NewMovieDataDetail) GetImDbRatingCountOk() (*string, bool)`

GetImDbRatingCountOk returns a tuple with the ImDbRatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbRatingCount

`func (o *NewMovieDataDetail) SetImDbRatingCount(v string)`

SetImDbRatingCount sets ImDbRatingCount field to given value.

### HasImDbRatingCount

`func (o *NewMovieDataDetail) HasImDbRatingCount() bool`

HasImDbRatingCount returns a boolean if a field has been set.

### SetImDbRatingCountNil

`func (o *NewMovieDataDetail) SetImDbRatingCountNil(b bool)`

 SetImDbRatingCountNil sets the value for ImDbRatingCount to be an explicit nil

### UnsetImDbRatingCount
`func (o *NewMovieDataDetail) UnsetImDbRatingCount()`

UnsetImDbRatingCount ensures that no value is present for ImDbRatingCount, not even an explicit nil
### GetMetacriticRating

`func (o *NewMovieDataDetail) GetMetacriticRating() string`

GetMetacriticRating returns the MetacriticRating field if non-nil, zero value otherwise.

### GetMetacriticRatingOk

`func (o *NewMovieDataDetail) GetMetacriticRatingOk() (*string, bool)`

GetMetacriticRatingOk returns a tuple with the MetacriticRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetacriticRating

`func (o *NewMovieDataDetail) SetMetacriticRating(v string)`

SetMetacriticRating sets MetacriticRating field to given value.

### HasMetacriticRating

`func (o *NewMovieDataDetail) HasMetacriticRating() bool`

HasMetacriticRating returns a boolean if a field has been set.

### SetMetacriticRatingNil

`func (o *NewMovieDataDetail) SetMetacriticRatingNil(b bool)`

 SetMetacriticRatingNil sets the value for MetacriticRating to be an explicit nil

### UnsetMetacriticRating
`func (o *NewMovieDataDetail) UnsetMetacriticRating()`

UnsetMetacriticRating ensures that no value is present for MetacriticRating, not even an explicit nil
### GetGenres

`func (o *NewMovieDataDetail) GetGenres() string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *NewMovieDataDetail) GetGenresOk() (*string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *NewMovieDataDetail) SetGenres(v string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *NewMovieDataDetail) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *NewMovieDataDetail) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *NewMovieDataDetail) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetGenreList

`func (o *NewMovieDataDetail) GetGenreList() []KeyValueItem`

GetGenreList returns the GenreList field if non-nil, zero value otherwise.

### GetGenreListOk

`func (o *NewMovieDataDetail) GetGenreListOk() (*[]KeyValueItem, bool)`

GetGenreListOk returns a tuple with the GenreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreList

`func (o *NewMovieDataDetail) SetGenreList(v []KeyValueItem)`

SetGenreList sets GenreList field to given value.

### HasGenreList

`func (o *NewMovieDataDetail) HasGenreList() bool`

HasGenreList returns a boolean if a field has been set.

### SetGenreListNil

`func (o *NewMovieDataDetail) SetGenreListNil(b bool)`

 SetGenreListNil sets the value for GenreList to be an explicit nil

### UnsetGenreList
`func (o *NewMovieDataDetail) UnsetGenreList()`

UnsetGenreList ensures that no value is present for GenreList, not even an explicit nil
### GetDirectors

`func (o *NewMovieDataDetail) GetDirectors() string`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *NewMovieDataDetail) GetDirectorsOk() (*string, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *NewMovieDataDetail) SetDirectors(v string)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *NewMovieDataDetail) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### SetDirectorsNil

`func (o *NewMovieDataDetail) SetDirectorsNil(b bool)`

 SetDirectorsNil sets the value for Directors to be an explicit nil

### UnsetDirectors
`func (o *NewMovieDataDetail) UnsetDirectors()`

UnsetDirectors ensures that no value is present for Directors, not even an explicit nil
### GetDirectorList

`func (o *NewMovieDataDetail) GetDirectorList() []StarShort`

GetDirectorList returns the DirectorList field if non-nil, zero value otherwise.

### GetDirectorListOk

`func (o *NewMovieDataDetail) GetDirectorListOk() (*[]StarShort, bool)`

GetDirectorListOk returns a tuple with the DirectorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorList

`func (o *NewMovieDataDetail) SetDirectorList(v []StarShort)`

SetDirectorList sets DirectorList field to given value.

### HasDirectorList

`func (o *NewMovieDataDetail) HasDirectorList() bool`

HasDirectorList returns a boolean if a field has been set.

### SetDirectorListNil

`func (o *NewMovieDataDetail) SetDirectorListNil(b bool)`

 SetDirectorListNil sets the value for DirectorList to be an explicit nil

### UnsetDirectorList
`func (o *NewMovieDataDetail) UnsetDirectorList()`

UnsetDirectorList ensures that no value is present for DirectorList, not even an explicit nil
### GetStars

`func (o *NewMovieDataDetail) GetStars() string`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *NewMovieDataDetail) GetStarsOk() (*string, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *NewMovieDataDetail) SetStars(v string)`

SetStars sets Stars field to given value.

### HasStars

`func (o *NewMovieDataDetail) HasStars() bool`

HasStars returns a boolean if a field has been set.

### SetStarsNil

`func (o *NewMovieDataDetail) SetStarsNil(b bool)`

 SetStarsNil sets the value for Stars to be an explicit nil

### UnsetStars
`func (o *NewMovieDataDetail) UnsetStars()`

UnsetStars ensures that no value is present for Stars, not even an explicit nil
### GetStarList

`func (o *NewMovieDataDetail) GetStarList() []StarShort`

GetStarList returns the StarList field if non-nil, zero value otherwise.

### GetStarListOk

`func (o *NewMovieDataDetail) GetStarListOk() (*[]StarShort, bool)`

GetStarListOk returns a tuple with the StarList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarList

`func (o *NewMovieDataDetail) SetStarList(v []StarShort)`

SetStarList sets StarList field to given value.

### HasStarList

`func (o *NewMovieDataDetail) HasStarList() bool`

HasStarList returns a boolean if a field has been set.

### SetStarListNil

`func (o *NewMovieDataDetail) SetStarListNil(b bool)`

 SetStarListNil sets the value for StarList to be an explicit nil

### UnsetStarList
`func (o *NewMovieDataDetail) UnsetStarList()`

UnsetStarList ensures that no value is present for StarList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


