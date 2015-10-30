'use strict';

// Main controller, starting point for app logic
angular.module('go-albums-app').controller('MainCtrl', function ($scope, $http) {

  	// ng-repeat renders this
    $scope.awesomeThings = [
		'HTML5 Boilerplate',
		'AngularJS',
		'Karma'
    ];

    $scope.albums = [];

    // Get some albums!
    $http.get('/albums').then(function(data){
    	$scope.albums = data.data;
    }, function(err){
    	console.log(err);
    })
});
