# CreateFleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FleetType** | [**FleetType**](FleetType.md) |  | 
**Expression** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFleet

`func NewCreateFleet(name string, fleetType FleetType, ) *CreateFleet`

NewCreateFleet instantiates a new CreateFleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFleetWithDefaults

`func NewCreateFleetWithDefaults() *CreateFleet`

NewCreateFleetWithDefaults instantiates a new CreateFleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFleet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFleet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFleet) SetName(v string)`

SetName sets Name field to given value.


### GetFleetType

`func (o *CreateFleet) GetFleetType() FleetType`

GetFleetType returns the FleetType field if non-nil, zero value otherwise.

### GetFleetTypeOk

`func (o *CreateFleet) GetFleetTypeOk() (*FleetType, bool)`

GetFleetTypeOk returns a tuple with the FleetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetType

`func (o *CreateFleet) SetFleetType(v FleetType)`

SetFleetType sets FleetType field to given value.


### GetExpression

`func (o *CreateFleet) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CreateFleet) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CreateFleet) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *CreateFleet) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


