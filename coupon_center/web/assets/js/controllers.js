/*
 *  Document   : controllers.js
 *  Author     : pixelcave
 *  Description: Our example controllers for demo pages
 *
 */

// Forms Validation Controller
App.controller('LoginValidationCtrl', ['$scope', '$localStorage', '$window', 'AuthService', '$state',
    function ($scope, $localStorage, $window, AuthService, $state) {
        $scope.user = {};
        // Init Material Forms Validation, for more examples you can check out https://github.com/jzaefferer/jquery-validation
        // var initValidationLogin = function(){
            var validation = jQuery('.js-validation-login').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parents('.form-group > div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'login-username': {
                        required: true,
                        minlength: 3
                    },
                    'login-password': {
                        required: true,
                        minlength: 5
                    }
                },
                messages: {
                    'login-username': {
                        required: '请输入用户名',
                        minlength: '用户名至少3个字符'
                    },
                    'login-password': {
                        required: '请输入密码',
                        minlength: '密码至少5个字符'
                    }
                }
            });
        // };

        // Init Bootstrap Forms Validation
        // initValidationLogin();
        $scope.login = function(){
            if(validation.form()){
                $scope.helpers.uiLoader("show");
                AuthService.login($scope.user).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal({
                            title: '登录成功',
                            text: res.msg,
                            type: 'success',
                            showCancelButton: false,
                            confirmButtonColor: '#d26a5c',
                            confirmButtonText: '确定',
                            closeOnConfirm: true,
                            html: false
                        }, function () {
                            $state.go('stats');
                        });

                    }else if(res.state == "error"){
                        swal('登录失败！', res.msg, 'error');
                    }
                });
            }
        }
    }
]);

// Forms Validation Controller
App.controller('RegisterValidationCtrl', ['$scope', '$localStorage', '$window', 'AuthService', '$state',
    function ($scope, $localStorage, $window, AuthService, $state) {
        $scope.user = {};
        // Init Material Forms Validation, for more examples you can check out https://github.com/jzaefferer/jquery-validation
        // var initValidationRegister = function(){
            var validation = jQuery('.js-validation-login').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parents('.form-group > div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'login-username': {
                        required: true,
                        minlength: 3
                    },
                    'login-password': {
                        required: true,
                        minlength: 5
                    }
                },
                messages: {
                    'login-username': {
                        required: '请输入用户名',
                        minlength: '用户名至少3个字符'
                    },
                    'login-password': {
                        required: '请输入密码',
                        minlength: '密码至少5个字符'
                    }
                }
            });
        // };

        // Init Bootstrap Forms Validation
        // initValidationRegister();
        $scope.register = function(){
            if(validation.form()){
                $scope.helpers.uiLoader("show");
                AuthService.register($scope.user).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal({
                            title: '注册成功',
                            text: res.msg,
                            type: 'success',
                            showCancelButton: false,
                            confirmButtonColor: '#d26a5c',
                            confirmButtonText: '确定',
                            closeOnConfirm: true,
                            html: false
                        }, function () {
                            $state.go('stats');
                        });

                    }else if(res.state == "error"){
                        swal('登录失败！', res.msg, 'error');
                    }
                });
            }
        }
    }
]);

App.controller('LogoutCtrl', ['$scope', '$localStorage', '$window', 'AuthService', '$state',
    function ($scope, $localStorage, $window, AuthService, $state) {
         var logout = function(){
            AuthService.logout().then(function(res){
                $scope.helpers.uiLoader("hide");
                if(res.state == "success"){
                    $state.go('login');
                }else if(res.state == "error"){
                    swal('退出失败！', res.msg, 'error');
                    $window.history.go(-1);
                }
            });
        };
        logout();
    }
]);



// Stats Controller
App.controller('StatsCtrl', ['$scope', '$localStorage', '$window', 'StatsService',
    function ($scope, $localStorage, $window, StatsService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        
        var nowDate = new Date();
        var n = nowDate.getTime()- 7 * 24 * 60 * 60 * 1000;
        $scope.startDate = new Date(n).toLocaleDateString();
        $scope.endDate = nowDate.toLocaleDateString();
        $scope.initChartsChartJS = function () {
            var chartLinesCon  = jQuery('.js-chartjs-lines')[0].getContext('2d');
            var chartLines;
            var chartLinesBarsRadarData;
            var globalOptions = {
                scaleFontFamily: "'Open Sans', 'Helvetica Neue', Helvetica, Arial, sans-serif",
                scaleFontColor: '#999',
                scaleFontStyle: '600',
                tooltipTitleFontFamily: "'Open Sans', 'Helvetica Neue', Helvetica, Arial, sans-serif",
                tooltipCornerRadius: 3,
                maintainAspectRatio: false,
                responsive: true
            };
            StatsService.getChartsData($scope.startDate, $scope.endDate).then(function(res){
                var chartLinesBarsRadarData = res.charts;
                chartLines = new Chart(chartLinesCon).Line(chartLinesBarsRadarData, globalOptions);
                $scope.loadCount(res.counts);
            });
        };
        $scope.loadCount = function(obj){
            jQuery("#stats-count1").attr('data-to', obj.createdCount);
            jQuery("#stats-count2").attr('data-to', obj.claimedCount);
            jQuery("#stats-count3").attr('data-to', obj.usedCount);
            jQuery('[data-toggle="countTo"]').each(function(){
                var $this       = jQuery(this);
                var $after      = $this.data('after');
                var $before     = $this.data('before');
                var $speed      = $this.data('speed') ? $this.data('speed') : 1500;
                var $interval   = $this.data('interval') ? $this.data('interval') : 15;
                $this.appear(function() {
                    $this.countTo({
                        speed: $speed,
                        refreshInterval: $interval
                    });
                });
            });
        };
        $scope.initChartsChartJS();

        StatsService.getStores().then(function(res){
            $scope.storeItems = res.storeritems;
        });

        StatsService.getCoupons().then(function(res){
            $scope.couponItems = res.couponitems;
        });
    }
]);

