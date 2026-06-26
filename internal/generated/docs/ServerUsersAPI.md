# \ServerUsersAPI

All URIs are relative to *https://api.torii.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BanUser**](ServerUsersAPI.md#BanUser) | **Post** /api/server/v1/users/{userId}/ban | Ban user
[**CreateUser**](ServerUsersAPI.md#CreateUser) | **Post** /api/server/v1/users | Create user
[**DeleteUser**](ServerUsersAPI.md#DeleteUser) | **Delete** /api/server/v1/users/{userId} | Delete user
[**GetUser**](ServerUsersAPI.md#GetUser) | **Get** /api/server/v1/users/{userId} | Get user
[**SearchUsers**](ServerUsersAPI.md#SearchUsers) | **Post** /api/server/v1/users/search | Search users
[**UnbanUser**](ServerUsersAPI.md#UnbanUser) | **Post** /api/server/v1/users/{userId}/unban | Unban user
[**UpdateUser**](ServerUsersAPI.md#UpdateUser) | **Patch** /api/server/v1/users/{userId} | Update user
[**UpdateUserMetadata**](ServerUsersAPI.md#UpdateUserMetadata) | **Patch** /api/server/v1/users/{userId}/metadata | Update user metadata



## BanUser

> ServerUserResponse BanUser(ctx, userId).Execute()

Ban user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to ban.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.BanUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.BanUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BanUser`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.BanUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to ban. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBanUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ServerUserResponse CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Create user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createUserRequest := *openapiclient.NewCreateUserRequest(map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}) // CreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerUsersAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> ServerUserResponse GetUser(ctx, userId).Execute()

Get user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> CursorPageResponseServerUserResponse SearchUsers(ctx).Limit(limit).Cursor(cursor).ServerUserSearchRequest(serverUserSearchRequest).Execute()

Search users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(50) // int32 | Maximum number of items in the returned page (default 20). (optional) (default to 20)
	cursor := "01931a73-8b00-7000-8000-000000000000" // string | Opaque cursor returned by the previous page's `nextCursor`. Omit to fetch the first page. (optional)
	serverUserSearchRequest := *openapiclient.NewServerUserSearchRequest() // ServerUserSearchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.SearchUsers(context.Background()).Limit(limit).Cursor(cursor).ServerUserSearchRequest(serverUserSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.SearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsers`: CursorPageResponseServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items in the returned page (default 20). | [default to 20]
 **cursor** | **string** | Opaque cursor returned by the previous page&#39;s &#x60;nextCursor&#x60;. Omit to fetch the first page. | 
 **serverUserSearchRequest** | [**ServerUserSearchRequest**](ServerUserSearchRequest.md) |  | 

### Return type

[**CursorPageResponseServerUserResponse**](CursorPageResponseServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbanUser

> ServerUserResponse UnbanUser(ctx, userId).Execute()

Unban user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to unban.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.UnbanUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.UnbanUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbanUser`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.UnbanUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to unban. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbanUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ServerUserResponse UpdateUser(ctx, userId).UpdateUserRequest(updateUserRequest).Execute()

Update user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to update.
	updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.UpdateUser(context.Background(), userId).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserMetadata

> ServerUserResponse UpdateUserMetadata(ctx, userId).UpdateUserMetadataRequest(updateUserMetadataRequest).Execute()

Update user metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "01931a73-8b00-7000-8000-000000000000" // string | Identifier of the user to update.
	updateUserMetadataRequest := *openapiclient.NewUpdateUserMetadataRequest() // UpdateUserMetadataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerUsersAPI.UpdateUserMetadata(context.Background(), userId).UpdateUserMetadataRequest(updateUserMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerUsersAPI.UpdateUserMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserMetadata`: ServerUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerUsersAPI.UpdateUserMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserMetadataRequest** | [**UpdateUserMetadataRequest**](UpdateUserMetadataRequest.md) |  | 

### Return type

[**ServerUserResponse**](ServerUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

