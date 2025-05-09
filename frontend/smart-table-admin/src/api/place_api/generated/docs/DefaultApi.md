# SmartTableAdminApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1PlaceCreatePost**](DefaultApi.md#adminV1PlaceCreatePost) | **POST** /admin/v1/place/create | Создание плейса
[**adminV1PlaceEmployeeAddPost**](DefaultApi.md#adminV1PlaceEmployeeAddPost) | **POST** /admin/v1/place/employee/add | Добавление сотрудника в плейс
[**adminV1PlaceEmployeeListPost**](DefaultApi.md#adminV1PlaceEmployeeListPost) | **POST** /admin/v1/place/employee/list | Получение списка сотрудников в плейса
[**adminV1PlaceListPost**](DefaultApi.md#adminV1PlaceListPost) | **POST** /admin/v1/place/list | Получение списка плейсов пользователя
[**adminV1PlaceMenuDishCreatePost**](DefaultApi.md#adminV1PlaceMenuDishCreatePost) | **POST** /admin/v1/place/menu/dish/create | Создание блюда в меню плейса
[**adminV1PlaceMenuDishListPost**](DefaultApi.md#adminV1PlaceMenuDishListPost) | **POST** /admin/v1/place/menu/dish/list | Получение списка позиций в меню плейса
[**adminV1PlaceTableDeeplinksListPost**](DefaultApi.md#adminV1PlaceTableDeeplinksListPost) | **POST** /admin/v1/place/table_deeplinks/list | Получение списка ссылок на столы для генерации QR



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


## adminV1PlaceEmployeeListPost

> AdminV1PlaceEmployeeListResponse adminV1PlaceEmployeeListPost(userUUID, jWTToken, adminV1PlaceEmployeeListRequest)

Получение списка сотрудников в плейса

Получение списка сотрудников в плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceEmployeeListRequest = new SmartTableAdminApi.AdminV1PlaceEmployeeListRequest(); // AdminV1PlaceEmployeeListRequest | 
apiInstance.adminV1PlaceEmployeeListPost(userUUID, jWTToken, adminV1PlaceEmployeeListRequest, (error, data, response) => {
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
 **adminV1PlaceEmployeeListRequest** | [**AdminV1PlaceEmployeeListRequest**](AdminV1PlaceEmployeeListRequest.md)|  | 

### Return type

[**AdminV1PlaceEmployeeListResponse**](AdminV1PlaceEmployeeListResponse.md)

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


## adminV1PlaceMenuDishCreatePost

> AdminV1PlaceMenuDishCreateResponse adminV1PlaceMenuDishCreatePost(userUUID, jWTToken, adminV1PlaceMenuDishCreateRequest)

Создание блюда в меню плейса

Создание блюда в меню плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceMenuDishCreateRequest = new SmartTableAdminApi.AdminV1PlaceMenuDishCreateRequest(); // AdminV1PlaceMenuDishCreateRequest | 
apiInstance.adminV1PlaceMenuDishCreatePost(userUUID, jWTToken, adminV1PlaceMenuDishCreateRequest, (error, data, response) => {
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
 **adminV1PlaceMenuDishCreateRequest** | [**AdminV1PlaceMenuDishCreateRequest**](AdminV1PlaceMenuDishCreateRequest.md)|  | 

### Return type

[**AdminV1PlaceMenuDishCreateResponse**](AdminV1PlaceMenuDishCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceMenuDishListPost

> File adminV1PlaceMenuDishListPost(userUUID, jWTToken, adminV1PlaceMenuDishListRequest)

Получение списка позиций в меню плейса

Получение списка позиций в меню плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceMenuDishListRequest = new SmartTableAdminApi.AdminV1PlaceMenuDishListRequest(); // AdminV1PlaceMenuDishListRequest | 
apiInstance.adminV1PlaceMenuDishListPost(userUUID, jWTToken, adminV1PlaceMenuDishListRequest, (error, data, response) => {
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
 **adminV1PlaceMenuDishListRequest** | [**AdminV1PlaceMenuDishListRequest**](AdminV1PlaceMenuDishListRequest.md)|  | 

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: multipart/mixed, application/json


## adminV1PlaceTableDeeplinksListPost

> AdminV1PlaceTableDeepLinksListResponse adminV1PlaceTableDeeplinksListPost(userUUID, jWTToken, adminV1PlaceTableDeepLinksListRequest)

Получение списка ссылок на столы для генерации QR

Получение списка ссылок на столы для генерации QR

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceTableDeepLinksListRequest = new SmartTableAdminApi.AdminV1PlaceTableDeepLinksListRequest(); // AdminV1PlaceTableDeepLinksListRequest | 
apiInstance.adminV1PlaceTableDeeplinksListPost(userUUID, jWTToken, adminV1PlaceTableDeepLinksListRequest, (error, data, response) => {
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
 **adminV1PlaceTableDeepLinksListRequest** | [**AdminV1PlaceTableDeepLinksListRequest**](AdminV1PlaceTableDeepLinksListRequest.md)|  | 

### Return type

[**AdminV1PlaceTableDeepLinksListResponse**](AdminV1PlaceTableDeepLinksListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

