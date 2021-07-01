const userRoutes = (app, fs) => {
  // variables
const dataPath = './data/catalogo.json';
  //
  //     // READ
app.get('/catalogo', (req, res) => {
fs.readFile(dataPath, 'utf8', (err, data) => {
if (err) {
throw err;
      }
  //
 res.send(JSON.parse(data));
  });
  });
  };
  //
module.exports = userRoutes;