// Stats Controller
App.controller('StoresCtrl', ['$scope', '$localStorage', '$window', 'StoreService',
    function ($scope, $localStorage, $window, StoreService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.list = [];
        $scope.loadData = function(page, size, callback){
            StoreService.getList(page, size).then(function(res){
                callback && callback(res);
            });
        };
        $scope.delete = function(storeId){
            swal({
                    title: '是否确定删除？',
                    text: '一旦删除将无法恢复!',
                    type: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#d26a5c',
                    confirmButtonText: '是',
                    cancelButtonText: "否",
                    closeOnConfirm: false,
                    html: false
                }, function () {
                    $scope.helpers.uiLoader("show");
                    StoreService.delete(storeId).then(function(res){
                        $scope.helpers.uiLoader("hide");
                        if(res.state == "success"){
                            swal('操作成功!', res.msg, 'success');
                            $scope.getData(1);
                        }else if(res.state == "error"){
                            swal('失败！', res.msg, 'error');
                        }
                    });
                });           
        }
    }
]);

// Stats Controller
App.controller('StoresAddCtrl', ['$scope', '$state', '$localStorage', '$window', 'AreaService', 'StoreService',
    function ($scope, $state,$localStorage, $window, AreaService, StoreService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.storemodel = {};
        $scope.provinces = [];
        $scope.citys = [];
        $scope.bindcitys = [];
        AreaService.getProvince().then(function(res){
            $scope.provinces = res.province;
            AreaService.getCity().then(function(res){
                $scope.citys = res.city;
                $scope.provinceChange = function(provincename) {
                    if(provincename != ""){
                        var selectedprovince = jQuery.grep($scope.provinces, function(item) {
                            return item.name == provincename;
                        });
                        if(selectedprovince.length >0 && selectedprovince[0] && selectedprovince[0].ProID){                        
                            $scope.bindcitys = jQuery.grep($scope.citys, function(item) {
                                return item.ProID == selectedprovince[0].ProID;
                            });
                        }else{
                            $scope.bindcitys = [];
                        }
                    }else{
                        $scope.bindcitys = [];
                    }
                }
            }); 
        });
        $scope.setMap = function(){
            $scope.addMarker();
        };
        // var initValidation = function(){
            var validation = jQuery('.js-validation-store').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'storeid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                storeid: function() {
                                    return $("#storeid").val();
                                }
                            }
                        }
                    },
                    'storename': {
                        required: true,
                        minlength: 2
                    },
                    'province': {
                        required: true
                    },
                    'city': {
                        required: true
                    },
                    'address': {
                        required: true,
                        minlength: 5
                    }
                },
                messages: {
                    'storeid': {
                        required: '请输入门店编号',
                        remote: '门店编号重复'
                    },
                    'storename': {
                        required: '请输入门店名',
                        minlength: '门店名至少2个字符'
                    },
                    'province': {
                        required: '请选择省'
                    },
                    'city': {
                        required: '请选择市'
                    },
                    'address': {
                        required: '请输入门店地址',
                        minlength: '门店地址至少5个字符'
                    }
                }
            });
        // };
        // Init Bootstrap Forms Validation
        // initValidation();
        $scope.saveStore = function(){
            if(validation.form()){//
                $scope.helpers.uiLoader("show");
                StoreService.add($scope.storemodel).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal('添加成功!', res.msg, 'success');
                        $scope.storemodel = {};
                        $scope.removeMarker();
                        validation.resetForm();
                    }else if(res.state == "error"){
                        swal('添加失败！', res.msg, 'error');
                    }
                });
            }
            //jQuery('#jsvalidationstore').trigger("submit");
        };
    }
]);

