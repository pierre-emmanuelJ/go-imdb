# TitleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**OriginalTitle** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**RuntimeMins** | Pointer to **NullableString** |  | [optional] 
**RuntimeStr** | Pointer to **NullableString** |  | [optional] 
**Plot** | Pointer to **NullableString** |  | [optional] 
**PlotLocal** | Pointer to **NullableString** |  | [optional] 
**PlotLocalIsRtl** | Pointer to **bool** |  | [optional] 
**Awards** | Pointer to **NullableString** |  | [optional] 
**Directors** | Pointer to **NullableString** |  | [optional] 
**DirectorList** | Pointer to [**[]StarShort**](StarShort.md) |  | [optional] 
**Writers** | Pointer to **NullableString** |  | [optional] 
**WriterList** | Pointer to [**[]StarShort**](StarShort.md) |  | [optional] 
**Stars** | Pointer to **NullableString** |  | [optional] 
**StarList** | Pointer to [**[]StarShort**](StarShort.md) |  | [optional] 
**ActorList** | Pointer to [**[]ActorShort**](ActorShort.md) |  | [optional] 
**FullCast** | Pointer to [**FullCastData**](FullCastData.md) |  | [optional] 
**Genres** | Pointer to **NullableString** |  | [optional] 
**GenreList** | Pointer to [**[]KeyValueItem**](KeyValueItem.md) |  | [optional] 
**Companies** | Pointer to **NullableString** |  | [optional] 
**CompanyList** | Pointer to [**[]CompanyShort**](CompanyShort.md) |  | [optional] 
**Countries** | Pointer to **NullableString** |  | [optional] 
**CountryList** | Pointer to [**[]KeyValueItem**](KeyValueItem.md) |  | [optional] 
**Languages** | Pointer to **NullableString** |  | [optional] 
**LanguageList** | Pointer to [**[]KeyValueItem**](KeyValueItem.md) |  | [optional] 
**ContentRating** | Pointer to **NullableString** |  | [optional] 
**ImDbRating** | Pointer to **NullableString** |  | [optional] 
**ImDbRatingVotes** | Pointer to **NullableString** |  | [optional] 
**MetacriticRating** | Pointer to **NullableString** |  | [optional] 
**Ratings** | Pointer to [**RatingData**](RatingData.md) |  | [optional] 
**Wikipedia** | Pointer to [**WikipediaData**](WikipediaData.md) |  | [optional] 
**Posters** | Pointer to [**PosterData**](PosterData.md) |  | [optional] 
**Images** | Pointer to [**ImageData**](ImageData.md) |  | [optional] 
**Trailer** | Pointer to [**TrailerData**](TrailerData.md) |  | [optional] 
**BoxOffice** | Pointer to [**BoxOfficeShort**](BoxOfficeShort.md) |  | [optional] 
**Tagline** | Pointer to **NullableString** |  | [optional] 
**Keywords** | Pointer to **NullableString** |  | [optional] 
**KeywordList** | Pointer to **[]string** |  | [optional] 
**Similars** | Pointer to [**[]SimilarShort**](SimilarShort.md) |  | [optional] 
**TvSeriesInfo** | Pointer to [**TvSeriesInfo**](TvSeriesInfo.md) |  | [optional] 
**TvEpisodeInfo** | Pointer to [**TvEpisodeInfo**](TvEpisodeInfo.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTitleData

`func NewTitleData() *TitleData`

NewTitleData instantiates a new TitleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleDataWithDefaults

`func NewTitleDataWithDefaults() *TitleData`

NewTitleDataWithDefaults instantiates a new TitleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TitleData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TitleData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TitleData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TitleData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TitleData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TitleData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *TitleData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TitleData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TitleData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TitleData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TitleData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TitleData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetOriginalTitle

`func (o *TitleData) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *TitleData) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *TitleData) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *TitleData) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *TitleData) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *TitleData) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetFullTitle

`func (o *TitleData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *TitleData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *TitleData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *TitleData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *TitleData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *TitleData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *TitleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TitleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TitleData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TitleData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TitleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TitleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *TitleData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TitleData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TitleData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *TitleData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *TitleData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *TitleData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetImage

`func (o *TitleData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TitleData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TitleData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *TitleData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *TitleData) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *TitleData) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetReleaseDate

`func (o *TitleData) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *TitleData) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *TitleData) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *TitleData) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *TitleData) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *TitleData) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetRuntimeMins

