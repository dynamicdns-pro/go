# \SubdomainAPI

All URIs are relative to *https://dynamicdns.pro/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Update**](SubdomainAPI.md#Update) | **Post** /{subdomain}/record | 
[**Updateip**](SubdomainAPI.md#Updateip) | **Post** /{subdomain} | update the ip address with the connecting ip address



## Update

> Update200Response Update(ctx, subdomain).UpdateRequest(updateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dynamicdns-pro/GIT_REPO_ID"
)

func main() {
	subdomain := "subdomain_example" // string | 
	updateRequest := *openapiclient.NewUpdateRequest() // UpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.Update(context.Background(), subdomain).UpdateRequest(updateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: Update200Response
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subdomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) |  | 

### Return type

[**Update200Response**](Update200Response.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updateip

> Updateip200Response Updateip(ctx, subdomain).Body(body).Execute()

update the ip address with the connecting ip address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dynamicdns-pro/GIT_REPO_ID"
)

func main() {
	subdomain := "subdomain_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.Updateip(context.Background(), subdomain).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.Updateip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updateip`: Updateip200Response
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.Updateip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subdomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**Updateip200Response**](Updateip200Response.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

