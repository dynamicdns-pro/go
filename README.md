# Go API client for dynamicdnsapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dynamicdnsapi "github.com/dynamicdns-pro/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dynamicdnsapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dynamicdnsapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dynamicdnsapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dynamicdnsapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dynamicdnsapi.ContextOperationServerIndices` and `dynamicdnsapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dynamicdnsapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dynamicdnsapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://dynamicdns.pro/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SubdomainAPI* | [**Update**](docs/SubdomainAPI.md#update) | **Post** /{subdomain}/record | 
*SubdomainAPI* | [**Updateip**](docs/SubdomainAPI.md#updateip) | **Post** /{subdomain} | update the ip address with the connecting ip address


## Documentation For Models

 - [Update200Response](docs/Update200Response.md)
 - [Update200ResponseAnyOf](docs/Update200ResponseAnyOf.md)
 - [Update400Response](docs/Update400Response.md)
 - [Update403Response](docs/Update403Response.md)
 - [UpdateRequest](docs/UpdateRequest.md)
 - [Updateip200Response](docs/Updateip200Response.md)
 - [Updateip200ResponseAnyOf](docs/Updateip200ResponseAnyOf.md)
 - [Updateip400Response](docs/Updateip400Response.md)
 - [Updateip403Response](docs/Updateip403Response.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### http

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), dynamicdnsapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