// Stats Controller
App.controller('StoresEditCtrl', ['$scope', '$state', '$stateParams', '$localStorage', '$window', 'AreaService', 'StoreService',
    function ($scope, $state, $stateParams, $localStorage, $window, AreaService, StoreService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.storemodel = {};
        $scope.provinces = [];
        $scope.citys = [];
        $scope.bindcitys = [];
        $scope.storeid = $stateParams.storeId;
        AreaService.getProvince().then(function(res){
            $scope.provinces = res.province;
            AreaService.getCity().then(function(res){
                $scope.citys = res.city;
                $scope.provinceChange = function(provincename) {
                    if(provincename != ""){
                        var selectedprovince = jQuery.grep($scope.provinces, function(item) {
                            return item.name == provincename;
                        });
                        if(selectedprovince.length >0 && selectedprovince[0] && selectedprovince[0].ProID){                        
                            $scope.bindcitys = jQuery.grep($scope.citys, function(item) {
                                return item.ProID == selectedprovince[0].ProID;
                            });
                        }else{
                            $scope.bindcitys = [];
                        }
                    }else{
                        $scope.bindcitys = [];
                    }
                };                
                StoreService.get($scope.storeid).then(function(res){
                    if(res && res.storeid && res.storeid == $scope.storeid){
                        var selectedprovince = jQuery.grep($scope.provinces, function(item) {
                            return item.name == res.province;
                        });
                        $scope.bindcitys = jQuery.grep($scope.citys, function(item) {
                            return item.ProID == selectedprovince[0].ProID;
                        });
                        $scope.storemodel = res;
                        $scope.addMarker();
                    }else{
                        $state.go("stores");
                        jQuery.notify({
                            icon: 'fa fa-warning',
                            message: '该条记录不存在！',
                            url: ''
                        },
                        {
                            element: 'body',
                            type: 'warning',
                            allow_dismiss: true,
                            newest_on_top: true,
                            showProgressbar: false,
                            placement: {
                                from: 'top',
                                align: 'center'
                            },
                            offset: 20,
                            spacing: 10,
                            z_index: 1033,
                            delay: 5000,
                            timer: 1000,
                            animate: {
                                enter: 'animated fadeIn',
                                exit: 'animated fadeOutDown'
                            }
                        });
                    }
                });
            }); 
        });
        $scope.setMap = function(){
            $scope.addMarker();
        };
        // var initValidation = function(){
            var validation = jQuery('.js-validation-store').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'storeid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                storeid: function() {
                                    return $("#storeid").val();
                                }
                            }
                        }
                    },
                    'storename': {
                        required: true,
                        minlength: 2
                    },
                    'province': {
                        required: true
                    },
                    'city': {
                        required: true
                    },
                    'address': {
                        required: true,
                        minlength: 5
                    }
                },
                messages: {
                    'storeid': {
                        required: '请输入门店编号',
                        remote: '门店编号重复'
                    },
                    'storename': {
                        required: '请输入门店名',
                        minlength: '门店名至少2个字符'
                    },
                    'province': {
                        required: '请选择省'
                    },
                    'city': {
                        required: '请选择市'
                    },
                    'address': {
                        required: '请输入门店地址',
                        minlength: '门店地址至少5个字符'
                    }
                }
            });
        // };
        // Init Bootstrap Forms Validation
        // initValidation();
        $scope.saveStore = function(){
            if(validation.form()){//
                $scope.helpers.uiLoader("show");
                StoreService.edit($scope.storemodel).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal('修改成功!', res.msg, 'success');
                        validation.resetForm();
                    }else if(res.state == "error"){
                        swal('修改失败！', res.msg, 'error');
                    }
                });
            }
            //jQuery('#jsvalidationstore').trigger("submit");
        };
    }
]);

// Stats Controller
App.controller('StoresMapCtrl', ['$scope', '$localStorage', '$window', 'StoreService',
    function ($scope, $localStorage, $window, StoreService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        StoreService.getMaps().then(function(res){
            $scope.addMarkers(res.list);
        });
    }
]);

// Stats Controller
App.controller('CampaignsCtrl', ['$scope', '$localStorage', '$window', 'CampaignService',
    function ($scope, $localStorage, $window, CampaignService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.list = [];
        $scope.loadData = function(page, size, callback){
            CampaignService.getList(page, size).then(function(res){
                callback && callback(res);
            });
        };
        $scope.pause = function(campaignId){
            CampaignService.pause(campaignId).then(function(res){
                if(res.state == "success"){
                    swal('操作成功!', res.msg, 'success');
                    $scope.getData(1);
                }else if(res.state == "error"){
                    swal('操作失败！', res.msg, 'error');
                }
            });
        };
        $scope.resume = function(campaignId){
            CampaignService.resume(campaignId).then(function(res){
                if(res.state == "success"){
                    swal('操作成功!', res.msg, 'success');
                    $scope.getData(1);
                }else if(res.state == "error"){
                    swal('操作失败！', res.msg, 'error');
                }
            });
        };
    }
]);

// Stats Controller
App.controller('CampaignsAddCtrl', ['$scope', '$localStorage', '$window', 'CampaignService',
    function ($scope, $localStorage, $window, CampaignService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.campaignmodel = {};

        // var initValidation = function(){
            var validation = jQuery('.js-validation-campaign').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'campaignid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                campaignid: function() {
                                    return $("#storeid").val();
                                }
                            }
                        }
                    },
                    'campaignname': {
                        required: true,
                        minlength: 2
                    },
                    'description': {
                        required: true
                    },
                    'startdate': {
                        required: true,
                        dateISO:true
                    },
                    'enddate': {
                        required: true,
                        dateISO:true
                    }
                },
                messages: {
                    'campaignid': {
                        required: '请输入活动编号',
                        remote: '活动编号重复'
                    },
                    'campaignname': {
                        required: '请输入活动名称',
                        minlength: '活动名称至少2个字符'
                    },
                    'description': {
                        required: '请输入活动描述'
                    },
                    'startdate': {
                        required: '请输入开始时间',
                        dateISO: '请输入正确的时间格式'
                    },
                    'enddate': {
                        required: '请输入结束时间',
                        dateISO: '请输入正确的时间格式'
                    }
                }
            });
        // };
        // Init Bootstrap Forms Validation
        // initValidation();
        $scope.saveCampaign = function(){
            if(validation.form()){//
                $scope.helpers.uiLoader("show");
                CampaignService.add($scope.campaignmodel).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal('添加成功!', res.msg, 'success');
                        $scope.campaignmodel = {};
                        validation.resetForm();
                    }else if(res.state == "error"){
                        swal('添加失败！', res.msg, 'error');
                    }
                });
            }
        }
    }
]);

