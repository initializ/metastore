# \PackagesApi

All URIs are relative to *http://localhost:8443/metastore/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImagePackages**](PackagesApi.md#GetImagePackages) | **Get** /images/{IDorDigest}/packages | List the packages in an image.
[**GetImagePackagesQuery**](PackagesApi.md#GetImagePackagesQuery) | **Get** /images/packages | List packages of the given image.
[**GetPackages**](PackagesApi.md#GetPackages) | **Get** /packages | Search packages by id, name and/or version.
[**GetSourcePackagesQuery**](PackagesApi.md#GetSourcePackagesQuery) | **Get** /sources/packages | List packages of the given source.
[**GetVulnerabilityPackages**](PackagesApi.md#GetVulnerabilityPackages) | **Get** /vulnerabilities/{CVEID}/packages | List packages that contain the given CVE id.


# **GetImagePackages**
> []ModelPackage GetImagePackages(ctx, iDorDigest)
List the packages in an image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iDorDigest** | **string**|  | 

### Return type

[**[]ModelPackage**](Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImagePackagesQuery**
> []ModelPackage GetImagePackagesQuery(ctx, optional)
List packages of the given image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackagesApiGetImagePackagesQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackagesApiGetImagePackagesQueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**|  | 
 **digest** | **optional.String**|  | 
 **name** | **optional.String**|  | 

### Return type

[**[]ModelPackage**](Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackages**
> []ModelPackage GetPackages(ctx, optional)
Search packages by id, name and/or version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackagesApiGetPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackagesApiGetPackagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| Any of id or name must be provided | 
 **name** | **optional.String**| Any of id or name must be provided | 
 **version** | **optional.String**|  | 

### Return type

[**[]ModelPackage**](Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcePackagesQuery**
> []ModelPackage GetSourcePackagesQuery(ctx, optional)
List packages of the given source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackagesApiGetSourcePackagesQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackagesApiGetSourcePackagesQueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**|  | 
 **repo** | **optional.String**|  | 
 **sha** | **optional.String**|  | 

### Return type

[**[]ModelPackage**](Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilityPackages**
> []ModelPackage GetVulnerabilityPackages(ctx, cVEID)
List packages that contain the given CVE id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cVEID** | **string**|  | 

### Return type

[**[]ModelPackage**](Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

