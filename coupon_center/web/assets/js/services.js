App.factory("AjaxService" ,['$http', '$q' ,function($http, $q){
    var action = function(method, url, params){
        console.log(params);
        var delay = $q.defer();
        $http({method: method, url: url, params: params})
            .success(function(response) {
                delay.resolve(response); 
            });
        return delay.promise;
    };
    return {
        get: function(url, params){
            var method = 'GET';
            return action(method, url, params);
        },
        post: function(url, params){
            var method = 'Post';
            return action(method, url, params);
        }
    }
}]);

App.factory("StatsService" ,['AjaxService' ,function(AjaxService){ 
    return {
        getStores: function(){
            var url = 'api/statsstore.json';
            return AjaxService.get(url);
        },
        getCoupons: function(){
            var url = 'api/statscoupon.json';
            return AjaxService.get(url);
        },
        getChartsData: function(startDate, endDate){
            var data={'startDate': startDate, 'endDate': endDate};
            var url = 'api/charts.json';
            return  AjaxService.get(url, data)
        }
    };
}]);

App.factory("StoreService" ,['AjaxService' ,function(AjaxService){
    return {
        getList: function(currentpage,pagesize){
            var data={'pagesize': pagesize, 'currentpage': currentpage};
            var url = 'api/storelist.json';
            return AjaxService.get(url, data);
        },
        getMaps: function(){
            var url = 'api/storelistmap.json';
            return AjaxService.get(url);
        },
        get: function(storeId){
            var data={'storeId': storeId};
            var url = 'api/store.json';
            return AjaxService.get(url, data);
        },
        delete: function(storeId){
            var data={'storeId': storeId};
            var url = 'api/deletestore.json';
            return AjaxService.get(url, data);
        },
        add: function(data){
            var url = 'api/addstore.json';
            return AjaxService.get(url, data);
        },
        edit: function(data){
            var url = 'api/editstore.json';
            return AjaxService.get(url, data);
        },
        checkId: function(storeId){
            var data={'storeId': storeId};
            var url = 'api/checkstoreid.json';
            return AjaxService.get(url, data);
        }
    };
}]);

App.factory("AreaService" ,['AjaxService' ,function(AjaxService){
    return {
        getProvince: function(){
            var url = "assets/js/plugins/jquery-city/json/province.json";
            return AjaxService.get(url);
        },
        getCity: function(){
            var url = "assets/js/plugins/jquery-city/json/city.json";
            return AjaxService.get(url);
        },
        getArea: function(){
            var url = "assets/js/plugins/jquery-city/json/area.json";
            return AjaxService.get(url);
        }
    };
}]);


App.factory("CampaignService" ,['AjaxService' ,function(AjaxService){
    return {
        getList: function(currentpage,pagesize){
            var data={'pagesize': pagesize, 'currentpage': currentpage};
            var url = 'api/campaignlist.json';
            return AjaxService.get(url, data);
        },
        get: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/campaign.json';
            return AjaxService.get(url, data);
        },
        add: function(data){
            var url = 'api/campaignadd.json';
            return AjaxService.get(url, data);
        },
        edit: function(data){
            var url = 'api/campaignedit.json';
            return AjaxService.get(url, data);
        },
        getStores: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/campaignstores.json';
            return AjaxService.get(url, data);
        },
        checkId: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/checkstoreid.json';
            return AjaxService.get(url, data);
        },
        getPromotions: function(campaignId, currentpage, pagesize){
            var data={'campaignId': campaignId, 'pagesize': pagesize, 'currentpage': currentpage};
            var url = 'api/campaignpromotions.json';
            return AjaxService.get(url, data);
        },
        getPromotion: function(promotionId){
            var data={'promotionId': promotionId};
            var url = 'api/promotion.json';
            return AjaxService.get(url, data);
        },
        addCoupon: function(data){
            var url = 'api/promotionadd.json';
            return AjaxService.get(url, data);
        },
        addVideo: function(data){
            var url = 'api/promotionadd.json';
            return AjaxService.get(url, data);
        },
        editCoupon: function(data){
            var url = 'api/promotionedit.json';
            return AjaxService.get(url, data);
        },
        editVideo: function(data){
            var url = 'api/promotionedit.json';
            return AjaxService.get(url, data);
        },
        checkPromotionId: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/checkstoreid.json';
            return AjaxService.get(url, data);
        },
        pause: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/campaignpause.json';
            return AjaxService.get(url, data);
        },
        resume: function(campaignId){
            var data={'campaignId': campaignId};
            var url = 'api/campaignresume.json';
            return AjaxService.get(url, data);
        },
        pausePromotion: function(promotionId){
            var data={'promotionId': promotionId};
            var url = 'api/promotionpause.json';
            return AjaxService.get(url, data);
        },
        resumePromotion: function(promotionId){
            var data={'promotionId': promotionId};
            var url = 'api/promotionresume.json';
            return AjaxService.get(url, data);
        }
    };
}]);

App.factory("AccountService" ,['AjaxService' ,function(AjaxService){
    return {
        getList: function(currentpage,pagesize){
            var data={'pagesize': pagesize, 'currentpage': currentpage};
            var url = 'api/accountlist.json';
            return AjaxService.get(url, data);
        },
        add: function(data){
            var url = 'api/accountadd.json';
            return AjaxService.get(url, data);
        },
        test: function(data){
            var url = 'api/test.php';
            return AjaxService.get(url, data);
        },
        edit: function(accountId, fun, permission){
            var data={'accountId': accountId, 'fun': fun, 'permission': permission};
            var url = 'api/accountedit.json';
            return AjaxService.get(url, data);
        },
        delete: function(accountId){
            var data={'accountId': accountId};
            var url = 'api/accountdelete.json';
            return AjaxService.get(url, data);
        }
    };
}]);