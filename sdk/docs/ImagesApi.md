# \ImagesApi

All URIs are relative to *http://localhost:8443/metastore/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageReport**](ImagesApi.md#CreateImageReport) | **Post** /images | Create a new image report. Related packages and vulnerabilities are also created.
[**GetImages**](ImagesApi.md#GetImages) | **Get** /images | Search image by id, name or digest .
[**GetPackageImages**](ImagesApi.md#GetPackageImages) | **Get** /packages/{IDorName}/images | List the images that contain the given package.
[**GetVulnerabilityImages**](ImagesApi.md#GetVulnerabilityImages) | **Get** /vulnerabilities/{CVEID}/images | List the images that contain the given vulnerability.


# **CreateImageReport**
> Image CreateImageReport(ctx, image)
Create a new image report. Related packages and vulnerabilities are also created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **image** | [**Image**](Image.md)|  | 

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImages**
> Image GetImages(ctx, )
Search image by id, name or digest .

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackageImages**
> []Image GetPackageImages(ctx, iDorName)
List the images that contain the given package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iDorName** | **string**|  | 

### Return type

[**[]Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilityImages**
> []Image GetVulnerabilityImages(ctx, cVEID)
List the images that contain the given vulnerability.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cVEID** | **string**|  | 

### Return type

[**[]Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

