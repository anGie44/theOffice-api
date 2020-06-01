var express = require('express')
var app = express();
var http = require('http');
require('dotenv').config();

var port = process.env.PORT || 8080;

app.use('/api', require('./api'))

// Setup server.
http.createServer(app).listen(port, function () {
    console.log(`Magic is happening on port ${port}`);
  });