`func (o *TitleData) GetRuntimeMins() string`

GetRuntimeMins returns the RuntimeMins field if non-nil, zero value otherwise.

### GetRuntimeMinsOk

`func (o *TitleData) GetRuntimeMinsOk() (*string, bool)`

GetRuntimeMinsOk returns a tuple with the RuntimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMins

`func (o *TitleData) SetRuntimeMins(v string)`

SetRuntimeMins sets RuntimeMins field to given value.

### HasRuntimeMins

`func (o *TitleData) HasRuntimeMins() bool`

HasRuntimeMins returns a boolean if a field has been set.

### SetRuntimeMinsNil

`func (o *TitleData) SetRuntimeMinsNil(b bool)`

 SetRuntimeMinsNil sets the value for RuntimeMins to be an explicit nil

### UnsetRuntimeMins
`func (o *TitleData) UnsetRuntimeMins()`

UnsetRuntimeMins ensures that no value is present for RuntimeMins, not even an explicit nil
### GetRuntimeStr

`func (o *TitleData) GetRuntimeStr() string`

GetRuntimeStr returns the RuntimeStr field if non-nil, zero value otherwise.

### GetRuntimeStrOk

`func (o *TitleData) GetRuntimeStrOk() (*string, bool)`

GetRuntimeStrOk returns a tuple with the RuntimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStr

`func (o *TitleData) SetRuntimeStr(v string)`

SetRuntimeStr sets RuntimeStr field to given value.

### HasRuntimeStr

`func (o *TitleData) HasRuntimeStr() bool`

HasRuntimeStr returns a boolean if a field has been set.

### SetRuntimeStrNil

`func (o *TitleData) SetRuntimeStrNil(b bool)`

 SetRuntimeStrNil sets the value for RuntimeStr to be an explicit nil

### UnsetRuntimeStr
`func (o *TitleData) UnsetRuntimeStr()`

UnsetRuntimeStr ensures that no value is present for RuntimeStr, not even an explicit nil
### GetPlot

`func (o *TitleData) GetPlot() string`

GetPlot returns the Plot field if non-nil, zero value otherwise.

### GetPlotOk

`func (o *TitleData) GetPlotOk() (*string, bool)`

GetPlotOk returns a tuple with the Plot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlot

`func (o *TitleData) SetPlot(v string)`

SetPlot sets Plot field to given value.

### HasPlot

`func (o *TitleData) HasPlot() bool`

HasPlot returns a boolean if a field has been set.

### SetPlotNil

`func (o *TitleData) SetPlotNil(b bool)`

 SetPlotNil sets the value for Plot to be an explicit nil

### UnsetPlot
`func (o *TitleData) UnsetPlot()`

UnsetPlot ensures that no value is present for Plot, not even an explicit nil
### GetPlotLocal

`func (o *TitleData) GetPlotLocal() string`

GetPlotLocal returns the PlotLocal field if non-nil, zero value otherwise.

### GetPlotLocalOk

`func (o *TitleData) GetPlotLocalOk() (*string, bool)`

GetPlotLocalOk returns a tuple with the PlotLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotLocal

`func (o *TitleData) SetPlotLocal(v string)`

SetPlotLocal sets PlotLocal field to given value.

### HasPlotLocal

`func (o *TitleData) HasPlotLocal() bool`

HasPlotLocal returns a boolean if a field has been set.

### SetPlotLocalNil

`func (o *TitleData) SetPlotLocalNil(b bool)`

 SetPlotLocalNil sets the value for PlotLocal to be an explicit nil

### UnsetPlotLocal
`func (o *TitleData) UnsetPlotLocal()`

UnsetPlotLocal ensures that no value is present for PlotLocal, not even an explicit nil
### GetPlotLocalIsRtl

`func (o *TitleData) GetPlotLocalIsRtl() bool`

GetPlotLocalIsRtl returns the PlotLocalIsRtl field if non-nil, zero value otherwise.

### GetPlotLocalIsRtlOk

`func (o *TitleData) GetPlotLocalIsRtlOk() (*bool, bool)`

GetPlotLocalIsRtlOk returns a tuple with the PlotLocalIsRtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotLocalIsRtl

`func (o *TitleData) SetPlotLocalIsRtl(v bool)`

SetPlotLocalIsRtl sets PlotLocalIsRtl field to given value.

