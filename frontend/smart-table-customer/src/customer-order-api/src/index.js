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


import ApiClient from './ApiClient';
import CustomerV1OrderCreateRequest from './model/CustomerV1OrderCreateRequest';
import CustomerV1OrderCreateResponse from './model/CustomerV1OrderCreateResponse';
import ErrorResponse from './model/ErrorResponse';
import DefaultApi from './api/DefaultApi';


/**
* API для управления заказами..<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var SmartTableMobileApi = require('index'); // See note below*.
* var xxxSvc = new SmartTableMobileApi.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new SmartTableMobileApi.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new SmartTableMobileApi.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new SmartTableMobileApi.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 1.0.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The CustomerV1OrderCreateRequest model constructor.
     * @property {module:model/CustomerV1OrderCreateRequest}
     */
    CustomerV1OrderCreateRequest,

    /**
     * The CustomerV1OrderCreateResponse model constructor.
     * @property {module:model/CustomerV1OrderCreateResponse}
     */
    CustomerV1OrderCreateResponse,

    /**
     * The ErrorResponse model constructor.
     * @property {module:model/ErrorResponse}
     */
    ErrorResponse,

    /**
    * The DefaultApi service constructor.
    * @property {module:api/DefaultApi}
    */
    DefaultApi
};
