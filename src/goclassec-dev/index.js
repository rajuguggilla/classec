/**
 * Created by bhanu.mokkala on 1/17/2017.
 */
var express   = require("express");
//var mysql     = require('mysql');
var path      = require('path');

var app = express();
/*
 var server = require('http').Server(app);
 var io = require('socket.io')(server);
 */
//server.listen(80);

var bodyParser = require('body-parser');
var jsonParser = bodyParser.json();

var app = express();
clist = require('./server/oscomplist');
zlist = require('./server/azcomplist');
app.get('/oscomplist', clist.getlist);
app.get('/azcomplist', zlist.getlist);

app.use('/scripts', express.static(__dirname + '/node_modules/'));
app.use('/bower_components', express.static(__dirname + '/bower_components/'));

app.get("/",function(req,res){

    res.sendFile(path.join(__dirname + '/public/index.html'));
});

// Set port
port = process.env.PORT || 2200;

// Use public directory for static files
app.use(express.static(__dirname + '/public'));


// Include the routes module
//require('./app/routes')(app);

// Your code here
app.listen(port);



