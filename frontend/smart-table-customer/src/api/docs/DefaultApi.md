# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**customerV1OrderCreatePost**](DefaultApi.md#customerV1OrderCreatePost) | **POST** /customer/v1/order/create | Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.
[**customerV1OrderCustomerSignInPost**](DefaultApi.md#customerV1OrderCustomerSignInPost) | **POST** /customer/v1/order/customer/sign-in | Авторизация пользователя в приложении
[**customerV1OrderCustomerSignUpPost**](DefaultApi.md#customerV1OrderCustomerSignUpPost) | **POST** /customer/v1/order/customer/sign-up | Регистрация пользователя в приложении



## customerV1OrderCreatePost

> CustomerV1OrderCreateResponse customerV1OrderCreatePost(customerV1OrderCreateRequest)

Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.

Создаёт новый заказ

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerV1OrderCreateRequest = new SmartTableMobileApi.CustomerV1OrderCreateRequest(); // CustomerV1OrderCreateRequest | 
apiInstance.customerV1OrderCreatePost(customerV1OrderCreateRequest, (error, data, response) => {
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
 **customerV1OrderCreateRequest** | [**CustomerV1OrderCreateRequest**](CustomerV1OrderCreateRequest.md)|  | 

### Return type

[**CustomerV1OrderCreateResponse**](CustomerV1OrderCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## customerV1OrderCustomerSignInPost

> CustomerV1OrderCustomerSignInResponse customerV1OrderCustomerSignInPost(customerV1OrderCustomerSignInRequest)

Авторизация пользователя в приложении

Авторизует пользователя

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerV1OrderCustomerSignInRequest = new SmartTableMobileApi.CustomerV1OrderCustomerSignInRequest(); // CustomerV1OrderCustomerSignInRequest | 
apiInstance.customerV1OrderCustomerSignInPost(customerV1OrderCustomerSignInRequest, (error, data, response) => {
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
 **customerV1OrderCustomerSignInRequest** | [**CustomerV1OrderCustomerSignInRequest**](CustomerV1OrderCustomerSignInRequest.md)|  | 

### Return type

[**CustomerV1OrderCustomerSignInResponse**](CustomerV1OrderCustomerSignInResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## customerV1OrderCustomerSignUpPost

> CustomerV1OrderCustomerSignUpResponse customerV1OrderCustomerSignUpPost(tgId, tgLogin, chatId, avatar)

Регистрация пользователя в приложении

Регистрирует пользователя

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let tgId = "tgId_example"; // String | 
let tgLogin = "tgLogin_example"; // String | 
let chatId = "chatId_example"; // String | 
let avatar = "/path/to/file"; // File | 
apiInstance.customerV1OrderCustomerSignUpPost(tgId, tgLogin, chatId, avatar, (error, data, response) => {
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
 **tgId** | **String**|  | 
 **tgLogin** | **String**|  | 
 **chatId** | **String**|  | 
 **avatar** | **File**|  | 

### Return type

[**CustomerV1OrderCustomerSignUpResponse**](CustomerV1OrderCustomerSignUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

