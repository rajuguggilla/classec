/**
 * Created by bhanu.mokkala on 1/31/2017.
 */
angular.module('goclassec')
    .controller('menuctrl', function($scope, $mdSidenav) {
        $scope.showMobileMainHeader = true;
        $scope.openSideNavPanel = function() {
            $mdSidenav('left').open();
        };
        $scope.closeSideNavPanel = function() {
            $mdSidenav('left').close();
        };
    });