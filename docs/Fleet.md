# Fleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**FleetType** | [**FleetType**](FleetType.md) |  | 
**Expression** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFleet

`func NewFleet(id string, name string, createdAt time.Time, fleetType FleetType, ) *Fleet`

NewFleet instantiates a new Fleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetWithDefaults

`func NewFleetWithDefaults() *Fleet`

NewFleetWithDefaults instantiates a new Fleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fleet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fleet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fleet) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Fleet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fleet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fleet) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Fleet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Fleet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Fleet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFleetType

`func (o *Fleet) GetFleetType() FleetType`

GetFleetType returns the FleetType field if non-nil, zero value otherwise.

### GetFleetTypeOk

`func (o *Fleet) GetFleetTypeOk() (*FleetType, bool)`

GetFleetTypeOk returns a tuple with the FleetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetType

`func (o *Fleet) SetFleetType(v FleetType)`

SetFleetType sets FleetType field to given value.


### GetExpression

`func (o *Fleet) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Fleet) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Fleet) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Fleet) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *Fleet) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *Fleet) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


