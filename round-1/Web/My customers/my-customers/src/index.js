import fs from "fs";
import express from "express";

const users = JSON.parse(fs.readFileSync("users.json"));

const app = express();
app.get("/user", async (_, res) => res.send("Use /user/:id endpoint"));

app.get("/user/:id", async (req, res) => {
  const id = parseInt(req.params.id);
  if (Number.isNaN(id)) {
    res.status(400).send("Invalid id");
    return;
  }
  const user = users[id];
  if (!user) {
    res.status(404).send("User not found");
    return;
  }
  res.send(user);
});

const port = Number(process.env.PORT) || 3000;
app.listen(port, () =>
  console.log(`🚀 Server running on http://localhost:${port}`)
);
