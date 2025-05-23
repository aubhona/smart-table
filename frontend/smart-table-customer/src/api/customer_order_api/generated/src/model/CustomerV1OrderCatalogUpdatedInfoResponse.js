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
import MenuDishItemUpdatedInfo from './MenuDishItemUpdatedInfo';

/**
 * The CustomerV1OrderCatalogUpdatedInfoResponse model module.
 * @module model/CustomerV1OrderCatalogUpdatedInfoResponse
 * @version 1.0.0
 */
class CustomerV1OrderCatalogUpdatedInfoResponse {
    /**
     * Constructs a new <code>CustomerV1OrderCatalogUpdatedInfoResponse</code>.
     * @alias module:model/CustomerV1OrderCatalogUpdatedInfoResponse
     * @param totalPrice {String} 
     * @param menuUpdatedInfo {Array.<module:model/MenuDishItemUpdatedInfo>} 
     */
    constructor(totalPrice, menuUpdatedInfo) { 
        
        CustomerV1OrderCatalogUpdatedInfoResponse.initialize(this, totalPrice, menuUpdatedInfo);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, totalPrice, menuUpdatedInfo) { 
        obj['total_price'] = totalPrice;
        obj['menu_updated_info'] = menuUpdatedInfo;
    }

    /**
     * Constructs a <code>CustomerV1OrderCatalogUpdatedInfoResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CustomerV1OrderCatalogUpdatedInfoResponse} obj Optional instance to populate.
     * @return {module:model/CustomerV1OrderCatalogUpdatedInfoResponse} The populated <code>CustomerV1OrderCatalogUpdatedInfoResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CustomerV1OrderCatalogUpdatedInfoResponse();

            if (data.hasOwnProperty('total_price')) {
                obj['total_price'] = ApiClient.convertToType(data['total_price'], 'String');
            }
            if (data.hasOwnProperty('menu_updated_info')) {
                obj['menu_updated_info'] = ApiClient.convertToType(data['menu_updated_info'], [MenuDishItemUpdatedInfo]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>CustomerV1OrderCatalogUpdatedInfoResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>CustomerV1OrderCatalogUpdatedInfoResponse</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of CustomerV1OrderCatalogUpdatedInfoResponse.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['total_price'] && !(typeof data['total_price'] === 'string' || data['total_price'] instanceof String)) {
            throw new Error("Expected the field `total_price` to be a primitive type in the JSON string but got " + data['total_price']);
        }
        if (data['menu_updated_info']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['menu_updated_info'])) {
                throw new Error("Expected the field `menu_updated_info` to be an array in the JSON data but got " + data['menu_updated_info']);
            }
            // validate the optional field `menu_updated_info` (array)
            for (const item of data['menu_updated_info']) {
                MenuDishItemUpdatedInfo.validateJSON(item);
            };
        }

        return true;
    }


}

CustomerV1OrderCatalogUpdatedInfoResponse.RequiredProperties = ["total_price", "menu_updated_info"];

/**
 * @member {String} total_price
 */
CustomerV1OrderCatalogUpdatedInfoResponse.prototype['total_price'] = undefined;

/**
 * @member {Array.<module:model/MenuDishItemUpdatedInfo>} menu_updated_info
 */
CustomerV1OrderCatalogUpdatedInfoResponse.prototype['menu_updated_info'] = undefined;






export default CustomerV1OrderCatalogUpdatedInfoResponse;

