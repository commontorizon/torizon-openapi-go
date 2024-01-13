# MetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]Series**](Series.md) |  | [optional] 

## Methods

### NewMetricsResponse

`func NewMetricsResponse() *MetricsResponse`

NewMetricsResponse instantiates a new MetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResponseWithDefaults

`func NewMetricsResponseWithDefaults() *MetricsResponse`

NewMetricsResponseWithDefaults instantiates a new MetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *MetricsResponse) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MetricsResponse) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *MetricsResponse) SetSeries(v []Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *MetricsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


