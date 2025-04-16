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

import ApiClient from '../ApiClient';

/**
 * The AdminV1UserSignUp409ErrorResponse model module.
 * @module model/AdminV1UserSignUp409ErrorResponse
 * @version 1.0.0
 */
class AdminV1UserSignUp409ErrorResponse {
    /**
     * Constructs a new <code>AdminV1UserSignUp409ErrorResponse</code>.
     * @alias module:model/AdminV1UserSignUp409ErrorResponse
     * @param code {module:model/AdminV1UserSignUp409ErrorResponse.CodeEnum} Код ошибки
     * @param message {String} Описание ошибки
     */
    constructor(code, message) { 
        
        AdminV1UserSignUp409ErrorResponse.initialize(this, code, message);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, code, message) { 
        obj['code'] = code;
        obj['message'] = message;
    }

    /**
     * Constructs a <code>AdminV1UserSignUp409ErrorResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AdminV1UserSignUp409ErrorResponse} obj Optional instance to populate.
     * @return {module:model/AdminV1UserSignUp409ErrorResponse} The populated <code>AdminV1UserSignUp409ErrorResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AdminV1UserSignUp409ErrorResponse();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AdminV1UserSignUp409ErrorResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AdminV1UserSignUp409ErrorResponse</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of AdminV1UserSignUp409ErrorResponse.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['code'] && !(typeof data['code'] === 'string' || data['code'] instanceof String)) {
            throw new Error("Expected the field `code` to be a primitive type in the JSON string but got " + data['code']);
        }
        // ensure the json data is a string
        if (data['message'] && !(typeof data['message'] === 'string' || data['message'] instanceof String)) {
            throw new Error("Expected the field `message` to be a primitive type in the JSON string but got " + data['message']);
        }

        return true;
    }


}

AdminV1UserSignUp409ErrorResponse.RequiredProperties = ["code", "message"];

/**
 * Код ошибки
 * @member {module:model/AdminV1UserSignUp409ErrorResponse.CodeEnum} code
 */
AdminV1UserSignUp409ErrorResponse.prototype['code'] = undefined;

/**
 * Описание ошибки
 * @member {String} message
 */
AdminV1UserSignUp409ErrorResponse.prototype['message'] = undefined;





/**
 * Allowed values for the <code>code</code> property.
 * @enum {String}
 * @readonly
 */
AdminV1UserSignUp409ErrorResponse['CodeEnum'] = {

    /**
     * value: "already_exist"
     * @const
     */
    "already_exist": "already_exist"
};



export default AdminV1UserSignUp409ErrorResponse;

