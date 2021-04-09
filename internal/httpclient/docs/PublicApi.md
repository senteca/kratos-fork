# \PublicApi

All URIs are relative to *https://demo.tenants.oryapis.com/api/kratos/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteSelfServiceBrowserSettingsOIDCSettingsFlow**](PublicApi.md#CompleteSelfServiceBrowserSettingsOIDCSettingsFlow) | **Post** /self-service/browser/flows/registration/strategies/oidc/settings/connections | Complete the Browser-Based Settings Flow for the OpenID Connect Strategy
[**CompleteSelfServiceLoginFlowWithPasswordMethod**](PublicApi.md#CompleteSelfServiceLoginFlowWithPasswordMethod) | **Post** /self-service/login/methods/password | Complete Login Flow with Username/Email Password Method
[**CompleteSelfServiceRecoveryFlowWithLinkMethod**](PublicApi.md#CompleteSelfServiceRecoveryFlowWithLinkMethod) | **Post** /self-service/recovery/methods/link | Complete Recovery Flow with Link Method
[**CompleteSelfServiceRegistrationFlowWithPasswordMethod**](PublicApi.md#CompleteSelfServiceRegistrationFlowWithPasswordMethod) | **Post** /self-service/registration/methods/password | Complete Registration Flow with Username/Email Password Method
[**CompleteSelfServiceSettingsFlowWithPasswordMethod**](PublicApi.md#CompleteSelfServiceSettingsFlowWithPasswordMethod) | **Post** /self-service/settings/methods/password | Complete Settings Flow with Username/Email Password Method
[**CompleteSelfServiceSettingsFlowWithProfileMethod**](PublicApi.md#CompleteSelfServiceSettingsFlowWithProfileMethod) | **Post** /self-service/settings/methods/profile | Complete Settings Flow with Profile Method
[**CompleteSelfServiceVerificationFlowWithLinkMethod**](PublicApi.md#CompleteSelfServiceVerificationFlowWithLinkMethod) | **Post** /self-service/verification/methods/link | Complete Verification Flow with Link Method
[**GetSchema**](PublicApi.md#GetSchema) | **Get** /schemas/{id} | 
[**GetSelfServiceError**](PublicApi.md#GetSelfServiceError) | **Get** /self-service/errors | Get User-Facing Self-Service Errors
[**GetSelfServiceLoginFlow**](PublicApi.md#GetSelfServiceLoginFlow) | **Get** /self-service/login/flows | Get Login Flow
[**GetSelfServiceRecoveryFlow**](PublicApi.md#GetSelfServiceRecoveryFlow) | **Get** /self-service/recovery/flows | Get information about a recovery flow
[**GetSelfServiceRegistrationFlow**](PublicApi.md#GetSelfServiceRegistrationFlow) | **Get** /self-service/registration/flows | Get Registration Flow
[**GetSelfServiceSettingsFlow**](PublicApi.md#GetSelfServiceSettingsFlow) | **Get** /self-service/settings/flows | Get Settings Flow
[**GetSelfServiceVerificationFlow**](PublicApi.md#GetSelfServiceVerificationFlow) | **Get** /self-service/verification/flows | Get Verification Flow
[**InitializeSelfServiceBrowserLogoutFlow**](PublicApi.md#InitializeSelfServiceBrowserLogoutFlow) | **Get** /self-service/browser/flows/logout | Initialize Browser-Based Logout User Flow
[**InitializeSelfServiceLoginViaAPIFlow**](PublicApi.md#InitializeSelfServiceLoginViaAPIFlow) | **Get** /self-service/login/api | Initialize Login Flow for API clients
[**InitializeSelfServiceLoginViaBrowserFlow**](PublicApi.md#InitializeSelfServiceLoginViaBrowserFlow) | **Get** /self-service/login/browser | Initialize Login Flow for browsers
[**InitializeSelfServiceRecoveryViaAPIFlow**](PublicApi.md#InitializeSelfServiceRecoveryViaAPIFlow) | **Get** /self-service/recovery/api | Initialize Recovery Flow for API Clients
[**InitializeSelfServiceRecoveryViaBrowserFlow**](PublicApi.md#InitializeSelfServiceRecoveryViaBrowserFlow) | **Get** /self-service/recovery/browser | Initialize Recovery Flow for Browser Clients
[**InitializeSelfServiceRegistrationViaAPIFlow**](PublicApi.md#InitializeSelfServiceRegistrationViaAPIFlow) | **Get** /self-service/registration/api | Initialize Registration Flow for API clients
[**InitializeSelfServiceRegistrationViaBrowserFlow**](PublicApi.md#InitializeSelfServiceRegistrationViaBrowserFlow) | **Get** /self-service/registration/browser | Initialize Registration Flow for browsers
[**InitializeSelfServiceSettingsViaAPIFlow**](PublicApi.md#InitializeSelfServiceSettingsViaAPIFlow) | **Get** /self-service/settings/api | Initialize Settings Flow for API Clients
[**InitializeSelfServiceSettingsViaBrowserFlow**](PublicApi.md#InitializeSelfServiceSettingsViaBrowserFlow) | **Get** /self-service/settings/browser | Initialize Settings Flow for Browsers
[**InitializeSelfServiceVerificationViaAPIFlow**](PublicApi.md#InitializeSelfServiceVerificationViaAPIFlow) | **Get** /self-service/verification/api | Initialize Verification Flow for API Clients
[**InitializeSelfServiceVerificationViaBrowserFlow**](PublicApi.md#InitializeSelfServiceVerificationViaBrowserFlow) | **Get** /self-service/verification/browser | Initialize Verification Flow for Browser Clients
[**RevokeSession**](PublicApi.md#RevokeSession) | **Delete** /sessions | Revoke and Invalidate a Session
[**Whoami**](PublicApi.md#Whoami) | **Get** /sessions/whoami | Check Who the Current HTTP Session Belongs To



## CompleteSelfServiceBrowserSettingsOIDCSettingsFlow

> CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(ctx).Execute()

Complete the Browser-Based Settings Flow for the OpenID Connect Strategy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceBrowserSettingsOIDCSettingsFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceBrowserSettingsOIDCSettingsFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceLoginFlowWithPasswordMethod

> LoginViaApiResponse CompleteSelfServiceLoginFlowWithPasswordMethod(ctx).Flow(flow).CompleteSelfServiceLoginFlowWithPasswordMethod(completeSelfServiceLoginFlowWithPasswordMethod).Execute()

Complete Login Flow with Username/Email Password Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flow := "flow_example" // string | The Flow ID
    completeSelfServiceLoginFlowWithPasswordMethod := *openapiclient.NewCompleteSelfServiceLoginFlowWithPasswordMethod() // CompleteSelfServiceLoginFlowWithPasswordMethod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceLoginFlowWithPasswordMethod(context.Background()).Flow(flow).CompleteSelfServiceLoginFlowWithPasswordMethod(completeSelfServiceLoginFlowWithPasswordMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceLoginFlowWithPasswordMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteSelfServiceLoginFlowWithPasswordMethod`: LoginViaApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.CompleteSelfServiceLoginFlowWithPasswordMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceLoginFlowWithPasswordMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Flow ID | 
 **completeSelfServiceLoginFlowWithPasswordMethod** | [**CompleteSelfServiceLoginFlowWithPasswordMethod**](CompleteSelfServiceLoginFlowWithPasswordMethod.md) |  | 

### Return type

[**LoginViaApiResponse**](loginViaApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceRecoveryFlowWithLinkMethod

> CompleteSelfServiceRecoveryFlowWithLinkMethod(ctx).Token(token).Flow(flow).CompleteSelfServiceRecoveryFlowWithLinkMethod(completeSelfServiceRecoveryFlowWithLinkMethod).Execute()

Complete Recovery Flow with Link Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | Recovery Token  The recovery token which completes the recovery request. If the token is invalid (e.g. expired) an error will be shown to the end-user. (optional)
    flow := "flow_example" // string | The Flow ID  format: uuid (optional)
    completeSelfServiceRecoveryFlowWithLinkMethod := *openapiclient.NewCompleteSelfServiceRecoveryFlowWithLinkMethod() // CompleteSelfServiceRecoveryFlowWithLinkMethod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceRecoveryFlowWithLinkMethod(context.Background()).Token(token).Flow(flow).CompleteSelfServiceRecoveryFlowWithLinkMethod(completeSelfServiceRecoveryFlowWithLinkMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceRecoveryFlowWithLinkMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceRecoveryFlowWithLinkMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Recovery Token  The recovery token which completes the recovery request. If the token is invalid (e.g. expired) an error will be shown to the end-user. | 
 **flow** | **string** | The Flow ID  format: uuid | 
 **completeSelfServiceRecoveryFlowWithLinkMethod** | [**CompleteSelfServiceRecoveryFlowWithLinkMethod**](CompleteSelfServiceRecoveryFlowWithLinkMethod.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceRegistrationFlowWithPasswordMethod

> RegistrationViaApiResponse CompleteSelfServiceRegistrationFlowWithPasswordMethod(ctx).Flow(flow).Body(body).Execute()

Complete Registration Flow with Username/Email Password Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flow := "flow_example" // string | Flow is flow ID. (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceRegistrationFlowWithPasswordMethod(context.Background()).Flow(flow).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceRegistrationFlowWithPasswordMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteSelfServiceRegistrationFlowWithPasswordMethod`: RegistrationViaApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.CompleteSelfServiceRegistrationFlowWithPasswordMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceRegistrationFlowWithPasswordMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | Flow is flow ID. | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**RegistrationViaApiResponse**](registrationViaApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceSettingsFlowWithPasswordMethod

> SettingsViaApiResponse CompleteSelfServiceSettingsFlowWithPasswordMethod(ctx).Flow(flow).CompleteSelfServiceSettingsFlowWithPasswordMethod(completeSelfServiceSettingsFlowWithPasswordMethod).Execute()

Complete Settings Flow with Username/Email Password Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flow := "flow_example" // string | Flow is flow ID. (optional)
    completeSelfServiceSettingsFlowWithPasswordMethod := *openapiclient.NewCompleteSelfServiceSettingsFlowWithPasswordMethod("Password_example") // CompleteSelfServiceSettingsFlowWithPasswordMethod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceSettingsFlowWithPasswordMethod(context.Background()).Flow(flow).CompleteSelfServiceSettingsFlowWithPasswordMethod(completeSelfServiceSettingsFlowWithPasswordMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceSettingsFlowWithPasswordMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteSelfServiceSettingsFlowWithPasswordMethod`: SettingsViaApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.CompleteSelfServiceSettingsFlowWithPasswordMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceSettingsFlowWithPasswordMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | Flow is flow ID. | 
 **completeSelfServiceSettingsFlowWithPasswordMethod** | [**CompleteSelfServiceSettingsFlowWithPasswordMethod**](CompleteSelfServiceSettingsFlowWithPasswordMethod.md) |  | 

### Return type

[**SettingsViaApiResponse**](settingsViaApiResponse.md)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceSettingsFlowWithProfileMethod

> SettingsFlow CompleteSelfServiceSettingsFlowWithProfileMethod(ctx).Flow(flow).Body(body).Execute()

Complete Settings Flow with Profile Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flow := "flow_example" // string | Flow is flow ID. (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceSettingsFlowWithProfileMethod(context.Background()).Flow(flow).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceSettingsFlowWithProfileMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteSelfServiceSettingsFlowWithProfileMethod`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.CompleteSelfServiceSettingsFlowWithProfileMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceSettingsFlowWithProfileMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | Flow is flow ID. | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**SettingsFlow**](settingsFlow.md)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSelfServiceVerificationFlowWithLinkMethod

> CompleteSelfServiceVerificationFlowWithLinkMethod(ctx).Token(token).Flow(flow).CompleteSelfServiceVerificationFlowWithLinkMethod(completeSelfServiceVerificationFlowWithLinkMethod).Execute()

Complete Verification Flow with Link Method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | Verification Token  The verification token which completes the verification request. If the token is invalid (e.g. expired) an error will be shown to the end-user. (optional)
    flow := "flow_example" // string | The Flow ID  format: uuid (optional)
    completeSelfServiceVerificationFlowWithLinkMethod := *openapiclient.NewCompleteSelfServiceVerificationFlowWithLinkMethod() // CompleteSelfServiceVerificationFlowWithLinkMethod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.CompleteSelfServiceVerificationFlowWithLinkMethod(context.Background()).Token(token).Flow(flow).CompleteSelfServiceVerificationFlowWithLinkMethod(completeSelfServiceVerificationFlowWithLinkMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.CompleteSelfServiceVerificationFlowWithLinkMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSelfServiceVerificationFlowWithLinkMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Verification Token  The verification token which completes the verification request. If the token is invalid (e.g. expired) an error will be shown to the end-user. | 
 **flow** | **string** | The Flow ID  format: uuid | 
 **completeSelfServiceVerificationFlowWithLinkMethod** | [**CompleteSelfServiceVerificationFlowWithLinkMethod**](CompleteSelfServiceVerificationFlowWithLinkMethod.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> map[string]interface{} GetSchema(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID must be set to the ID of schema you want to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSchema(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID must be set to the ID of schema you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceError

> ErrorContainer GetSelfServiceError(ctx).Error_(error_).Execute()

Get User-Facing Self-Service Errors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    error_ := "error__example" // string | Error is the container's ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceError(context.Background()).Error_(error_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceError`: ErrorContainer
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **error_** | **string** | Error is the container&#39;s ID | 

### Return type

[**ErrorContainer**](errorContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceLoginFlow

> LoginFlow GetSelfServiceLoginFlow(ctx).Id(id).Execute()

Get Login Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The Login Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/login?flow=abcde`).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceLoginFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceLoginFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceLoginFlow`: LoginFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceLoginFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceLoginFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Login Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/login?flow&#x3D;abcde&#x60;). | 

### Return type

[**LoginFlow**](loginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceRecoveryFlow

> RecoveryFlow GetSelfServiceRecoveryFlow(ctx).Id(id).Execute()

Get information about a recovery flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The Flow ID  The value for this parameter comes from `request` URL Query parameter sent to your application (e.g. `/recovery?flow=abcde`).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceRecoveryFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceRecoveryFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceRecoveryFlow`: RecoveryFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceRecoveryFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceRecoveryFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Flow ID  The value for this parameter comes from &#x60;request&#x60; URL Query parameter sent to your application (e.g. &#x60;/recovery?flow&#x3D;abcde&#x60;). | 

### Return type

[**RecoveryFlow**](recoveryFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceRegistrationFlow

> RegistrationFlow GetSelfServiceRegistrationFlow(ctx).Id(id).Execute()

Get Registration Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The Registration Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/registration?flow=abcde`).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceRegistrationFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceRegistrationFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceRegistrationFlow`: RegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceRegistrationFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceRegistrationFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Registration Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/registration?flow&#x3D;abcde&#x60;). | 

### Return type

[**RegistrationFlow**](registrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceSettingsFlow

> SettingsFlow GetSelfServiceSettingsFlow(ctx).Id(id).Execute()

Get Settings Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID is the Settings Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/settings?flow=abcde`).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceSettingsFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceSettingsFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceSettingsFlow`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceSettingsFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceSettingsFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID is the Settings Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/settings?flow&#x3D;abcde&#x60;). | 

### Return type

[**SettingsFlow**](settingsFlow.md)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServiceVerificationFlow

> VerificationFlow GetSelfServiceVerificationFlow(ctx).Id(id).Execute()

Get Verification Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The Flow ID  The value for this parameter comes from `request` URL Query parameter sent to your application (e.g. `/verification?flow=abcde`).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.GetSelfServiceVerificationFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetSelfServiceVerificationFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfServiceVerificationFlow`: VerificationFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetSelfServiceVerificationFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServiceVerificationFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Flow ID  The value for this parameter comes from &#x60;request&#x60; URL Query parameter sent to your application (e.g. &#x60;/verification?flow&#x3D;abcde&#x60;). | 

### Return type

[**VerificationFlow**](verificationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceBrowserLogoutFlow

> InitializeSelfServiceBrowserLogoutFlow(ctx).Execute()

Initialize Browser-Based Logout User Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceBrowserLogoutFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceBrowserLogoutFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceBrowserLogoutFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceLoginViaAPIFlow

> LoginFlow InitializeSelfServiceLoginViaAPIFlow(ctx).Refresh(refresh).Execute()

Initialize Login Flow for API clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refresh := true // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceLoginViaAPIFlow(context.Background()).Refresh(refresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceLoginViaAPIFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceLoginViaAPIFlow`: LoginFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.InitializeSelfServiceLoginViaAPIFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceLoginViaAPIFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. | 

### Return type

[**LoginFlow**](loginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceLoginViaBrowserFlow

> InitializeSelfServiceLoginViaBrowserFlow(ctx).Execute()

Initialize Login Flow for browsers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceLoginViaBrowserFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceLoginViaBrowserFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceLoginViaBrowserFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceRecoveryViaAPIFlow

> RecoveryFlow InitializeSelfServiceRecoveryViaAPIFlow(ctx).Execute()

Initialize Recovery Flow for API Clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceRecoveryViaAPIFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceRecoveryViaAPIFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceRecoveryViaAPIFlow`: RecoveryFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.InitializeSelfServiceRecoveryViaAPIFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceRecoveryViaAPIFlowRequest struct via the builder pattern


### Return type

[**RecoveryFlow**](recoveryFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceRecoveryViaBrowserFlow

> InitializeSelfServiceRecoveryViaBrowserFlow(ctx).Execute()

Initialize Recovery Flow for Browser Clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceRecoveryViaBrowserFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceRecoveryViaBrowserFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceRecoveryViaBrowserFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceRegistrationViaAPIFlow

> RegistrationFlow InitializeSelfServiceRegistrationViaAPIFlow(ctx).Execute()

Initialize Registration Flow for API clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceRegistrationViaAPIFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceRegistrationViaAPIFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceRegistrationViaAPIFlow`: RegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.InitializeSelfServiceRegistrationViaAPIFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceRegistrationViaAPIFlowRequest struct via the builder pattern


### Return type

[**RegistrationFlow**](registrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceRegistrationViaBrowserFlow

> InitializeSelfServiceRegistrationViaBrowserFlow(ctx).Execute()

Initialize Registration Flow for browsers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceRegistrationViaBrowserFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceRegistrationViaBrowserFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceRegistrationViaBrowserFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceSettingsViaAPIFlow

> SettingsFlow InitializeSelfServiceSettingsViaAPIFlow(ctx).Execute()

Initialize Settings Flow for API Clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceSettingsViaAPIFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceSettingsViaAPIFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceSettingsViaAPIFlow`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.InitializeSelfServiceSettingsViaAPIFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceSettingsViaAPIFlowRequest struct via the builder pattern


### Return type

[**SettingsFlow**](settingsFlow.md)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceSettingsViaBrowserFlow

> InitializeSelfServiceSettingsViaBrowserFlow(ctx).Execute()

Initialize Settings Flow for Browsers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceSettingsViaBrowserFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceSettingsViaBrowserFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceSettingsViaBrowserFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceVerificationViaAPIFlow

> VerificationFlow InitializeSelfServiceVerificationViaAPIFlow(ctx).Execute()

Initialize Verification Flow for API Clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceVerificationViaAPIFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceVerificationViaAPIFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceVerificationViaAPIFlow`: VerificationFlow
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.InitializeSelfServiceVerificationViaAPIFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceVerificationViaAPIFlowRequest struct via the builder pattern


### Return type

[**VerificationFlow**](verificationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSelfServiceVerificationViaBrowserFlow

> InitializeSelfServiceVerificationViaBrowserFlow(ctx).Execute()

Initialize Verification Flow for Browser Clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.InitializeSelfServiceVerificationViaBrowserFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.InitializeSelfServiceVerificationViaBrowserFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSelfServiceVerificationViaBrowserFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSession

> RevokeSession(ctx).RevokeSession(revokeSession).Execute()

Revoke and Invalidate a Session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    revokeSession := *openapiclient.NewRevokeSession("SessionToken_example") // RevokeSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.RevokeSession(context.Background()).RevokeSession(revokeSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.RevokeSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeSession** | [**RevokeSession**](RevokeSession.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Whoami

> Session Whoami(ctx).Cookie(cookie).Authorization(authorization).Execute()

Check Who the Current HTTP Session Belongs To



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cookie := "cookie_example" // string |  (optional)
    authorization := "authorization_example" // string | in: authorization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.Whoami(context.Background()).Cookie(cookie).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.Whoami``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Whoami`: Session
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.Whoami`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string** |  | 
 **authorization** | **string** | in: authorization | 

### Return type

[**Session**](session.md)

### Authorization

[sessionToken](../README.md#sessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

