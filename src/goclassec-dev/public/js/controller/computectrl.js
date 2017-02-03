/**
 * Created by bhanu.mokkala on 1/18/2017.
 */
angular.module('goclassec')
    .controller('computectrl', function($scope, $http) {
        $scope.tab1 = function () {
            $scope.loading = true;
            $scope.predicate = 'InstanceID';
            $scope.reverse = true;
            $scope.currentPage = 1;
            $scope.order = function (predicate) {
                $scope.reverse = ($scope.predicate === predicate) ? !$scope.reverse : false;
                $scope.predicate = predicate;
            };

            $http.get('/oscomplist')
                .then(function (data) {
                   // console.log(data.data.Value);
                    $scope.loading = false;
                    $scope.oscs = data.data.Value;
                    $scope.totalItems = $scope.oscs.length;
                    $scope.paginate = function (value) {
                        var begin, end, index;
                        begin = ($scope.currentPage - 1) * $scope.numPerPage;
                        end = begin + $scope.numPerPage;
                        index = $scope.oscs.indexOf(value);
                        return (begin <= index && index < end);
                    };
                }, function (error) {
                    console.log('Error: ' + data);
                });

            $scope.numPerPage = 5;
        };
        $scope.tab2 = function () {
            $scope.loading = true;
            $scope.predicate = 'VmId';
            $scope.reverse = true;
            $scope.currentPage = 1;
            $scope.order = function (predicate) {
                $scope.reverse = ($scope.predicate === predicate) ? !$scope.reverse : false;
                $scope.predicate = predicate;
            };

            $http.get('/azcomplist')
                .then(function (data) {
                    console.log(data.data.Value);
                    $scope.loading = false;
                    $scope.oscs = data.data.Value;
                    $scope.totalItems = $scope.oscs.length;
                    $scope.paginate = function (value) {
                        var begin, end, index;
                        begin = ($scope.currentPage - 1) * $scope.numPerPage;
                        end = begin + $scope.numPerPage;
                        index = $scope.oscs.indexOf(value);
                        return (begin <= index && index < end);
                    };
                }, function (error) {
                    console.log('Error: ' + data);
                });

            $scope.numPerPage = 5;
        };
    });