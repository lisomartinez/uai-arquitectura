"use strict";

const userRoutes = require("./catalogo");

const appRouter = (app, fs) => {
  app.get("/", (req, res) => {
    res.sendStatus(200);
  });
  userRoutes(app, fs);
};

module.exports = appRouter;