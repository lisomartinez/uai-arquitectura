"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.Menu = void 0;

var _mongoose = _interopRequireDefault(require("mongoose"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const menuItemSchema = new _mongoose.default.Schema({
  item_id: {
    type: Number,
    required: true,
    trim: true
  },
  name: {
    type: String,
    required: true,
    trim: true
  },
  ingredients: {
    type: [String]
  }
});
const menuSchema = new _mongoose.default.Schema({
  restaurant_id: {
    type: Number,
    required: true
  },
  items: {
    type: [menuItemSchema]
  }
});
menuSchema.index({
  restaurant_id: 1
}, {
  unique: true
});

const Menu = _mongoose.default.model('menu', menuSchema);

exports.Menu = Menu;