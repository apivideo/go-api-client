# VideoUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerId** | Pointer to **NullableString** | The unique ID for the player you want to associate with your video. | [optional] 
**Title** | Pointer to **string** | The title you want to use for your video. | [optional] 
**Description** | Pointer to **string** | A brief description of the video. | [optional] 
**Public** | Pointer to **bool** | Whether the video is publicly available or not. False means it is set to private. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos/). | [optional] 
**Panoramic** | Pointer to **bool** | Whether the video is a 360 degree or immersive video. | [optional] 
**Mp4Support** | Pointer to **bool** | Whether the player supports the mp4 format. | [optional] 
**Tags** | Pointer to **[]string** | A list of terms or words you want to tag the video with. Make sure the list includes all the tags you want as whatever you send in this list will overwrite the existing list for the video. | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | A list (array) of dictionaries where each dictionary contains a key value pair that describes the video. As with tags, you must send the complete list of metadata you want as whatever you send here will overwrite the existing metadata for the video. | [optional] 
**Language** | Pointer to **NullableString** | Use this parameter to set the language of the video. When this parameter is set, the API creates a transcript of the video using the language you specify. You must use the [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) format.  &#x60;language&#x60; is a permanent attribute of the video. You can update it to another language using the [&#x60;PATCH /videos/{videoId}&#x60;](https://docs.api.video/reference/api/Videos#update-a-video-object) operation. This triggers the API to generate a new transcript using a different language. | [optional] 
**Transcript** | Pointer to **bool** | Use this parameter to enable transcription.   - When &#x60;true&#x60;, the API generates a transcript for the video. - The default value is &#x60;false&#x60;. - If you define a video language using the &#x60;language&#x60; parameter, the API uses that language to transcribe the video. If you do not define a language, the API detects it based on the video.  - When the API generates a transcript, it will be available as a caption for the video. | [optional] 
**TranscriptSummary** | Pointer to **bool** | Use this parameter to enable summarization.   - When &#x60;true&#x60;, the API generates a summary for the video, based on the transcription. - The default value is &#x60;false&#x60;. - If you define a video language using the &#x60;language&#x60; parameter, the API uses that language to summarize the video. If you do not define a language, the API detects it based on the video. | [optional] 
**TranscriptSummaryAttributes** | Pointer to **[]string** | Use this parameter to define the elements of a summary that you want to generate. If you do not define this parameter, the API generates a full summary with all attributes. The possible values are &#x60;abstract&#x60; and &#x60;takeaways&#x60;. | [optional] 

## Methods

### NewVideoUpdatePayload

`func NewVideoUpdatePayload() *VideoUpdatePayload`

NewVideoUpdatePayload instantiates a new VideoUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoUpdatePayloadWithDefaults

`func NewVideoUpdatePayloadWithDefaults() *VideoUpdatePayload`

NewVideoUpdatePayloadWithDefaults instantiates a new VideoUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerId

`func (o *VideoUpdatePayload) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *VideoUpdatePayload) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *VideoUpdatePayload) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *VideoUpdatePayload) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### SetPlayerIdNil

`func (o *VideoUpdatePayload) SetPlayerIdNil(b bool)`

 SetPlayerIdNil sets the value for PlayerId to be an explicit nil

### UnsetPlayerId
`func (o *VideoUpdatePayload) UnsetPlayerId()`

UnsetPlayerId ensures that no value is present for PlayerId, not even an explicit nil
### GetTitle

`func (o *VideoUpdatePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoUpdatePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoUpdatePayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VideoUpdatePayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *VideoUpdatePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoUpdatePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoUpdatePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoUpdatePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublic

`func (o *VideoUpdatePayload) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *VideoUpdatePayload) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *VideoUpdatePayload) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *VideoUpdatePayload) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPanoramic

`func (o *VideoUpdatePayload) GetPanoramic() bool`

GetPanoramic returns the Panoramic field if non-nil, zero value otherwise.

### GetPanoramicOk

`func (o *VideoUpdatePayload) GetPanoramicOk() (*bool, bool)`

GetPanoramicOk returns a tuple with the Panoramic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramic

`func (o *VideoUpdatePayload) SetPanoramic(v bool)`

SetPanoramic sets Panoramic field to given value.

### HasPanoramic

`func (o *VideoUpdatePayload) HasPanoramic() bool`

HasPanoramic returns a boolean if a field has been set.

### GetMp4Support

`func (o *VideoUpdatePayload) GetMp4Support() bool`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *VideoUpdatePayload) GetMp4SupportOk() (*bool, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *VideoUpdatePayload) SetMp4Support(v bool)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *VideoUpdatePayload) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.

### GetTags

`func (o *VideoUpdatePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VideoUpdatePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VideoUpdatePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VideoUpdatePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *VideoUpdatePayload) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VideoUpdatePayload) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VideoUpdatePayload) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VideoUpdatePayload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLanguage

`func (o *VideoUpdatePayload) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VideoUpdatePayload) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VideoUpdatePayload) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VideoUpdatePayload) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *VideoUpdatePayload) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *VideoUpdatePayload) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTranscript

`func (o *VideoUpdatePayload) GetTranscript() bool`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *VideoUpdatePayload) GetTranscriptOk() (*bool, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *VideoUpdatePayload) SetTranscript(v bool)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *VideoUpdatePayload) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetTranscriptSummary

`func (o *VideoUpdatePayload) GetTranscriptSummary() bool`

GetTranscriptSummary returns the TranscriptSummary field if non-nil, zero value otherwise.

### GetTranscriptSummaryOk

`func (o *VideoUpdatePayload) GetTranscriptSummaryOk() (*bool, bool)`

GetTranscriptSummaryOk returns a tuple with the TranscriptSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSummary

`func (o *VideoUpdatePayload) SetTranscriptSummary(v bool)`

SetTranscriptSummary sets TranscriptSummary field to given value.

### HasTranscriptSummary

`func (o *VideoUpdatePayload) HasTranscriptSummary() bool`

HasTranscriptSummary returns a boolean if a field has been set.

### GetTranscriptSummaryAttributes

`func (o *VideoUpdatePayload) GetTranscriptSummaryAttributes() []string`

GetTranscriptSummaryAttributes returns the TranscriptSummaryAttributes field if non-nil, zero value otherwise.

### GetTranscriptSummaryAttributesOk

`func (o *VideoUpdatePayload) GetTranscriptSummaryAttributesOk() (*[]string, bool)`

GetTranscriptSummaryAttributesOk returns a tuple with the TranscriptSummaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSummaryAttributes

`func (o *VideoUpdatePayload) SetTranscriptSummaryAttributes(v []string)`

SetTranscriptSummaryAttributes sets TranscriptSummaryAttributes field to given value.

### HasTranscriptSummaryAttributes

`func (o *VideoUpdatePayload) HasTranscriptSummaryAttributes() bool`

HasTranscriptSummaryAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