// Stats Controller
App.controller('CampaignsEditCtrl', ['$scope', '$localStorage', '$window', 'CampaignService', '$stateParams', '$state',
    function ($scope, $localStorage, $window, CampaignService, $stateParams, $state) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.campaignmodel = {};
        $scope.campaignid = $stateParams.campaignId;
        CampaignService.get($scope.campaignid).then(function(res){
            if(res && res.campaignid && res.campaignid == $scope.campaignid){
                $scope.campaignmodel = res;
            }else{
                $state.go("campaigns");
                jQuery.notify({
                    icon: 'fa fa-warning',
                    message: '该条记录不存在！',
                    url: ''
                },
                {
                    element: 'body',
                    type: 'warning',
                    allow_dismiss: true,
                    newest_on_top: true,
                    showProgressbar: false,
                    placement: {
                        from: 'top',
                        align: 'center'
                    },
                    offset: 20,
                    spacing: 10,
                    z_index: 1033,
                    delay: 5000,
                    timer: 1000,
                    animate: {
                        enter: 'animated fadeIn',
                        exit: 'animated fadeOutDown'
                    }
                });
            }
        });

        // var initValidation = function(){
            var validation = jQuery('.js-validation-campaign').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'campaignid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                campaignid: function() {
                                    return $("#storeid").val();
                                }
                            }
                        }
                    },
                    'campaignname': {
                        required: true,
                        minlength: 2
                    },
                    'description': {
                        required: true
                    },
                    'startdate': {
                        required: true,
                        dateISO:true
                    },
                    'enddate': {
                        required: true,
                        dateISO:true
                    }
                },
                messages: {
                    'campaignid': {
                        required: '请输入活动编号',
                        remote: '活动编号重复'
                    },
                    'campaignname': {
                        required: '请输入活动名称',
                        minlength: '活动名称至少2个字符'
                    },
                    'description': {
                        required: '请输入活动描述'
                    },
                    'startdate': {
                        required: '请输入开始时间',
                        dateISO: '请输入正确的时间格式'
                    },
                    'enddate': {
                        required: '请输入结束时间',
                        dateISO: '请输入正确的时间格式'
                    }
                }
            });
        // };
        // Init Bootstrap Forms Validation
        // initValidation();
        $scope.saveCampaign = function(){
            if(validation.form()){//
                $scope.helpers.uiLoader("show");
                CampaignService.edit($scope.campaignmodel).then(function(res){
                    $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal('修改成功!', res.msg, 'success');
                        validation.resetForm();
                    }else if(res.state == "error"){
                        swal('修改失败！', res.msg, 'error');
                    }
                });
            }
        }
    }
]);

// Stats Controller
App.controller('CampaignsShowCtrl', ['$scope', '$localStorage', '$window', 'CampaignService', '$stateParams', '$state',
    function ($scope, $localStorage, $window, CampaignService, $stateParams, $state) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.campaignmodel = {};
        $scope.campaignid = $stateParams.campaignId;
        $scope.list = [];
        $scope.loadData = function(page, size, callback){
            CampaignService.getPromotions($scope.campaignid, page, size).then(function(res){
                callback && callback(res);
            });
        };
        CampaignService.get($scope.campaignid).then(function(res){
            if(res && res.campaignid && res.campaignid == $scope.campaignid){
                $scope.campaignmodel = res; 
            }else{
                $state.go("campaigns");
                jQuery.notify({
                    icon: 'fa fa-warning',
                    message: '该条记录不存在！',
                    url: ''
                },
                {
                    element: 'body',
                    type: 'warning',
                    allow_dismiss: true,
                    newest_on_top: true,
                    showProgressbar: false,
                    placement: {
                        from: 'top',
                        align: 'center'
                    },
                    offset: 20,
                    spacing: 10,
                    z_index: 1033,
                    delay: 5000,
                    timer: 1000,
                    animate: {
                        enter: 'animated fadeIn',
                        exit: 'animated fadeOutDown'
                    }
                });
            }
        });
    }
]);

