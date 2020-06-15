const express = require('express');
const url = require('url');
const path = require('path');

const Invoke_Hyperledger = require('./Invoke_Hyperledger');
const Query_Hyperledger = require('./Query_Hyperledger');



const app = express();


app.listen(8000, console.log(`App listening at http://localhost:8000`));

app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname + '/index.html'));
})

app.get('/query_hyperledger', function (req, res) {
  q = url.parse(req.url, true).query;
  Query_Hyperledger('retrieve',q.key).then((response) => { res.send(response); });
});

app.get('/invoke_hyperledger', function (req, res) {
  q = url.parse(req.url, true).query;
  Invoke_Hyperledger('store', q.key, q.value).then((response) => res.send({ "status": "success" }));
});
