# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | **string** | The unique identifier of the video object. | 
**CreatedAt** | Pointer to **string** | When a video was created, presented in ATOM UTC format. | [optional] 
**Title** | Pointer to **string** | The title of the video content.  | [optional] 
**Description** | Pointer to **string** | A description for the video content.  | [optional] 
**PublishedAt** | Pointer to **string** | The date and time the API created the video. Date and time are provided using ATOM UTC format. | [optional] 
**UpdatedAt** | Pointer to **string** | The date and time the video was updated. Date and time are provided using ATOM UTC format. | [optional] 
**DiscardedAt** | Pointer to **NullableString** | The date and time the video was discarded. The API populates this field only if you have the Video Restore feature enabled and discard a video. Date and time are provided using ATOM UTC format. | [optional] 
**DeletesAt** | Pointer to **NullableString** | The date and time the video will be permanently deleted. The API populates this field only if you have the Video Restore feature enabled and discard a video. Discarded videos are pemanently deleted after 90 days. Date and time are provided using ATOM UTC format. | [optional] 
**Discarded** | Pointer to **bool** | Returns &#x60;true&#x60; for videos you discarded when you have the Video Restore feature enabled. Returns &#x60;false&#x60; for every other video. | [optional] 
**Language** | Pointer to **string** | Returns the language of a video in [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) format. You can set the language during video creation via the API, otherwise it is detected automatically. | [optional] 
**LanguageOrigin** | Pointer to **NullableString** | Returns the origin of the last update on the video&#39;s &#x60;language&#x60; attribute.  - &#x60;api&#x60; means that the last update was requested from the API. - &#x60;auto&#x60; means that the last update was done automatically by the API. | [optional] 
**Tags** | Pointer to **[]string** | One array of tags (each tag is a string) in order to categorize a video. Tags may include spaces.   | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Metadata you can use to categorise and filter videos. Metadata is a list of dictionaries, where each dictionary represents a key value pair for categorising a video.  | [optional] 
**Source** | Pointer to [**VideoSource**](VideoSource.md) |  | [optional] 
**Assets** | Pointer to [**VideoAssets**](VideoAssets.md) |  | [optional] 
**PlayerId** | Pointer to **string** | The id of the player that will be applied on the video.  | [optional] 
**Public** | Pointer to **bool** | Defines if the content is publicly reachable or if a unique token is needed for each play session. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos/).  | [optional] 
**Panoramic** | Pointer to **bool** | Defines if video is panoramic.  | [optional] 
**Mp4Support** | Pointer to **bool** | This lets you know whether mp4 is supported. If enabled, an mp4 URL will be provided in the response for the video.  | [optional] 

## Methods

### NewVideo

`func NewVideo(videoId string, ) *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *Video) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *Video) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *Video) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.


### GetCreatedAt

`func (o *Video) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Video) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Video) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Video) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *Video) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Video) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Video) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Video) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Video) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Video) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Video) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Video) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Video) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Video) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Video) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Video) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Video) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Video) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Video) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Video) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDiscardedAt

`func (o *Video) GetDiscardedAt() string`

GetDiscardedAt returns the DiscardedAt field if non-nil, zero value otherwise.

### GetDiscardedAtOk

`func (o *Video) GetDiscardedAtOk() (*string, bool)`

GetDiscardedAtOk returns a tuple with the DiscardedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedAt

`func (o *Video) SetDiscardedAt(v string)`

SetDiscardedAt sets DiscardedAt field to given value.

### HasDiscardedAt

`func (o *Video) HasDiscardedAt() bool`

HasDiscardedAt returns a boolean if a field has been set.

### SetDiscardedAtNil

`func (o *Video) SetDiscardedAtNil(b bool)`

 SetDiscardedAtNil sets the value for DiscardedAt to be an explicit nil

### UnsetDiscardedAt
`func (o *Video) UnsetDiscardedAt()`

