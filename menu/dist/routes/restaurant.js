"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _express = require("express");

var _restaurant = _interopRequireDefault(require("../controllers/restaurant"));

var _menu = _interopRequireDefault(require("../controllers/menu"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const router = (0, _express.Router)();
router.route('/:id').get(_restaurant.default.getOne);
router.route('/:id/menu').get(_menu.default.getOne);
var _default = router;
exports.default = _default;