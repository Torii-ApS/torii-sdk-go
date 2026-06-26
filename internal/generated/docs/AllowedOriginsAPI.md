# \AllowedOriginsAPI

All URIs are relative to *https://api.torii.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](AllowedOriginsAPI.md#List) | **Get** /api/server/v1/allowed-origins | List escape-hatch origins for this environment
[**Set**](AllowedOriginsAPI.md#Set) | **Put** /api/server/v1/allowed-origins | Replace the escape-hatch origins for this environment



## List

> AllowedOriginsResponse List(ctx).Execute()

List escape-hatch origins for this environment

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedOriginsAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: AllowedOriginsResponse
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**AllowedOriginsResponse**](AllowedOriginsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set

> AllowedOriginsResponse Set(ctx).SetAllowedOriginsRequest(setAllowedOriginsRequest).Execute()

Replace the escape-hatch origins for this environment



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
	setAllowedOriginsRequest := *openapiclient.NewSetAllowedOriginsRequest([]string{"Origins_example"}) // SetAllowedOriginsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedOriginsAPI.Set(context.Background()).SetAllowedOriginsRequest(setAllowedOriginsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.Set``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set`: AllowedOriginsResponse
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.Set`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setAllowedOriginsRequest** | [**SetAllowedOriginsRequest**](SetAllowedOriginsRequest.md) |  | 

### Return type

[**AllowedOriginsResponse**](AllowedOriginsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

