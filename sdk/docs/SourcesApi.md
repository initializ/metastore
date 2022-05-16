# \SourcesApi

All URIs are relative to *http://localhost:8443/metastore/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSource**](SourcesApi.md#CreateSource) | **Post** /sources | Create a new source report. Related packages and vulnerabilities are also created.
[**GetPackageSources**](SourcesApi.md#GetPackageSources) | **Get** /packages/{IDorName}/sources | List the sources containing the given package.
[**GetSources**](SourcesApi.md#GetSources) | **Get** /sources | Search for sources by ID, repository, commit sha and/or organization.
[**GetVulnerabilitySources**](SourcesApi.md#GetVulnerabilitySources) | **Get** /vulnerabilities/{CVEID}/sources | List sources that contain the given vulnerability.


# **CreateSource**
> Source CreateSource(ctx, image)
Create a new source report. Related packages and vulnerabilities are also created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **image** | [**Source**](Source.md)|  | 

### Return type

[**Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackageSources**
> []Source GetPackageSources(ctx, iDorName)
List the sources containing the given package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iDorName** | **string**|  | 

### Return type

[**[]Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSources**
> []Source GetSources(ctx, optional)
Search for sources by ID, repository, commit sha and/or organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SourcesApiGetSourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourcesApiGetSourcesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**|  | 
 **repo** | **optional.String**|  | 
 **sha** | **optional.String**|  | 
 **org** | **optional.String**|  | 

### Return type

[**[]Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilitySources**
> []Source GetVulnerabilitySources(ctx, cVEID)
List sources that contain the given vulnerability.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cVEID** | **string**|  | 

### Return type

[**[]Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

