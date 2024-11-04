# AnalyticsMetricsOverTimeResponseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | Returns the metric and relevant parameters you selected. | [optional] 
**Interval** | Pointer to **string** | Returns the interval you selected. | [optional] 
**Timeframe** | Pointer to [**AnalyticsAggregatedMetricsResponseContextTimeframe**](AnalyticsAggregatedMetricsResponseContextTimeframe.md) |  | [optional] 

## Methods

### NewAnalyticsMetricsOverTimeResponseContext

`func NewAnalyticsMetricsOverTimeResponseContext() *AnalyticsMetricsOverTimeResponseContext`

NewAnalyticsMetricsOverTimeResponseContext instantiates a new AnalyticsMetricsOverTimeResponseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsOverTimeResponseContextWithDefaults

`func NewAnalyticsMetricsOverTimeResponseContextWithDefaults() *AnalyticsMetricsOverTimeResponseContext`

NewAnalyticsMetricsOverTimeResponseContextWithDefaults instantiates a new AnalyticsMetricsOverTimeResponseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AnalyticsMetricsOverTimeResponseContext) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AnalyticsMetricsOverTimeResponseContext) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AnalyticsMetricsOverTimeResponseContext) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AnalyticsMetricsOverTimeResponseContext) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetInterval

`func (o *AnalyticsMetricsOverTimeResponseContext) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AnalyticsMetricsOverTimeResponseContext) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AnalyticsMetricsOverTimeResponseContext) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AnalyticsMetricsOverTimeResponseContext) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeframe

`func (o *AnalyticsMetricsOverTimeResponseContext) GetTimeframe() AnalyticsAggregatedMetricsResponseContextTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *AnalyticsMetricsOverTimeResponseContext) GetTimeframeOk() (*AnalyticsAggregatedMetricsResponseContextTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *AnalyticsMetricsOverTimeResponseContext) SetTimeframe(v AnalyticsAggregatedMetricsResponseContextTimeframe)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *AnalyticsMetricsOverTimeResponseContext) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


