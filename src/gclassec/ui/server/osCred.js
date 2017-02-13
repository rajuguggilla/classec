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

exports.postOpenstack = function(req, res) {
    console.log("POST oPENsTACK function called");
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
                     Host: req.body.openstack.osHost,
                     Username: req.body.openstack.osUserName,
                     Password: req.body.openstack.osPassword,
                     ProjectID:req.body.openstack.osTenantID,
                     ProjectName: req.body.openstack.osTenantName,
                     Container: req.body.openstack.osContainerName,
                     ImageRegion: req.body.openstack.osImageRegion,
                     Controller: req.body.openstack.osController
                     });

    var options = {
        host: '110.110.110.233',
        port: 4567,
        path: '/openstack/updatecredentials',
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