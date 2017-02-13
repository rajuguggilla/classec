/**
 * Created by bhanu.mokkala on 11/18/2016.
 */
var conn = require('./db');
var mysql      = require('mysql');
var connection = mysql.createConnection(conn.db);

exports.getconfigdtls = function(req, res) {

   // connection.query("Select * from user_config" , function (err, rows, fields) {
    connection.query("Select * from user_config where service_type='" + req.query.service + "'" , function (err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            // console.log(err.message);
            res.send(rows);
        }
        else
            console.log(err.message);
    });
};

exports.updateparam = function(req, res) {
    //console.log(req.query.svc_param);
    //console.log(req.query.svc_threshold);
    // res.send("got it");
    var curdate = new Date();
    //console.log("Insert into config (param_name, param_value, param_date) Values('secretkey','" + req.query.secretkey + "','" + curdate + "')");
    connection.query("Update user_config set service_threshold='" + req.query.svc_threshold +"' where service_param='" + req.query.svc_param +"'", function (err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            // console.log(err.message);
        }
        else
            console.log(err.message);
    });
};