# AnalyticsMetricsOverTimeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**AnalyticsMetricsOverTimeResponseContext**](AnalyticsMetricsOverTimeResponseContext.md) |  | 
**Data** | [**[]AnalyticsMetricsOverTimeResponseData**](AnalyticsMetricsOverTimeResponseData.md) | Returns an array of metrics and the timestamps. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewAnalyticsMetricsOverTimeResponse

`func NewAnalyticsMetricsOverTimeResponse(context AnalyticsMetricsOverTimeResponseContext, data []AnalyticsMetricsOverTimeResponseData, pagination Pagination, ) *AnalyticsMetricsOverTimeResponse`

NewAnalyticsMetricsOverTimeResponse instantiates a new AnalyticsMetricsOverTimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsOverTimeResponseWithDefaults

`func NewAnalyticsMetricsOverTimeResponseWithDefaults() *AnalyticsMetricsOverTimeResponse`

NewAnalyticsMetricsOverTimeResponseWithDefaults instantiates a new AnalyticsMetricsOverTimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AnalyticsMetricsOverTimeResponse) GetContext() AnalyticsMetricsOverTimeResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AnalyticsMetricsOverTimeResponse) GetContextOk() (*AnalyticsMetricsOverTimeResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AnalyticsMetricsOverTimeResponse) SetContext(v AnalyticsMetricsOverTimeResponseContext)`

SetContext sets Context field to given value.


### GetData

`func (o *AnalyticsMetricsOverTimeResponse) GetData() []AnalyticsMetricsOverTimeResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsMetricsOverTimeResponse) GetDataOk() (*[]AnalyticsMetricsOverTimeResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsMetricsOverTimeResponse) SetData(v []AnalyticsMetricsOverTimeResponseData)`

SetData sets Data field to given value.


### GetPagination

`func (o *AnalyticsMetricsOverTimeResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnalyticsMetricsOverTimeResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnalyticsMetricsOverTimeResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


