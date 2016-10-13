/*
 *  Document   : app.js
 *  Author     : pixelcave
 *  Description: Setting up and vital functionality for our App
 *
 */  

var userSession; 

// Create our angular module
var App = angular.module('app', [
    'ngStorage',
    'ui.router',
    'ui.bootstrap',
    'oc.lazyLoad',
    'pascalprecht.translate'
]);


App.constant('USER_ROLES', {
  all: '*',
  admin: 'admin',
  guest: 'guest'
});

App.constant('USER_Permissions', {
  stats: 'stats',
  stores: 'stores',
  campaigns: 'campaigns',
  account: 'account'
});

App.constant('AUTH_Permission', {
  noentry: 'NoEntry',
  view: 'View',
  edit: 'Edit'
});

App.constant('AUTH_State', {
  loggedin: 'auth-login-loggedin',
  adminlogin: 'auth-logout-admin',
  notAuthenticated: 'auth-not-authenticated',
  notAuthorized: 'auth-not-Authorized'
});

// Router configuration
App.config(['$stateProvider', '$urlRouterProvider', 'USER_ROLES', 'USER_Permissions', 'AUTH_Permission',
    function ($stateProvider, $urlRouterProvider, USER_ROLES, USER_Permissions, AUTH_Permission) {
        $urlRouterProvider.otherwise('/stats');
        $stateProvider
            .state('noauth', {
                url: '/noauth',
                data: { authorizedRoles : USER_ROLES.all },
                views:{
                    '':{
                        templateUrl: 'assets/views/main.html'
                    },
                    'main@noauth':{
                        templateUrl: 'assets/views/noauth.html'
                    }
                }
            })
            .state('login', {
                url: '/login',
                data: { authorizedRoles : USER_ROLES.guest },
                views:{
                    '':{
                        templateUrl: 'assets/views/user.html'
                    },
                    'main@login':{
                        templateUrl: 'assets/views/login.html',
                        controller: 'LoginValidationCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('register', {
                url: '/register',
                data: { authorizedRoles : USER_ROLES.guest },
                views:{
                    '':{
                        templateUrl: 'assets/views/user.html'
                    },
                    'main@register':{
                        templateUrl: 'assets/views/register.html',
                        controller: 'RegisterValidationCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('logout', {
                url: '/logout',
                data: { authorizedRoles : USER_ROLES.admin },
                views:{
                    '':{
                        template: '',
                        controller: 'LogoutCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('stats', {
                url: '/stats',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.stats ,permission : AUTH_Permission.view},
                views:{
                    '':{
                        templateUrl: 'assets/views/main.html'
                    },
                    'main@stats':{
                        templateUrl: 'assets/views/stats.html',
                        controller: 'StatsCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker3.min.css',
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker.min.js',
                                        'assets/js/plugins/bootstrap-datepicker/locales/bootstrap-datepicker.zh-CN.min.js',
                                        'assets/js/core/jquery.countTo.min.js',
                                        'assets/js/plugins/chartjs/Chart.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('stores', {
                url: '/stores',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.stores ,permission : AUTH_Permission.view},
                views:{
                    '':{
                        templateUrl: 'assets/views/main.html'
                    },
                    'main@stores':{
                        templateUrl: 'assets/views/stores_index.html',
                        controller: 'StoresCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('stores.add', {
                url: '/add',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.stores ,permission : AUTH_Permission.edit},
                views:{
                    'main@stores':{
                        templateUrl: 'assets/views/stores_add.html',
                        controller: 'StoresAddCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js',
                                        'assets/js/plugins/bmapsjs/css/bmaps.css'
                                        //,'assets/js/plugins/bmapsjs/bmaps.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('stores.edit', {
                url: '/edit/:storeId',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.stores ,permission : AUTH_Permission.edit},
                views:{
                    'main@stores':{
                        templateUrl: 'assets/views/stores_edit.html',
                        controller: 'StoresEditCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/bootstrap-notify/bootstrap-notify.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js',
                                        'assets/js/plugins/bmapsjs/css/bmaps.css'
                                        //,'assets/js/plugins/bmapsjs/bmaps.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('stores.map', {
                url: '/map',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.stores ,permission : AUTH_Permission.view},
                views:{
                    'main@stores':{
                        templateUrl: 'assets/views/stores_map.html',
                        controller: 'StoresMapCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/bmapsjs/css/bmaps.css'
                                        //,'assets/js/plugins/bmapsjs/bmaps.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns', {
                url: '/campaigns',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.view},
                views:{
                    '':{
                        templateUrl: 'assets/views/main.html'
                    },
                    'main@campaigns':{
                        templateUrl: 'assets/views/campaigns_index.html',
                        controller: 'CampaignsCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns.add', {
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.edit},
                url: '/add',
                views:{
                    'main@campaigns':{
                        templateUrl: 'assets/views/campaigns_add.html',
                        controller: 'CampaignsAddCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker3.min.css',
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker.min.js',
                                        'assets/js/plugins/bootstrap-datepicker/locales/bootstrap-datepicker.zh-CN.min.js',
                                        'assets/js/plugins/bootstrap-notify/bootstrap-notify.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns.edit', {
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.edit},
                url: '/edit/:campaignId',
                views:{
                    'main@campaigns':{
                        templateUrl: 'assets/views/campaigns_edit.html',
                        controller: 'CampaignsEditCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker3.min.css',
                                        'assets/js/plugins/bootstrap-datepicker/bootstrap-datepicker.min.js',
                                        'assets/js/plugins/bootstrap-datepicker/locales/bootstrap-datepicker.zh-CN.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns.show', {
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.edit},
                url: '/show/:campaignId',
                views:{
                    'main@campaigns':{
                        templateUrl: 'assets/views/campaigns_show.html',
                        controller: 'CampaignsShowCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/bootstrap-notify/bootstrap-notify.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns.addpromotions', {
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.edit},
                url: '/show/:campaignId',
                views:{
                    'main@campaigns':{
                        templateUrl: 'assets/views/promotion_add.html',
                        controller: 'CampaignsAddPromotionsCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/bootstrap-notify/bootstrap-notify.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js',
                                        'assets/js/plugins/ng-file-upload/ng-file-upload.min.js',
                                        'assets/js/plugins/ng-file-upload/ng-file-upload-shim.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('campaigns.editpromotions', {
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.campaigns ,permission : AUTH_Permission.edit},
                url: '/show/:campaignId/:promotionId',
                views:{
                    'main@campaigns':{
                        templateUrl: 'assets/views/promotion_edit.html',
                        controller: 'CampaignsEditPromotionsCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/bootstrap-notify/bootstrap-notify.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js',
                                        'assets/js/plugins/ng-file-upload/ng-file-upload.min.js',
                                        'assets/js/plugins/ng-file-upload/ng-file-upload-shim.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('account', {
                url: '/account',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.account ,permission : AUTH_Permission.view},
                views:{
                    '':{
                        templateUrl: 'assets/views/main.html'
                    },
                    'main@account':{
                        templateUrl: 'assets/views/account_index.html',
                        controller: 'AccountCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
            .state('account.add', {
                url: '/add',
                data: { authorizedRoles : USER_ROLES.admin ,authorizedPermission : USER_Permissions.account ,permission : AUTH_Permission.edit},
                views:{
                    'main@account':{
                        templateUrl: 'assets/views/account_add.html',
                        controller: 'AccountAddCtrl',
                        resolve: {
                            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                                return $ocLazyLoad.load({
                                    insertBefore: '#css-bootstrap',
                                    serie: true,
                                    files: [
                                        'assets/js/plugins/jquery-validation/jquery.validate.min.js',
                                        'assets/js/plugins/jquery-validation/additional-methods.min.js',
                                        'assets/js/plugins/sweetalert/sweetalert.min.css',
                                        'assets/js/plugins/sweetalert/sweetalert.min.js'
                                    ]
                                });
                            }]
                        }
                    }
                }
            })
    }
]);

// Tooltips and Popovers configuration
App.config(['$uibTooltipProvider',
    function ($uibTooltipProvider) {
        $uibTooltipProvider.options({
            appendToBody: true
        });
    }
]);

App.config(['$translateProvider',
    function($translateProvider){  
        var lang = window.localStorage.lang||'zh-CN';  
        $translateProvider.preferredLanguage(lang);  
        $translateProvider.useStaticFilesLoader({  
            prefix: 'assets/lang/',  
            suffix: '.json'  
        });  
    }
]); 

// Custom UI helper functions
App.factory('uiHelpers', function () {
    return {
        // Handles active color theme
        uiHandleColorTheme: function (themeName) {
            var colorTheme = jQuery('#css-theme');

            if (themeName) {
                if (colorTheme.length && (colorTheme.prop('href') !== 'assets/css/themes/' + themeName + '.min.css')) {
                    jQuery('#css-theme').prop('href', 'assets/css/themes/' + themeName + '.min.css');
                } else if (!colorTheme.length) {
                    jQuery('#css-main').after('<link rel="stylesheet" id="css-theme" href="assets/css/themes/' + themeName + '.min.css">');
                }
            } else {
                if (colorTheme.length) {
                    colorTheme.remove();
                }
            }
        },
        // Handles #main-container height resize to push footer to the bottom of the page
        uiHandleMain: function () {
            var lMain       = jQuery('#main-container');
            var hWindow     = jQuery(window).height();
            var hHeader     = jQuery('#header-navbar').outerHeight();
            var hFooter     = jQuery('#page-footer').outerHeight();

            if (jQuery('#page-container').hasClass('header-navbar-fixed')) {
                lMain.css('min-height', hWindow - hFooter);
            } else {
                lMain.css('min-height', hWindow - (hHeader + hFooter));
            }
        },
        // Handles transparent header functionality (solid on scroll - used in frontend pages)
        uiHandleHeader: function () {
            var lPage = jQuery('#page-container');

            if (lPage.hasClass('header-navbar-fixed') && lPage.hasClass('header-navbar-transparent')) {
                jQuery(window).on('scroll', function(){
                    if (jQuery(this).scrollTop() > 20) {
                        lPage.addClass('header-navbar-scroll');
                    } else {
                        lPage.removeClass('header-navbar-scroll');
                    }
                });
            }
        },
        // Handles sidebar and side overlay custom scrolling functionality
        uiHandleScroll: function(mode) {
            var windowW            = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
            var lPage              = jQuery('#page-container');
            var lSidebar           = jQuery('#sidebar');
            var lSidebarScroll     = jQuery('#sidebar-scroll');
            var lSideOverlay       = jQuery('#side-overlay');
            var lSideOverlayScroll = jQuery('#side-overlay-scroll');

            // Init scrolling
            if (mode === 'init') {
                // Init scrolling only if required the first time
                uiHandleScroll();
            } else {
                // If screen width is greater than 991 pixels and .side-scroll is added to #page-container
                if (windowW > 991 && lPage.hasClass('side-scroll') && (lSidebar.length || lSideOverlay.length)) {
                    // If sidebar exists
                    if (lSidebar.length) {
                        // Turn sidebar's scroll lock off (slimScroll will take care of it)
                        jQuery(lSidebar).scrollLock('disable');

                        // If sidebar scrolling does not exist init it..
                        if (lSidebarScroll.length && (!lSidebarScroll.parent('.slimScrollDiv').length)) {
                            lSidebarScroll.slimScroll({
                                height: lSidebar.outerHeight(),
                                color: '#fff',
                                size: '5px',
                                opacity : .35,
                                wheelStep : 15,
                                distance : '2px',
                                railVisible: false,
                                railOpacity: 1
                            });
                        }
                        else { // ..else resize scrolling height
                            lSidebarScroll
                                .add(lSidebarScroll.parent())
                                .css('height', lSidebar.outerHeight());
                        }
                    }

                    // If side overlay exists
                    if (lSideOverlay.length) {
                        // Turn side overlay's scroll lock off (slimScroll will take care of it)
                        jQuery(lSideOverlay).scrollLock('disable');

                        // If side overlay scrolling does not exist init it..
                        if (lSideOverlayScroll.length && (!lSideOverlayScroll.parent('.slimScrollDiv').length)) {
                            lSideOverlayScroll.slimScroll({
                                height: lSideOverlay.outerHeight(),
                                color: '#000',
                                size: '5px',
                                opacity : .35,
                                wheelStep : 15,
                                distance : '2px',
                                railVisible: false,
                                railOpacity: 1
                            });
                        }
                        else { // ..else resize scrolling height
                            lSideOverlayScroll
                                .add(lSideOverlayScroll.parent())
                                .css('height', lSideOverlay.outerHeight());
                        }
                    }
                } else {
                    // If sidebar exists
                    if (lSidebar.length) {
                        // If sidebar scrolling exists destroy it..
                        if (lSidebarScroll.length && lSidebarScroll.parent('.slimScrollDiv').length) {
                            lSidebarScroll
                                .slimScroll({destroy: true});
                            lSidebarScroll
                                .attr('style', '');
                        }

                        // Turn sidebars's scroll lock on
                        jQuery(lSidebar).scrollLock();
                    }

                    // If side overlay exists
                    if (lSideOverlay.length) {
                        // If side overlay scrolling exists destroy it..
                        if (lSideOverlayScroll.length && lSideOverlayScroll.parent('.slimScrollDiv').length) {
                            lSideOverlayScroll
                                .slimScroll({destroy: true});
                            lSideOverlayScroll
                                .attr('style', '');
                        }

                        // Turn side overlay's scroll lock on
                        jQuery(lSideOverlay).scrollLock();
                    }
                }
            }
        },
        // Handles page loader functionality
        uiLoader: function (mode) {
            var lBody       = jQuery('body');
            var lpageLoader = jQuery('#page-loader');

            if (mode === 'show') {
                if (lpageLoader.length) {
                    lpageLoader.fadeIn(250);
                } else {
                    lBody.prepend('<div id="page-loader"></div>');
                }
            } else if (mode === 'hide') {
                if (lpageLoader.length) {
                    lpageLoader.fadeOut(250);
                }
            }
        },
        uiSidebarHide: function (mode) {
            var lBody       = jQuery('body');
            var lpageLoader = jQuery('#sidebar');
            var lpageLoader2 = jQuery('#page-container');

            if (mode === true) {
                if (lpageLoader.length) {
                    lpageLoader.hide();
                }
                if (lpageLoader.length) {
                    lpageLoader2.addClass('sidebar-h');
                }
            }
            else{
                if (lpageLoader.length) {
                    lpageLoader.show();
                }
                if (lpageLoader.length) {
                    lpageLoader2.removeClass('sidebar-h');
                }
            }
        },
        // Handles blocks API functionality
        uiBlocks: function (block, mode, button) {
            // Set default icons for fullscreen and content toggle buttons
            var iconFullscreen         = 'si si-size-fullscreen';
            var iconFullscreenActive   = 'si si-size-actual';
            var iconContent            = 'si si-arrow-up';
            var iconContentActive      = 'si si-arrow-down';

            if (mode === 'init') {
                // Auto add the default toggle icons
                switch(button.data('action')) {
                    case 'fullscreen_toggle':
                        button.html('<i class="' + (button.closest('.block').hasClass('block-opt-fullscreen') ? iconFullscreenActive : iconFullscreen) + '"></i>');
                        break;
                    case 'content_toggle':
                        button.html('<i class="' + (button.closest('.block').hasClass('block-opt-hidden') ? iconContentActive : iconContent) + '"></i>');
                        break;
                    default:
                        return false;
                }
            } else {
                // Get block element
                var elBlock = (block instanceof jQuery) ? block : jQuery(block);

                // If element exists, procceed with blocks functionality
                if (elBlock.length) {
                    // Get block option buttons if exist (need them to update their icons)
                    var btnFullscreen  = jQuery('[data-js-block-option][data-action="fullscreen_toggle"]', elBlock);
                    var btnToggle      = jQuery('[data-js-block-option][data-action="content_toggle"]', elBlock);

                    // Mode selection
                    switch(mode) {
                        case 'fullscreen_toggle':
                            elBlock.toggleClass('block-opt-fullscreen');

                            // Enable/disable scroll lock to block
                            if (elBlock.hasClass('block-opt-fullscreen')) {
                                jQuery(elBlock).scrollLock();
                            } else {
                                jQuery(elBlock).scrollLock('disable');
                            }

                            // Update block option icon
                            if (btnFullscreen.length) {
                                if (elBlock.hasClass('block-opt-fullscreen')) {
                                    jQuery('i', btnFullscreen)
                                        .removeClass(iconFullscreen)
                                        .addClass(iconFullscreenActive);
                                } else {
                                    jQuery('i', btnFullscreen)
                                        .removeClass(iconFullscreenActive)
                                        .addClass(iconFullscreen);
                                }
                            }
                            break;
                        case 'fullscreen_on':
                            elBlock.addClass('block-opt-fullscreen');

                            // Enable scroll lock to block
                            jQuery(elBlock).scrollLock();

                            // Update block option icon
                            if (btnFullscreen.length) {
                                jQuery('i', btnFullscreen)
                                    .removeClass(iconFullscreen)
                                    .addClass(iconFullscreenActive);
                            }
                            break;
                        case 'fullscreen_off':
                            elBlock.removeClass('block-opt-fullscreen');

                            // Disable scroll lock to block
                            jQuery(elBlock).scrollLock('disable');

                            // Update block option icon
                            if (btnFullscreen.length) {
                                jQuery('i', btnFullscreen)
                                    .removeClass(iconFullscreenActive)
                                    .addClass(iconFullscreen);
                            }
                            break;
                        case 'content_toggle':
                            elBlock.toggleClass('block-opt-hidden');

                            // Update block option icon
                            if (btnToggle.length) {
                                if (elBlock.hasClass('block-opt-hidden')) {
                                    jQuery('i', btnToggle)
                                        .removeClass(iconContent)
                                        .addClass(iconContentActive);
                                } else {
                                    jQuery('i', btnToggle)
                                        .removeClass(iconContentActive)
                                        .addClass(iconContent);
                                }
                            }
                            break;
                        case 'content_hide':
                            elBlock.addClass('block-opt-hidden');

                            // Update block option icon
                            if (btnToggle.length) {
                                jQuery('i', btnToggle)
                                    .removeClass(iconContent)
                                    .addClass(iconContentActive);
                            }
                            break;
                        case 'content_show':
                            elBlock.removeClass('block-opt-hidden');

                            // Update block option icon
                            if (btnToggle.length) {
                                jQuery('i', btnToggle)
                                    .removeClass(iconContentActive)
                                    .addClass(iconContent);
                            }
                            break;
                        case 'refresh_toggle':
                            elBlock.toggleClass('block-opt-refresh');

                            // Return block to normal state if the demostration mode is on in the refresh option button - data-action-mode="demo"
                            if (jQuery('[data-js-block-option][data-action="refresh_toggle"][data-action-mode="demo"]', elBlock).length) {
                                setTimeout(function(){
                                    elBlock.removeClass('block-opt-refresh');
                                }, 2000);
                            }
                            break;
                        case 'state_loading':
                            elBlock.addClass('block-opt-refresh');
                            break;
                        case 'state_normal':
                            elBlock.removeClass('block-opt-refresh');
                            break;
                        case 'close':
                            elBlock.hide();
                            break;
                        case 'open':
                            elBlock.show();
                            break;
                        default:
                            return false;
                    }
                }
            }
        }
    };
});

// Run our App
App.run(function($rootScope, uiHelpers) {
    // Access uiHelpers easily from all controllers
    $rootScope.helpers = uiHelpers;

    // On window resize or orientation change resize #main-container & Handle scrolling
    var resizeTimeout;

    jQuery(window).on('resize orientationchange', function () {
        clearTimeout(resizeTimeout);

        resizeTimeout = setTimeout(function(){
            $rootScope.helpers.uiHandleScroll();
            $rootScope.helpers.uiHandleMain();
        }, 150);
    });
});

App.service('Session', function () {
  this.create = function (sessionId, userId, userRole, userPermission) {
    this.id = sessionId;
    this.userId = userId;
    this.userRole = userRole;
    this.userPermission = userPermission;
  };
  this.destroy = function () {
    this.id = null;
    this.userId = null;
    this.userRole = null;
    this.userPermission = [];
  };
  return this;
});

App.run(function(Session){ 
    var sessionId = null;
    var userId = null;
    var userRole = null;
    var userPermission = [];
    if(userSession.id){sessionId = userSession.id;}
    if(userSession.userId){userId = userSession.userId;}
    if(userSession.userRole){userRole = userSession.userRole;}
    if(userSession.userPermission){userPermission = userSession.userPermission;}
    Session.create(sessionId, userId, userRole, userPermission);
});

App.factory("AuthService",['Session', 'AUTH_State', 'USER_ROLES', 'AUTH_Permission', '$http', '$q' ,
    function(Session, AUTH_State, USER_ROLES, AUTH_Permission, $http, $q){
    return {
        isAuthenticated: function(){
            return !!Session.userId;
        },
        isAuthorized: function(isAuthenticated, authorizedRoles, authorizedPermission, authPermission){
        	if(authorizedRoles != USER_ROLES.all){
	            if(isAuthenticated){
	                if(authorizedRoles == USER_ROLES.guest){
	                    return AUTH_State.loggedin;
	                }else{
	                	if(authorizedPermission){
							var Permissions = jQuery.grep(Session.userPermission, function(item) {
	                            return item.hasOwnProperty(authorizedPermission);
	                        });
	                        if(Permissions.length > 0){
		                        for (var i = 0,j = Permissions.length; i < j; i++) {
		                        	var permission = null;
		                        	for(item in Permissions[i]) { 
		                        		if(item == authorizedPermission) {
		                        			permission = Permissions[i][authorizedPermission];
		                        		}
		                        	}
		                    		if(permission == AUTH_Permission.noentry){
		                    			return AUTH_State.notAuthorized;
		                    		}else if(permission == AUTH_Permission.view && authPermission == AUTH_Permission.edit){
		                    			return AUTH_State.notAuthorized;
		                    		}
		                    	}
	                        }
	                        else{
								return AUTH_State.notAuthorized;
	                        }
	                	}
	                }
	            }
	            else{
	                if(authorizedRoles == USER_ROLES.admin){
	                    return AUTH_State.notAuthenticated;
	                }
	            }
        	}
        },
        isAuthorizedEl: function(authorizedPermission, authPermission){
            if(authorizedPermission){
				var Permissions = jQuery.grep(Session.userPermission, function(item) {
                    return item.hasOwnProperty(authorizedPermission);
                });
                if(Permissions.length > 0){
                    for (var i = 0,j = Permissions.length; i < j; i++) {
                    	var permission = null;
                    	for(item in Permissions[i]) { 
                    		if(item == authorizedPermission) {
                    			permission = Permissions[i][authorizedPermission];
                    		}
                    	}
                		if(permission == AUTH_Permission.noentry){
                			return false;
                		}else if(permission == AUTH_Permission.view && authPermission == AUTH_Permission.edit){
                			return false;
                		}
                	}
                }
                else{
					return false;
                }
                return true;
        	}
        	return false;
        },
        login: function(params){
            var url = 'api/login.json';
            var delay = $q.defer();
            $http({method: "Get", url: url, params: params})
                .success(function(response) {
                    if(response.state == 'success'){
                        Session.create(response.user.sessionId, response.user.userId, response.user.userRole, response.user.userPermission);
                    }
                    delay.resolve(response); 
                });
            return delay.promise;
        },
        register: function(params){
            var url = 'api/register.json';
            var delay = $q.defer();
            $http({method: "Get", url: url, params: params})
                .success(function(response) {
                    if(response.state == 'success'){
                        Session.create(response.user.sessionId, response.user.userId, response.user.userRole, response.user.userPermission);
                    }
                    delay.resolve(response); 
                });
            return delay.promise;
        },
        logout: function(){
            var params = { "userId": Session.userId, "sessionId": Session.id};
            var url = 'api/logout.json';
            var delay = $q.defer();
            $http({method: "Get", url: url, params: params})
                .success(function(response) {
                    if(response.state == 'success'){
                        Session.destroy();
                    }
                    delay.resolve(response); 
                });
            return delay.promise;
        }
    };
}]);

App.run(function($rootScope, $state, AuthService, AUTH_State){ 
    $rootScope.$on('$stateChangeStart', function (event, next, current){
        var authorizedRoles = next.data.authorizedRoles;
        var authorizedPermission = next.data.authorizedPermission;
        var authPermission = next.data.permission;
        var isAuthenticated = AuthService.isAuthenticated();        
        var authstate = AuthService.isAuthorized(isAuthenticated, authorizedRoles, authorizedPermission, authPermission);
        switch(authstate){
            case AUTH_State.loggedin:
                event.preventDefault();
                $state.go('stats');
                break;
            case AUTH_State.notAuthorized:
                event.preventDefault();
                $state.go('noauth');
                break;
            case AUTH_State.notAuthenticated:
                event.preventDefault();
                $state.go('login');
                break;
        }
    });
});

// Application Main Controller
App.controller('AppCtrl', ['$scope', '$localStorage', '$window',
    function ($scope, $localStorage, $window) {
        // Template Settings
        $scope.oneui = {
            version: '2.2', // Template version
            localStorage: false, // Enable/Disable local storage
            settings: {
                activeColorTheme: false, // Set a color theme of your choice, available: 'amethyst', 'city, 'flat', 'modern' and 'smooth'
                sidebarLeft: true, // true: Left Sidebar and right Side Overlay, false: Right Sidebar and left Side Overlay
                sidebarOpen: true, // Visible Sidebar by default (> 991px)
                sidebarOpenXs: false, // Visible Sidebar by default (< 992px)
                sidebarMini: false, // Mini hoverable Sidebar (> 991px)
                sideScroll: true, // Enables custom scrolling on Sidebar and Side Overlay instead of native scrolling (> 991px)
                headerFixed: true, // Enables fixed header
            }
        };

        // If local storage setting is enabled
        if ($scope.oneui.localStorage) {
            // Save/Restore local storage settings
            if ($scope.oneui.localStorage) {
                if (angular.isDefined($localStorage.oneuiSettings)) {
                    $scope.oneui.settings = $localStorage.oneuiSettings;
                } else {
                    $localStorage.oneuiSettings = $scope.oneui.settings;
                }
            }

            // Watch for settings changes
            $scope.$watch('oneui.settings', function () {
                // If settings are changed then save them to localstorage
                $localStorage.oneuiSettings = $scope.oneui.settings;
            }, true);
        }

        $scope.$watch('oneui.settings.sidebarHide', function () {
            // Handle Color Theme
            $scope.helpers.uiSidebarHide($scope.oneui.settings.sidebarHide);
        }, true);

        // Watch for activeColorTheme variable update
        $scope.$watch('oneui.settings.activeColorTheme', function () {
            // Handle Color Theme
            $scope.helpers.uiHandleColorTheme($scope.oneui.settings.activeColorTheme);
        }, true);

        // Watch for sideScroll variable update
        $scope.$watch('oneui.settings.sideScroll', function () {
            // Handle Scrolling
            setTimeout(function () {
                $scope.helpers.uiHandleScroll();
            }, 150);
        }, true);

        // When view content is loaded
        $scope.$on('$viewContentLoaded', function () {
            // Hide page loader
            $scope.helpers.uiLoader('hide');

            // Resize #main-container
            $scope.helpers.uiHandleMain();
        });
        
    }
]);


/*
 * Partial views controllers
 *
 */

// Side Overlay Controller
App.controller('SideOverlayCtrl', ['$scope', '$localStorage', '$window',
    function ($scope, $localStorage, $window, USER_Permissions) {
        // When view content is loaded
        $scope.$on('$includeContentLoaded', function () {
            // Handle Scrolling
            $scope.helpers.uiHandleScroll();
        });
    }
]);

// Sidebar Controller
App.controller('SidebarCtrl', ['$scope', '$localStorage', '$window', '$location', 'USER_Permissions', 'AuthService', 'AUTH_Permission',
    function ($scope, $localStorage, $window, $location, USER_Permissions, AuthService, AUTH_Permission) {
        $scope.statsMenu = $scope.permissionView = AuthService.isAuthorizedEl(USER_Permissions.stats, AUTH_Permission.view);
        $scope.storesMenu = $scope.permissionView = AuthService.isAuthorizedEl(USER_Permissions.stores, AUTH_Permission.view);
        $scope.campaignsMenu = $scope.permissionView = AuthService.isAuthorizedEl(USER_Permissions.campaigns, AUTH_Permission.view);
        $scope.accountMenu = $scope.permissionView = AuthService.isAuthorizedEl(USER_Permissions.account, AUTH_Permission.view);
        // When view content is loaded
        $scope.$on('$includeContentLoaded', function () {
            // Handle Scrolling
            $scope.helpers.uiHandleScroll();

            // Get current path to use it for adding active classes to our submenus
            $scope.path = $location.path();
        });
    }
]);

// Header Controller
App.controller('HeaderCtrl', ['$scope', '$localStorage', '$window',
    function ($scope, $localStorage, $window) {
        // When view content is loaded
        $scope.$on('$includeContentLoaded', function () {
            // Transparent header functionality
            $scope.helpers.uiHandleHeader();
        });
    }
]);

// LanguageSwitching Controller
App.controller('LanguageSwitchingCtrl', ['$scope', '$translate',
    function ($scope, $translate) {
        $scope.switching = function(lang){  
            $translate.use(lang);  
            window.localStorage.lang = lang;  
            window.location.reload();  
        };  
        $scope.cur_lang = $translate.use();
    }
]);