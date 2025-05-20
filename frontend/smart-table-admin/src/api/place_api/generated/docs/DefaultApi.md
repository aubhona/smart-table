# SmartTableAdminApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1PlaceCreatePost**](DefaultApi.md#adminV1PlaceCreatePost) | **POST** /admin/v1/place/create | Создание плейса
[**adminV1PlaceDeletePost**](DefaultApi.md#adminV1PlaceDeletePost) | **POST** /admin/v1/place/delete | Удаление плейса
[**adminV1PlaceEditPost**](DefaultApi.md#adminV1PlaceEditPost) | **POST** /admin/v1/place/edit | Редактирование плейса
[**adminV1PlaceEmployeeAddPost**](DefaultApi.md#adminV1PlaceEmployeeAddPost) | **POST** /admin/v1/place/employee/add | Добавление сотрудника в плейс
[**adminV1PlaceEmployeeDeletePost**](DefaultApi.md#adminV1PlaceEmployeeDeletePost) | **POST** /admin/v1/place/employee/delete | Удаление сотрудника из плейса
[**adminV1PlaceEmployeeEditPost**](DefaultApi.md#adminV1PlaceEmployeeEditPost) | **POST** /admin/v1/place/employee/edit | Редактирование роли сотрудника плейса
[**adminV1PlaceEmployeeListPost**](DefaultApi.md#adminV1PlaceEmployeeListPost) | **POST** /admin/v1/place/employee/list | Получение списка сотрудников в плейса
[**adminV1PlaceListPost**](DefaultApi.md#adminV1PlaceListPost) | **POST** /admin/v1/place/list | Получение списка плейсов пользователя
[**adminV1PlaceMenuDishCreatePost**](DefaultApi.md#adminV1PlaceMenuDishCreatePost) | **POST** /admin/v1/place/menu/dish/create | Создание блюда в меню плейса
[**adminV1PlaceMenuDishDeletePost**](DefaultApi.md#adminV1PlaceMenuDishDeletePost) | **POST** /admin/v1/place/menu/dish/delete | Удаление блюда в меню плейса
[**adminV1PlaceMenuDishEditPost**](DefaultApi.md#adminV1PlaceMenuDishEditPost) | **POST** /admin/v1/place/menu/dish/edit | Редактирование блюда в меню плейса
[**adminV1PlaceMenuDishListPost**](DefaultApi.md#adminV1PlaceMenuDishListPost) | **POST** /admin/v1/place/menu/dish/list | Получение списка позиций в меню плейса
[**adminV1PlaceOrderEditPost**](DefaultApi.md#adminV1PlaceOrderEditPost) | **POST** /admin/v1/place/order/edit | Редкатирование заказа
[**adminV1PlaceOrderInfoPost**](DefaultApi.md#adminV1PlaceOrderInfoPost) | **POST** /admin/v1/place/order/info | Получение подробной информации о заказе
[**adminV1PlaceOrderListPost**](DefaultApi.md#adminV1PlaceOrderListPost) | **POST** /admin/v1/place/order/list | Получение списка заказов плейса
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


## adminV1PlaceDeletePost

> adminV1PlaceDeletePost(userUUID, jWTToken, adminV1PlaceDeleteRequest)

Удаление плейса

Удаление плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceDeleteRequest = new SmartTableAdminApi.AdminV1PlaceDeleteRequest(); // AdminV1PlaceDeleteRequest | 
apiInstance.adminV1PlaceDeletePost(userUUID, jWTToken, adminV1PlaceDeleteRequest, (error, data, response) => {
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
 **adminV1PlaceDeleteRequest** | [**AdminV1PlaceDeleteRequest**](AdminV1PlaceDeleteRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceEditPost

> adminV1PlaceEditPost(userUUID, jWTToken, adminV1PlaceEditRequest)

Редактирование плейса

Редактирование плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceEditRequest = new SmartTableAdminApi.AdminV1PlaceEditRequest(); // AdminV1PlaceEditRequest | 
apiInstance.adminV1PlaceEditPost(userUUID, jWTToken, adminV1PlaceEditRequest, (error, data, response) => {
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
 **adminV1PlaceEditRequest** | [**AdminV1PlaceEditRequest**](AdminV1PlaceEditRequest.md)|  | 

### Return type

null (empty response body)

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


## adminV1PlaceEmployeeDeletePost

> adminV1PlaceEmployeeDeletePost(userUUID, jWTToken, adminV1PlaceEmployeeDeleteRequest)

Удаление сотрудника из плейса

Удаление сотрудника из плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceEmployeeDeleteRequest = new SmartTableAdminApi.AdminV1PlaceEmployeeDeleteRequest(); // AdminV1PlaceEmployeeDeleteRequest | 
apiInstance.adminV1PlaceEmployeeDeletePost(userUUID, jWTToken, adminV1PlaceEmployeeDeleteRequest, (error, data, response) => {
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
 **adminV1PlaceEmployeeDeleteRequest** | [**AdminV1PlaceEmployeeDeleteRequest**](AdminV1PlaceEmployeeDeleteRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceEmployeeEditPost

> adminV1PlaceEmployeeEditPost(userUUID, jWTToken, adminV1PlaceEmployeeEditRequest)

Редактирование роли сотрудника плейса

Редактирование роли сотрудника плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceEmployeeEditRequest = new SmartTableAdminApi.AdminV1PlaceEmployeeEditRequest(); // AdminV1PlaceEmployeeEditRequest | 
apiInstance.adminV1PlaceEmployeeEditPost(userUUID, jWTToken, adminV1PlaceEmployeeEditRequest, (error, data, response) => {
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
 **adminV1PlaceEmployeeEditRequest** | [**AdminV1PlaceEmployeeEditRequest**](AdminV1PlaceEmployeeEditRequest.md)|  | 

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


## adminV1PlaceMenuDishDeletePost

> adminV1PlaceMenuDishDeletePost(userUUID, jWTToken, adminV1PlaceMenuDishCreateRequest)

Удаление блюда в меню плейса

Удаление блюда в меню плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceMenuDishCreateRequest = new SmartTableAdminApi.AdminV1PlaceMenuDishCreateRequest(); // AdminV1PlaceMenuDishCreateRequest | 
apiInstance.adminV1PlaceMenuDishDeletePost(userUUID, jWTToken, adminV1PlaceMenuDishCreateRequest, (error, data, response) => {
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
 **adminV1PlaceMenuDishCreateRequest** | [**AdminV1PlaceMenuDishCreateRequest**](AdminV1PlaceMenuDishCreateRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceMenuDishEditPost

> adminV1PlaceMenuDishEditPost(userUUID, jWTToken, adminV1PlaceMenuDishDeleteRequest)

Редактирование блюда в меню плейса

Редактирование блюда в меню плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceMenuDishDeleteRequest = new SmartTableAdminApi.AdminV1PlaceMenuDishDeleteRequest(); // AdminV1PlaceMenuDishDeleteRequest | 
apiInstance.adminV1PlaceMenuDishEditPost(userUUID, jWTToken, adminV1PlaceMenuDishDeleteRequest, (error, data, response) => {
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
 **adminV1PlaceMenuDishDeleteRequest** | [**AdminV1PlaceMenuDishDeleteRequest**](AdminV1PlaceMenuDishDeleteRequest.md)|  | 

### Return type

null (empty response body)

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


## adminV1PlaceOrderEditPost

> adminV1PlaceOrderEditPost(userUUID, jWTToken, adminV1PlaceOrderEditRequest)

Редкатирование заказа

Редкатирование заказа

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceOrderEditRequest = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest(); // AdminV1PlaceOrderEditRequest | 
apiInstance.adminV1PlaceOrderEditPost(userUUID, jWTToken, adminV1PlaceOrderEditRequest, (error, data, response) => {
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
 **adminV1PlaceOrderEditRequest** | [**AdminV1PlaceOrderEditRequest**](AdminV1PlaceOrderEditRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceOrderInfoPost

> AdminV1PlaceOrderInfoResponse adminV1PlaceOrderInfoPost(userUUID, jWTToken, adminV1PlaceOrderInfoRequest)

Получение подробной информации о заказе

Получение подробной информации о заказе

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceOrderInfoRequest = new SmartTableAdminApi.AdminV1PlaceOrderInfoRequest(); // AdminV1PlaceOrderInfoRequest | 
apiInstance.adminV1PlaceOrderInfoPost(userUUID, jWTToken, adminV1PlaceOrderInfoRequest, (error, data, response) => {
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
 **adminV1PlaceOrderInfoRequest** | [**AdminV1PlaceOrderInfoRequest**](AdminV1PlaceOrderInfoRequest.md)|  | 

### Return type

[**AdminV1PlaceOrderInfoResponse**](AdminV1PlaceOrderInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1PlaceOrderListPost

> AdminV1PlaceOrderListResponse adminV1PlaceOrderListPost(userUUID, jWTToken, adminV1PlaceOrderListRequest)

Получение списка заказов плейса

Получение списка заказов плейса

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1PlaceOrderListRequest = new SmartTableAdminApi.AdminV1PlaceOrderListRequest(); // AdminV1PlaceOrderListRequest | 
apiInstance.adminV1PlaceOrderListPost(userUUID, jWTToken, adminV1PlaceOrderListRequest, (error, data, response) => {
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
 **adminV1PlaceOrderListRequest** | [**AdminV1PlaceOrderListRequest**](AdminV1PlaceOrderListRequest.md)|  | 

### Return type

[**AdminV1PlaceOrderListResponse**](AdminV1PlaceOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


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

