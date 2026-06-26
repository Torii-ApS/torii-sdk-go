# UserSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this session. | 
**UserId** | **string** | Identifier of the end-user this session belongs to. | 
**EnvironmentId** | **string** | Identifier of the environment this session belongs to. | 
**UserAgent** | Pointer to **NullableString** | Raw User-Agent string captured when the session was created. | [optional] 
**IpAddress** | Pointer to **NullableString** | IP address captured when the session was created. | [optional] 
**CreatedAt** | **time.Time** | When this session was created (ISO-8601 UTC). | 
**ExpiresAt** | **time.Time** | When this session expires (ISO-8601 UTC). | 
**LastUsedAt** | **time.Time** | When this session was last seen by the API (ISO-8601 UTC). | 
**ActiveOrganizationId** | Pointer to **NullableString** | Active organization pinned to this session (&#x60;org_id&#x60; claim on re-mint). | [optional] 
**ImpersonatedBy** | Pointer to **NullableString** | Platform user behind this session when it was established via impersonation; null for normal sign-ins. | [optional] 

## Methods

### NewUserSessionResponse

`func NewUserSessionResponse(id string, userId string, environmentId string, createdAt time.Time, expiresAt time.Time, lastUsedAt time.Time, ) *UserSessionResponse`

NewUserSessionResponse instantiates a new UserSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionResponseWithDefaults

`func NewUserSessionResponseWithDefaults() *UserSessionResponse`

NewUserSessionResponseWithDefaults instantiates a new UserSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserSessionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSessionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSessionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEnvironmentId

`func (o *UserSessionResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserSessionResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserSessionResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetUserAgent

`func (o *UserSessionResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *UserSessionResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *UserSessionResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *UserSessionResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *UserSessionResponse) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *UserSessionResponse) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetIpAddress

`func (o *UserSessionResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserSessionResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserSessionResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserSessionResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *UserSessionResponse) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *UserSessionResponse) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetCreatedAt

`func (o *UserSessionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSessionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSessionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *UserSessionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserSessionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserSessionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetLastUsedAt

`func (o *UserSessionResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *UserSessionResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *UserSessionResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### GetActiveOrganizationId

`func (o *UserSessionResponse) GetActiveOrganizationId() string`

GetActiveOrganizationId returns the ActiveOrganizationId field if non-nil, zero value otherwise.

### GetActiveOrganizationIdOk

`func (o *UserSessionResponse) GetActiveOrganizationIdOk() (*string, bool)`

GetActiveOrganizationIdOk returns a tuple with the ActiveOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOrganizationId

`func (o *UserSessionResponse) SetActiveOrganizationId(v string)`

SetActiveOrganizationId sets ActiveOrganizationId field to given value.

### HasActiveOrganizationId

`func (o *UserSessionResponse) HasActiveOrganizationId() bool`

HasActiveOrganizationId returns a boolean if a field has been set.

### SetActiveOrganizationIdNil

`func (o *UserSessionResponse) SetActiveOrganizationIdNil(b bool)`

 SetActiveOrganizationIdNil sets the value for ActiveOrganizationId to be an explicit nil

### UnsetActiveOrganizationId
`func (o *UserSessionResponse) UnsetActiveOrganizationId()`

UnsetActiveOrganizationId ensures that no value is present for ActiveOrganizationId, not even an explicit nil
### GetImpersonatedBy

`func (o *UserSessionResponse) GetImpersonatedBy() string`

GetImpersonatedBy returns the ImpersonatedBy field if non-nil, zero value otherwise.

### GetImpersonatedByOk

`func (o *UserSessionResponse) GetImpersonatedByOk() (*string, bool)`

GetImpersonatedByOk returns a tuple with the ImpersonatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatedBy

`func (o *UserSessionResponse) SetImpersonatedBy(v string)`

SetImpersonatedBy sets ImpersonatedBy field to given value.

### HasImpersonatedBy

`func (o *UserSessionResponse) HasImpersonatedBy() bool`

HasImpersonatedBy returns a boolean if a field has been set.

### SetImpersonatedByNil

`func (o *UserSessionResponse) SetImpersonatedByNil(b bool)`

 SetImpersonatedByNil sets the value for ImpersonatedBy to be an explicit nil

### UnsetImpersonatedBy
`func (o *UserSessionResponse) UnsetImpersonatedBy()`

UnsetImpersonatedBy ensures that no value is present for ImpersonatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


