# \InferenceEndpointsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInferenceEndpointJob**](InferenceEndpointsAPI.md#CreateInferenceEndpointJob) | **Post** /organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs | Create a New Inference Endpoint Job
[**DeleteInferenceEndpointJob**](InferenceEndpointsAPI.md#DeleteInferenceEndpointJob) | **Delete** /organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id} | Cancel an Inference Endpoint Job
[**GetInferenceEndpoint**](InferenceEndpointsAPI.md#GetInferenceEndpoint) | **Get** /organizations/{organization_name}/inference-endpoints/{inference_endpoint_name} | Get an Inference Endpoint
[**GetInferenceEndpointJob**](InferenceEndpointsAPI.md#GetInferenceEndpointJob) | **Get** /organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id} | Get an Inference Endpoint Job
[**ListInferenceEndpointJobs**](InferenceEndpointsAPI.md#ListInferenceEndpointJobs) | **Get** /organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs | List Inference Endpoint Jobs
[**ListInferenceEndpoints**](InferenceEndpointsAPI.md#ListInferenceEndpoints) | **Get** /organizations/{organization_name}/inference-endpoints | List Inference Endpoints



## CreateInferenceEndpointJob

> InferenceEndpointJob CreateInferenceEndpointJob(ctx, organizationName, inferenceEndpointName).InferenceEndpointJobPrototype(inferenceEndpointJobPrototype).Execute()

Create a New Inference Endpoint Job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	inferenceEndpointName := "inferenceEndpointName_example" // string | The inference endpoint name.
	inferenceEndpointJobPrototype := *openapiclient.NewInferenceEndpointJobPrototype(interface{}(123)) // InferenceEndpointJobPrototype | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InferenceEndpointsAPI.CreateInferenceEndpointJob(context.Background(), organizationName, inferenceEndpointName).InferenceEndpointJobPrototype(inferenceEndpointJobPrototype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.CreateInferenceEndpointJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInferenceEndpointJob`: InferenceEndpointJob
	fmt.Fprintf(os.Stdout, "Response from `InferenceEndpointsAPI.CreateInferenceEndpointJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**inferenceEndpointName** | **string** | The inference endpoint name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInferenceEndpointJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inferenceEndpointJobPrototype** | [**InferenceEndpointJobPrototype**](InferenceEndpointJobPrototype.md) |  | 

### Return type

[**InferenceEndpointJob**](InferenceEndpointJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInferenceEndpointJob

> DeleteInferenceEndpointJob(ctx, organizationName, inferenceEndpointName, inferenceEndpointJobId).Execute()

Cancel an Inference Endpoint Job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	inferenceEndpointName := "inferenceEndpointName_example" // string | The inference endpoint name.
	inferenceEndpointJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The inference endpoint job identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InferenceEndpointsAPI.DeleteInferenceEndpointJob(context.Background(), organizationName, inferenceEndpointName, inferenceEndpointJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.DeleteInferenceEndpointJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**inferenceEndpointName** | **string** | The inference endpoint name. | 
**inferenceEndpointJobId** | **string** | The inference endpoint job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInferenceEndpointJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceEndpoint

> InferenceEndpoint GetInferenceEndpoint(ctx, organizationName, inferenceEndpointName).Execute()

Get an Inference Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	inferenceEndpointName := "inferenceEndpointName_example" // string | The inference endpoint name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InferenceEndpointsAPI.GetInferenceEndpoint(context.Background(), organizationName, inferenceEndpointName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.GetInferenceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceEndpoint`: InferenceEndpoint
	fmt.Fprintf(os.Stdout, "Response from `InferenceEndpointsAPI.GetInferenceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**inferenceEndpointName** | **string** | The inference endpoint name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InferenceEndpoint**](InferenceEndpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceEndpointJob

> InferenceEndpointJob GetInferenceEndpointJob(ctx, organizationName, inferenceEndpointName, inferenceEndpointJobId).Execute()

Get an Inference Endpoint Job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	inferenceEndpointName := "inferenceEndpointName_example" // string | The inference endpoint name.
	inferenceEndpointJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The inference endpoint job identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InferenceEndpointsAPI.GetInferenceEndpointJob(context.Background(), organizationName, inferenceEndpointName, inferenceEndpointJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.GetInferenceEndpointJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceEndpointJob`: InferenceEndpointJob
	fmt.Fprintf(os.Stdout, "Response from `InferenceEndpointsAPI.GetInferenceEndpointJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**inferenceEndpointName** | **string** | The inference endpoint name. | 
**inferenceEndpointJobId** | **string** | The inference endpoint job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceEndpointJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InferenceEndpointJob**](InferenceEndpointJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInferenceEndpointJobs

> InferenceEndpointJobCollection ListInferenceEndpointJobs(ctx, organizationName, inferenceEndpointName).Page(page).PageSize(pageSize).Execute()

List Inference Endpoint Jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	inferenceEndpointName := "inferenceEndpointName_example" // string | The inference endpoint name.
	page := int32(56) // int32 | The page number. (optional)
	pageSize := int32(56) // int32 | The maximum number of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InferenceEndpointsAPI.ListInferenceEndpointJobs(context.Background(), organizationName, inferenceEndpointName).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.ListInferenceEndpointJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInferenceEndpointJobs`: InferenceEndpointJobCollection
	fmt.Fprintf(os.Stdout, "Response from `InferenceEndpointsAPI.ListInferenceEndpointJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**inferenceEndpointName** | **string** | The inference endpoint name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInferenceEndpointJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page number. | 
 **pageSize** | **int32** | The maximum number of items per page. | 

### Return type

[**InferenceEndpointJobCollection**](InferenceEndpointJobCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInferenceEndpoints

> InferenceEndpointCollection ListInferenceEndpoints(ctx, organizationName).Page(page).PageSize(pageSize).Execute()

List Inference Endpoints



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	page := int32(56) // int32 | The page number. (optional)
	pageSize := int32(56) // int32 | The maximum number of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InferenceEndpointsAPI.ListInferenceEndpoints(context.Background(), organizationName).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InferenceEndpointsAPI.ListInferenceEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInferenceEndpoints`: InferenceEndpointCollection
	fmt.Fprintf(os.Stdout, "Response from `InferenceEndpointsAPI.ListInferenceEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInferenceEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. | 
 **pageSize** | **int32** | The maximum number of items per page. | 

### Return type

[**InferenceEndpointCollection**](InferenceEndpointCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

