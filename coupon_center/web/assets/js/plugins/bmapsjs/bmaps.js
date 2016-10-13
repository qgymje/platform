var BMaps=function(){
    if (!this) return new GMaps();
    this.map = null;
    this.options = {"el": "bmap", "centerlng": 104.030464, "centerlat": 36.060507, "icon":"assets/js/plugins/bmapsjs/img/markerico.png", "zoomlevel": 5};
    this.marker = null;
    this.lngel = null;
    this.lngel = null;
};

BMaps.prototype.
creatMap = function(){
    if(!this.map){
        this.map = new BMap.Map(this.options.el);
        var point = new BMap.Point(this.options.centerlng, this.options.centerlat);
        this.map.centerAndZoom(point, this.options.zoomlevel);
        this.map.setCenter(point);
        this.map.enableScrollWheelZoom();
        this.map.setDefaultCursor("pointer");
        this.map.setMinZoom(5); 
    }
};

BMaps.prototype.
addMarkers = function(markerArr){
    if(!this.map){
      this.creatMap();
    }
    var pointarr = [];
    var myIcon = new BMap.Icon(this.options.icon, new BMap.Size(21, 33));   
    for (var i = 0, j = markerArr.length; i < j; i++) {
      var json = markerArr[i];
      var p0 = json.lng;
      var p1 = json.lat;
      var point = new BMap.Point(p0, p1);
      pointarr[i] = point;
      var marker = new BMap.Marker(point, {"icon": myIcon, offset: new BMap.Size(0, -16)});
      var iw = new BMap.InfoWindow("<b class='iw_poi_title' title='" + json.title + "'>" + json.title + "</b><div class='iw_poi_content'>" + json.content + "</div>");
      var label = new BMap.Label(json.title, { "offset": new BMap.Size(0, -20) });
      marker.setLabel(label);
      marker.addEventListener("click", function () {
          this.openInfoWindow(iw);
      });
      marker.addEventListener("infowindowopen", function () {
          this.getLabel().hide();
      });
      marker.addEventListener("infowindowclose", function () {
          this.getLabel().show();
      });
      label.addEventListener("click", function () {
          marker.openInfoWindow(iw);
      });
      label.setStyle({
          borderColor: "#808080",
          color: "#333",
          cursor: "pointer",
          display: "block",
          maxWidth: "none"
      });
      this.map.addOverlay(marker);
      (function () {
          var _iw = new BMap.InfoWindow("<b class='iw_poi_title' title='" + json.title + "'>" + json.title + "</b><div class='iw_poi_content'>" + json.content + "</div>");
          var _marker = marker;
          _marker.addEventListener("click", function () {
              this.openInfoWindow(_iw);
          });
          _iw.addEventListener("open", function () {
              _marker.getLabel().hide();
          })
          _iw.addEventListener("close", function () {
              _marker.getLabel().show();
          })
          label.addEventListener("click", function () {
              _marker.openInfoWindow(_iw);
          })
      })()
    }
    this.map.setViewport(pointarr);
};

BMaps.prototype.
addMarker = function(lngel, latel){
    if(!this.map){
      this.creatMap();
    }
    // if(callback && typeof callback == 'function'){
    //   this.callback = callback;
    // }
    this.lngel = lngel;
    this.latel = latel;
    var self = this;
    if(!this.marker){
        var defaultpoint = this.map.getCenter();
        if(lngel.val() && lngel.val() != "" && latel.val() && latel.val() != "null"){
            defaultpoint = new BMap.Point(lngel.val(), latel.val());
        }
        var myIcon = new BMap.Icon(self.options.icon, new BMap.Size(21, 33));   
        var marker = new BMap.Marker(defaultpoint, {"icon": myIcon, offset: new BMap.Size(0, -16)});
        marker.enableDragging();
        marker.addEventListener("dragend", function(e){
            self.lngel.val(e.point.lng);
            self.latel.val(e.point.lat);
        })
        self.marker = marker;
        self.map.addOverlay(marker); 
        self.lngel.val(defaultpoint.lng);
        self.latel.val(defaultpoint.lat);
    }
    this.map.addEventListener("click", function(e){ 
        var point = e.point;  
        if(!self.marker){
            var myIcon = new BMap.Icon(self.options.icon, new BMap.Size(21, 33));   
            var marker = new BMap.Marker(point, {"icon": myIcon, offset: new BMap.Size(0, -16)});
            marker.enableDragging();
            marker.addEventListener("dragend", function(e){
                self.lngel.val(e.point.lng);
                self.latel.val(e.point.lat);
            })
            self.marker = marker;
            self.map.addOverlay(marker); 
        }else{
            self.marker.setPosition(point);
        }
        self.lngel.val(point.lng);
        self.latel.val(point.lat); 
    });
};

BMaps.prototype.removeMarker = function(){
    if(!this.map){
      this.creatMap();
    }
    if(!this.marker){
      removeOverlay(this.marker);
    }
};