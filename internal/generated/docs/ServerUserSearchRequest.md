# ServerUserSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Filter by name (case-insensitive substring match). Send null to require users with no name. | [optional] 
**Email** | Pointer to **NullableString** | Filter by primary email (case-insensitive substring match). Send null to require users with no email. | [optional] 
**Statuses** | Pointer to **[]string** | Filter by user status. Returns users matching any of the supplied statuses. | [optional] 
**CreatedAfter** | Pointer to **NullableTime** | Only return users created at or after this instant (ISO-8601 UTC). | [optional] 
**CreatedBefore** | Pointer to **NullableTime** | Only return users created at or before this instant (ISO-8601 UTC). | [optional] 

## Methods

### NewServerUserSearchRequest

`func NewServerUserSearchRequest() *ServerUserSearchRequest`

NewServerUserSearchRequest instantiates a new ServerUserSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerUserSearchRequestWithDefaults

`func NewServerUserSearchRequestWithDefaults() *ServerUserSearchRequest`

NewServerUserSearchRequestWithDefaults instantiates a new ServerUserSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerUserSearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerUserSearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerUserSearchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerUserSearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServerUserSearchRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServerUserSearchRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *ServerUserSearchRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServerUserSearchRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServerUserSearchRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServerUserSearchRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ServerUserSearchRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ServerUserSearchRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatuses

`func (o *ServerUserSearchRequest) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ServerUserSearchRequest) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ServerUserSearchRequest) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ServerUserSearchRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetCreatedAfter

`func (o *ServerUserSearchRequest) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *ServerUserSearchRequest) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *ServerUserSearchRequest) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *ServerUserSearchRequest) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### SetCreatedAfterNil

`func (o *ServerUserSearchRequest) SetCreatedAfterNil(b bool)`

 SetCreatedAfterNil sets the value for CreatedAfter to be an explicit nil

### UnsetCreatedAfter
`func (o *ServerUserSearchRequest) UnsetCreatedAfter()`

UnsetCreatedAfter ensures that no value is present for CreatedAfter, not even an explicit nil
### GetCreatedBefore

`func (o *ServerUserSearchRequest) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *ServerUserSearchRequest) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *ServerUserSearchRequest) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *ServerUserSearchRequest) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### SetCreatedBeforeNil

`func (o *ServerUserSearchRequest) SetCreatedBeforeNil(b bool)`

 SetCreatedBeforeNil sets the value for CreatedBefore to be an explicit nil

### UnsetCreatedBefore
`func (o *ServerUserSearchRequest) UnsetCreatedBefore()`

UnsetCreatedBefore ensures that no value is present for CreatedBefore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