UnsetDiscardedAt ensures that no value is present for DiscardedAt, not even an explicit nil
### GetDeletesAt

`func (o *Video) GetDeletesAt() string`

GetDeletesAt returns the DeletesAt field if non-nil, zero value otherwise.

### GetDeletesAtOk

`func (o *Video) GetDeletesAtOk() (*string, bool)`

GetDeletesAtOk returns a tuple with the DeletesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletesAt

`func (o *Video) SetDeletesAt(v string)`

SetDeletesAt sets DeletesAt field to given value.

### HasDeletesAt

`func (o *Video) HasDeletesAt() bool`

HasDeletesAt returns a boolean if a field has been set.

### SetDeletesAtNil

`func (o *Video) SetDeletesAtNil(b bool)`

 SetDeletesAtNil sets the value for DeletesAt to be an explicit nil

### UnsetDeletesAt
`func (o *Video) UnsetDeletesAt()`

UnsetDeletesAt ensures that no value is present for DeletesAt, not even an explicit nil
### GetDiscarded

`func (o *Video) GetDiscarded() bool`

GetDiscarded returns the Discarded field if non-nil, zero value otherwise.

### GetDiscardedOk

`func (o *Video) GetDiscardedOk() (*bool, bool)`

GetDiscardedOk returns a tuple with the Discarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarded

`func (o *Video) SetDiscarded(v bool)`

SetDiscarded sets Discarded field to given value.

### HasDiscarded

`func (o *Video) HasDiscarded() bool`

HasDiscarded returns a boolean if a field has been set.

### GetLanguage

`func (o *Video) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Video) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Video) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Video) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguageOrigin

`func (o *Video) GetLanguageOrigin() string`

GetLanguageOrigin returns the LanguageOrigin field if non-nil, zero value otherwise.

### GetLanguageOriginOk

`func (o *Video) GetLanguageOriginOk() (*string, bool)`

GetLanguageOriginOk returns a tuple with the LanguageOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOrigin

`func (o *Video) SetLanguageOrigin(v string)`

SetLanguageOrigin sets LanguageOrigin field to given value.

### HasLanguageOrigin

`func (o *Video) HasLanguageOrigin() bool`

HasLanguageOrigin returns a boolean if a field has been set.

### SetLanguageOriginNil

`func (o *Video) SetLanguageOriginNil(b bool)`

 SetLanguageOriginNil sets the value for LanguageOrigin to be an explicit nil

### UnsetLanguageOrigin
`func (o *Video) UnsetLanguageOrigin()`

UnsetLanguageOrigin ensures that no value is present for LanguageOrigin, not even an explicit nil
### GetTags

`func (o *Video) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Video) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Video) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Video) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *Video) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Video) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Video) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Video) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSource

`func (o *Video) GetSource() VideoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Video) GetSourceOk() (*VideoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Video) SetSource(v VideoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Video) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAssets

`func (o *Video) GetAssets() VideoAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Video) GetAssetsOk() (*VideoAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Video) SetAssets(v VideoAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Video) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPlayerId

`func (o *Video) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *Video) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *Video) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *Video) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetPublic

`func (o *Video) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Video) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Video) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Video) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPanoramic

`func (o *Video) GetPanoramic() bool`

GetPanoramic returns the Panoramic field if non-nil, zero value otherwise.

### GetPanoramicOk

`func (o *Video) GetPanoramicOk() (*bool, bool)`

GetPanoramicOk returns a tuple with the Panoramic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramic

`func (o *Video) SetPanoramic(v bool)`

SetPanoramic sets Panoramic field to given value.

### HasPanoramic

`func (o *Video) HasPanoramic() bool`

HasPanoramic returns a boolean if a field has been set.

### GetMp4Support

`func (o *Video) GetMp4Support() bool`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *Video) GetMp4SupportOk() (*bool, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *Video) SetMp4Support(v bool)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *Video) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


