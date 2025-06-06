/**
 * SmartTable Mobile API
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
    factory(root.expect, root.SmartTableMobileApi);
  }
}(this, function(expect, SmartTableMobileApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new SmartTableMobileApi.PlaceInfo();
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

  describe('PlaceInfo', function() {
    it('should create an instance of PlaceInfo', function() {
      // uncomment below and update the code to test PlaceInfo
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be.a(SmartTableMobileApi.PlaceInfo);
    });

    it('should have the property uuid (base name: "uuid")', function() {
      // uncomment below and update the code to test the property uuid
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be();
    });

    it('should have the property address (base name: "address")', function() {
      // uncomment below and update the code to test the property address
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be();
    });

    it('should have the property tableCount (base name: "table_count")', function() {
      // uncomment below and update the code to test the property tableCount
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be();
    });

    it('should have the property openingTime (base name: "opening_time")', function() {
      // uncomment below and update the code to test the property openingTime
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be();
    });

    it('should have the property closingTime (base name: "closing_time")', function() {
      // uncomment below and update the code to test the property closingTime
      //var instance = new SmartTableMobileApi.PlaceInfo();
      //expect(instance).to.be();
    });

  });

}));
