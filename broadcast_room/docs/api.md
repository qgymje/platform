# room API list

## release notes
1. v0.0.1 
    1. 2016.10.13
    2. mock api data

2. v0.0.2 
    1. 2016.10.18
    2. add types API

3. v0.0.3 
    1. 2016.10.18
    2. add fullow/unfollw APIs

4. v0.0.4
    1. 2016.10.19
    2. update list/info return struct, add broadcast info if is playing

5. v0.0.5
    1. 2016.10.19
    2. add start/end APIs

6. v0.0.6
    1. 2016.10.20
    2. add enter/leave APIs
    3. update list/info/start/end API

7. v0.0.7
    1. 2016.10.23
    2. add barrage send/list APIs


## list

name|desc|dev status
---|---|---
/room/ | [room list](#room_list) | [YES]
/room/types | [room types](#room_types) | [YES]
/room/info| [room info](#room_info) | [YES]
/room/ | [create/update room](#create_room) | [YES]
/room/follow | [room_follow](#room_follow) | [YES]
/room/unfollow | [room_unfollow](#room_unfollow) | [YES]
/live/start | [start to broadcast](#broadcast_start) | [YES]
/live/end | [end broadcast](#broadcast_end) | [YES]
/live/enter | [enter a room](#broadcast_enter) | [YES]
/live/leave | [leave a room](#broadcast_leave) | [YES]
/live/join | [join](#broadcast_join) | [YES]
/barrage/| [send](#barrage_send) | [YES]
/barrage/| [list](#barrage_list) | [YES]

---


<div id="room_list"></div>

## room list

URL: /room/

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
  "data": {
    "list": [
      {
        "roomID": "5806feb1c86ab4a4ad50ec5b",
        "name": "这是一个很爽的直播哟123",
        "userName": "hello",
        "cover": "http://oaa75dzf2.bkt.clouddn.com/hellscreen.jpg",
        "isPlaying": true,
        "followNum": 88,
        "broadcast": {
          "broadcastID": "58070070c86ab4a8e2fd23eb",
          "roomID": "5806feb1c86ab4a4ad50ec5b",
          "startTime": 1476853872,
          "duration": 1880,
          "totalAudience": 1,
          "currentAudience": 1
        }
      }
    ],
    "page": 0,
    "pageSize": 20,
    "totalPage": 0
  },
  "msg": "success"
}
```
---


<div id="create_room"></div>

## create/update room

URL: /room/

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
name|string| yes| room name
cover | string | no | room cover, if empty, use old one

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true,
    "roomID": "58043dc4c86ab47026f6e04c"
  }
}
```
---

<div id="room_info"></div>

## room info

URL: /room/info

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
    "roomID": "5806feb1c86ab4a4ad50ec5b",
    "name": "这是一个很爽的直播哟123",
    "userName": "hello",
    "cover": "http://oaa75dzf2.bkt.clouddn.com/hellscreen.jpg",
    "isPlaying": true,
    "followNum": 88,
    "broadcast": {
      "broadcastID": "58070070c86ab4a8e2fd23eb",
      "roomID": "5806feb1c86ab4a4ad50ec5b",
      "startTime": 1476853872,
      "duration": 1880,
      "totalAudience": 1,
      "currentAudience": 1
    }
  },
  "msg": "success"
}
```
---

<div id="room_types"></div>

## room types

URL: /room/types

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
        "id": 1,
        "name": "英雄联盟"
      },
      {
        "id": 2,
        "name": "守望先锋"
      },
      {
        "id": 3,
        "name": "炉石传说"
      },
      {
        "id": 4,
        "name": "DOTA2"
      },
      {
        "id": 5,
        "name": "魔兽世界"
      }
    ]
  },
  "msg": "success"
}
```
---

<div id="follow"></div>

## follow

URL: /room/follow

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
room_id | string | yes | room id

RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true,
    "roomID": "58043d36c86ab46e7557177c"
  },
  "msg": "success"
}
```
---

<div id="unfollow"></div>

## unfollow

URL: /room/unfollow

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
room_id | string | yes | room id

RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true,
    "roomID": "58043d36c86ab46e7557177c"
  },
  "msg": "success"
}
```
---

<div id="start"></div>

## start

URL: /live/start

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---

RETURN:
```json
{
  "code": "200",
  "data": {
    "broadcastID": "580715bac86ab4e1a9ed1c25",
    "roomID": "5806feb1c86ab4a4ad50ec5b",
    "startTime": 1476859322
  },
  "msg": "success"
}
```
---


<div id="end"></div>

## end

URL: /live/end

AUTH: YES

METHOD: PUT

PARAMETERS:

name|type|required|description
---|---|---|---

RETURN:
```json
{
  "code": "200",
  "data": {
    "broadcastID": "580715bac86ab4e1a9ed1c25",
    "roomID": "5806feb1c86ab4a4ad50ec5b",
    "totalAudience:" 20,
    "startTime": 1476859322
  },
  "msg": "success"
}
```
---

<div id="enter"></div>

## enter

URL: /live/enter

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
broadcast_id | string | yes | broadcast id

RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true,
    "broadcastID": "58086b43c86ab49358b8c265"
  },
  "msg": "success"
}
```
---


<div id="leave"></div>

## leave

URL: /live/leave

AUTH: YES

METHOD: PUT

PARAMETERS:

name|type|required|description
---|---|---|---
broadcast_id | string | yes | broadcast id

RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true,
    "broadcastID": "58086b43c86ab49358b8c265"
  },
  "msg": "success"
}
```
---


<div id="broadcast_join"></div>

## join

URL: /live/join/:broadcast_id/:token

AUTH: YES

METHOD: GET

PARAMETERS:

name|type|required|description
---|---|---|---
broadcast_id | string | yes | broadcast id
token | string | no| token

RETURN:

> a websocket connection

---

<div id="barrage_send"></div>

## send

URL: /barrage/

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|required|description
---|---|---|---
broadcast_id | string | yes | broadcast id
text | string | yes | can not be empty
type_id| int | yes | type_id

RETURN:
```json
{
  "code": "200",
  "data": {
    "success": true
  },
  "msg": "success"
}
```

---

<div id="barrage_list"></div>

## list

URL: /barrage/

AUTH: YES

METHOD: GET

PARAMETERS:

name|type|required|description
---|---|---|---
broadcast_id | string | yes | broadcast id
start_time| string | yes | 2016-10-22T15:00:00
end_time| string | yes | 2016-10-22T15:30:00
page_size| int | no | default 50
page | int | no | default 1

RETURN:
```json
{
  "code": "200",
  "data": {
    "list": [
      {
        "broadcastID": "580b1254c86ab4ded0c401b3",
        "userID": "57e2267ec86ab45af3d14806",
        "createdAt": 1477120606,
        "username": "hello",
        "level": 108
      },
      {
        "broadcastID": "580b1254c86ab4ded0c401b3",
        "userID": "57e2267ec86ab45af3d14806",
        "createdAt": 1477120675,
        "username": "hello",
        "level": 108
      },
      {
        "broadcastID": "580b1254c86ab4ded0c401b3",
        "userID": "57e2267ec86ab45af3d14806",
        "text": "hello there",
        "createdAt": 1477120713,
        "username": "hello",
        "level": 108
      }
    ]
  },
  "msg": "success"
}
```

----
