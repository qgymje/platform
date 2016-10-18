# room API list

## release notes
1. v0.0.1 
    1. 2016.10.13
    2. mock api data

1. v0.0.2 
    1. 2016.10.18
    2. add types API

1. v0.0.3 
    1. 2016.10.18
    2. add fullow/unfollw APIs


## list

name|desc|dev status
---|---|---
/room/ | [room list](#room_list) | [YES]
/room/types | [room types](#room_types) | [YES]
/room/info| [room info](#room_info) | [YES]
/room/ | [create/update room](#create_room) | [YES]
/room/follow | [room_follow](#room_follow) | [YES]
/room/unfollow | [room_unfollow](#room_unfollow) | [YES]
/live/start | [start to broadcast](#broadcast_start) | [NO]
/live/end | [end broadcast](#broadcast_end) | [NO]
/live/enter | [enter a room](#broadcast_enter) | [NO]
/live/leave | [leave a room](#broadcast_leave) | [NO]

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
  "msg": "success",
  "data": {
      "list": [
            { 
             "roomID": "58043dc4c86ab47026f6e04c",
             "name": "hello world",
             "userName": "hello world",
             "cover": "http://oaa75dzf2.bkt.clouddn.com/hellscreen.jpg",
             "isPlaying": true,
             "followNum": 17
            }
        ],
       "page":1,
       "pageSize":20,
       "totalPage":5,
  }
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
  "msg": "success",
  "data": {
        "roomID": "58043d36c86ab46e7557177c",
        "name": "hello world",
        "userName": "hello world",
        "cover": "http://oaa75dzf2.bkt.clouddn.com/hellscreen.jpg",
        "isPlaying": true,
        "followNum": 12
  }
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
