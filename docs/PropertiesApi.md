# \PropertiesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommitHostedPropertyValue**](PropertiesApi.md#DeleteCommitHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**DeletePullRequestHostedPropertyValue**](PropertiesApi.md#DeletePullRequestHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**DeleteRepositoryHostedPropertyValue**](PropertiesApi.md#DeleteRepositoryHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**DeleteUserHostedPropertyValue**](PropertiesApi.md#DeleteUserHostedPropertyValue) | **Delete** /users/{username}/properties/{app_key}/{property_name} | 
[**GetCommitHostedPropertyValue**](PropertiesApi.md#GetCommitHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**GetPullRequestHostedPropertyValue**](PropertiesApi.md#GetPullRequestHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**GetRepositoryHostedPropertyValue**](PropertiesApi.md#GetRepositoryHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**RetrieveUserHostedPropertyValue**](PropertiesApi.md#RetrieveUserHostedPropertyValue) | **Get** /users/{username}/properties/{app_key}/{property_name} | 
[**UpdateCommitHostedPropertyValue**](PropertiesApi.md#UpdateCommitHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**UpdatePullRequestHostedPropertyValue**](PropertiesApi.md#UpdatePullRequestHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**UpdateRepositoryHostedPropertyValue**](PropertiesApi.md#UpdateRepositoryHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**UpdateUserHostedPropertyValue**](PropertiesApi.md#UpdateUserHostedPropertyValue) | **Put** /users/{username}/properties/{app_key}/{property_name} | 


# **DeleteCommitHostedPropertyValue**
> DeleteCommitHostedPropertyValue(ctx, username, repoSlug, commit, appKey, propertyName)


Delete an application property value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePullRequestHostedPropertyValue**
> DeletePullRequestHostedPropertyValue(ctx, username, repoSlug, pullrequestId, appKey, propertyName)


Delete an application property value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryHostedPropertyValue**
> DeleteRepositoryHostedPropertyValue(ctx, username, repoSlug, appKey, propertyName)


Delete an application property value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserHostedPropertyValue**
> DeleteUserHostedPropertyValue(ctx, username, appKey, propertyName)


Delete an application property value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The user. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitHostedPropertyValue**
> GetCommitHostedPropertyValue(ctx, username, repoSlug, commit, appKey, propertyName)


Retrieve an application property value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestHostedPropertyValue**
> GetPullRequestHostedPropertyValue(ctx, username, repoSlug, pullrequestId, appKey, propertyName)


Retrieve an application property value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHostedPropertyValue**
> GetRepositoryHostedPropertyValue(ctx, username, repoSlug, appKey, propertyName)


Retrieve an application property value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveUserHostedPropertyValue**
> RetrieveUserHostedPropertyValue(ctx, username, appKey, propertyName)


Retrieve an application property value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The user. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommitHostedPropertyValue**
> UpdateCommitHostedPropertyValue(ctx, username, repoSlug, commit, appKey, propertyName)


Update an application property value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullRequestHostedPropertyValue**
> UpdatePullRequestHostedPropertyValue(ctx, username, repoSlug, pullrequestId, appKey, propertyName)


Update an application property value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryHostedPropertyValue**
> UpdateRepositoryHostedPropertyValue(ctx, username, repoSlug, appKey, propertyName)


Update an application property value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserHostedPropertyValue**
> UpdateUserHostedPropertyValue(ctx, username, appKey, propertyName)


Update an application property value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The user. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

