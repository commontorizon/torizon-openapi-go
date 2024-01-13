# QueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | **string** |  | 
**Targets** | [**map[string]TargetImage**](TargetImage.md) |  | 
**InFlight** | **bool** |  | 

## Methods

### NewQueueResponse

`func NewQueueResponse(correlationId string, targets map[string]TargetImage, inFlight bool, ) *QueueResponse`

NewQueueResponse instantiates a new QueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueResponseWithDefaults

`func NewQueueResponseWithDefaults() *QueueResponse`

NewQueueResponseWithDefaults instantiates a new QueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *QueueResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *QueueResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *QueueResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetTargets

`func (o *QueueResponse) GetTargets() map[string]TargetImage`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *QueueResponse) GetTargetsOk() (*map[string]TargetImage, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *QueueResponse) SetTargets(v map[string]TargetImage)`

SetTargets sets Targets field to given value.


### GetInFlight

`func (o *QueueResponse) GetInFlight() bool`

GetInFlight returns the InFlight field if non-nil, zero value otherwise.

### GetInFlightOk

`func (o *QueueResponse) GetInFlightOk() (*bool, bool)`

GetInFlightOk returns a tuple with the InFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFlight

`func (o *QueueResponse) SetInFlight(v bool)`

SetInFlight sets InFlight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


