# SeriesMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**TimeAggregation**](TimeAggregation.md) |  | 
**DeviceIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSeriesMeta

`func NewSeriesMeta(aggregation TimeAggregation, ) *SeriesMeta`

NewSeriesMeta instantiates a new SeriesMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesMetaWithDefaults

`func NewSeriesMetaWithDefaults() *SeriesMeta`

NewSeriesMetaWithDefaults instantiates a new SeriesMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *SeriesMeta) GetAggregation() TimeAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *SeriesMeta) GetAggregationOk() (*TimeAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *SeriesMeta) SetAggregation(v TimeAggregation)`

SetAggregation sets Aggregation field to given value.


### GetDeviceIds

`func (o *SeriesMeta) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SeriesMeta) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *SeriesMeta) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *SeriesMeta) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


