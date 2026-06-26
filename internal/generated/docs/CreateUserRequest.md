# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Primary email for the new user. If omitted, the user is created without a sign-in identity. | [optional] 
**Password** | Pointer to **NullableString** | Initial password. Subject to the environment&#39;s password policy. Omit to create a passwordless user (e.g. social-only). | [optional] 
**FirstName** | Pointer to **NullableString** | First (given) name to seed on the profile. | [optional] 
**LastName** | Pointer to **NullableString** | Last (family) name to seed on the profile. | [optional] 
**PublicMetadata** | **map[string]interface{}** | Initial public metadata (SDK-readable, server-written). Max 512 bytes. | 
**PrivateMetadata** | **map[string]interface{}** | Initial private metadata (server-only). Max 4096 bytes. | 
**UnsafeMetadata** | **map[string]interface{}** | Initial unsafe metadata (end-user writable). Max 512 bytes. | 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(publicMetadata map[string]interface{}, privateMetadata map[string]interface{}, unsafeMetadata map[string]interface{}, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateUserRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateUserRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *CreateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateUserRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUserRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFirstName

`func (o *CreateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CreateUserRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CreateUserRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CreateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CreateUserRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CreateUserRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPublicMetadata

`func (o *CreateUserRequest) GetPublicMetadata() map[string]interface{}`

GetPublicMetadata returns the PublicMetadata field if non-nil, zero value otherwise.

### GetPublicMetadataOk

`func (o *CreateUserRequest) GetPublicMetadataOk() (*map[string]interface{}, bool)`

GetPublicMetadataOk returns a tuple with the PublicMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetadata

`func (o *CreateUserRequest) SetPublicMetadata(v map[string]interface{})`

SetPublicMetadata sets PublicMetadata field to given value.


### GetPrivateMetadata

`func (o *CreateUserRequest) GetPrivateMetadata() map[string]interface{}`

GetPrivateMetadata returns the PrivateMetadata field if non-nil, zero value otherwise.

### GetPrivateMetadataOk

`func (o *CreateUserRequest) GetPrivateMetadataOk() (*map[string]interface{}, bool)`

GetPrivateMetadataOk returns a tuple with the PrivateMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMetadata

`func (o *CreateUserRequest) SetPrivateMetadata(v map[string]interface{})`

SetPrivateMetadata sets PrivateMetadata field to given value.


### GetUnsafeMetadata

`func (o *CreateUserRequest) GetUnsafeMetadata() map[string]interface{}`

GetUnsafeMetadata returns the UnsafeMetadata field if non-nil, zero value otherwise.

### GetUnsafeMetadataOk

`func (o *CreateUserRequest) GetUnsafeMetadataOk() (*map[string]interface{}, bool)`

GetUnsafeMetadataOk returns a tuple with the UnsafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeMetadata

`func (o *CreateUserRequest) SetUnsafeMetadata(v map[string]interface{})`

SetUnsafeMetadata sets UnsafeMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


