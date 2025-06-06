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
    instance = new SmartTableAdminApi.ItemGroupInfo();
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

  describe('ItemGroupInfo', function() {
    it('should create an instance of ItemGroupInfo', function() {
      // uncomment below and update the code to test ItemGroupInfo
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be.a(SmartTableAdminApi.ItemGroupInfo);
    });

    it('should have the property menuDishUuid (base name: "menu_dish_uuid")', function() {
      // uncomment below and update the code to test the property menuDishUuid
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property itemUuidList (base name: "item_uuid_list")', function() {
      // uncomment below and update the code to test the property itemUuidList
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property comment (base name: "comment")', function() {
      // uncomment below and update the code to test the property comment
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property resolution (base name: "resolution")', function() {
      // uncomment below and update the code to test the property resolution
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property itemPrice (base name: "item_price")', function() {
      // uncomment below and update the code to test the property itemPrice
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property resultPrice (base name: "result_price")', function() {
      // uncomment below and update the code to test the property resultPrice
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

    it('should have the property count (base name: "count")', function() {
      // uncomment below and update the code to test the property count
      //var instance = new SmartTableAdminApi.ItemGroupInfo();
      //expect(instance).to.be();
    });

  });

}));
