# gift API list

## release notes
1. v1.0.0 
    1. 2016.11.12
    2. API finished, missing the broadcast info

2. v1.0.1 
    1. 2016.11.14
    2. add send gift broadcast info

3. v1.0.2 
    1. 2016.11.21
    2. broadcast info update: add gift info
    3. update gift image abs path
    4. remove unused API doc gift/info
    5. rename ammount to amount

4. v1.1.0 
    1. 2016.11.22
    2. add gift rank broadcast msg

5. v1.1.1 
    1. 2016.11.22
    2. add user_id field to broadcast msg

## list

name|desc|dev status
---|---|---
/gift/ | [gift_list](#gift_list) | [YES]
/gift/ | [gift_send](#gift_send) | [YES]

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
```json
{"type":10003,"data":{"broadcast_id":"5816fb65c86ab4a629fa5a20","user_id":"57e2267ec86ab45af3d14806","username":"hello0","gift_id":"3","gift_name":"手里剑","gift_image":"http://localhost:8000/v1/gift/uploads/shoulijian.png","combo":1,"amount":1,"total_price":10,"last_send_time":1479875822}}
```

> type is defined by the server, in this case, type id is 10003

ALSO ATTENTION:

```
{"type":10004,"data":[{"broadcast_id":"5816fb65c86ab4a629fa5a20","user_id":"57e2267ec86ab45af3d14806","username":"hello0","gift_id":"3","gift_name":"手里剑","gift_image":"http://localhost:8000/v1/gift/uploads/shoulijian.png","combo":1,"amount":1,"total_price":10,"last_send_time":1479875822}]}
```

> type is defined by the server, in this case, type id is 10004

---

