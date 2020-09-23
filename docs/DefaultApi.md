# \DefaultApi

All URIs are relative to *http://openapi.tuling123.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Chat**](DefaultApi.md#Chat) | **Post** /openapi/api/v2 | 问答



## Chat

> ChatRsp Chat(ctx, body)

问答

问答

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ChatReq**](ChatReq.md)|  | 

### Return type

[**ChatRsp**](ChatRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

