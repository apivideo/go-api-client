# AnalyticsMetricsBreakdownResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | Returns the metric and relevant parameters you selected. | [optional] 
**Breakdown** | Pointer to **string** | Returns the dimension you selected. | [optional] 
**Timeframe** | Pointer to [**AnalyticsAggregatedMetricsResponseContextTimeframe**](AnalyticsAggregatedMetricsResponseContextTimeframe.md) |  | [optional] 

## Methods

### NewAnalyticsMetricsBreakdownResponseContext

`func NewAnalyticsMetricsBreakdownResponseContext() *AnalyticsMetricsBreakdownResponseContext`

NewAnalyticsMetricsBreakdownResponseContext instantiates a new AnalyticsMetricsBreakdownResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsBreakdownResponseContextWithDefaults

`func NewAnalyticsMetricsBreakdownResponseContextWithDefaults() *AnalyticsMetricsBreakdownResponseContext`

NewAnalyticsMetricsBreakdownResponseContextWithDefaults instantiates a new AnalyticsMetricsBreakdownResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AnalyticsMetricsBreakdownResponseContext) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AnalyticsMetricsBreakdownResponseContext) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AnalyticsMetricsBreakdownResponseContext) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AnalyticsMetricsBreakdownResponseContext) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetBreakdown

`func (o *AnalyticsMetricsBreakdownResponseContext) GetBreakdown() string`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *AnalyticsMetricsBreakdownResponseContext) GetBreakdownOk() (*string, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *AnalyticsMetricsBreakdownResponseContext) SetBreakdown(v string)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *AnalyticsMetricsBreakdownResponseContext) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetTimeframe

`func (o *AnalyticsMetricsBreakdownResponseContext) GetTimeframe() AnalyticsAggregatedMetricsResponseContextTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *AnalyticsMetricsBreakdownResponseContext) GetTimeframeOk() (*AnalyticsAggregatedMetricsResponseContextTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *AnalyticsMetricsBreakdownResponseContext) SetTimeframe(v AnalyticsAggregatedMetricsResponseContextTimeframe)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *AnalyticsMetricsBreakdownResponseContext) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


