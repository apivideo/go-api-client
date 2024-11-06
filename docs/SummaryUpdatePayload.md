# SummaryUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstract** | Pointer to **string** | A short outline of the contents of the video. | [optional] 
**Takeaways** | Pointer to **[]string** | A list of 3 key points from the video, in chronological order. | [optional] 

## Methods

### NewSummaryUpdatePayload

`func NewSummaryUpdatePayload() *SummaryUpdatePayload`

NewSummaryUpdatePayload instantiates a new SummaryUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryUpdatePayloadWithDefaults

`func NewSummaryUpdatePayloadWithDefaults() *SummaryUpdatePayload`

NewSummaryUpdatePayloadWithDefaults instantiates a new SummaryUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstract

`func (o *SummaryUpdatePayload) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SummaryUpdatePayload) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SummaryUpdatePayload) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SummaryUpdatePayload) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### GetTakeaways

`func (o *SummaryUpdatePayload) GetTakeaways() []string`

GetTakeaways returns the Takeaways field if non-nil, zero value otherwise.

### GetTakeawaysOk

`func (o *SummaryUpdatePayload) GetTakeawaysOk() (*[]string, bool)`

GetTakeawaysOk returns a tuple with the Takeaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeaways

`func (o *SummaryUpdatePayload) SetTakeaways(v []string)`

SetTakeaways sets Takeaways field to given value.

### HasTakeaways

`func (o *SummaryUpdatePayload) HasTakeaways() bool`

HasTakeaways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


