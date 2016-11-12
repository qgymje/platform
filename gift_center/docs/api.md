# gift API list

## release notes
1. v1.0.0 
    1. 2016.11.12
    2. API finished, missing the broadcast info

## list

name|desc|dev status
---|---|---
/gift/ | [gift_list](#gift_list) | [YES]
/gift/ | [gift_send](#gift_send) | [YES]
/gift/info/:id | [gift_info](#gift_send) | [YES]

---

<div id="gift_list"></div>

## gift list

URL: /gift/

AUTH: YES

METHOD: GET

PARAMETERS:

name|type|required|description
---|---|---|---

RETURN:
```json
{
  "code": "200",
  "data": {
    "list": [
      {
        "giftID": "1",
        "name": "神龙",
        "image": "/v1/gift/uploads/xianglong.png",
        "snowBall": 600,
        "combo": 2
      },
      {
        "giftID": "2",
        "name": "仙丹",
        "image": "/v1/gift/uploads/xiandang.png",
        "snowFlake": 1,
        "combo": 1
      },
      {
        "giftID": "3",
        "name": "手里剑",
        "image": "/v1/gift/uploads/shoulijian.png",
        "snowFlake": 10,
        "combo": 1
      },
      {
        "giftID": "4",
        "name": "神甲",
        "image": "/v1/gift/uploads/shenjia.png",
        "snowBall": 1,
        "combo": 1
      },
      {
        "giftID": "5",
        "name": "葫芦",
        "image": "/v1/gift/uploads/hulu.png",
        "snowBall": 200,
        "combo": 2
      },
      {
        "giftID": "6",
        "name": "飞毛腿",
        "image": "/v1/gift/uploads/feimaotui.png",
        "snowBall": 50,
        "combo": 2
      }
    ]
  },
  "msg": "success"
}
```

---

<div id="gift_send"></div>

## gift list

URL: /gift/

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
gift_id | string | yes | gift id 
room_id | string | yes | room id: broadcast room id
type_id | int | yes | type id is the broadcast msg's type field


RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true,
    "sendGiftID": "58"
  },
  "msg": "success"
}

```

ATTENTION:

* This request will cause a broadcast to the room
```
// still being formulated
```

ALSO ATTENTION:

* send gift action will cause the gift rand update, gift rand message will broadcast every 5 second
```
// still being formulated
```

> type is defined by the server, in this case, type id is 10003

---

