# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**customerV1SignInPost**](DefaultApi.md#customerV1SignInPost) | **POST** /customer/v1/sign-in | Авторизация пользователя в приложении
[**customerV1SignUpPost**](DefaultApi.md#customerV1SignUpPost) | **POST** /customer/v1/sign-up | Регистрация пользователя в приложении



## customerV1SignInPost

> CustomerV1OrderCustomerSignInResponse customerV1SignInPost(customerV1OrderCustomerSignInRequest)

Авторизация пользователя в приложении

Авторизует пользователя

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerV1OrderCustomerSignInRequest = new SmartTableMobileApi.CustomerV1OrderCustomerSignInRequest(); // CustomerV1OrderCustomerSignInRequest | 
apiInstance.customerV1SignInPost(customerV1OrderCustomerSignInRequest, (error, data, response) => {
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


## customerV1SignUpPost

> CustomerV1OrderCustomerSignUpResponse customerV1SignUpPost(customerV1OrderCustomerSignUpRequest)

Регистрация пользователя в приложении

Регистрирует пользователя

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerV1OrderCustomerSignUpRequest = new SmartTableMobileApi.CustomerV1OrderCustomerSignUpRequest(); // CustomerV1OrderCustomerSignUpRequest | 
apiInstance.customerV1SignUpPost(customerV1OrderCustomerSignUpRequest, (error, data, response) => {
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
 **customerV1OrderCustomerSignUpRequest** | [**CustomerV1OrderCustomerSignUpRequest**](CustomerV1OrderCustomerSignUpRequest.md)|  | 

### Return type

[**CustomerV1OrderCustomerSignUpResponse**](CustomerV1OrderCustomerSignUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

