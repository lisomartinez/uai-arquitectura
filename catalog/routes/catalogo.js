const userRoutes = (app, fs) => {
  const catalog = "./data/catalogo.json";
  const restaurants = "./data/restaurants.json";

  app.get("/restaurants", (req, res) => {
    fs.readFile(restaurants, "utf8", (err, data) => {
      if (err) {
        throw err;
      }
      res.send(JSON.parse(restaurants));
    });
  });

  app.get("/restaurants/:restaurantId/menu", (req, res) => {
    if (req.params.restaurantId != 1) {
      res.status(404);
      res.send("restaurant not found");
      return;
    }

    fs.readFile(catalog, "utf8", (err, data) => {
      if (err) {
        throw err;
      }
      res.send(JSON.parse(data));
    });
  });
};

module.exports = userRoutes;
