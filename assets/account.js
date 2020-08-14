angular.module('myApp', []).controller('userCtrl', function($scope) {
    $scope.username = '';
    $scope.password = '';
    $scope.confirmpassword = '';

    $scope.edit = true;
    $scope.error = false;
    $scope.incomplete = false;
    $scope.hideform = true;
    $scope.editUser = function(id) {
        if (id == 'new') {
            $scope.username = '';
            $scope.password = '';
            $scope.confirmpassword = '';
        } else {
            $scope.edit = false;
            $scope.username = $scope.editUser(id).username;
            $scope.password = $scope.editUser(id).password;
            $scope.confirmpassword =$scope.editUser(id).confirmpassword;
        }
    };

    $scope.$watch('username',function() {$scope.test();});
    $scope.$watch('password',function() {$scope.test();});
    $scope.$watch('confirmpassword', function() {$scope.test();});

    $scope.test = function() {
        if ($scope.password !== $scope.confirmpassword) {
            $scope.error = true;
        } else {
            $scope.error = false;
        }
        $scope.incomplete = false;
        if ($scope.edit && (!$scope.username.length || !$scope.password.length || !$scope.confirmpassword.length)) {
            $scope.incomplete = true;
        }
    };

});