# UpdateUserMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicMetadata** | Pointer to **map[string]interface{}** | Public metadata bag: SDK-readable, server-written. Max 512 bytes. | [optional] 
**PrivateMetadata** | Pointer to **map[string]interface{}** | Private metadata bag: server-only, never exposed to the SDK or in a JWT. Max 4096 bytes. | [optional] 
**UnsafeMetadata** | Pointer to **map[string]interface{}** | Unsafe metadata bag: readable and writable by the end-user via the SDK. Max 512 bytes. | [optional] 

## Methods

### NewUpdateUserMetadataRequest

`func NewUpdateUserMetadataRequest() *UpdateUserMetadataRequest`

NewUpdateUserMetadataRequest instantiates a new UpdateUserMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserMetadataRequestWithDefaults

`func NewUpdateUserMetadataRequestWithDefaults() *UpdateUserMetadataRequest`

NewUpdateUserMetadataRequestWithDefaults instantiates a new UpdateUserMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicMetadata

`func (o *UpdateUserMetadataRequest) GetPublicMetadata() map[string]interface{}`

GetPublicMetadata returns the PublicMetadata field if non-nil, zero value otherwise.

### GetPublicMetadataOk

`func (o *UpdateUserMetadataRequest) GetPublicMetadataOk() (*map[string]interface{}, bool)`

GetPublicMetadataOk returns a tuple with the PublicMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetadata

`func (o *UpdateUserMetadataRequest) SetPublicMetadata(v map[string]interface{})`

SetPublicMetadata sets PublicMetadata field to given value.

### HasPublicMetadata

`func (o *UpdateUserMetadataRequest) HasPublicMetadata() bool`

HasPublicMetadata returns a boolean if a field has been set.

### GetPrivateMetadata

`func (o *UpdateUserMetadataRequest) GetPrivateMetadata() map[string]interface{}`

GetPrivateMetadata returns the PrivateMetadata field if non-nil, zero value otherwise.

### GetPrivateMetadataOk

`func (o *UpdateUserMetadataRequest) GetPrivateMetadataOk() (*map[string]interface{}, bool)`

GetPrivateMetadataOk returns a tuple with the PrivateMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMetadata

`func (o *UpdateUserMetadataRequest) SetPrivateMetadata(v map[string]interface{})`

SetPrivateMetadata sets PrivateMetadata field to given value.

### HasPrivateMetadata

`func (o *UpdateUserMetadataRequest) HasPrivateMetadata() bool`

HasPrivateMetadata returns a boolean if a field has been set.

### GetUnsafeMetadata

`func (o *UpdateUserMetadataRequest) GetUnsafeMetadata() map[string]interface{}`

GetUnsafeMetadata returns the UnsafeMetadata field if non-nil, zero value otherwise.

### GetUnsafeMetadataOk

`func (o *UpdateUserMetadataRequest) GetUnsafeMetadataOk() (*map[string]interface{}, bool)`

GetUnsafeMetadataOk returns a tuple with the UnsafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeMetadata

`func (o *UpdateUserMetadataRequest) SetUnsafeMetadata(v map[string]interface{})`

SetUnsafeMetadata sets UnsafeMetadata field to given value.

### HasUnsafeMetadata

`func (o *UpdateUserMetadataRequest) HasUnsafeMetadata() bool`

HasUnsafeMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


