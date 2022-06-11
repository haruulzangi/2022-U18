// HZU18{Wh0's>N3xT?--5802a3042}
import fs from "fs";
import path from "path";
import mime from "mime-types";
import express from "express";

const app = express();

app.get("*", async (req, res) => {
  const requestPath = req.path.endsWith("/")
    ? req.path + "index.html"
    : req.path;
  const mimeType = mime.contentType(path.extname(requestPath)) || "text/plain";
  res.setHeader("Content-Type", mimeType);
  fs.createReadStream(path.join("public", requestPath))
    .on("error", () => {
      res.status(404).send("Not found");
    })
    .pipe(res);
});

const port = Number(process.env.PORT) || 3000;
app.listen(port, () => console.log(`Server is listening on ${port}`));
