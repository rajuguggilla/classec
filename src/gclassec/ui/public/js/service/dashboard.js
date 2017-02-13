/**
 * Created by bhanu.mokkala on 11/16/2016.
 */
'use strict';

angular.module('goclassec')
    .factory('dbservice', function($http){
    return {
        getdashboard: function () {
           return $http.get('/dashboard');
        },
        getdbcountstate: function() {
            return $http.get('/dbcountstate');
        },
        getdbcountvolstate: function() {
            return $http.get('/dbcountvolstate');
        }
    }

});
