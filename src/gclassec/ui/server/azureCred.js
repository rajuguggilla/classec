/**
 * Created by bhanu.mokkala on 12/22/2016.
 */
/**
 * Created by bhanu.mokkala on 12/16/2016.
 */
var http = require('http');
var https = require('https');
var querystring = require('querystring');
var fs = require('fs');
var request = require("request");

exports.postAzure = function(req, res) {
    console.log("POSTAzure function called");
    console.log("req.body===", req.body);
    console.log("req.body.aws===", req.body);
    // var options = {
    //     method: 'POST',
    //     url: 'localhost:4567/aws/updatecredentials',
    //     body: req.body
    // };

    // request(options,function(error,res,body){
    //     if(error) throw new Error(error);
    //     console.log("##########from Go Server######");
    //     console.log(body);
    //     res.send(body);

    // })

    var data = querystring.stringify({
                     clientid: req.body.azure.clientID,
                     clientsecret: req.body.azure.clientSecret,
                     subscriptionid: req.body.azure.subscriptionID,
                     tenantid: req.body.azure.tenantID
                     });

    var options = {
        host: '110.110.110.233',
        port: 4567,
        path: '/azure/updatecredentials',
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Content-Length': Buffer.byteLength(data)
        }
    };

    var req = http.request(options, function(res) {
        res.setEncoding('utf8');
        res.on('data', function (chunk) {
            console.log("body: " + chunk);
        });
});

req.write(data);
req.end();

};