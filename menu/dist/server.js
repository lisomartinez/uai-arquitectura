"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.start = exports.app = void 0;

var _express = _interopRequireDefault(require("express"));

var _bodyParser = require("body-parser");

var _restaurant = _interopRequireDefault(require("./routes/restaurant"));

var _db = require("./utils/db");

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const app = (0, _express.default)();
exports.app = app;
app.use((0, _bodyParser.json)());
app.use((0, _bodyParser.urlencoded)({
  extended: true
}));
app.get('/', (req, res) => {
  res.sendStatus(200);
});
app.use('/restaurants', _restaurant.default);

const start = async () => {
  try {
    await (0, _db.connect)();
    app.listen(8082, () => {
      console.log(`REST API on http://localhost:8082/`);
    });
  } catch (e) {
    console.error(e);
  }
};

exports.start = start;