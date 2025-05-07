# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1RestaurantCreatePost**](DefaultApi.md#adminV1RestaurantCreatePost) | **POST** /admin/v1/restaurant/create | Создание ресторана
[**adminV1RestaurantDishCreatePost**](DefaultApi.md#adminV1RestaurantDishCreatePost) | **POST** /admin/v1/restaurant/dish/create | Создание блюда ресторана
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


## adminV1RestaurantDishCreatePost

> AdminV1RestaurantDishCreateResponse adminV1RestaurantDishCreatePost(userUUID, restaurantUuid, dishName, description, category, calories, weight, dishPictureFile)

Создание блюда ресторана

Создание блюда ресторана

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
let restaurantUuid = "restaurantUuid_example"; // String | Уникальный идентификатор ресторана
let dishName = "dishName_example"; // String | 
let description = "description_example"; // String | 
let category = "category_example"; // String | 
let calories = 56; // Number | 
let weight = 56; // Number | 
let dishPictureFile = "/path/to/file"; // File | 
apiInstance.adminV1RestaurantDishCreatePost(userUUID, restaurantUuid, dishName, description, category, calories, weight, dishPictureFile, (error, data, response) => {
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

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
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

