# WebhooksCreationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** | An array of webhook events that you want to subscribe to. | 
**Url** | **string** | The the url to which HTTP notifications are sent. It could be any http or https URL. | 

## Methods

### NewWebhooksCreationPayload

`func NewWebhooksCreationPayload(events []string, url string, ) *WebhooksCreationPayload`

NewWebhooksCreationPayload instantiates a new WebhooksCreationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksCreationPayloadWithDefaults

`func NewWebhooksCreationPayloadWithDefaults() *WebhooksCreationPayload`

NewWebhooksCreationPayloadWithDefaults instantiates a new WebhooksCreationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhooksCreationPayload) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhooksCreationPayload) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhooksCreationPayload) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetUrl

`func (o *WebhooksCreationPayload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhooksCreationPayload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhooksCreationPayload) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


