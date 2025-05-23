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
 * The MenuDishItem model module.
 * @module model/MenuDishItem
 * @version 1.0.0
 */
class MenuDishItem {
    /**
     * Constructs a new <code>MenuDishItem</code>.
     * @alias module:model/MenuDishItem
     * @param id {String} 
     * @param price {String} 
     * @param name {String} 
     * @param category {String} 
     * @param calories {Number} 
     * @param weight {Number} 
     * @param count {Number} 
     */
    constructor(id, price, name, category, calories, weight, count) { 
        
        MenuDishItem.initialize(this, id, price, name, category, calories, weight, count);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, price, name, category, calories, weight, count) { 
        obj['id'] = id;
        obj['price'] = price;
        obj['name'] = name;
        obj['category'] = category;
        obj['calories'] = calories;
        obj['weight'] = weight;
        obj['count'] = count;
    }

    /**
     * Constructs a <code>MenuDishItem</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MenuDishItem} obj Optional instance to populate.
     * @return {module:model/MenuDishItem} The populated <code>MenuDishItem</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MenuDishItem();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('calories')) {
                obj['calories'] = ApiClient.convertToType(data['calories'], 'Number');
            }
            if (data.hasOwnProperty('weight')) {
                obj['weight'] = ApiClient.convertToType(data['weight'], 'Number');
            }
            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>MenuDishItem</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>MenuDishItem</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of MenuDishItem.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['id'] && !(typeof data['id'] === 'string' || data['id'] instanceof String)) {
            throw new Error("Expected the field `id` to be a primitive type in the JSON string but got " + data['id']);
        }
        // ensure the json data is a string
        if (data['price'] && !(typeof data['price'] === 'string' || data['price'] instanceof String)) {
            throw new Error("Expected the field `price` to be a primitive type in the JSON string but got " + data['price']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['category'] && !(typeof data['category'] === 'string' || data['category'] instanceof String)) {
            throw new Error("Expected the field `category` to be a primitive type in the JSON string but got " + data['category']);
        }

        return true;
    }


}

MenuDishItem.RequiredProperties = ["id", "price", "name", "category", "calories", "weight", "count"];

/**
 * @member {String} id
 */
MenuDishItem.prototype['id'] = undefined;

/**
 * @member {String} price
 */
MenuDishItem.prototype['price'] = undefined;

/**
 * @member {String} name
 */
MenuDishItem.prototype['name'] = undefined;

/**
 * @member {String} category
 */
MenuDishItem.prototype['category'] = undefined;

/**
 * @member {Number} calories
 */
MenuDishItem.prototype['calories'] = undefined;

/**
 * @member {Number} weight
 */
MenuDishItem.prototype['weight'] = undefined;

/**
 * @member {Number} count
 */
MenuDishItem.prototype['count'] = undefined;






export default MenuDishItem;