### HasPlotLocalIsRtl

`func (o *TitleData) HasPlotLocalIsRtl() bool`

HasPlotLocalIsRtl returns a boolean if a field has been set.

### GetAwards

`func (o *TitleData) GetAwards() string`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *TitleData) GetAwardsOk() (*string, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *TitleData) SetAwards(v string)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *TitleData) HasAwards() bool`

HasAwards returns a boolean if a field has been set.

### SetAwardsNil

`func (o *TitleData) SetAwardsNil(b bool)`

 SetAwardsNil sets the value for Awards to be an explicit nil

### UnsetAwards
`func (o *TitleData) UnsetAwards()`

UnsetAwards ensures that no value is present for Awards, not even an explicit nil
### GetDirectors

`func (o *TitleData) GetDirectors() string`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *TitleData) GetDirectorsOk() (*string, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *TitleData) SetDirectors(v string)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *TitleData) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### SetDirectorsNil

`func (o *TitleData) SetDirectorsNil(b bool)`

 SetDirectorsNil sets the value for Directors to be an explicit nil

### UnsetDirectors
`func (o *TitleData) UnsetDirectors()`

UnsetDirectors ensures that no value is present for Directors, not even an explicit nil
### GetDirectorList

`func (o *TitleData) GetDirectorList() []StarShort`

GetDirectorList returns the DirectorList field if non-nil, zero value otherwise.

### GetDirectorListOk

`func (o *TitleData) GetDirectorListOk() (*[]StarShort, bool)`

GetDirectorListOk returns a tuple with the DirectorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorList

`func (o *TitleData) SetDirectorList(v []StarShort)`

SetDirectorList sets DirectorList field to given value.

### HasDirectorList

`func (o *TitleData) HasDirectorList() bool`

HasDirectorList returns a boolean if a field has been set.

### SetDirectorListNil

`func (o *TitleData) SetDirectorListNil(b bool)`

 SetDirectorListNil sets the value for DirectorList to be an explicit nil

### UnsetDirectorList
`func (o *TitleData) UnsetDirectorList()`

UnsetDirectorList ensures that no value is present for DirectorList, not even an explicit nil
### GetWriters

`func (o *TitleData) GetWriters() string`

GetWriters returns the Writers field if non-nil, zero value otherwise.

### GetWritersOk

`func (o *TitleData) GetWritersOk() (*string, bool)`

GetWritersOk returns a tuple with the Writers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriters

`func (o *TitleData) SetWriters(v string)`

SetWriters sets Writers field to given value.

### HasWriters

`func (o *TitleData) HasWriters() bool`

HasWriters returns a boolean if a field has been set.

### SetWritersNil

`func (o *TitleData) SetWritersNil(b bool)`

 SetWritersNil sets the value for Writers to be an explicit nil

### UnsetWriters
`func (o *TitleData) UnsetWriters()`

UnsetWriters ensures that no value is present for Writers, not even an explicit nil
### GetWriterList

`func (o *TitleData) GetWriterList() []StarShort`

GetWriterList returns the WriterList field if non-nil, zero value otherwise.

### GetWriterListOk

`func (o *TitleData) GetWriterListOk() (*[]StarShort, bool)`

GetWriterListOk returns a tuple with the WriterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterList

`func (o *TitleData) SetWriterList(v []StarShort)`

SetWriterList sets WriterList field to given value.

### HasWriterList

`func (o *TitleData) HasWriterList() bool`

HasWriterList returns a boolean if a field has been set.

### SetWriterListNil

`func (o *TitleData) SetWriterListNil(b bool)`

 SetWriterListNil sets the value for WriterList to be an explicit nil

### UnsetWriterList
`func (o *TitleData) UnsetWriterList()`

UnsetWriterList ensures that no value is present for WriterList, not even an explicit nil
### GetStars

`func (o *TitleData) GetStars() string`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *TitleData) GetStarsOk() (*string, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *TitleData) SetStars(v string)`

SetStars sets Stars field to given value.

### HasStars

`func (o *TitleData) HasStars() bool`

HasStars returns a boolean if a field has been set.

### SetStarsNil

`func (o *TitleData) SetStarsNil(b bool)`

 SetStarsNil sets the value for Stars to be an explicit nil

### UnsetStars
`func (o *TitleData) UnsetStars()`

UnsetStars ensures that no value is present for Stars, not even an explicit nil
### GetStarList

