/**
 * SmartTable Admin API
 * API для управления плейсом.
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
import OrderResolution from './OrderResolution';
import OrderStatus from './OrderStatus';

/**
 * The OrderMainInfo model module.
 * @module model/OrderMainInfo
 * @version 1.0.0
 */
class OrderMainInfo {
    /**
     * Constructs a new <code>OrderMainInfo</code>.
     * @alias module:model/OrderMainInfo
     * @param uuid {String} Уникальный идентификатор заказа
     * @param status {module:model/OrderStatus} 
     * @param tableNumber {Number} Номер стола
     * @param guestsCount {Number} Количество клиентов в заказе
     * @param totalPrice {String} Итоговая стоимость заказа
     * @param createdAt {Date} Дата создания заказа
     */
    constructor(uuid, status, tableNumber, guestsCount, totalPrice, createdAt) { 
        
        OrderMainInfo.initialize(this, uuid, status, tableNumber, guestsCount, totalPrice, createdAt);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, uuid, status, tableNumber, guestsCount, totalPrice, createdAt) { 
        obj['uuid'] = uuid;
        obj['status'] = status;
        obj['table_number'] = tableNumber;
        obj['guests_count'] = guestsCount;
        obj['total_price'] = totalPrice;
        obj['created_at'] = createdAt;
    }

    /**
     * Constructs a <code>OrderMainInfo</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OrderMainInfo} obj Optional instance to populate.
     * @return {module:model/OrderMainInfo} The populated <code>OrderMainInfo</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OrderMainInfo();

            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = OrderStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('resolution')) {
                obj['resolution'] = OrderResolution.constructFromObject(data['resolution']);
            }
            if (data.hasOwnProperty('table_number')) {
                obj['table_number'] = ApiClient.convertToType(data['table_number'], 'Number');
            }
            if (data.hasOwnProperty('guests_count')) {
                obj['guests_count'] = ApiClient.convertToType(data['guests_count'], 'Number');
            }
            if (data.hasOwnProperty('total_price')) {
                obj['total_price'] = ApiClient.convertToType(data['total_price'], 'String');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>OrderMainInfo</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>OrderMainInfo</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of OrderMainInfo.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['uuid'] && !(typeof data['uuid'] === 'string' || data['uuid'] instanceof String)) {
            throw new Error("Expected the field `uuid` to be a primitive type in the JSON string but got " + data['uuid']);
        }
        // ensure the json data is a string
        if (data['total_price'] && !(typeof data['total_price'] === 'string' || data['total_price'] instanceof String)) {
            throw new Error("Expected the field `total_price` to be a primitive type in the JSON string but got " + data['total_price']);
        }

        return true;
    }


}

OrderMainInfo.RequiredProperties = ["uuid", "status", "table_number", "guests_count", "total_price", "created_at"];

/**
 * Уникальный идентификатор заказа
 * @member {String} uuid
 */
OrderMainInfo.prototype['uuid'] = undefined;

/**
 * @member {module:model/OrderStatus} status
 */
OrderMainInfo.prototype['status'] = undefined;

/**
 * @member {module:model/OrderResolution} resolution
 */
OrderMainInfo.prototype['resolution'] = undefined;

/**
 * Номер стола
 * @member {Number} table_number
 */
OrderMainInfo.prototype['table_number'] = undefined;

/**
 * Количество клиентов в заказе
 * @member {Number} guests_count
 */
OrderMainInfo.prototype['guests_count'] = undefined;

/**
 * Итоговая стоимость заказа
 * @member {String} total_price
 */
OrderMainInfo.prototype['total_price'] = undefined;

/**
 * Дата создания заказа
 * @member {Date} created_at
 */
OrderMainInfo.prototype['created_at'] = undefined;






export default OrderMainInfo;

