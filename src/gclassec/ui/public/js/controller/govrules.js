/**
 * Created by bhanu.mokkala on 2/2/2017.
 */
angular.module('goclassec')
    .controller('govrules', function($scope, $http) {
        $http.get('/userconfig', {params: { service: 'compute' }})
            .then(function (data) {
                $scope.configdata = data;
            });
        $scope.selected = {};
        $scope.getTemplate = function (param) {
            if (param.service_param === $scope.selected.service_param) return 'edit';
            else return 'display';
        };

        $scope.editparam = function (param) {
            $scope.selected = angular.copy(param);
        };

        $scope.saveparam = function (param) {
            console.log("Saving contact");
            //$scope.model.contacts[idx] = angular.copy($scope.model.selected);
            var data = $.param({
                svc_param: param.service_param,
                svc_threshold: param.service_threshold
            });
            $http.put('/userconfig?' + data)
                .then(function (data) {
                    $scope.configdata = data;
                });
            $scope.reset();
        };

        $scope.reset = function () {
            $scope.selected = {};
        };
    });