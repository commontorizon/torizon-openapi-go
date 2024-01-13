# TimeAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Method** | [**TimeAggregationMethod**](TimeAggregationMethod.md) |  | 

## Methods

### NewTimeAggregation

`func NewTimeAggregation(bucket string, method TimeAggregationMethod, ) *TimeAggregation`

NewTimeAggregation instantiates a new TimeAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeAggregationWithDefaults

`func NewTimeAggregationWithDefaults() *TimeAggregation`

NewTimeAggregationWithDefaults instantiates a new TimeAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *TimeAggregation) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *TimeAggregation) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *TimeAggregation) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetMethod

`func (o *TimeAggregation) GetMethod() TimeAggregationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TimeAggregation) GetMethodOk() (*TimeAggregationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TimeAggregation) SetMethod(v TimeAggregationMethod)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