`func (o *TitleData) GetStarList() []StarShort`

GetStarList returns the StarList field if non-nil, zero value otherwise.

### GetStarListOk

`func (o *TitleData) GetStarListOk() (*[]StarShort, bool)`

GetStarListOk returns a tuple with the StarList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarList

`func (o *TitleData) SetStarList(v []StarShort)`

SetStarList sets StarList field to given value.

### HasStarList

`func (o *TitleData) HasStarList() bool`

HasStarList returns a boolean if a field has been set.

### SetStarListNil

`func (o *TitleData) SetStarListNil(b bool)`

 SetStarListNil sets the value for StarList to be an explicit nil

### UnsetStarList
`func (o *TitleData) UnsetStarList()`

UnsetStarList ensures that no value is present for StarList, not even an explicit nil
### GetActorList

`func (o *TitleData) GetActorList() []ActorShort`

GetActorList returns the ActorList field if non-nil, zero value otherwise.

### GetActorListOk

`func (o *TitleData) GetActorListOk() (*[]ActorShort, bool)`

GetActorListOk returns a tuple with the ActorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorList

`func (o *TitleData) SetActorList(v []ActorShort)`

SetActorList sets ActorList field to given value.

### HasActorList

`func (o *TitleData) HasActorList() bool`

HasActorList returns a boolean if a field has been set.

### SetActorListNil

`func (o *TitleData) SetActorListNil(b bool)`

 SetActorListNil sets the value for ActorList to be an explicit nil

### UnsetActorList
`func (o *TitleData) UnsetActorList()`

UnsetActorList ensures that no value is present for ActorList, not even an explicit nil
### GetFullCast

`func (o *TitleData) GetFullCast() FullCastData`

GetFullCast returns the FullCast field if non-nil, zero value otherwise.

### GetFullCastOk

`func (o *TitleData) GetFullCastOk() (*FullCastData, bool)`

GetFullCastOk returns a tuple with the FullCast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCast

`func (o *TitleData) SetFullCast(v FullCastData)`

SetFullCast sets FullCast field to given value.

### HasFullCast

`func (o *TitleData) HasFullCast() bool`

HasFullCast returns a boolean if a field has been set.

### GetGenres

`func (o *TitleData) GetGenres() string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *TitleData) GetGenresOk() (*string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *TitleData) SetGenres(v string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *TitleData) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *TitleData) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *TitleData) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetGenreList

`func (o *TitleData) GetGenreList() []KeyValueItem`

GetGenreList returns the GenreList field if non-nil, zero value otherwise.

### GetGenreListOk

`func (o *TitleData) GetGenreListOk() (*[]KeyValueItem, bool)`

GetGenreListOk returns a tuple with the GenreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreList

`func (o *TitleData) SetGenreList(v []KeyValueItem)`

SetGenreList sets GenreList field to given value.

### HasGenreList

`func (o *TitleData) HasGenreList() bool`

HasGenreList returns a boolean if a field has been set.

### SetGenreListNil

`func (o *TitleData) SetGenreListNil(b bool)`

 SetGenreListNil sets the value for GenreList to be an explicit nil

### UnsetGenreList
`func (o *TitleData) UnsetGenreList()`

UnsetGenreList ensures that no value is present for GenreList, not even an explicit nil
### GetCompanies

`func (o *TitleData) GetCompanies() string`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *TitleData) GetCompaniesOk() (*string, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *TitleData) SetCompanies(v string)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *TitleData) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### SetCompaniesNil

`func (o *TitleData) SetCompaniesNil(b bool)`

 SetCompaniesNil sets the value for Companies to be an explicit nil

### UnsetCompanies
`func (o *TitleData) UnsetCompanies()`

UnsetCompanies ensures that no value is present for Companies, not even an explicit nil
### GetCompanyList

`func (o *TitleData) GetCompanyList() []CompanyShort`

GetCompanyList returns the CompanyList field if non-nil, zero value otherwise.

### GetCompanyListOk

`func (o *TitleData) GetCompanyListOk() (*[]CompanyShort, bool)`

GetCompanyListOk returns a tuple with the CompanyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyList

`func (o *TitleData) SetCompanyList(v []CompanyShort)`

SetCompanyList sets CompanyList field to given value.

### HasCompanyList

`func (o *TitleData) HasCompanyList() bool`

HasCompanyList returns a boolean if a field has been set.

### SetCompanyListNil

`func (o *TitleData) SetCompanyListNil(b bool)`

 SetCompanyListNil sets the value for CompanyList to be an explicit nil

### UnsetCompanyList
`func (o *TitleData) UnsetCompanyList()`

UnsetCompanyList ensures that no value is present for CompanyList, not even an explicit nil
### GetCountries

`func (o *TitleData) GetCountries() string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *TitleData) GetCountriesOk() (*string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *TitleData) SetCountries(v string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *TitleData) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### SetCountriesNil