// Stats Controller
App.controller('CampaignsAddPromotionsCtrl', ['$scope', '$localStorage', '$window', 'CampaignService', '$stateParams', '$state',
    function ($scope, $localStorage, $window, CampaignService, $stateParams, $state) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.campaignmodel = {};
        $scope.campaignid = $stateParams.campaignId;
        $scope.promotionmodel = {};
        $scope.couponmodel = {};
        $scope.videmodel = {};
        $scope.storesmodel = [];
        $scope.selected = [];
        var updateSelected = function (action, id) {
            if (action == 'add' & $scope.selected.indexOf(id) == -1){
                $scope.selected.push(id);
            }
            if (action == 'remove' && $scope.selected.indexOf(id) != -1){
                $scope.selected.splice($scope.selected.indexOf(id), 1);
            }
        };
        $scope.updateSelection = function ($event, id) {
            var checkbox = $event.target;
            var action = (checkbox.checked ? 'add' : 'remove');
            updateSelected(action, id);
        };
        $scope.selectAll = function ($event) {
            var checkbox = $event.target;
            var action = (checkbox.checked ? 'add' : 'remove');
            for (var i = 0,j = $scope.storesmodel.length; i < j; i++) {
                var store = $scope.storesmodel[i];
                updateSelected(action, store.storeid);
            }
        };
        $scope.isSelected = function (id) {
            return $scope.selected.indexOf(id) >= 0;
        };
        $scope.isSelectedAll = function () {
            return $scope.selected.length === $scope.storesmodel.length;
        };
        CampaignService.getStores($scope.campaignid).then(function(res){
            $scope.storesmodel = res.list;
        });

        CampaignService.get($scope.campaignid).then(function(res){
            if(res && res.campaignid && res.campaignid == $scope.campaignid){
                $scope.campaignmodel = res;
                $scope.couponmodel.campaignid = $scope.campaignid;
                $scope.videmodel.campaignid = $scope.campaignid;
            }else{
                $state.go("campaigns");
                jQuery.notify({
                    icon: 'fa fa-warning',
                    message: '该条记录不存在！',
                    url: ''
                },
                {
                    element: 'body',
                    type: 'warning',
                    allow_dismiss: true,
                    newest_on_top: true,
                    showProgressbar: false,
                    placement: {
                        from: 'top',
                        align: 'center'
                    },
                    offset: 20,
                    spacing: 10,
                    z_index: 1033,
                    delay: 5000,
                    timer: 1000,
                    animate: {
                        enter: 'animated fadeIn',
                        exit: 'animated fadeOutDown'
                    }
                });
            }
        });        

        $scope.uploadImg = function (file) {
            var progress = jQuery("#couponpictureprogress");
            var progressbar = progress.find(".progress-bar");
            progressbar.fadeIn(1);
            Upload.upload({
                url: 'upload/url',
                data: {file: file}
            }).then(function (resp) {
                $scope.couponmodel.couponpicture = resp.filepath;
            }, function (resp) {
                swal('上传失败！', resp.status, 'error');
            }, function (evt) {
                var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                if(progressPercentage > 5){
                    progressbar.attr('aria-valuenow',progressPercentage)
                    .css('width',progressPercentage + '%')
                    .text(progressPercentage + '%');
                }
                if(progressPercentage >= 100){
                    progress.fadeOut(1);
                }
            });
        };
        $scope.uploadVideo = function (file) {
            var progress = jQuery("#videoprogress");
            var progressbar = progress.find(".progress-bar");
            progressbar.fadeIn(1);
            Upload.upload({
                url: 'upload/url',
                data: {file: file}
            }).then(function (resp) {
                $scope.videmodel.uploadvideo = resp.filepath;
            }, function (resp) {
                swal('上传失败！', resp.status, 'error');
            }, function (evt) {
                var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                if(progressPercentage > 5){
                    progressbar.attr('aria-valuenow',progressPercentage)
                    .css('width',progressPercentage + '%')
                    .text(progressPercentage + '%');
                }
                if(progressPercentage >= 100){
                    progress.fadeOut(1);
                }
            });
        };

        var validation = jQuery('.js-validation-promotion').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'promotionid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                promotionid: function() {
                                    return $("#promotionid").val();
                                }
                            }
                        }
                    },
                    'promotiontype': {
                        required: true
                    },
                    'couponname': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'coupondescription': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponvalue': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponcount': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        digits: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponunitpriceforpromotion': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'videoname': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'bouncelink': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'videocount': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        },
                        digits: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'videounitpriceforpromotion': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    }
                },
                messages: {
                    'promotionid': {
                        required: "请输入推广编号",
                        remote: "推广编号重复"
                    },
                    'promotiontype': {
                        required: "请选择推广类型"
                    },
                    'couponname': {
                        required: "请输入优惠券名称"
                    },
                    'coupondescription': {
                        required: "请输入优惠券描述"
                    },
                    'couponvalue': {
                        required: "请输入优惠券价值",
                        number: "请输入数字"
                    },
                    'couponcount': {
                        required: "请输入优惠券数量",
                        digits: "请输入整数"
                    },
                    'couponunitpriceforpromotion': {
                        required: "请输入优惠券推广单价",
                        number: "请输入数字"
                    },
                    'videoname': {
                        required: "请输入视频名称"
                    },
                    'bouncelink': {
                        required: "请输入视频跳转链接"
                    },
                    'videocount': {
                        required: "请输入视频播放次数预期",
                        digits: "请输入整数"
                    },
                    'videounitpriceforpromotion': {
                        required: "请输入视频播放单价",
                        number:"请输入数字"
                    }
                }
            });

        $scope.savePromotion = function(){
            if(validation.form()){
                if($scope.promotionmodel.promotiontype == "coupon"){
                    if($scope.couponmodel.couponpicture && $scope.couponmodel.couponpicture != ""){
                        if($scope.selected.length > 0){
                            $scope.couponmodel.selectedstores = $scope.selected;
                            $scope.helpers.uiLoader("show");
                            $scope.couponmodel.campaignid = $scope.campaignmodel.campaignid;
                            $scope.couponmodel.promotionid = $scope.promotionmodel.promotionid;
                            $scope.couponmodel.promotiontype = $scope.promotionmodel.promotiontype;
                            CampaignService.addCoupon($scope.couponmodel).then(function(res){
                                $scope.helpers.uiLoader("hide");
                                if(res.state == "success"){
                                    swal('添加成功!', res.msg, 'success');
                                    $scope.promotionmodel = {};
                                    $scope.couponmodel = {};
                                    $scope.videmodel = {};
                                    validation.resetForm();
                                }else if(res.state == "error"){
                                    swal('添加失败！', res.msg, 'error');
                                }
                            });         
                        }else{
                            swal('保存失败！', "你至少选择一个可用门店", 'warning');
                        }                 
                    }else{
                        swal('保存失败！', "请上传优惠券图片", 'warning');
                    }
                }else if($scope.promotionmodel.promotiontype == "video"){
                    if($scope.videmodel.uploadvideo && $scope.videmodel.uploadvideo != ""){
                        $scope.helpers.uiLoader("show");
                        $scope.videmodel.campaignid = $scope.campaignmodel.campaignid;
                        $scope.videmodel.promotionid = $scope.promotionmodel.promotionid;
                        $scope.videmodel.promotiontype = $scope.promotionmodel.promotiontype;
                        CampaignService.addVideo($scope.videmodel).then(function(res){
                            $scope.helpers.uiLoader("hide");
                            if(res.state == "success"){
                                swal('添加成功!', res.msg, 'success');
                                $scope.promotionmodel = {};
                                $scope.couponmodel = {};
                                $scope.videmodel = {};
                                validation.resetForm();
                            }else if(res.state == "error"){
                                swal('添加失败！', res.msg, 'error');
                            }
                        });
                    }else{
                        swal('保存失败！', "请上传视频", 'warning');
                    }
                }
            }
        }
        
    }
]);

