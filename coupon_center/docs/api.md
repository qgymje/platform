# coupon API list

## release notes
1. v1.0.0 
    1. 2016.11.01
    2. initial API

## list

name|desc|dev status
---|---|---
/coupon/my | [my_coupon list](#my_coupon_list) | [YES]
/coupon/broadcast/send| [copon_send](#coupon_send) | [YES]
/coupon/broadcast/take| [coupn_take](#coupon_take) | [YES]
/coupon/broadcast/stop| [coupn_stop](#coupon_stop) | [YES]

---


<div id="my_coupon_list"></div>

## my coupon list

URL: /coupon/my

AUTH: YES

METHOD: GET

PARAMETERS:

name|type|required|description
---|---|---|---
page|int| no| default value: 1
page_size| int| no | default value: 20

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "list": [
      {
        "couponID": "1",
        "name": "oppo ad1",
        "image": "https://swsdl.vivo.com.cn/vivoshop/web/dist/img/help/coupon-declare_22aa6c1.png",
        "number": 13,
        "description": "优惠券有使用期限限制，过了有效期不能使用",
        "price": 10
      },
      {
        "couponID": "2",
        "name": "kfc ad1",
        "image": "http://kfc.m.xixik.net/c/kfc/xixik_kfc_a9a2326f9b5db604.jpg",
        "number": 20,
        "description": "优惠券有使用期限限制，过了有效期不能使用",
        "price": 10
      }
    ],
    "page": 0,
    "pageSize": 20,
    "totalPage": 0
  }
}
```
---

<div id="my_coupon_list"></div>

## send coupon

URL: /coupon/broadcast/send

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
coupon_id|int|yes | coupon id from list
number| int| yes |  coupon number need to be sent
duration| int| yes |  how long ? unit is "second"

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true,
    "sendCouponID": "12"
  }
}
```

ATTENTION:

* THIS API will cause a broadcast info every 5 seconds:

```json
{"type":10001,"data":{"send_coupon_id":"7","broadcast_id":"5816fb65c86ab4a629fa5a20","remain_amount":2,"remain_time":174,"coupon_id":"1","name":"oppo ad1","image":"https://swsdl.vivo.com.cn/vivoshop/web/dist/img/help/coupon-declare_22aa6c1.png","description":"优惠券有使用期限限制，过了有效期不能使用"}}
```

> type is defined by the server

* also, it will be stopped autimatically when the duration is drained and send a broadcast info just like the `stop` API

---

<div id="coupon_take"></div>

## take coupon

URL: /coupon/broadcast/take

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
sendcoupon_id|int|yes | sendcoupon id from send action result

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true,
    "sendCouponID": "12"
  }
}
```

---


<div id="coupon_stop"></div>

## stop coupon

URL: /coupon/broadcast/stop

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
sendcoupon_id|int|yes | sendcoupon id from send action result

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true,
    "sendCouponID": "12"
  }
}
```

ATTENTION:

* STOP will cause a broadcast, msg like this:

```json
{"type":10002,"data":{"send_coupon_id":"12","broadcast_id":"5816fb65c86ab4a629fa5a20","stop_time":1477971797}}
```

> type is defined by the server
---


