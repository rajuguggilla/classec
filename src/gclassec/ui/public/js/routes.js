/**
 * Created by bhanu.mokkala on 1/17/2017.
 */
'use strict';

angular.module("goclassec")
    .config(['$routeProvider', '$locationProvider', '$mdThemingProvider',function($routeProvider, $locationProvider, $mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('deep-purple')
            .accentPalette('blue');
        $locationProvider
            .html5Mode(false)
            .hashPrefix('!');
        $routeProvider
            .when('/dashboard', {
                templateUrl : 'view/dashboard.html',
                controller : 'dbctrl'
            })
            .when('/', {
                templateUrl : 'view/dashboard.html',
                controller : 'dbctrl'
            })
            .when('/compute', {
                templateUrl : 'view/compute.html',
                controller : 'computectrl'
            })
            .when('/governance', {
                templateUrl : 'view/governance.html',
                controller : 'govctrl'
            })
            .when('/govrules', {
                templateUrl : 'view/govrules.html',
                controller : 'govrules'
            })
            .when('/config', {
                templateUrl : 'view/config.html',
                controller : 'configcontroller'
            });
    }]);