# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Meta** | [**SeriesMeta**](SeriesMeta.md) |  | 
**Points** | Pointer to [**[]Tuple2LongOptionDouble**](Tuple2LongOptionDouble.md) |  | [optional] 

## Methods

### NewSeries

`func NewSeries(name string, meta SeriesMeta, ) *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Series) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Series) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Series) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *Series) GetMeta() SeriesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Series) GetMetaOk() (*SeriesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Series) SetMeta(v SeriesMeta)`

SetMeta sets Meta field to given value.


### GetPoints

`func (o *Series) GetPoints() []Tuple2LongOptionDouble`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Series) GetPointsOk() (*[]Tuple2LongOptionDouble, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Series) SetPoints(v []Tuple2LongOptionDouble)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *Series) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


