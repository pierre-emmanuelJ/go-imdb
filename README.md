# Go API client for client

The IMDb-API Documentation. You need a <a href='/Identity/Account/Manage' target='_blank'><code>API Key</code></a> for testing APIs.<br/><a class='link' href='/API'>Back to API Tester</a>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.8.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://imdb-api.com](https://imdb-api.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MoviesApiApi* | [**APIBoxOfficeAllTimeApiKeyGet**](docs/MoviesApiApi.md#apiboxofficealltimeapikeyget) | **Get** /API/BoxOfficeAllTime/{apiKey} | 
*MoviesApiApi* | [**APIBoxOfficeApiKeyGet**](docs/MoviesApiApi.md#apiboxofficeapikeyget) | **Get** /API/BoxOffice/{apiKey} | 
*MoviesApiApi* | [**APIComingSoonApiKeyGet**](docs/MoviesApiApi.md#apicomingsoonapikeyget) | **Get** /API/ComingSoon/{apiKey} | 
*MoviesApiApi* | [**APICompanyApiKeyIdGet**](docs/MoviesApiApi.md#apicompanyapikeyidget) | **Get** /API/Company/{apiKey}/{id} | 
*MoviesApiApi* | [**APIExternalSitesApiKeyIdGet**](docs/MoviesApiApi.md#apiexternalsitesapikeyidget) | **Get** /API/ExternalSites/{apiKey}/{id} | 
*MoviesApiApi* | [**APIFAQApiKeyIdGet**](docs/MoviesApiApi.md#apifaqapikeyidget) | **Get** /API/FAQ/{apiKey}/{id} | 
*MoviesApiApi* | [**APIFullCastApiKeyIdGet**](docs/MoviesApiApi.md#apifullcastapikeyidget) | **Get** /API/FullCast/{apiKey}/{id} | 
*MoviesApiApi* | [**APIIMDbListApiKeyIdGet**](docs/MoviesApiApi.md#apiimdblistapikeyidget) | **Get** /API/IMDbList/{apiKey}/{id} | 
*MoviesApiApi* | [**APIImagesApiKeyIdGet**](docs/MoviesApiApi.md#apiimagesapikeyidget) | **Get** /API/Images/{apiKey}/{id} | 
*MoviesApiApi* | [**APIImagesApiKeyIdOptionsGet**](docs/MoviesApiApi.md#apiimagesapikeyidoptionsget) | **Get** /API/Images/{apiKey}/{id}/{options} | 
*MoviesApiApi* | [**APIInTheatersApiKeyGet**](docs/MoviesApiApi.md#apiintheatersapikeyget) | **Get** /API/InTheaters/{apiKey} | 
*MoviesApiApi* | [**APIKeywordApiKeyIdGet**](docs/MoviesApiApi.md#apikeywordapikeyidget) | **Get** /API/Keyword/{apiKey}/{id} | 
*MoviesApiApi* | [**APIMetacriticReviewsApiKeyIdGet**](docs/MoviesApiApi.md#apimetacriticreviewsapikeyidget) | **Get** /API/MetacriticReviews/{apiKey}/{id} | 
*MoviesApiApi* | [**APIMostPopularMoviesApiKeyGet**](docs/MoviesApiApi.md#apimostpopularmoviesapikeyget) | **Get** /API/MostPopularMovies/{apiKey} | 
*MoviesApiApi* | [**APIMostPopularTVsApiKeyGet**](docs/MoviesApiApi.md#apimostpopulartvsapikeyget) | **Get** /API/MostPopularTVs/{apiKey} | 
*MoviesApiApi* | [**APINameApiKeyIdGet**](docs/MoviesApiApi.md#apinameapikeyidget) | **Get** /API/Name/{apiKey}/{id} | 
*MoviesApiApi* | [**APIPostersApiKeyIdGet**](docs/MoviesApiApi.md#apipostersapikeyidget) | **Get** /API/Posters/{apiKey}/{id} | 
*MoviesApiApi* | [**APIRatingsApiKeyIdGet**](docs/MoviesApiApi.md#apiratingsapikeyidget) | **Get** /API/Ratings/{apiKey}/{id} | 
*MoviesApiApi* | [**APIReviewsApiKeyIdGet**](docs/MoviesApiApi.md#apireviewsapikeyidget) | **Get** /API/Reviews/{apiKey}/{id} | 
*MoviesApiApi* | [**APISearchAllApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchallapikeyexpressionget) | **Get** /API/SearchAll/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchapikeyexpressionget) | **Get** /API/Search/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchCompanyApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchcompanyapikeyexpressionget) | **Get** /API/SearchCompany/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchEpisodeApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchepisodeapikeyexpressionget) | **Get** /API/SearchEpisode/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchKeywordApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchkeywordapikeyexpressionget) | **Get** /API/SearchKeyword/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchMovieApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchmovieapikeyexpressionget) | **Get** /API/SearchMovie/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchNameApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchnameapikeyexpressionget) | **Get** /API/SearchName/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchSeriesApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchseriesapikeyexpressionget) | **Get** /API/SearchSeries/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISearchTitleApiKeyExpressionGet**](docs/MoviesApiApi.md#apisearchtitleapikeyexpressionget) | **Get** /API/SearchTitle/{apiKey}/{expression} | 
*MoviesApiApi* | [**APISeasonEpisodesApiKeyIdSeasonNumberGet**](docs/MoviesApiApi.md#apiseasonepisodesapikeyidseasonnumberget) | **Get** /API/SeasonEpisodes/{apiKey}/{id}/{seasonNumber} | 
*MoviesApiApi* | [**APISubtitlesApiKeyIdGet**](docs/MoviesApiApi.md#apisubtitlesapikeyidget) | **Get** /API/Subtitles/{apiKey}/{id} | 
*MoviesApiApi* | [**APISubtitlesApiKeyIdSeasonNumberGet**](docs/MoviesApiApi.md#apisubtitlesapikeyidseasonnumberget) | **Get** /API/Subtitles/{apiKey}/{id}/{seasonNumber} | 
*MoviesApiApi* | [**APITop250MoviesApiKeyGet**](docs/MoviesApiApi.md#apitop250moviesapikeyget) | **Get** /API/Top250Movies/{apiKey} | 
*MoviesApiApi* | [**APITop250TVsApiKeyGet**](docs/MoviesApiApi.md#apitop250tvsapikeyget) | **Get** /API/Top250TVs/{apiKey} | 
*MoviesApiApi* | [**APITrailerApiKeyIdGet**](docs/MoviesApiApi.md#apitrailerapikeyidget) | **Get** /API/Trailer/{apiKey}/{id} | 
*MoviesApiApi* | [**APIUsageApiKeyGet**](docs/MoviesApiApi.md#apiusageapikeyget) | **Get** /API/Usage/{apiKey} | 
*MoviesApiApi* | [**APIUserRatingsApiKeyIdGet**](docs/MoviesApiApi.md#apiuserratingsapikeyidget) | **Get** /API/UserRatings/{apiKey}/{id} | 
*MoviesApiApi* | [**APIYouTubeApiKeyVGet**](docs/MoviesApiApi.md#apiyoutubeapikeyvget) | **Get** /API/YouTube/{apiKey}/{v} | 
*MoviesApiApi* | [**APIYouTubePlaylistApiKeyListGet**](docs/MoviesApiApi.md#apiyoutubeplaylistapikeylistget) | **Get** /API/YouTubePlaylist/{apiKey}/{list} | 
*MoviesApiApi* | [**APIYouTubeTrailerApiKeyIdGet**](docs/MoviesApiApi.md#apiyoutubetrailerapikeyidget) | **Get** /API/YouTubeTrailer/{apiKey}/{id} | 
*MoviesApiApi* | [**LangAPIReportApiKeyIdGet**](docs/MoviesApiApi.md#langapireportapikeyidget) | **Get** /{lang}/API/Report/{apiKey}/{id} | 
*MoviesApiApi* | [**LangAPIReportApiKeyIdOptionsGet**](docs/MoviesApiApi.md#langapireportapikeyidoptionsget) | **Get** /{lang}/API/Report/{apiKey}/{id}/{options} | 
*MoviesApiApi* | [**LangAPITitleApiKeyIdGet**](docs/MoviesApiApi.md#langapititleapikeyidget) | **Get** /{lang}/API/Title/{apiKey}/{id} | 
*MoviesApiApi* | [**LangAPITitleApiKeyIdOptionsGet**](docs/MoviesApiApi.md#langapititleapikeyidoptionsget) | **Get** /{lang}/API/Title/{apiKey}/{id}/{options} | 
*MoviesApiApi* | [**LangAPIWikipediaApiKeyIdGet**](docs/MoviesApiApi.md#langapiwikipediaapikeyidget) | **Get** /{lang}/API/Wikipedia/{apiKey}/{id} | 


## Documentation For Models

 - [ActorShort](docs/ActorShort.md)
 - [AwardData](docs/AwardData.md)
 - [AwardEvent](docs/AwardEvent.md)
 - [AwardOutcome](docs/AwardOutcome.md)
 - [AwardOutcomeDetail](docs/AwardOutcomeDetail.md)
 - [BoxOfficeAllTimeData](docs/BoxOfficeAllTimeData.md)
 - [BoxOfficeAllTimeDataDetail](docs/BoxOfficeAllTimeDataDetail.md)
 - [BoxOfficeShort](docs/BoxOfficeShort.md)
 - [BoxOfficeWeekendData](docs/BoxOfficeWeekendData.md)
 - [BoxOfficeWeekendDataDetail](docs/BoxOfficeWeekendDataDetail.md)
 - [CastMovie](docs/CastMovie.md)
 - [CastShort](docs/CastShort.md)
 - [CastShortItem](docs/CastShortItem.md)
 - [CompanyData](docs/CompanyData.md)
 - [CompanyShort](docs/CompanyShort.md)
 - [EpisodeShortDetail](docs/EpisodeShortDetail.md)
 - [ExternalSiteData](docs/ExternalSiteData.md)
 - [ExternalSiteItem](docs/ExternalSiteItem.md)
 - [FAQData](docs/FAQData.md)
 - [FAQDetail](docs/FAQDetail.md)
 - [FullCastData](docs/FullCastData.md)
 - [IMDbListData](docs/IMDbListData.md)
 - [IMDbListDataDetail](docs/IMDbListDataDetail.md)
 - [ImageData](docs/ImageData.md)
 - [ImageDataDetail](docs/ImageDataDetail.md)
 - [KeyValueItem](docs/KeyValueItem.md)
 - [KeywordData](docs/KeywordData.md)
 - [KnownFor](docs/KnownFor.md)
 - [LanguageUrl](docs/LanguageUrl.md)
 - [MetacriticReviewData](docs/MetacriticReviewData.md)
 - [MetacriticReviewDetail](docs/MetacriticReviewDetail.md)
 - [MostPopularData](docs/MostPopularData.md)
 - [MostPopularDataDetail](docs/MostPopularDataDetail.md)
 - [MovieShort](docs/MovieShort.md)
 - [NameAwardData](docs/NameAwardData.md)
 - [NameAwardEvent](docs/NameAwardEvent.md)
 - [NameAwardOutcome](docs/NameAwardOutcome.md)
 - [NameAwardOutcomeDetail](docs/NameAwardOutcomeDetail.md)
 - [NameData](docs/NameData.md)
 - [NewMovieData](docs/NewMovieData.md)
 - [NewMovieDataDetail](docs/NewMovieDataDetail.md)
 - [PosterData](docs/PosterData.md)
 - [PosterDataItem](docs/PosterDataItem.md)
 - [RatingData](docs/RatingData.md)
 - [ReviewData](docs/ReviewData.md)
 - [ReviewDetail](docs/ReviewDetail.md)
 - [SearchData](docs/SearchData.md)
 - [SearchResult](docs/SearchResult.md)
 - [SeasonEpisodeData](docs/SeasonEpisodeData.md)
 - [SimilarShort](docs/SimilarShort.md)
 - [StarShort](docs/StarShort.md)
 - [SubtitleData](docs/SubtitleData.md)
 - [SubtitleDataDetail](docs/SubtitleDataDetail.md)
 - [TitleData](docs/TitleData.md)
 - [Top250Data](docs/Top250Data.md)
 - [Top250DataDetail](docs/Top250DataDetail.md)
 - [TrailerData](docs/TrailerData.md)
 - [TvEpisodeInfo](docs/TvEpisodeInfo.md)
 - [TvSeriesInfo](docs/TvSeriesInfo.md)
 - [UsageData](docs/UsageData.md)
 - [UserRatingData](docs/UserRatingData.md)
 - [UserRatingDataDemographic](docs/UserRatingDataDemographic.md)
 - [UserRatingDataDemographicDetail](docs/UserRatingDataDemographicDetail.md)
 - [UserRatingDataDetail](docs/UserRatingDataDetail.md)
 - [WikipediaData](docs/WikipediaData.md)
 - [WikipediaDataPlot](docs/WikipediaDataPlot.md)
 - [YouTubeData](docs/YouTubeData.md)
 - [YouTubeDataItem](docs/YouTubeDataItem.md)
 - [YouTubePlaylistData](docs/YouTubePlaylistData.md)
 - [YouTubePlaylistDataItem](docs/YouTubePlaylistDataItem.md)
 - [YouTubeTrailerData](docs/YouTubeTrailerData.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



