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
 * The AdminV1UserSignUpRequest model module.
 * @module model/AdminV1UserSignUpRequest
 * @version 1.0.0
 */
class AdminV1UserSignUpRequest {
    /**
     * Constructs a new <code>AdminV1UserSignUpRequest</code>.
     * @alias module:model/AdminV1UserSignUpRequest
     * @param login {String} Логин пользователя
     * @param tgLogin {String} Логин пользователя
     * @param firstName {String} Имя пользователя на латинице
     * @param lastName {String} Фамилия пользователя на латинице
     * @param password {String} Пароль пользователя
     */
    constructor(login, tgLogin, firstName, lastName, password) { 
        
        AdminV1UserSignUpRequest.initialize(this, login, tgLogin, firstName, lastName, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, login, tgLogin, firstName, lastName, password) { 
        obj['login'] = login;
        obj['tg_login'] = tgLogin;
        obj['first_name'] = firstName;
        obj['last_name'] = lastName;
        obj['password'] = password;
    }

    /**
     * Constructs a <code>AdminV1UserSignUpRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AdminV1UserSignUpRequest} obj Optional instance to populate.
     * @return {module:model/AdminV1UserSignUpRequest} The populated <code>AdminV1UserSignUpRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AdminV1UserSignUpRequest();

            if (data.hasOwnProperty('login')) {
                obj['login'] = ApiClient.convertToType(data['login'], 'String');
            }
            if (data.hasOwnProperty('tg_login')) {
                obj['tg_login'] = ApiClient.convertToType(data['tg_login'], 'String');
            }
            if (data.hasOwnProperty('first_name')) {
                obj['first_name'] = ApiClient.convertToType(data['first_name'], 'String');
            }
            if (data.hasOwnProperty('last_name')) {
                obj['last_name'] = ApiClient.convertToType(data['last_name'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AdminV1UserSignUpRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AdminV1UserSignUpRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of AdminV1UserSignUpRequest.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['login'] && !(typeof data['login'] === 'string' || data['login'] instanceof String)) {
            throw new Error("Expected the field `login` to be a primitive type in the JSON string but got " + data['login']);
        }
        // ensure the json data is a string
        if (data['tg_login'] && !(typeof data['tg_login'] === 'string' || data['tg_login'] instanceof String)) {
            throw new Error("Expected the field `tg_login` to be a primitive type in the JSON string but got " + data['tg_login']);
        }
        // ensure the json data is a string
        if (data['first_name'] && !(typeof data['first_name'] === 'string' || data['first_name'] instanceof String)) {
            throw new Error("Expected the field `first_name` to be a primitive type in the JSON string but got " + data['first_name']);
        }
        // ensure the json data is a string
        if (data['last_name'] && !(typeof data['last_name'] === 'string' || data['last_name'] instanceof String)) {
            throw new Error("Expected the field `last_name` to be a primitive type in the JSON string but got " + data['last_name']);
        }
        // ensure the json data is a string
        if (data['password'] && !(typeof data['password'] === 'string' || data['password'] instanceof String)) {
            throw new Error("Expected the field `password` to be a primitive type in the JSON string but got " + data['password']);
        }

        return true;
    }


}

AdminV1UserSignUpRequest.RequiredProperties = ["login", "tg_login", "first_name", "last_name", "password"];

/**
 * Логин пользователя
 * @member {String} login
 */
AdminV1UserSignUpRequest.prototype['login'] = undefined;

/**
 * Логин пользователя
 * @member {String} tg_login
 */
AdminV1UserSignUpRequest.prototype['tg_login'] = undefined;

/**
 * Имя пользователя на латинице
 * @member {String} first_name
 */
AdminV1UserSignUpRequest.prototype['first_name'] = undefined;

/**
 * Фамилия пользователя на латинице
 * @member {String} last_name
 */
AdminV1UserSignUpRequest.prototype['last_name'] = undefined;

/**
 * Пароль пользователя
 * @member {String} password
 */
AdminV1UserSignUpRequest.prototype['password'] = undefined;






export default AdminV1UserSignUpRequest;

