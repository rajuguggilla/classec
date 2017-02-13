/**
 * Created by bhanu.mokkala on 1/17/2017.
 */
angular.module('goclassec')
    .controller('configcontroller', function($scope, $http) {
console.log("Updating Credentials");

// console.log("Updating AWS Credentials");
// console.log("Updating AWS Credentials");
// console.log("Updating AWS Credentials");
        $scope.uploadAws = function(aws){
            console.log("Inside uploadAws function");
            console.log("AWS in Config.js===", aws);
            $http.post('/awsCred',{"aws":aws}).then(function(data){

            console.log(data);
            })
        }

        $scope.uploadAzure = function(azure){
            console.log("Inside uploadAzure function");
            console.log("Azure in Config.js===", azure);
            $http.post('/azureCred',{"azure":azure}).then(function(data){

            console.log(data);
            })
        }

        $scope.uploadOpenstack = function(openstack){
            console.log("Inside uploadOpensatck function");
            console.log("Openstack in Config.js===", openstack);
            $http.post('/osCred',{"openstack":openstack}).then(function(data){

            console.log(data);
            })
        }

        $scope.uploadHelion = function(helion){
            console.log("Inside uploadHelion function");
            console.log("Helion in Config.js===", helion);
            $http.post('/hosCred',{"helion":helion}).then(function(data){

            console.log(data);
            })
        }

        $scope.uploadVmware = function(vmware){
            console.log("Inside uploadVmware function");
            console.log("Vmware in Config.js===", vmware);
            $http.post('/vmwareCred',{"vmware":vmware}).then(function(data){

            console.log(data);
            })
        }

    });
    