`func (o *TitleData) SetCountriesNil(b bool)`

 SetCountriesNil sets the value for Countries to be an explicit nil

### UnsetCountries
`func (o *TitleData) UnsetCountries()`

UnsetCountries ensures that no value is present for Countries, not even an explicit nil
### GetCountryList

`func (o *TitleData) GetCountryList() []KeyValueItem`

GetCountryList returns the CountryList field if non-nil, zero value otherwise.

### GetCountryListOk

`func (o *TitleData) GetCountryListOk() (*[]KeyValueItem, bool)`

GetCountryListOk returns a tuple with the CountryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryList

`func (o *TitleData) SetCountryList(v []KeyValueItem)`

SetCountryList sets CountryList field to given value.

### HasCountryList

`func (o *TitleData) HasCountryList() bool`

HasCountryList returns a boolean if a field has been set.

### SetCountryListNil

`func (o *TitleData) SetCountryListNil(b bool)`

 SetCountryListNil sets the value for CountryList to be an explicit nil

### UnsetCountryList
`func (o *TitleData) UnsetCountryList()`

UnsetCountryList ensures that no value is present for CountryList, not even an explicit nil
### GetLanguages

`func (o *TitleData) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *TitleData) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *TitleData) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *TitleData) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *TitleData) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *TitleData) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetLanguageList

`func (o *TitleData) GetLanguageList() []KeyValueItem`

GetLanguageList returns the LanguageList field if non-nil, zero value otherwise.

### GetLanguageListOk

`func (o *TitleData) GetLanguageListOk() (*[]KeyValueItem, bool)`

GetLanguageListOk returns a tuple with the LanguageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageList

`func (o *TitleData) SetLanguageList(v []KeyValueItem)`

SetLanguageList sets LanguageList field to given value.

### HasLanguageList

`func (o *TitleData) HasLanguageList() bool`

HasLanguageList returns a boolean if a field has been set.

### SetLanguageListNil

`func (o *TitleData) SetLanguageListNil(b bool)`

 SetLanguageListNil sets the value for LanguageList to be an explicit nil

### UnsetLanguageList
`func (o *TitleData) UnsetLanguageList()`

UnsetLanguageList ensures that no value is present for LanguageList, not even an explicit nil
### GetContentRating

`func (o *TitleData) GetContentRating() string`

GetContentRating returns the ContentRating field if non-nil, zero value otherwise.

### GetContentRatingOk

`func (o *TitleData) GetContentRatingOk() (*string, bool)`

GetContentRatingOk returns a tuple with the ContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRating

`func (o *TitleData) SetContentRating(v string)`

SetContentRating sets ContentRating field to given value.

### HasContentRating

`func (o *TitleData) HasContentRating() bool`

HasContentRating returns a boolean if a field has been set.

### SetContentRatingNil

`func (o *TitleData) SetContentRatingNil(b bool)`

 SetContentRatingNil sets the value for ContentRating to be an explicit nil

### UnsetContentRating
`func (o *TitleData) UnsetContentRating()`

UnsetContentRating ensures that no value is present for ContentRating, not even an explicit nil
### GetImDbRating

`func (o *TitleData) GetImDbRating() string`

GetImDbRating returns the ImDbRating field if non-nil, zero value otherwise.

### GetImDbRatingOk

`func (o *TitleData) GetImDbRatingOk() (*string, bool)`

GetImDbRatingOk returns a tuple with the ImDbRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbRating

`func (o *TitleData) SetImDbRating(v string)`

SetImDbRating sets ImDbRating field to given value.

### HasImDbRating

`func (o *TitleData) HasImDbRating() bool`

HasImDbRating returns a boolean if a field has been set.

### SetImDbRatingNil

`func (o *TitleData) SetImDbRatingNil(b bool)`

 SetImDbRatingNil sets the value for ImDbRating to be an explicit nil

### UnsetImDbRating
`func (o *TitleData) UnsetImDbRating()`

UnsetImDbRating ensures that no value is present for ImDbRating, not even an explicit nil
### GetImDbRatingVotes

`func (o *TitleData) GetImDbRatingVotes() string`

GetImDbRatingVotes returns the ImDbRatingVotes field if non-nil, zero value otherwise.

### GetImDbRatingVotesOk

`func (o *TitleData) GetImDbRatingVotesOk() (*string, bool)`

GetImDbRatingVotesOk returns a tuple with the ImDbRatingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbRatingVotes

`func (o *TitleData) SetImDbRatingVotes(v string)`

SetImDbRatingVotes sets ImDbRatingVotes field to given value.

### HasImDbRatingVotes

`func (o *TitleData) HasImDbRatingVotes() bool`

HasImDbRatingVotes returns a boolean if a field has been set.

### SetImDbRatingVotesNil

`func (o *TitleData) SetImDbRatingVotesNil(b bool)`

 SetImDbRatingVotesNil sets the value for ImDbRatingVotes to be an explicit nil

### UnsetImDbRatingVotes
`func (o *TitleData) UnsetImDbRatingVotes()`

UnsetImDbRatingVotes ensures that no value is present for ImDbRatingVotes, not even an explicit nil
### GetMetacriticRating

`func (o *TitleData) GetMetacriticRating() string`

GetMetacriticRating returns the MetacriticRating field if non-nil, zero value otherwise.

### GetMetacriticRatingOk

`func (o *TitleData) GetMetacriticRatingOk() (*string, bool)`

GetMetacriticRatingOk returns a tuple with the MetacriticRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetacriticRating

`func (o *TitleData) SetMetacriticRating(v string)`

SetMetacriticRating sets MetacriticRating field to given value.

### HasMetacriticRating

`func (o *TitleData) HasMetacriticRating() bool`

HasMetacriticRating returns a boolean if a field has been set.

### SetMetacriticRatingNil

`func (o *TitleData) SetMetacriticRatingNil(b bool)`

 SetMetacriticRatingNil sets the value for MetacriticRating to be an explicit nil

### UnsetMetacriticRating
`func (o *TitleData) UnsetMetacriticRating()`

UnsetMetacriticRating ensures that no value is present for MetacriticRating, not even an explicit nil
### GetRatings

`func (o *TitleData) GetRatings() RatingData`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *TitleData) GetRatingsOk() (*RatingData, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *TitleData) SetRatings(v RatingData)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *TitleData) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetWikipedia

`func (o *TitleData) GetWikipedia() WikipediaData`

GetWikipedia returns the Wikipedia field if non-nil, zero value otherwise.

### GetWikipediaOk

`func (o *TitleData) GetWikipediaOk() (*WikipediaData, bool)`

GetWikipediaOk returns a tuple with the Wikipedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikipedia

`func (o *TitleData) SetWikipedia(v WikipediaData)`

SetWikipedia sets Wikipedia field to given value.

### HasWikipedia

`func (o *TitleData) HasWikipedia() bool`

HasWikipedia returns a boolean if a field has been set.

### GetPosters

`func (o *TitleData) GetPosters() PosterData`

GetPosters returns the Posters field if non-nil, zero value otherwise.

### GetPostersOk

`func (o *TitleData) GetPostersOk() (*PosterData, bool)`

GetPostersOk returns a tuple with the Posters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosters

`func (o *TitleData) SetPosters(v PosterData)`

SetPosters sets Posters field to given value.

### HasPosters

`func (o *TitleData) HasPosters() bool`

HasPosters returns a boolean if a field has been set.

### GetImages

`func (o *TitleData) GetImages() ImageData`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *TitleData) GetImagesOk() (*ImageData, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *TitleData) SetImages(v ImageData)`

