# VideoCreationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of your new video. | 
**Description** | Pointer to **string** | A brief description of your video. | [optional] 
**Source** | Pointer to **string** | You can either add a video already on the web, by entering the URL of the video, or you can also enter the &#x60;videoId&#x60; of one of the videos you already have on your api.video acccount, and this will generate a copy of your video. Creating a copy of a video can be especially useful if you want to keep your original video and trim or apply a watermark onto the copy you would create. | [optional] 
**Public** | Pointer to **bool** | Default: True. If set to &#x60;false&#x60; the video will become private. More information on private videos can be found [here](https://docs.api.video/delivery/video-privacy-access-management) | [optional] [default to true]
**Panoramic** | Pointer to **bool** | Indicates if your video is a 360/immersive video. | [optional] [default to false]
**Mp4Support** | Pointer to **bool** | Enables mp4 version in addition to streamed version. | [optional] [default to true]
**PlayerId** | Pointer to **string** | The unique identification number for your video player. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags you want to use to describe your video. | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | A list of key value pairs that you use to provide metadata for your video. | [optional] 
**Clip** | Pointer to [**VideoClip**](VideoClip.md) |  | [optional] 
**Watermark** | Pointer to [**VideoWatermark**](VideoWatermark.md) |  | [optional] 
**Language** | Pointer to **NullableString** | Use this parameter to set the language of the video. When this parameter is set, the API creates a transcript of the video using the language you specify. You must use the [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) format.  &#x60;language&#x60; is a permanent attribute of the video. You can update it to another language using the [&#x60;PATCH /videos/{videoId}&#x60;](https://docs.api.video/reference/api/Videos#update-a-video-object) operation. This triggers the API to generate a new transcript using a different language. | [optional] 
**Transcript** | Pointer to **bool** | Use this parameter to enable transcription.   - When &#x60;true&#x60;, the API generates a transcript for the video. - The default value is &#x60;false&#x60;. - If you define a video language using the &#x60;language&#x60; parameter, the API uses that language to transcribe the video. If you do not define a language, the API detects it based on the video.  - When the API generates a transcript, it will be available as a caption for the video. | [optional] 
**TranscriptSummary** | Pointer to **bool** | Use this parameter to enable summarization. We recommend using this parameter together with &#x60;transcript: true&#x60;.  - When &#x60;true&#x60;, the API generates a summary for the video, based on the transcription. - The default value is &#x60;false&#x60;. - If you define a video language using the &#x60;language&#x60; parameter, the API uses that language to summarize the video. If you do not define a language, the API detects it based on the video. | [optional] 
**TranscriptSummaryAttributes** | Pointer to **[]string** | Use this parameter to define the elements of a summary that you want to generate. If you do not define this parameter, the API generates a full summary with all attributes. The possible values are &#x60;abstract&#x60; and &#x60;takeaways&#x60;. | [optional] 

## Methods

### NewVideoCreationPayload

`func NewVideoCreationPayload(title string, ) *VideoCreationPayload`

NewVideoCreationPayload instantiates a new VideoCreationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCreationPayloadWithDefaults

`func NewVideoCreationPayloadWithDefaults() *VideoCreationPayload`

NewVideoCreationPayloadWithDefaults instantiates a new VideoCreationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *VideoCreationPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoCreationPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoCreationPayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *VideoCreationPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoCreationPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoCreationPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoCreationPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *VideoCreationPayload) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VideoCreationPayload) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VideoCreationPayload) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VideoCreationPayload) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPublic

`func (o *VideoCreationPayload) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *VideoCreationPayload) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *VideoCreationPayload) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *VideoCreationPayload) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPanoramic

`func (o *VideoCreationPayload) GetPanoramic() bool`

GetPanoramic returns the Panoramic field if non-nil, zero value otherwise.

### GetPanoramicOk

`func (o *VideoCreationPayload) GetPanoramicOk() (*bool, bool)`

GetPanoramicOk returns a tuple with the Panoramic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramic

`func (o *VideoCreationPayload) SetPanoramic(v bool)`

SetPanoramic sets Panoramic field to given value.

### HasPanoramic

`func (o *VideoCreationPayload) HasPanoramic() bool`

HasPanoramic returns a boolean if a field has been set.

### GetMp4Support

`func (o *VideoCreationPayload) GetMp4Support() bool`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *VideoCreationPayload) GetMp4SupportOk() (*bool, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *VideoCreationPayload) SetMp4Support(v bool)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *VideoCreationPayload) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.

### GetPlayerId

`func (o *VideoCreationPayload) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *VideoCreationPayload) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *VideoCreationPayload) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *VideoCreationPayload) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetTags

`func (o *VideoCreationPayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VideoCreationPayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VideoCreationPayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VideoCreationPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *VideoCreationPayload) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VideoCreationPayload) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VideoCreationPayload) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VideoCreationPayload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetClip

`func (o *VideoCreationPayload) GetClip() VideoClip`

GetClip returns the Clip field if non-nil, zero value otherwise.

### GetClipOk

`func (o *VideoCreationPayload) GetClipOk() (*VideoClip, bool)`

GetClipOk returns a tuple with the Clip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClip

`func (o *VideoCreationPayload) SetClip(v VideoClip)`

SetClip sets Clip field to given value.

### HasClip

`func (o *VideoCreationPayload) HasClip() bool`

HasClip returns a boolean if a field has been set.

### GetWatermark

`func (o *VideoCreationPayload) GetWatermark() VideoWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *VideoCreationPayload) GetWatermarkOk() (*VideoWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *VideoCreationPayload) SetWatermark(v VideoWatermark)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *VideoCreationPayload) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLanguage

`func (o *VideoCreationPayload) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VideoCreationPayload) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VideoCreationPayload) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VideoCreationPayload) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *VideoCreationPayload) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *VideoCreationPayload) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTranscript

`func (o *VideoCreationPayload) GetTranscript() bool`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *VideoCreationPayload) GetTranscriptOk() (*bool, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *VideoCreationPayload) SetTranscript(v bool)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *VideoCreationPayload) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetTranscriptSummary

`func (o *VideoCreationPayload) GetTranscriptSummary() bool`

GetTranscriptSummary returns the TranscriptSummary field if non-nil, zero value otherwise.

### GetTranscriptSummaryOk

`func (o *VideoCreationPayload) GetTranscriptSummaryOk() (*bool, bool)`

GetTranscriptSummaryOk returns a tuple with the TranscriptSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSummary

`func (o *VideoCreationPayload) SetTranscriptSummary(v bool)`

SetTranscriptSummary sets TranscriptSummary field to given value.

### HasTranscriptSummary

`func (o *VideoCreationPayload) HasTranscriptSummary() bool`

HasTranscriptSummary returns a boolean if a field has been set.

### GetTranscriptSummaryAttributes

`func (o *VideoCreationPayload) GetTranscriptSummaryAttributes() []string`

GetTranscriptSummaryAttributes returns the TranscriptSummaryAttributes field if non-nil, zero value otherwise.

### GetTranscriptSummaryAttributesOk

`func (o *VideoCreationPayload) GetTranscriptSummaryAttributesOk() (*[]string, bool)`

GetTranscriptSummaryAttributesOk returns a tuple with the TranscriptSummaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSummaryAttributes

`func (o *VideoCreationPayload) SetTranscriptSummaryAttributes(v []string)`

SetTranscriptSummaryAttributes sets TranscriptSummaryAttributes field to given value.

### HasTranscriptSummaryAttributes

`func (o *VideoCreationPayload) HasTranscriptSummaryAttributes() bool`

HasTranscriptSummaryAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


