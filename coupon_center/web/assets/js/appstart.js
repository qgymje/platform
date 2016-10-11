 
angular.element(document).ready(function() { 
  jQuery.getJSON('api/user.json', function(data) { 
    userSession = angular.fromJson(data); 
    angular.bootstrap(document, ['app']); 
  });  
}); 