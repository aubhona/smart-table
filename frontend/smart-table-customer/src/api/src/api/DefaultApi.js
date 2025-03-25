/**
 * SmartTable Mobile API
 * API для управления заказами.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import CustomerV1OrderCreateRequest from '../model/CustomerV1OrderCreateRequest';
import CustomerV1OrderCreateResponse from '../model/CustomerV1OrderCreateResponse';
import CustomerV1OrderCustomerSignInRequest from '../model/CustomerV1OrderCustomerSignInRequest';
import CustomerV1OrderCustomerSignInResponse from '../model/CustomerV1OrderCustomerSignInResponse';
import CustomerV1OrderCustomerSignUpResponse from '../model/CustomerV1OrderCustomerSignUpResponse';
import ErrorResponse from '../model/ErrorResponse';

/**
* Default service.
* @module api/DefaultApi
* @version 1.0.0
*/
export default class DefaultApi {

    /**
    * Constructs a new DefaultApi. 
    * @alias module:api/DefaultApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the customerV1OrderCreatePost operation.
     * @callback module:api/DefaultApi~customerV1OrderCreatePostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/CustomerV1OrderCreateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.
     * Создаёт новый заказ
     * @param {module:model/CustomerV1OrderCreateRequest} customerV1OrderCreateRequest 
     * @param {module:api/DefaultApi~customerV1OrderCreatePostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/CustomerV1OrderCreateResponse}
     */
    customerV1OrderCreatePost(customerV1OrderCreateRequest, callback) {
      let postBody = customerV1OrderCreateRequest;
      // verify the required parameter 'customerV1OrderCreateRequest' is set
      if (customerV1OrderCreateRequest === undefined || customerV1OrderCreateRequest === null) {
        throw new Error("Missing the required parameter 'customerV1OrderCreateRequest' when calling customerV1OrderCreatePost");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = CustomerV1OrderCreateResponse;
      return this.apiClient.callApi(
        '/customer/v1/order/create', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the customerV1OrderCustomerSignInPost operation.
     * @callback module:api/DefaultApi~customerV1OrderCustomerSignInPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/CustomerV1OrderCustomerSignInResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Авторизация пользователя в приложении
     * Авторизует пользователя
     * @param {module:model/CustomerV1OrderCustomerSignInRequest} customerV1OrderCustomerSignInRequest 
     * @param {module:api/DefaultApi~customerV1OrderCustomerSignInPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/CustomerV1OrderCustomerSignInResponse}
     */
    customerV1OrderCustomerSignInPost(customerV1OrderCustomerSignInRequest, callback) {
      let postBody = customerV1OrderCustomerSignInRequest;
      // verify the required parameter 'customerV1OrderCustomerSignInRequest' is set
      if (customerV1OrderCustomerSignInRequest === undefined || customerV1OrderCustomerSignInRequest === null) {
        throw new Error("Missing the required parameter 'customerV1OrderCustomerSignInRequest' when calling customerV1OrderCustomerSignInPost");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = CustomerV1OrderCustomerSignInResponse;
      return this.apiClient.callApi(
        '/customer/v1/order/customer/sign-in', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the customerV1OrderCustomerSignUpPost operation.
     * @callback module:api/DefaultApi~customerV1OrderCustomerSignUpPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/CustomerV1OrderCustomerSignUpResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Регистрация пользователя в приложении
     * Регистрирует пользователя
     * @param {String} tgId 
     * @param {String} tgLogin 
     * @param {String} chatId 
     * @param {File} avatar 
     * @param {module:api/DefaultApi~customerV1OrderCustomerSignUpPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/CustomerV1OrderCustomerSignUpResponse}
     */
    customerV1OrderCustomerSignUpPost(tgId, tgLogin, chatId, avatar, callback) {
      let postBody = null;
      // verify the required parameter 'tgId' is set
      if (tgId === undefined || tgId === null) {
        throw new Error("Missing the required parameter 'tgId' when calling customerV1OrderCustomerSignUpPost");
      }
      // verify the required parameter 'tgLogin' is set
      if (tgLogin === undefined || tgLogin === null) {
        throw new Error("Missing the required parameter 'tgLogin' when calling customerV1OrderCustomerSignUpPost");
      }
      // verify the required parameter 'chatId' is set
      if (chatId === undefined || chatId === null) {
        throw new Error("Missing the required parameter 'chatId' when calling customerV1OrderCustomerSignUpPost");
      }
      // verify the required parameter 'avatar' is set
      if (avatar === undefined || avatar === null) {
        throw new Error("Missing the required parameter 'avatar' when calling customerV1OrderCustomerSignUpPost");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
        'tg_id': tgId,
        'tg_login': tgLogin,
        'chat_id': chatId,
        'avatar': avatar
      };

      let authNames = [];
      let contentTypes = ['multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = CustomerV1OrderCustomerSignUpResponse;
      return this.apiClient.callApi(
        '/customer/v1/order/customer/sign-up', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
