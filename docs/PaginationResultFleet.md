# PaginationResultFleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]Fleet**](Fleet.md) |  | [optional] 
**Total** | **int64** |  | 
**Offset** | **int64** |  | 
**Limit** | **int64** |  | 

## Methods

### NewPaginationResultFleet

`func NewPaginationResultFleet(total int64, offset int64, limit int64, ) *PaginationResultFleet`

NewPaginationResultFleet instantiates a new PaginationResultFleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResultFleetWithDefaults

`func NewPaginationResultFleetWithDefaults() *PaginationResultFleet`

NewPaginationResultFleetWithDefaults instantiates a new PaginationResultFleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *PaginationResultFleet) GetValues() []Fleet`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PaginationResultFleet) GetValuesOk() (*[]Fleet, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PaginationResultFleet) SetValues(v []Fleet)`

SetValues sets Values field to given value.

### HasValues

`func (o *PaginationResultFleet) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetTotal

`func (o *PaginationResultFleet) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationResultFleet) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationResultFleet) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *PaginationResultFleet) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationResultFleet) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationResultFleet) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *PaginationResultFleet) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationResultFleet) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationResultFleet) SetLimit(v int64)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


