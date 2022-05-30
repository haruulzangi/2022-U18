import pg from "pg";
import express from "express";

const { Client } = pg;

const client = new Client();
await client.connect();

const app = express();
app.get("/vaccine", async (_, res) => {
  const result = await client.query('SELECT id, name FROM "vaccines"');
  res.json(result.rows);
});
app.get("/vaccine/:id", async (req, res) => {
  try {
    const result = await client.query(
      `SELECT id, name FROM "vaccines" WHERE id = ${req.params.id}`
    );
    res.json(result);
  } catch (e) {
    res.send(`${e.message}\n${e.stack}`);
  }
});

const port = Number(process.env.PORT) || 3000;
app.listen(port, () =>
  console.log(`🚀 Server running on http://localhost:${port}`)
);
