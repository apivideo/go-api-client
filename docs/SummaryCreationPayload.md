# SummaryCreationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | **string** | Create a summary of a video using the video ID. | 
**Origin** | Pointer to **string** | Use this parameter to define how the API generates the summary. The only allowed value is &#x60;auto&#x60;, which means that the API generates a summary automatically.  If you do not set this parameter, **the API will not generate a summary automatically**.  In this case, &#x60;sourceStatus&#x60; will return &#x60;missing&#x60;, and you have to manually add a summary using the &#x60;PATCH /summaries/{summaryId}/source&#x60; endpoint operation. | [optional] 
**Attributes** | Pointer to **[]string** | Use this parameter to define the elements of a summary that you want to generate. If you do not define this parameter, the API generates a full summary with all attributes. The possible values are &#x60;abstract&#x60; and &#x60;takeaways&#x60;. | [optional] 

## Methods

### NewSummaryCreationPayload

`func NewSummaryCreationPayload(videoId string, ) *SummaryCreationPayload`

NewSummaryCreationPayload instantiates a new SummaryCreationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryCreationPayloadWithDefaults

`func NewSummaryCreationPayloadWithDefaults() *SummaryCreationPayload`

NewSummaryCreationPayloadWithDefaults instantiates a new SummaryCreationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *SummaryCreationPayload) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *SummaryCreationPayload) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *SummaryCreationPayload) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.


### GetOrigin

`func (o *SummaryCreationPayload) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SummaryCreationPayload) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SummaryCreationPayload) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SummaryCreationPayload) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetAttributes

`func (o *SummaryCreationPayload) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SummaryCreationPayload) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SummaryCreationPayload) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SummaryCreationPayload) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


