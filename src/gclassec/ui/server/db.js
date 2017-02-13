/**
 * Created by bhanu.mokkala on 11/18/2016.
 */
/*
var mysql      = require('mysql');

var connection = mysql.createConnection({
    host     : '110.110.110.170',
    user     : 'root',
    password : 'root',
    database : 'cloud_assessment'
});

connection.connect(function(err) {
    if (err) throw err;
});

module.exports = connection;
    */

module.exports={
    db: {
        host     : '110.110.110.164',
        user     : 'root',
        password : 'root',
        database : 'cloud_assessment'
    }
};