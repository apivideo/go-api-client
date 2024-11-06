# SummarySource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstract** | Pointer to **string** | A short outline of the contents of the video. The length of an &#x60;abstract&#x60; depends on the amount of content in a video that can be transcribed. The API condenses the contents into minimum 20, maximum 300 words. | [optional] 
**Takeaways** | Pointer to **[]string** | A list of 3 key points from the video, in chronological order. | [optional] 

## Methods

### NewSummarySource

`func NewSummarySource() *SummarySource`

NewSummarySource instantiates a new SummarySource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarySourceWithDefaults

`func NewSummarySourceWithDefaults() *SummarySource`

NewSummarySourceWithDefaults instantiates a new SummarySource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstract

`func (o *SummarySource) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SummarySource) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SummarySource) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SummarySource) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### GetTakeaways

`func (o *SummarySource) GetTakeaways() []string`

GetTakeaways returns the Takeaways field if non-nil, zero value otherwise.

### GetTakeawaysOk

`func (o *SummarySource) GetTakeawaysOk() (*[]string, bool)`

GetTakeawaysOk returns a tuple with the Takeaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeaways

`func (o *SummarySource) SetTakeaways(v []string)`

SetTakeaways sets Takeaways field to given value.

### HasTakeaways

`func (o *SummarySource) HasTakeaways() bool`

HasTakeaways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


