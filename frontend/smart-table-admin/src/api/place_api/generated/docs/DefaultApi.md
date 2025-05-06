# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1PlaceCreatePost**](DefaultApi.md#adminV1PlaceCreatePost) | **POST** /admin/v1/place/create | Создание плейса



## adminV1PlaceCreatePost

> AdminV1PlaceCreateResponse adminV1PlaceCreatePost(userUUID, adminV1PlaceCreateRequest)

Создание плейса

Создание плейса

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';
let defaultClient = SmartTableMobileApi.ApiClient.instance;
// Configure API key authorization: CookieAuth
let CookieAuth = defaultClient.authentications['CookieAuth'];
CookieAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CookieAuth.apiKeyPrefix = 'Token';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let adminV1PlaceCreateRequest = new SmartTableMobileApi.AdminV1PlaceCreateRequest(); // AdminV1PlaceCreateRequest | 
apiInstance.adminV1PlaceCreatePost(userUUID, adminV1PlaceCreateRequest, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUUID** | **String**| Уникальный идентификатор пользователя | 
 **adminV1PlaceCreateRequest** | [**AdminV1PlaceCreateRequest**](AdminV1PlaceCreateRequest.md)|  | 

### Return type

[**AdminV1PlaceCreateResponse**](AdminV1PlaceCreateResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

