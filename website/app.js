const express = require('express');
const url = require('url');
const path = require('path');

const Invoke_Hyperledger = require('./Invoke_Hyperledger');
const Query_Hyperledger = require('./Query_Hyperledger');



const app = express();


app.listen(8000, console.log('Listen at 8000'));

app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname + '/index.html'));
})

app.get('/query_hyperledger', function (req, res) {
  q = url.parse(req.url, true).query;
  Query_Hyperledger(q.key).then((data) => { return res.send(data); });
});

app.get('/invoke_hyperledger', function (req, res) {
  q = url.parse(req.url, true).query;
  Invoke_Hyperledger(q.key, q.value).then((data) => res.send({ "status": "success" }));
});