// Stats Controller
App.controller('CampaignsEditPromotionsCtrl', ['$scope', '$localStorage', '$window', 'CampaignService', '$stateParams', '$state',
    function ($scope, $localStorage, $window, CampaignService, $stateParams, $state) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.campaignmodel = {};
        $scope.campaignid = $stateParams.campaignId;
        $scope.promotionid = $stateParams.promotionId;
        $scope.promotionmodel = {};
        $scope.couponmodel = {};
        $scope.videmodel = {};
        $scope.storesmodel = [];
        $scope.selected = [];
        var updateSelected = function (action, id) {
            if (action == 'add' & $scope.selected.indexOf(id) == -1){
                $scope.selected.push(id);
            }
            if (action == 'remove' && $scope.selected.indexOf(id) != -1){
                $scope.selected.splice($scope.selected.indexOf(id), 1);
            }
        };
        $scope.updateSelection = function ($event, id) {
            var checkbox = $event.target;
            var action = (checkbox.checked ? 'add' : 'remove');
            updateSelected(action, id);
        };
        $scope.selectAll = function ($event) {
            var checkbox = $event.target;
            var action = (checkbox.checked ? 'add' : 'remove');
            for (var i = 0,j = $scope.storesmodel.length; i < j; i++) {
                var store = $scope.storesmodel[i];
                updateSelected(action, store.storeid);
            }
        };
        $scope.isSelected = function (id) {
            return $scope.selected.indexOf(id) >= 0;
        };
        $scope.isSelectedAll = function () {
            return $scope.selected.length === $scope.storesmodel.length;
        };
        $scope.loadSelect = function () {
            for (var i = 0,j = $scope.selected.length; i < j; i++) {
                var storeid = $scope.selected[i];
                updateSelected('add', storeid);
            }
        };
        CampaignService.get($scope.campaignid).then(function(res){
            if(res && res.campaignid && res.campaignid == $scope.campaignid){
                $scope.campaignmodel = res;
                $scope.couponmodel.campaignid = $scope.campaignid;
                $scope.videmodel.campaignid = $scope.campaignid;
            }else{
                $state.go("campaigns");
                jQuery.notify({
                    icon: 'fa fa-warning',
                    message: '该条记录不存在！',
                    url: ''
                },
                {
                    element: 'body',
                    type: 'warning',
                    allow_dismiss: true,
                    newest_on_top: true,
                    showProgressbar: false,
                    placement: {
                        from: 'top',
                        align: 'center'
                    },
                    offset: 20,
                    spacing: 10,
                    z_index: 1033,
                    delay: 5000,
                    timer: 1000,
                    animate: {
                        enter: 'animated fadeIn',
                        exit: 'animated fadeOutDown'
                    }
                });
            }
        });

        CampaignService.getStores($scope.campaignid).then(function(res){
            $scope.storesmodel = res.list;
            CampaignService.getPromotion($scope.promotionid).then(function(res){
                if(res && res.campaignid && res.campaignid == $scope.campaignid && res.promotionid && res.promotionid == $scope.promotionid){
                    $scope.promotionmodel.promotionid = res.promotionid;
                    $scope.promotionmodel.promotiontype = res.promotiontype;
                    if(res.promotiontype == "coupon"){
                        $scope.couponmodel = res;
                        $scope.selected = res.selectedstores;
                        $scope.loadSelect();
                    }else if(res.promotiontype == "video"){
                        $scope.videmodel = res;
                    }
                }else{
                    $state.go("campaigns");
                    jQuery.notify({
                        icon: 'fa fa-warning',
                        message: '该条记录不存在！',
                        url: ''
                    },
                    {
                        element: 'body',
                        type: 'warning',
                        allow_dismiss: true,
                        newest_on_top: true,
                        showProgressbar: false,
                        placement: {
                            from: 'top',
                            align: 'center'
                        },
                        offset: 20,
                        spacing: 10,
                        z_index: 1033,
                        delay: 5000,
                        timer: 1000,
                        animate: {
                            enter: 'animated fadeIn',
                            exit: 'animated fadeOutDown'
                        }
                    });
                }
            });
        });

        $scope.uploadImg = function (file) {
            var progress = jQuery("#couponpictureprogress");
            var progressbar = progress.find(".progress-bar");
            progressbar.fadeIn(1);
            Upload.upload({
                url: 'upload/url',
                data: {file: file}
            }).then(function (resp) {
                $scope.couponmodel.couponpicture = resp.filepath;
            }, function (resp) {
                swal('上传失败！', resp.status, 'error');
            }, function (evt) {
                var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                if(progressPercentage > 5){
                    progressbar.attr('aria-valuenow',progressPercentage)
                    .css('width',progressPercentage + '%')
                    .text(progressPercentage + '%');
                }
                if(progressPercentage >= 100){
                    progress.fadeOut(1);
                }
            });
        };
        $scope.uploadVideo = function (file) {
            var progress = jQuery("#videoprogress");
            var progressbar = progress.find(".progress-bar");
            progressbar.fadeIn(1);
            Upload.upload({
                url: 'upload/url',
                data: {file: file}
            }).then(function (resp) {
                $scope.videmodel.uploadvideo = resp.filepath;
            }, function (resp) {
                swal('上传失败！', resp.status, 'error');
            }, function (evt) {
                var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                if(progressPercentage > 5){
                    progressbar.attr('aria-valuenow',progressPercentage)
                    .css('width',progressPercentage + '%')
                    .text(progressPercentage + '%');
                }
                if(progressPercentage >= 100){
                    progress.fadeOut(1);
                }
            });
        };

        var validation = jQuery('.js-validation-promotion').validate({
                ignore: [],
                errorClass: 'help-block text-right animated fadeInDown',
                errorElement: 'div',
                errorPlacement: function(error, e) {
                    jQuery(e).parent('div').append(error);
                },
                highlight: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error').addClass('has-error');
                    elem.closest('.help-block').remove();
                },
                success: function(e) {
                    var elem = jQuery(e);

                    elem.closest('.form-group').removeClass('has-error');
                    elem.closest('.help-block').remove();
                },
                rules: {
                    'promotionid': {
                        required: true,
                        remote: {
                            url: "api/checkstoreid.json",
                            type: "post",
                            dataType: "json",
                            data: {
                                promotionid: function() {
                                    return $("#promotionid").val();
                                }
                            }
                        }
                    },
                    'promotiontype': {
                        required: true
                    },
                    'couponname': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'coupondescription': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponvalue': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponcount': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        digits: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'couponunitpriceforpromotion': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "coupon";
                        }
                    },
                    'videoname': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'bouncelink': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'videocount': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        },
                        digits: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    },
                    'videounitpriceforpromotion': {
                        required: function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        },
                        number:function(){
                            return $scope.promotionmodel.promotiontype == "video";
                        }
                    }
                },
                messages: {
                    'promotionid': {
                        required: "请输入推广编号",
                        remote: "推广编号重复"
                    },
                    'promotiontype': {
                        required: "请选择推广类型"
                    },
                    'couponname': {
                        required: "请输入优惠券名称"
                    },
                    'coupondescription': {
                        required: "请输入优惠券描述"
                    },
                    'couponvalue': {
                        required: "请输入优惠券价值",
                        number: "请输入数字"
                    },
                    'couponcount': {
                        required: "请输入优惠券数量",
                        digits: "请输入整数"
                    },
                    'couponunitpriceforpromotion': {
                        required: "请输入优惠券推广单价",
                        number: "请输入数字"
                    },
                    'videoname': {
                        required: "请输入视频名称"
                    },
                    'bouncelink': {
                        required: "请输入视频跳转链接"
                    },
                    'videocount': {
                        required: "请输入视频播放次数预期",
                        digits: "请输入整数"
                    },
                    'videounitpriceforpromotion': {
                        required: "请输入视频播放单价",
                        number:"请输入数字"
                    }
                }
            });

        $scope.savePromotion = function(){
            if(validation.form()){
                if($scope.promotionmodel.promotiontype == "coupon"){
                    if($scope.couponmodel.couponpicture && $scope.couponmodel.couponpicture != ""){
                        if($scope.selected.length > 0){
                            $scope.couponmodel.selectedstores = $scope.selected;
                            $scope.helpers.uiLoader("show");
                            $scope.couponmodel.campaignid = $scope.campaignmodel.campaignid;
                            $scope.couponmodel.promotionid = $scope.promotionmodel.promotionid;
                            $scope.couponmodel.promotiontype = $scope.promotionmodel.promotiontype;
                            CampaignService.editCoupon($scope.couponmodel).then(function(res){
                                $scope.helpers.uiLoader("hide");
                                if(res.state == "success"){
                                    swal('修改成功!', res.msg, 'success');
                                    validation.resetForm();
                                }else if(res.state == "error"){
                                    swal('修改失败！', res.msg, 'error');
                                }
                            }); 
                        }else{
                            swal('保存失败！', "你至少选择一个可用门店", 'warning');
                        }                                    
                    }else{
                        swal('保存失败！', "请上传优惠券图片", 'warning');
                    }
                }else if($scope.promotionmodel.promotiontype == "video"){
                    if($scope.videmodel.uploadvideo && $scope.videmodel.uploadvideo != ""){
                        $scope.helpers.uiLoader("show");
                        $scope.videmodel.campaignid = $scope.campaignmodel.campaignid;
                        $scope.videmodel.promotionid = $scope.promotionmodel.promotionid;
                        $scope.videmodel.promotiontype = $scope.promotionmodel.promotiontype;
                        CampaignService.editVideo($scope.videmodel).then(function(res){
                            $scope.helpers.uiLoader("hide");
                            if(res.state == "success"){
                                swal('修改成功!', res.msg, 'success');
                                $scope.promotionmodel = {};
                                $scope.couponmodel = {};
                                $scope.videmodel = {};
                                validation.resetForm();
                            }else if(res.state == "error"){
                                swal('修改失败！', res.msg, 'error');
                            }
                        });
                    }else{
                        swal('保存失败！', "请上传视频", 'warning');
                    }
                }
            }
        }
        
    }
]);

