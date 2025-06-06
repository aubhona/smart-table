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

/**
 * The AdminV1PlaceTableDeepLinksListRequest model module.
 * @module model/AdminV1PlaceTableDeepLinksListRequest
 * @version 1.0.0
 */
class AdminV1PlaceTableDeepLinksListRequest {
    /**
     * Constructs a new <code>AdminV1PlaceTableDeepLinksListRequest</code>.
     * @alias module:model/AdminV1PlaceTableDeepLinksListRequest
     * @param placeUuid {String} Уникальный идентификатор плейса
     */
    constructor(placeUuid) { 
        
        AdminV1PlaceTableDeepLinksListRequest.initialize(this, placeUuid);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, placeUuid) { 
        obj['place_uuid'] = placeUuid;
    }

    /**
     * Constructs a <code>AdminV1PlaceTableDeepLinksListRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AdminV1PlaceTableDeepLinksListRequest} obj Optional instance to populate.
     * @return {module:model/AdminV1PlaceTableDeepLinksListRequest} The populated <code>AdminV1PlaceTableDeepLinksListRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AdminV1PlaceTableDeepLinksListRequest();

            if (data.hasOwnProperty('place_uuid')) {
                obj['place_uuid'] = ApiClient.convertToType(data['place_uuid'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AdminV1PlaceTableDeepLinksListRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AdminV1PlaceTableDeepLinksListRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of AdminV1PlaceTableDeepLinksListRequest.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['place_uuid'] && !(typeof data['place_uuid'] === 'string' || data['place_uuid'] instanceof String)) {
            throw new Error("Expected the field `place_uuid` to be a primitive type in the JSON string but got " + data['place_uuid']);
        }

        return true;
    }


}

AdminV1PlaceTableDeepLinksListRequest.RequiredProperties = ["place_uuid"];

/**
 * Уникальный идентификатор плейса
 * @member {String} place_uuid
 */
AdminV1PlaceTableDeepLinksListRequest.prototype['place_uuid'] = undefined;






export default AdminV1PlaceTableDeepLinksListRequest;

