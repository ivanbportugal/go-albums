'use strict';

// Starting point for the Angular app
angular.module('go-albums-app', []).config(function($routeProvider) {

    // Angular's Dependency Injection gets us routing for free
    $routeProvider
        .when('/', {
            templateUrl: 'views/main.html',
            controller: 'MainCtrl'
        })
        .otherwise({
            redirectTo: '/'
        });
});