"use strict";

const userRoutes = (app, fs) => {
  app.get('/restaurants', (req, res) => {
    res.sendStatus(200);
  });
  app.get('/restaurants/:restaurantId/menu', (req, res) => {
    if (req.params.restaurantId !== 1) {
      res.status(404);
      res.send('restaurant not found');
      return;
    }

    res.status(200);
  });
};

module.exports = userRoutes;