const express = require("express");
const bodyParser = require("body-parser");
const notifier = require("node-notifier");
const path = require("path");

const app = express();
const port = process.env.PORT || 9000;

//middleware funtion which takes request tranforms to JSON & pass it to controller
app.use(bodyParser.json());

app.get("/health", (req, res) => {
  res.status(200).send();
});

app.post("/notify", (req, res) => {
  notify(req.body, (reply) => res.send(reply));
});

app.listen(port, () =>
  console.log(`Server is running at https://localhost:${port}`)
);

const notify = ({ title, message }, cb) => {
  notifier.notify(
    {
      title: title || "Unknown title",
      message: message || "Unknown message",
      icon: path.join(__dirname, './notification-img.png'),
      sound: true,
      wait: true,
      reply: true,
      closeLabel: "Completed?",
      timeout: 15,
    },
    (err, response, reply) => {
      cb(reply);
    }
  );
};
