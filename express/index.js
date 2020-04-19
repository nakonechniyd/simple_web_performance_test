const express = require("express");

const uuid = require("uuid");

const app = express();

function prefix(req, res) {
  return res.send({ prefix: uuid.v4() });
}

app.use("/prefix", prefix);

const start = async () => {
  const port = 8080;
  try {
    app.listen(port, () => {
      console.log(`Started on http://localhost:${port}`);
    });
  } catch (e) {
    console.error(e);
  }
};

start();
