# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1RestaurantCreatePost**](DefaultApi.md#adminV1RestaurantCreatePost) | **POST** /admin/v1/restaurant/create | Создание ресторана
[**adminV1RestaurantListGet**](DefaultApi.md#adminV1RestaurantListGet) | **GET** /admin/v1/restaurant/list | Получение списка ресторанов пользователя



## adminV1RestaurantCreatePost

> AdminV1RestaurantCreateResponse adminV1RestaurantCreatePost(userUUID, adminV1RestaurantCreateRequest)

Создание ресторана

Создание ресторана

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
let adminV1RestaurantCreateRequest = new SmartTableMobileApi.AdminV1RestaurantCreateRequest(); // AdminV1RestaurantCreateRequest | 
apiInstance.adminV1RestaurantCreatePost(userUUID, adminV1RestaurantCreateRequest, (error, data, response) => {
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
 **adminV1RestaurantCreateRequest** | [**AdminV1RestaurantCreateRequest**](AdminV1RestaurantCreateRequest.md)|  | 

### Return type

[**AdminV1RestaurantCreateResponse**](AdminV1RestaurantCreateResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1RestaurantListGet

> AdminV1RestaurantListResponse adminV1RestaurantListGet(userUUID)

Получение списка ресторанов пользователя

Получение списка ресторанов пользователя

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
apiInstance.adminV1RestaurantListGet(userUUID, (error, data, response) => {
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

### Return type

[**AdminV1RestaurantListResponse**](AdminV1RestaurantListResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

