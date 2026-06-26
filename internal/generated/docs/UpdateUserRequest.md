# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | New first (given) name. Send null to clear; omit to leave unchanged. | [optional] 
**LastName** | Pointer to **NullableString** | New last (family) name. Send null to clear; omit to leave unchanged. | [optional] 
**Locale** | Pointer to **NullableString** | New preferred locale. Send null to clear; omit to leave unchanged. | [optional] 
**UnsafeMetadata** | Pointer to **map[string]interface{}** | Deep-merges into the user&#39;s unsafe metadata (a key set to null removes it); omit to leave unchanged. Merged result max 512 bytes. | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UpdateUserRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UpdateUserRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UpdateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UpdateUserRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UpdateUserRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLocale

`func (o *UpdateUserRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateUserRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateUserRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateUserRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UpdateUserRequest) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UpdateUserRequest) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetUnsafeMetadata

`func (o *UpdateUserRequest) GetUnsafeMetadata() map[string]interface{}`

GetUnsafeMetadata returns the UnsafeMetadata field if non-nil, zero value otherwise.

### GetUnsafeMetadataOk

`func (o *UpdateUserRequest) GetUnsafeMetadataOk() (*map[string]interface{}, bool)`

GetUnsafeMetadataOk returns a tuple with the UnsafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeMetadata

`func (o *UpdateUserRequest) SetUnsafeMetadata(v map[string]interface{})`

SetUnsafeMetadata sets UnsafeMetadata field to given value.

### HasUnsafeMetadata

`func (o *UpdateUserRequest) HasUnsafeMetadata() bool`

HasUnsafeMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


