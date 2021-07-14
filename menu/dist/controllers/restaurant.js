"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _crud = require("../utils/crud");

var _restaurant = require("../models/restaurant");

var _default = (0, _crud.crudControllers)(Item);

exports.default = _default;