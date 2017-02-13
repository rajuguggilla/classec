/**
 * Created by bhanu.mokkala on 1/17/2017.
 */
var express   = require("express");
//var mysql     = require('mysql');
var path      = require('path');
//var bodyParser = require('body-parser')
var app = express();


/*
 var server = require('http').Server(app);
 var io = require('socket.io')(server);
 */
//server.listen(80);

var bodyParser = require('body-parser');
var jsonParser = bodyParser.json();
app.use(jsonParser);

clist = require('./server/oscomplist');
zlist = require('./server/azcomplist');
db = require('./server/dashboard');
ucfg = require('./server/userconfig');
ec2 = require('./server/ec2');
hlist = require('./server/hoscomplist')
vlist = require('./server/vmwarecomplist');
upAws = require('./server/awsCred');
upAzure = require('./server/azureCred');
upHos = require('./server/hosCred');
upOs = require('./server/osCred');
upVmware = require('./server/vmwareCred');




app.get('/oscomplist', clist.getlist);
app.get('/azcomplist', zlist.getlist);
app.get('/dashboard', db.getAll);
app.get('/hoscomplist', hlist.getlist);
app.get('/vmwarecomplist', vlist.getlist);
app.get('/dbcountstate', db.countbystate);
app.get('/dbcountvolstate', db.countvolbystate);
app.get('/userconfig', ucfg.getconfigdtls);
app.put('/userconfig', ucfg.updateparam);
app.get('/ec2utilcost', ec2.utilcost);
app.post('/awsCred', upAws.postAws);
app.post('/azureCred', upAzure.postAzure);
app.post('/hosCred', upHos.postHos);
app.post('/osCred', upOs.postOpenstack);
app.post('/vmwareCred', upVmware.postVmware);



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



