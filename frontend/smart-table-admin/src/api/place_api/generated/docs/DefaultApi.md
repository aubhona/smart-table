# SmartTableAdminApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1PlaceCreatePost**](DefaultApi.md#adminV1PlaceCreatePost) | **POST** /admin/v1/place/create | Создание плейса
[**adminV1PlaceEmployeeAddPost**](DefaultApi.md#adminV1PlaceEmployeeAddPost) | **POST** /admin/v1/place/employee/add | Добавление сотрудника в плейс
[**adminV1PlaceListPost**](DefaultApi.md#adminV1PlaceListPost) | **POST** /admin/v1/place/list | Получение списка плейсов пользователя



## adminV1PlaceCreatePost

> AdminV1PlaceCreateResponse adminV1PlaceCreatePost(userUUID, jWTToken, adminV1PlaceCreateRequest)

Создание плейса

Создание плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceCreateRequest = new SmartTableAdminApi.AdminV1PlaceCreateRequest(); // AdminV1PlaceCreateRequest | 
apiInstance.adminV1PlaceCreatePost(userUUID, jWTToken, adminV1PlaceCreateRequest, (error, data, response) => {
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
 **jWTToken** | **String**| jwt токен пользователя | 
 **adminV1PlaceCreateRequest** | [**AdminV1PlaceCreateRequest**](AdminV1PlaceCreateRequest.md)|  | 

### Return type

[**AdminV1PlaceCreateResponse**](AdminV1PlaceCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceEmployeeAddPost

> adminV1PlaceEmployeeAddPost(userUUID, jWTToken, adminV1PlaceEmployeeAddRequest)

Добавление сотрудника в плейс

Добавление сотрудника в плейс

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceEmployeeAddRequest = new SmartTableAdminApi.AdminV1PlaceEmployeeAddRequest(); // AdminV1PlaceEmployeeAddRequest | 
apiInstance.adminV1PlaceEmployeeAddPost(userUUID, jWTToken, adminV1PlaceEmployeeAddRequest, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **adminV1PlaceEmployeeAddRequest** | [**AdminV1PlaceEmployeeAddRequest**](AdminV1PlaceEmployeeAddRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceListPost

> AdminV1PlaceListResponse adminV1PlaceListPost(userUUID, jWTToken, adminV1PlaceListRequest)

Получение списка плейсов пользователя

Получение списка плейсов пользователя

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceListRequest = new SmartTableAdminApi.AdminV1PlaceListRequest(); // AdminV1PlaceListRequest | 
apiInstance.adminV1PlaceListPost(userUUID, jWTToken, adminV1PlaceListRequest, (error, data, response) => {
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
 **jWTToken** | **String**| jwt токен пользователя | 
 **adminV1PlaceListRequest** | [**AdminV1PlaceListRequest**](AdminV1PlaceListRequest.md)|  | 

### Return type

[**AdminV1PlaceListResponse**](AdminV1PlaceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

