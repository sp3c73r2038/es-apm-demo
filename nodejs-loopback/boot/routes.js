module.exports = (app) => {
  app.get("/ping", (req, resp) => {
    resp.send("pong");
  });
};
