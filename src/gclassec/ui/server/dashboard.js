/**
 * Created by bhanu.mokkala on 11/9/2016.
 */
/*
var mysql      = require('mysql');
var connection = mysql.createConnection({
    host     : '110.110.110.170',
    user     : 'root',
    password : 'root',
    database : 'cloud_assessment'
});

connection.connect(function(err){
    if(!err) {
        //console.log("Database is connected ... nn");
    } else {
        console.log(err.message);
    }
});
*/
var conn = require('./db');
var mysql      = require('mysql');
var connection = mysql.createConnection(conn.db);

exports.getAll = function(req, res) {
    // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query('call sp_dashboard', function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};

exports.countbystate = function(req, res) {
    // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query('Select * from vw_countbystate', function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};

exports.countvolbystate = function(req, res) {
    // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query('Select * from vw_countvolbystate', function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};