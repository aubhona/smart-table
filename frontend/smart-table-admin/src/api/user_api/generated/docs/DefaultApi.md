# SmartTableAdminApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**adminV1UserSignInPost**](DefaultApi.md#adminV1UserSignInPost) | **POST** /admin/v1/user/sign-in | Авторизация пользователя в админке
[**adminV1UserSignUpPost**](DefaultApi.md#adminV1UserSignUpPost) | **POST** /admin/v1/user/sign-up | Регистрация пользователя в админке



## adminV1UserSignInPost

> AdminV1UserSignInResponse adminV1UserSignInPost(adminV1UserSignInRequest)

Авторизация пользователя в админке

Авторизует пользователя в админке

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let adminV1UserSignInRequest = new SmartTableAdminApi.AdminV1UserSignInRequest(); // AdminV1UserSignInRequest | 
apiInstance.adminV1UserSignInPost(adminV1UserSignInRequest, (error, data, response) => {
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
 **adminV1UserSignInRequest** | [**AdminV1UserSignInRequest**](AdminV1UserSignInRequest.md)|  | 

### Return type

[**AdminV1UserSignInResponse**](AdminV1UserSignInResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## adminV1UserSignUpPost

> AdminV1UserSignUpResponse adminV1UserSignUpPost(adminV1UserSignUpRequest)

Регистрация пользователя в админке

Регистрирует пользователя в админке

### Example

```javascript
import SmartTableAdminApi from 'smart_table_admin_api';

let apiInstance = new SmartTableAdminApi.DefaultApi();
let adminV1UserSignUpRequest = new SmartTableAdminApi.AdminV1UserSignUpRequest(); // AdminV1UserSignUpRequest | 
apiInstance.adminV1UserSignUpPost(adminV1UserSignUpRequest, (error, data, response) => {
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
 **adminV1UserSignUpRequest** | [**AdminV1UserSignUpRequest**](AdminV1UserSignUpRequest.md)|  | 

### Return type

[**AdminV1UserSignUpResponse**](AdminV1UserSignUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

