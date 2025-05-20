# SmartTableMobileApi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**customerV1OrderCartGet**](DefaultApi.md#customerV1OrderCartGet) | **GET** /customer/v1/order/cart | Получить корзину
[**customerV1OrderCatalogGet**](DefaultApi.md#customerV1OrderCatalogGet) | **GET** /customer/v1/order/catalog | Получение каталога
[**customerV1OrderCatalogUpdatedInfoGet**](DefaultApi.md#customerV1OrderCatalogUpdatedInfoGet) | **GET** /customer/v1/order/catalog/updated-info | Получить обновленную информацию по каталогу
[**customerV1OrderCreatePost**](DefaultApi.md#customerV1OrderCreatePost) | **POST** /customer/v1/order/create | Создаёт новый заказ
[**customerV1OrderCustomerListGet**](DefaultApi.md#customerV1OrderCustomerListGet) | **GET** /customer/v1/order/customer/list | Получение списка пользователей заказа
[**customerV1OrderFinishPost**](DefaultApi.md#customerV1OrderFinishPost) | **POST** /customer/v1/order/finish | Запрос на завершение заказа
[**customerV1OrderItemStatePost**](DefaultApi.md#customerV1OrderItemStatePost) | **POST** /customer/v1/order/item/state | Получить карточку блюда
[**customerV1OrderItemsCommitPost**](DefaultApi.md#customerV1OrderItemsCommitPost) | **POST** /customer/v1/order/items/commit | Добавить блюда к чеку
[**customerV1OrderItemsDraftCountEditPost**](DefaultApi.md#customerV1OrderItemsDraftCountEditPost) | **POST** /customer/v1/order/items/draft/count/edit | Изменяет количество блюд в корзине.
[**customerV1OrderTipSavePost**](DefaultApi.md#customerV1OrderTipSavePost) | **POST** /customer/v1/order/tip/save | Сохранение чека



## customerV1OrderCartGet

> File customerV1OrderCartGet(customerUUID, jWTToken, orderUUID)

Получить корзину

Возвращает подробную информацию по текущей корзине пользователя в заказе

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderCartGet(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/mixed, application/json


## customerV1OrderCatalogGet

> File customerV1OrderCatalogGet(customerUUID, jWTToken, orderUUID)

Получение каталога

Отображение каталога плейса

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderCatalogGet(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/mixed, application/json


## customerV1OrderCatalogUpdatedInfoGet

> CustomerV1OrderCatalogUpdatedInfoResponse customerV1OrderCatalogUpdatedInfoGet(customerUUID, jWTToken, orderUUID)

Получить обновленную информацию по каталогу

Возвращает обновленную стоимость корзины, и количество блюд из меню, которые добавлены в корзину

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderCatalogUpdatedInfoGet(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

[**CustomerV1OrderCatalogUpdatedInfoResponse**](CustomerV1OrderCatalogUpdatedInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## customerV1OrderCreatePost

> CustomerV1OrderCreateResponse customerV1OrderCreatePost(customerUUID, jWTToken, customerV1OrderCreateRequest)

Создаёт новый заказ

Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let customerV1OrderCreateRequest = new SmartTableMobileApi.CustomerV1OrderCreateRequest(); // CustomerV1OrderCreateRequest | 
apiInstance.customerV1OrderCreatePost(customerUUID, jWTToken, customerV1OrderCreateRequest, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **customerV1OrderCreateRequest** | [**CustomerV1OrderCreateRequest**](CustomerV1OrderCreateRequest.md)|  | 

### Return type

[**CustomerV1OrderCreateResponse**](CustomerV1OrderCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## customerV1OrderCustomerListGet

> CustomerV1OrderCustomerListResponse customerV1OrderCustomerListGet(customerUUID, jWTToken, orderUUID)

Получение списка пользователей заказа

Получение списка пользователей заказа

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderCustomerListGet(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

[**CustomerV1OrderCustomerListResponse**](CustomerV1OrderCustomerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## customerV1OrderFinishPost

> customerV1OrderFinishPost(customerUUID, jWTToken, orderUUID)

Запрос на завершение заказа

Переводит заказ в статус ожидает оплаты

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderFinishPost(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## customerV1OrderItemStatePost

> File customerV1OrderItemStatePost(customerUUID, jWTToken, orderUUID, customerV1OrderItemsStateRequest)

Получить карточку блюда

Возвращает подробную информацию по выбранному блюду

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
let customerV1OrderItemsStateRequest = new SmartTableMobileApi.CustomerV1OrderItemsStateRequest(); // CustomerV1OrderItemsStateRequest | 
apiInstance.customerV1OrderItemStatePost(customerUUID, jWTToken, orderUUID, customerV1OrderItemsStateRequest, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 
 **customerV1OrderItemsStateRequest** | [**CustomerV1OrderItemsStateRequest**](CustomerV1OrderItemsStateRequest.md)|  | 

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: multipart/mixed, application/json


## customerV1OrderItemsCommitPost

> customerV1OrderItemsCommitPost(customerUUID, jWTToken, orderUUID)

Добавить блюда к чеку

Добавляет блюда из корзины в чек текущего заказа

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderItemsCommitPost(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## customerV1OrderItemsDraftCountEditPost

> customerV1OrderItemsDraftCountEditPost(customerUUID, jWTToken, orderUUID, customerV1OrderItemsDraftCountEditRequest)

Изменяет количество блюд в корзине.

Изменяет количество блюд в корзине. Передается число в запросе. Знак числа определяет добавить или удалить блюда.s 

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
let customerV1OrderItemsDraftCountEditRequest = new SmartTableMobileApi.CustomerV1OrderItemsDraftCountEditRequest(); // CustomerV1OrderItemsDraftCountEditRequest | 
apiInstance.customerV1OrderItemsDraftCountEditPost(customerUUID, jWTToken, orderUUID, customerV1OrderItemsDraftCountEditRequest, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 
 **customerV1OrderItemsDraftCountEditRequest** | [**CustomerV1OrderItemsDraftCountEditRequest**](CustomerV1OrderItemsDraftCountEditRequest.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## customerV1OrderTipSavePost

> customerV1OrderTipSavePost(customerUUID, jWTToken, orderUUID)

Сохранение чека

Сохраняет чек заказа

### Example

```javascript
import SmartTableMobileApi from 'smart_table_mobile_api';

let apiInstance = new SmartTableMobileApi.DefaultApi();
let customerUUID = "customerUUID_example"; // String | Уникальный идентификатор пользователя
let jWTToken = "jWTToken_example"; // String | jwt токен пользователя
let orderUUID = "orderUUID_example"; // String | Уникальный идентификатор заказа
apiInstance.customerV1OrderTipSavePost(customerUUID, jWTToken, orderUUID, (error, data, response) => {
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
 **customerUUID** | **String**| Уникальный идентификатор пользователя | 
 **jWTToken** | **String**| jwt токен пользователя | 
 **orderUUID** | **String**| Уникальный идентификатор заказа | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