SetImages sets Images field to given value.

### HasImages

`func (o *TitleData) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetTrailer

`func (o *TitleData) GetTrailer() TrailerData`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *TitleData) GetTrailerOk() (*TrailerData, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *TitleData) SetTrailer(v TrailerData)`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *TitleData) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.

### GetBoxOffice

`func (o *TitleData) GetBoxOffice() BoxOfficeShort`

GetBoxOffice returns the BoxOffice field if non-nil, zero value otherwise.

### GetBoxOfficeOk

`func (o *TitleData) GetBoxOfficeOk() (*BoxOfficeShort, bool)`

GetBoxOfficeOk returns a tuple with the BoxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxOffice

`func (o *TitleData) SetBoxOffice(v BoxOfficeShort)`

SetBoxOffice sets BoxOffice field to given value.

### HasBoxOffice

`func (o *TitleData) HasBoxOffice() bool`

HasBoxOffice returns a boolean if a field has been set.

### GetTagline

`func (o *TitleData) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *TitleData) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *TitleData) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *TitleData) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### SetTaglineNil

`func (o *TitleData) SetTaglineNil(b bool)`

 SetTaglineNil sets the value for Tagline to be an explicit nil

### UnsetTagline
`func (o *TitleData) UnsetTagline()`

UnsetTagline ensures that no value is present for Tagline, not even an explicit nil
### GetKeywords

`func (o *TitleData) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *TitleData) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *TitleData) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *TitleData) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *TitleData) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *TitleData) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetKeywordList

