/**
 * Created by bhanu.mokkala on 11/7/2016.
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
exports.findAll = function(req, res) {
   // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query('SELECT * from vw_ec2withcost Where ec2_env <> "production"', function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};

exports.stopped = function(req, res) {
    // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query("SELECT * from vw_ec2withcost WHERE ec2_state='stopped' and stopdays<>''", function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};

exports.ebsunattached = function(req, res) {
    // res.send([{name:'wine1'}, {name:'wine2'}, {name:'wine3'}]);
    connection.query("SELECT ebs_volumeid, ebs_size, ebs_iops, ebs_region, ebs_volume_type FROM cloud_assessment.ebs_static where ebs_attachment_state=''", function(err, rows, fields) {
        if (!err) {
            //console.log('The solution is: ', rows);
            res.send(rows);
        }
        else
            console.log(err.message);
    });

    //connection.end();
};

exports.findById = function(req, res) {
    connection.query("SELECT * from ec2_static where ec_id='" + req.params.id+"'", function(err, row, fields) {
 //       connection.end();
        if (!err) {
            res.send(row);
        }
        else
            console.log(err.message);
//        connection.end();
    });
};

exports.utilcost = function(req, res) {
    connection.query("call sp_ec2cpuutilcost();", function(err, row, fields) {
        //       connection.end();
        if (!err) {
            res.send(row);
        }
        else
            console.log(err.message);
//        connection.end();
    });
};