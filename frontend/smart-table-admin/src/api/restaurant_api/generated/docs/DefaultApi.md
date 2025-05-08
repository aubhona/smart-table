# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1RestaurantCreatePost**](DefaultApi.md#adminV1RestaurantCreatePost) | **POST** /admin/v1/restaurant/create | Создание ресторана
[**adminV1RestaurantDishCreatePost**](DefaultApi.md#adminV1RestaurantDishCreatePost) | **POST** /admin/v1/restaurant/dish/create | Создание блюда ресторана
[**adminV1RestaurantListGet**](DefaultApi.md#adminV1RestaurantListGet) | **GET** /admin/v1/restaurant/list | Получение списка ресторанов пользователя



## adminV1RestaurantCreatePost

> AdminV1RestaurantCreateResponse adminV1RestaurantCreatePost(userUUID, jWTToken, adminV1RestaurantCreateRequest)

Создание ресторана

Создание ресторана

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let userUUID = "userUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let adminV1RestaurantCreateRequest = new SmartTableMobileApi.AdminV1RestaurantCreateRequest(); // AdminV1RestaurantCreateRequest | 
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


## adminV1RestaurantDishCreatePost

> AdminV1RestaurantDishCreateResponse adminV1RestaurantDishCreatePost(userUUID, jWTToken, restaurantUuid, dishName, description, category, calories, weight, dishPictureFile)

Создание блюда ресторана

Создание блюда ресторана

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
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


## adminV1RestaurantListGet

> AdminV1RestaurantListResponse adminV1RestaurantListGet(userUUID, jWTToken)

Получение списка ресторанов пользователя

Получение списка ресторанов пользователя

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
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