`func (o *TitleData) GetKeywordList() []string`

GetKeywordList returns the KeywordList field if non-nil, zero value otherwise.

### GetKeywordListOk

`func (o *TitleData) GetKeywordListOk() (*[]string, bool)`

GetKeywordListOk returns a tuple with the KeywordList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordList

`func (o *TitleData) SetKeywordList(v []string)`

SetKeywordList sets KeywordList field to given value.

### HasKeywordList

`func (o *TitleData) HasKeywordList() bool`

HasKeywordList returns a boolean if a field has been set.

### SetKeywordListNil

`func (o *TitleData) SetKeywordListNil(b bool)`

 SetKeywordListNil sets the value for KeywordList to be an explicit nil

### UnsetKeywordList
`func (o *TitleData) UnsetKeywordList()`

UnsetKeywordList ensures that no value is present for KeywordList, not even an explicit nil
### GetSimilars

`func (o *TitleData) GetSimilars() []SimilarShort`

GetSimilars returns the Similars field if non-nil, zero value otherwise.

### GetSimilarsOk

`func (o *TitleData) GetSimilarsOk() (*[]SimilarShort, bool)`

GetSimilarsOk returns a tuple with the Similars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilars

`func (o *TitleData) SetSimilars(v []SimilarShort)`

SetSimilars sets Similars field to given value.

### HasSimilars

`func (o *TitleData) HasSimilars() bool`

HasSimilars returns a boolean if a field has been set.

### SetSimilarsNil

`func (o *TitleData) SetSimilarsNil(b bool)`

 SetSimilarsNil sets the value for Similars to be an explicit nil

### UnsetSimilars
`func (o *TitleData) UnsetSimilars()`

UnsetSimilars ensures that no value is present for Similars, not even an explicit nil
### GetTvSeriesInfo

`func (o *TitleData) GetTvSeriesInfo() TvSeriesInfo`

GetTvSeriesInfo returns the TvSeriesInfo field if non-nil, zero value otherwise.

### GetTvSeriesInfoOk

`func (o *TitleData) GetTvSeriesInfoOk() (*TvSeriesInfo, bool)`

GetTvSeriesInfoOk returns a tuple with the TvSeriesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvSeriesInfo

`func (o *TitleData) SetTvSeriesInfo(v TvSeriesInfo)`

SetTvSeriesInfo sets TvSeriesInfo field to given value.

### HasTvSeriesInfo

`func (o *TitleData) HasTvSeriesInfo() bool`

HasTvSeriesInfo returns a boolean if a field has been set.

### GetTvEpisodeInfo

`func (o *TitleData) GetTvEpisodeInfo() TvEpisodeInfo`

GetTvEpisodeInfo returns the TvEpisodeInfo field if non-nil, zero value otherwise.

### GetTvEpisodeInfoOk

`func (o *TitleData) GetTvEpisodeInfoOk() (*TvEpisodeInfo, bool)`

GetTvEpisodeInfoOk returns a tuple with the TvEpisodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvEpisodeInfo

`func (o *TitleData) SetTvEpisodeInfo(v TvEpisodeInfo)`

SetTvEpisodeInfo sets TvEpisodeInfo field to given value.

### HasTvEpisodeInfo

`func (o *TitleData) HasTvEpisodeInfo() bool`

HasTvEpisodeInfo returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TitleData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TitleData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TitleData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TitleData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TitleData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TitleData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


