# CursorPageResponseServerUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServerUserResponse**](ServerUserResponse.md) | Items in this page, in stable order. | 
**NextCursor** | Pointer to **NullableString** | Cursor to pass to fetch the next page. Null when this is the last page. | [optional] 
**HasMore** | **bool** | True if more pages are available (equivalent to &#x60;nextCursor !&#x3D; null&#x60;). | 

## Methods

### NewCursorPageResponseServerUserResponse

`func NewCursorPageResponseServerUserResponse(items []ServerUserResponse, hasMore bool, ) *CursorPageResponseServerUserResponse`

NewCursorPageResponseServerUserResponse instantiates a new CursorPageResponseServerUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageResponseServerUserResponseWithDefaults

`func NewCursorPageResponseServerUserResponseWithDefaults() *CursorPageResponseServerUserResponse`

NewCursorPageResponseServerUserResponseWithDefaults instantiates a new CursorPageResponseServerUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CursorPageResponseServerUserResponse) GetItems() []ServerUserResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPageResponseServerUserResponse) GetItemsOk() (*[]ServerUserResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPageResponseServerUserResponse) SetItems(v []ServerUserResponse)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *CursorPageResponseServerUserResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CursorPageResponseServerUserResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CursorPageResponseServerUserResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CursorPageResponseServerUserResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *CursorPageResponseServerUserResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CursorPageResponseServerUserResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *CursorPageResponseServerUserResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CursorPageResponseServerUserResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CursorPageResponseServerUserResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


