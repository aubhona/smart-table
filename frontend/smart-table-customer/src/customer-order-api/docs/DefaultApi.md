# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**customerV1OrderCreatePost**](DefaultApi.md#customerV1OrderCreatePost) | **POST** /customer/v1/order/create | Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.



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

