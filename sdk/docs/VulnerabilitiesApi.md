# \VulnerabilitiesApi

All URIs are relative to *http://localhost:8443/metastore/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImageVulnerabilities**](VulnerabilitiesApi.md#GetImageVulnerabilities) | **Get** /images/{IDorDigest}/vulnerabilities | List vulnerabilities from the given image.
[**GetPackageVulnerabilities**](VulnerabilitiesApi.md#GetPackageVulnerabilities) | **Get** /packages/{IDorName}/vulnerabilities | List vulnerabilities from the given package.
[**GetSourceVulnerabilitiesQuery**](VulnerabilitiesApi.md#GetSourceVulnerabilitiesQuery) | **Get** /sources/vulnerabilities | List vulnerabilities of the given source.
[**GetVulnerabilities**](VulnerabilitiesApi.md#GetVulnerabilities) | **Get** /vulnerabilities | Search for vulnerabilities by CVE id.


# **GetImageVulnerabilities**
> []Vulnerability GetImageVulnerabilities(ctx, iDorDigest, optional)
List vulnerabilities from the given image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iDorDigest** | **string**|  | 
 **optional** | ***VulnerabilitiesApiGetImageVulnerabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VulnerabilitiesApiGetImageVulnerabilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **optional.String**| Case insensitive vulnerabilities severity filter. Possible values are: low, medium, high, critical, unknown. | 

### Return type

[**[]Vulnerability**](Vulnerability.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackageVulnerabilities**
> []Vulnerability GetPackageVulnerabilities(ctx, iDorName, optional)
List vulnerabilities from the given package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iDorName** | **string**|  | 
 **optional** | ***VulnerabilitiesApiGetPackageVulnerabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VulnerabilitiesApiGetPackageVulnerabilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **optional.String**| Case insensitive vulnerabilities severity filter. Possible values are: low, medium, high, critical, unknown. | 

### Return type

[**[]Vulnerability**](Vulnerability.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceVulnerabilitiesQuery**
> []Vulnerability GetSourceVulnerabilitiesQuery(ctx, optional)
List vulnerabilities of the given source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VulnerabilitiesApiGetSourceVulnerabilitiesQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VulnerabilitiesApiGetSourceVulnerabilitiesQueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **severity** | **optional.String**| Case insensitive vulnerabilities severity filter. Possible values are: low, medium, high, critical, unknown. | 
 **id** | **optional.Int32**|  | 
 **repo** | **optional.String**|  | 
 **sha** | **optional.String**|  | 

### Return type

[**[]Vulnerability**](Vulnerability.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilities**
> []Vulnerability GetVulnerabilities(ctx, cVEID, optional)
Search for vulnerabilities by CVE id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cVEID** | **string**|  | 
 **optional** | ***VulnerabilitiesApiGetVulnerabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VulnerabilitiesApiGetVulnerabilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **optional.String**| Case insensitive vulnerabilities severity filter. Possible values are: low, medium, high, critical, unknown. | 

### Return type

[**[]Vulnerability**](Vulnerability.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

