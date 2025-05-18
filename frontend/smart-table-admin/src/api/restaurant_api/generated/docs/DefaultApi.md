# SmartTableAdminApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1RestaurantCreatePost**](DefaultApi.md#adminV1RestaurantCreatePost) | **POST** /admin/v1/restaurant/create | Создание ресторана
[**adminV1RestaurantDeletePost**](DefaultApi.md#adminV1RestaurantDeletePost) | **POST** /admin/v1/restaurant/delete | Удаление ресторана
[**adminV1RestaurantDishCreatePost**](DefaultApi.md#adminV1RestaurantDishCreatePost) | **POST** /admin/v1/restaurant/dish/create | Создание блюда ресторана
[**adminV1RestaurantDishDeletePost**](DefaultApi.md#adminV1RestaurantDishDeletePost) | **POST** /admin/v1/restaurant/dish/delete | Удаление блюда ресторана
[**adminV1RestaurantDishEditPost**](DefaultApi.md#adminV1RestaurantDishEditPost) | **POST** /admin/v1/restaurant/dish/edit | Редактирование блюда ресторана
[**adminV1RestaurantDishListPost**](DefaultApi.md#adminV1RestaurantDishListPost) | **POST** /admin/v1/restaurant/dish/list | Получение списка блюд ресторана
[**adminV1RestaurantEditPost**](DefaultApi.md#adminV1RestaurantEditPost) | **POST** /admin/v1/restaurant/edit | Редактирование ресторана
[**adminV1RestaurantListGet**](DefaultApi.md#adminV1RestaurantListGet) | **GET** /admin/v1/restaurant/list | Получение списка ресторанов пользователя



## adminV1RestaurantCreatePost

> AdminV1RestaurantCreateResponse adminV1RestaurantCreatePost(userUUID, jWTToken, adminV1RestaurantCreateRequest)

Создание ресторана

Создание ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1RestaurantCreateRequest = new SmartTableAdminApi.AdminV1RestaurantCreateRequest(); // AdminV1RestaurantCreateRequest | 
apiInstance.adminV1RestaurantCreatePost(userUUID, jWTToken, adminV1RestaurantCreateRequest, (error, data, response) => {
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
 **adminV1RestaurantCreateRequest** | [**AdminV1RestaurantCreateRequest**](AdminV1RestaurantCreateRequest.md)|  | 

### Return type

[**AdminV1RestaurantCreateResponse**](AdminV1RestaurantCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1RestaurantDeletePost

> adminV1RestaurantDeletePost(userUUID, jWTToken, adminV1RestaurantDeleteRequest)

Удаление ресторана

Удаление ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1RestaurantDeleteRequest = new SmartTableAdminApi.AdminV1RestaurantDeleteRequest(); // AdminV1RestaurantDeleteRequest | 
apiInstance.adminV1RestaurantDeletePost(userUUID, jWTToken, adminV1RestaurantDeleteRequest, (error, data, response) => {
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
 **adminV1RestaurantDeleteRequest** | [**AdminV1RestaurantDeleteRequest**](AdminV1RestaurantDeleteRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1RestaurantDishCreatePost

> AdminV1RestaurantDishCreateResponse adminV1RestaurantDishCreatePost(userUUID, jWTToken, restaurantUuid, dishName, description, category, calories, weight, dishPictureFile)

Создание блюда ресторана

Создание блюда ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let restaurantUuid = "restaurantUuid_example"; // String | Уникальный идентификатор ресторана
let dishName = "dishName_example"; // String | 
let description = "description_example"; // String | 
let category = "category_example"; // String | 
let calories = 56; // Number | 
let weight = 56; // Number | 
let dishPictureFile = "/path/to/file"; // File | 
apiInstance.adminV1RestaurantDishCreatePost(userUUID, jWTToken, restaurantUuid, dishName, description, category, calories, weight, dishPictureFile, (error, data, response) => {
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
 **restaurantUuid** | **String**| Уникальный идентификатор ресторана | 
 **dishName** | **String**|  | 
 **description** | **String**|  | 
 **category** | **String**|  | 
 **calories** | **Number**|  | 
 **weight** | **Number**|  | 
 **dishPictureFile** | **File**|  | 

### Return type

[**AdminV1RestaurantDishCreateResponse**](AdminV1RestaurantDishCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## adminV1RestaurantDishDeletePost

> adminV1RestaurantDishDeletePost(userUUID, jWTToken, dishUuid)

Удаление блюда ресторана

Удаление блюда ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let dishUuid = "dishUuid_example"; // String | Уникальный идентификатор блюда
apiInstance.adminV1RestaurantDishDeletePost(userUUID, jWTToken, dishUuid, (error, data, response) => {
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
 **dishUuid** | **String**| Уникальный идентификатор блюда | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## adminV1RestaurantDishEditPost

> adminV1RestaurantDishEditPost(userUUID, jWTToken, dishUuid, dishName, description, category, calories, weight, dishPictureFile)

Редактирование блюда ресторана

Редактирование блюда ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let dishUuid = "dishUuid_example"; // String | Уникальный идентификатор блюда
let dishName = "dishName_example"; // String | 
let description = "description_example"; // String | 
let category = "category_example"; // String | 
let calories = 56; // Number | 
let weight = 56; // Number | 
let dishPictureFile = "/path/to/file"; // File | 
apiInstance.adminV1RestaurantDishEditPost(userUUID, jWTToken, dishUuid, dishName, description, category, calories, weight, dishPictureFile, (error, data, response) => {
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
 **dishUuid** | **String**| Уникальный идентификатор блюда | 
 **dishName** | **String**|  | 
 **description** | **String**|  | 
 **category** | **String**|  | 
 **calories** | **Number**|  | 
 **weight** | **Number**|  | 
 **dishPictureFile** | **File**|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## adminV1RestaurantDishListPost

> File adminV1RestaurantDishListPost(userUUID, jWTToken, adminV1RestaurantDishListRequest)

Получение списка блюд ресторана

Получение списка блюд ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1RestaurantDishListRequest = new SmartTableAdminApi.AdminV1RestaurantDishListRequest(); // AdminV1RestaurantDishListRequest | 
apiInstance.adminV1RestaurantDishListPost(userUUID, jWTToken, adminV1RestaurantDishListRequest, (error, data, response) => {
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
 **adminV1RestaurantDishListRequest** | [**AdminV1RestaurantDishListRequest**](AdminV1RestaurantDishListRequest.md)|  | 

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: multipart/mixed, application/json


## adminV1RestaurantEditPost

> adminV1RestaurantEditPost(userUUID, jWTToken, adminV1RestaurantEditRequest)

Редактирование ресторана

Редактирование ресторана

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1RestaurantEditRequest = new SmartTableAdminApi.AdminV1RestaurantEditRequest(); // AdminV1RestaurantEditRequest | 
apiInstance.adminV1RestaurantEditPost(userUUID, jWTToken, adminV1RestaurantEditRequest, (error, data, response) => {
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
 **adminV1RestaurantEditRequest** | [**AdminV1RestaurantEditRequest**](AdminV1RestaurantEditRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1RestaurantListGet

> AdminV1RestaurantListResponse adminV1RestaurantListGet(userUUID, jWTToken)

Получение списка ресторанов пользователя

Получение списка ресторанов пользователя

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
apiInstance.adminV1RestaurantListGet(userUUID, jWTToken, (error, data, response) => {
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

### Return type

[**AdminV1RestaurantListResponse**](AdminV1RestaurantListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

