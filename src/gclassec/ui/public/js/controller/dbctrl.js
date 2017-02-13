/**
 * Created by bhanu.mokkala on 1/18/2017.
 */
angular.module('goclassec')
    .controller('dbctrl', function($scope, $http, dbservice) {
            $scope.loading = true;
            $scope.data = [];
            dbservice.getdashboard()
                .then(function (data) {
                    //console.log(data);
                    $scope.db = data;
                }), function (data) {
                console.log('Error: ' + data);
            };

            dbservice.getdbcountstate()
                .then(function (dbg1) {
                    // $scope.dbg1 = data;
                    var i;
                    var labels = [];
                    var graphdata = [];
                    var graphcost = [];
                    for (i = 0; i < dbg1.data.length; i++) {
                        graphdata.push(dbg1.data[i].countid);
                        graphcost.push(dbg1.data[i].sumcost);
                        labels.push(dbg1.data[i].ec2_state);
                    }
                    $scope.labels = labels;
                    $scope.data = graphdata;
                    $scope.labels1 = labels;
                    $scope.data1 = graphcost;
                }), function (data) {
                console.log('Error: ' + data);
            };

            dbservice.getdbcountvolstate()
                .then(function (dbg2) {
                    // $scope.dbg1 = data;
                    var i;
                    var labels = [];
                    var graphdata = [];
                    for (i = 0; i < dbg2.data.length; i++) {
                        graphdata.push(dbg2.data[i].countstate);
                        labels.push(dbg2.data[i].vol_state);
                    }
                    $scope.labels2 = labels;
                    $scope.data2 = graphdata;
                }), function (data) {
                console.log('Error: ' + data);
            };

            $scope.loading = false;
    });