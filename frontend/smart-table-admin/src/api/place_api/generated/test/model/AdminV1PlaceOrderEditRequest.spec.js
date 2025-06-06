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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.SmartTableAdminApi);
  }
}(this, function(expect, SmartTableAdminApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('AdminV1PlaceOrderEditRequest', function() {
    it('should create an instance of AdminV1PlaceOrderEditRequest', function() {
      // uncomment below and update the code to test AdminV1PlaceOrderEditRequest
      //var instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
      //expect(instance).to.be.a(SmartTableAdminApi.AdminV1PlaceOrderEditRequest);
    });

    it('should have the property orderUuid (base name: "order_uuid")', function() {
      // uncomment below and update the code to test the property orderUuid
      //var instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
      //expect(instance).to.be();
    });

    it('should have the property orderStatus (base name: "order_status")', function() {
      // uncomment below and update the code to test the property orderStatus
      //var instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
      //expect(instance).to.be();
    });

    it('should have the property itemUuid (base name: "item_uuid")', function() {
      // uncomment below and update the code to test the property itemUuid
      //var instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
      //expect(instance).to.be();
    });

    it('should have the property itemStatus (base name: "item_status")', function() {
      // uncomment below and update the code to test the property itemStatus
      //var instance = new SmartTableAdminApi.AdminV1PlaceOrderEditRequest();
      //expect(instance).to.be();
    });

  });

}));
