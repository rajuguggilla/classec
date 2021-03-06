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
                },function (error) {
                    console.log('Error: ' + data);
                });

            $scope.numPerPage = 5;
        };
		
		$scope.tab3 = function () {
            $scope.loading = true;
            $scope.predicate = 'InstanceID';
            $scope.reverse = true;
            $scope.currentPage = 1;
            $scope.order = function (predicate) {
                $scope.reverse = ($scope.predicate === predicate) ? !$scope.reverse : false;
                $scope.predicate = predicate;
            };

            $http.get('/hoscomplist')
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
		
        $scope.tab4 = function () {
            $scope.loading = true;
            $scope.predicate = 'ec_id';
            $scope.reverse = true;
            $scope.currentPage = 1;
            $scope.order = function (predicate) {
                $scope.reverse = ($scope.predicate === predicate) ? !$scope.reverse : false;
                $scope.predicate = predicate;
            };

            $http.get('/ec2utilcost')
                .then(function (data) {
                    $scope.loading=false;
                     console.log(data);
                    var total = 0;
                    for (var i = 0; i < data.data[0].length; i++) {
                        total = total + data.data[0][i].cost;
                    }
                    $scope.totalundertuilized = total;

                    //   }

                    $scope.ec2s = data.data[0];
                    $scope.totalItems = $scope.ec2s.length;
                    $scope.paginate = function (value) {
                        var begin, end, index;
                        begin = ($scope.currentPage - 1) * $scope.numPerPage;
                        end = begin + $scope.numPerPage;
                        index = $scope.ec2s.indexOf(value);
                        return (begin <= index && index < end);
                    };
                }),function (data) {
                    console.log('Error: ' + data);
                };

            $scope.numPerPage = 5;
        };
		
		$scope.tab5 = function () {
            $scope.loading = true;
            $scope.predicate = 'Id';
            $scope.reverse = true;
            $scope.currentPage = 1;
            $scope.order = function (predicate) {
                $scope.reverse = ($scope.predicate === predicate) ? !$scope.reverse : false;
                $scope.predicate = predicate;
            };

            $http.get('/vmwarecomplist')
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
        $scope.gettotal = function(filtered) {
            //console.log(filtered.length);
            if (typeof filtered != 'undefined') {
                var total = 0;
                for (var i = 0; i < filtered.length; i++) {
                    total = total + filtered[i].cost;
                }

                //$scope.totalstopped = total;
                return total;
            }
        };
    });