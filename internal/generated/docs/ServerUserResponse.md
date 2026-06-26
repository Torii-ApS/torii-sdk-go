# ServerUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this user. | 
**EnvironmentId** | **string** | Identifier of the environment this user belongs to. | 
**Name** | Pointer to **NullableString** | Full name on the profile, if any. | [optional] 
**FirstName** | Pointer to **NullableString** | First (given) name on the profile, if any. | [optional] 
**LastName** | Pointer to **NullableString** | Last (family) name on the profile, if any. | [optional] 
**Locale** | Pointer to **NullableString** | Preferred locale for emails and UI messages. | [optional] 
**Status** | **string** | Lifecycle status of the user (e.g. active, banned). | 
**CreatedAt** | **time.Time** | When this user was created (ISO-8601 UTC). | 
**UpdatedAt** | **time.Time** | When this user was last modified (ISO-8601 UTC). | 
**Email** | Pointer to **NullableString** | Primary email on the profile, if any. | [optional] 
**EmailVerifiedAt** | Pointer to **NullableTime** | When this user&#39;s primary email was verified, if it has been verified. | [optional] 
**DeletedAt** | Pointer to **NullableTime** | When this user was deleted, if soft-deleted. Null for active users. | [optional] 
**PublicMetadata** | **map[string]interface{}** | Public metadata: readable by the SDK, writable only server-side. | 
**PrivateMetadata** | **map[string]interface{}** | Private metadata: server-only. Never exposed to the SDK or in a JWT. | 
**UnsafeMetadata** | **map[string]interface{}** | Unsafe metadata: readable and writable by the end-user via the SDK. | 

## Methods

### NewServerUserResponse

`func NewServerUserResponse(id string, environmentId string, status string, createdAt time.Time, updatedAt time.Time, publicMetadata map[string]interface{}, privateMetadata map[string]interface{}, unsafeMetadata map[string]interface{}, ) *ServerUserResponse`

NewServerUserResponse instantiates a new ServerUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerUserResponseWithDefaults

`func NewServerUserResponseWithDefaults() *ServerUserResponse`

NewServerUserResponseWithDefaults instantiates a new ServerUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironmentId

`func (o *ServerUserResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ServerUserResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ServerUserResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *ServerUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerUserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerUserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServerUserResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServerUserResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFirstName

`func (o *ServerUserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServerUserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServerUserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ServerUserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ServerUserResponse) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ServerUserResponse) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ServerUserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServerUserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServerUserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ServerUserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ServerUserResponse) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ServerUserResponse) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLocale

`func (o *ServerUserResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ServerUserResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ServerUserResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ServerUserResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *ServerUserResponse) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *ServerUserResponse) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetStatus

`func (o *ServerUserResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerUserResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerUserResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ServerUserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerUserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerUserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ServerUserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerUserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerUserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmail

`func (o *ServerUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServerUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServerUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServerUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ServerUserResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ServerUserResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerifiedAt

`func (o *ServerUserResponse) GetEmailVerifiedAt() time.Time`

GetEmailVerifiedAt returns the EmailVerifiedAt field if non-nil, zero value otherwise.

### GetEmailVerifiedAtOk

`func (o *ServerUserResponse) GetEmailVerifiedAtOk() (*time.Time, bool)`

GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerifiedAt

`func (o *ServerUserResponse) SetEmailVerifiedAt(v time.Time)`

SetEmailVerifiedAt sets EmailVerifiedAt field to given value.

### HasEmailVerifiedAt

`func (o *ServerUserResponse) HasEmailVerifiedAt() bool`

HasEmailVerifiedAt returns a boolean if a field has been set.

### SetEmailVerifiedAtNil

`func (o *ServerUserResponse) SetEmailVerifiedAtNil(b bool)`

 SetEmailVerifiedAtNil sets the value for EmailVerifiedAt to be an explicit nil

### UnsetEmailVerifiedAt
`func (o *ServerUserResponse) UnsetEmailVerifiedAt()`

UnsetEmailVerifiedAt ensures that no value is present for EmailVerifiedAt, not even an explicit nil
### GetDeletedAt

`func (o *ServerUserResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ServerUserResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ServerUserResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ServerUserResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ServerUserResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ServerUserResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetPublicMetadata

`func (o *ServerUserResponse) GetPublicMetadata() map[string]interface{}`

GetPublicMetadata returns the PublicMetadata field if non-nil, zero value otherwise.

### GetPublicMetadataOk

`func (o *ServerUserResponse) GetPublicMetadataOk() (*map[string]interface{}, bool)`

GetPublicMetadataOk returns a tuple with the PublicMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetadata

`func (o *ServerUserResponse) SetPublicMetadata(v map[string]interface{})`

SetPublicMetadata sets PublicMetadata field to given value.


### GetPrivateMetadata

`func (o *ServerUserResponse) GetPrivateMetadata() map[string]interface{}`

GetPrivateMetadata returns the PrivateMetadata field if non-nil, zero value otherwise.

### GetPrivateMetadataOk

`func (o *ServerUserResponse) GetPrivateMetadataOk() (*map[string]interface{}, bool)`

GetPrivateMetadataOk returns a tuple with the PrivateMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMetadata

`func (o *ServerUserResponse) SetPrivateMetadata(v map[string]interface{})`

SetPrivateMetadata sets PrivateMetadata field to given value.


### GetUnsafeMetadata

`func (o *ServerUserResponse) GetUnsafeMetadata() map[string]interface{}`

GetUnsafeMetadata returns the UnsafeMetadata field if non-nil, zero value otherwise.

### GetUnsafeMetadataOk

`func (o *ServerUserResponse) GetUnsafeMetadataOk() (*map[string]interface{}, bool)`

GetUnsafeMetadataOk returns a tuple with the UnsafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeMetadata

`func (o *ServerUserResponse) SetUnsafeMetadata(v map[string]interface{})`

SetUnsafeMetadata sets UnsafeMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


