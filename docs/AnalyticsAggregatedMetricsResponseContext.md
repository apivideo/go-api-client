# AnalyticsAggregatedMetricsResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | Returns the metric and relevant parameters you selected. | [optional] 
**Aggregation** | Pointer to **string** | Returns the aggregation you selected. | [optional] 
**Timeframe** | Pointer to [**AnalyticsAggregatedMetricsResponseContextTimeframe**](AnalyticsAggregatedMetricsResponseContextTimeframe.md) |  | [optional] 

## Methods

### NewAnalyticsAggregatedMetricsResponseContext

`func NewAnalyticsAggregatedMetricsResponseContext() *AnalyticsAggregatedMetricsResponseContext`

NewAnalyticsAggregatedMetricsResponseContext instantiates a new AnalyticsAggregatedMetricsResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsAggregatedMetricsResponseContextWithDefaults

`func NewAnalyticsAggregatedMetricsResponseContextWithDefaults() *AnalyticsAggregatedMetricsResponseContext`

NewAnalyticsAggregatedMetricsResponseContextWithDefaults instantiates a new AnalyticsAggregatedMetricsResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AnalyticsAggregatedMetricsResponseContext) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AnalyticsAggregatedMetricsResponseContext) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AnalyticsAggregatedMetricsResponseContext) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AnalyticsAggregatedMetricsResponseContext) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetAggregation

`func (o *AnalyticsAggregatedMetricsResponseContext) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *AnalyticsAggregatedMetricsResponseContext) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *AnalyticsAggregatedMetricsResponseContext) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *AnalyticsAggregatedMetricsResponseContext) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetTimeframe

`func (o *AnalyticsAggregatedMetricsResponseContext) GetTimeframe() AnalyticsAggregatedMetricsResponseContextTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *AnalyticsAggregatedMetricsResponseContext) GetTimeframeOk() (*AnalyticsAggregatedMetricsResponseContextTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *AnalyticsAggregatedMetricsResponseContext) SetTimeframe(v AnalyticsAggregatedMetricsResponseContextTimeframe)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *AnalyticsAggregatedMetricsResponseContext) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