// Stats Controller
App.controller('AccountCtrl', ['$scope', '$localStorage', '$window', 'AccountService',
    function ($scope, $localStorage, $window, AccountService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.list = [];
        $scope.loadData = function(page, size, callback){
            AccountService.getList(page, size).then(function(res){
                callback && callback(res);
            });
        };
        $scope.delete = function(accountId){
            AccountService.delete(accountId).then(function(res){
                if(res.state == "success"){
                    swal('操作成功!', res.msg, 'success');
                    $scope.getData(1);
                }else if(res.state == "error"){
                    swal('操作失败！', res.msg, 'error');
                }
            });
        };
        $scope.selectChange = function(accountId, fun, permission){
            AccountService.edit(accountId, fun, permission).then(function(res){
                if(res.state == "success"){
                    swal('操作成功!', res.msg, 'success');
                }else if(res.state == "error"){
                    swal('操作失败！', res.msg, 'error');
                }
            });
        };
    }
]);

// Stats Controller
App.controller('AccountAddCtrl', ['$scope', '$localStorage', '$window', 'AccountService',
    function ($scope, $localStorage, $window, AccountService) {
        // Init jQuery AutoComplete example, for more examples you can check out https://github.com/Pixabay/jQuery-autoComplete
        $scope.accountmodel = {
            "like":{
                "boll": true,
                "fds": true
            }
        };
        // var validation = jQuery('.js-validation-account').validate({
        //         ignore: [],
        //         errorClass: 'help-block text-right animated fadeInDown',
        //         errorElement: 'div',
        //         errorPlacement: function(error, e) {
        //             jQuery(e).parent('div').append(error);
        //         },
        //         highlight: function(e) {
        //             var elem = jQuery(e);

        //             elem.closest('.form-group').removeClass('has-error').addClass('has-error');
        //             elem.closest('.help-block').remove();
        //         },
        //         success: function(e) {
        //             var elem = jQuery(e);

        //             elem.closest('.form-group').removeClass('has-error');
        //             elem.closest('.help-block').remove();
        //         },
        //         rules: {
        //             'accountid': {
        //                 required: true,
        //                 remote: {
        //                     url: "api/checkstoreid.json",
        //                     type: "post",
        //                     dataType: "json",
        //                     data: {
        //                         accountid: function() {
        //                             return $("#accountid").val();
        //                         }
        //                     }
        //                 }
        //             },
        //             'accountname': {
        //                 required: true
        //             },
        //             'accountpwd':{
        //                 required: true
        //             },
        //             'accountstats': {
        //                 required: true
        //             },
        //             'accountstores': {
        //                 required: true
        //             },
        //             'accountcampaign': {
        //                 required: true
        //             },
        //             'accountaccount': {
        //                 required: true
        //             }
        //         },
        //         messages: {
        //             'accountid': {
        //                 required: '请输入账户编号',
        //                 remote: '账户编号重复'
        //             },
        //             'accountname': {
        //                 required: '请输入账户用户名'
        //             },
        //             'accountpwd': {
        //                 required: '请输入账户密码'
        //             },
        //             'accountstats': {
        //                 required: '请选择数据功能权限'
        //             },
        //             'accountstores': {
        //                 required: '请选择门店功能权限'
        //             },
        //             'accountcampaign': {
        //                 required: '请选择活动功能权限'
        //             },
        //             'accountaccount': {
        //                 required: '请选择账户管理权限'
        //             }
        //         }
        //     });
        $scope.saveAccount = function(){
            // if(validation.form()){//
                // $scope.helpers.uiLoader("show");
                AccountService.test($scope.accountmodel).then(function(res){
                    // $scope.helpers.uiLoader("hide");
                    if(res.state == "success"){
                        swal('添加成功!', res.msg, 'success');
                    //     $scope.accountmodel = {};
                    //     validation.resetForm();
                    }else if(res.state == "error"){
                        swal('添加失败！', res.msg, 'error');
                    }
                });
            // }
        }
    }
]);