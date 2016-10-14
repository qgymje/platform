# room API list

## release notes
1. v1.0.0 
    1. 2016.10.13
    2. mock api data


## list

name|desc|dev status
---|---|---
/room/ | [room list](#room_list) | [NO]
/room/ | [create/update room](#create_room) | [NO]
/room/:room_id | [room info](#room_info) | [NO]
/room/follow | [room_follow](#room_follow) | [NO]
/room/unfollow | [room_unfollow](#room_unfollow) | [NO]
/live/start | [start to broadcast](#broadcast_start) | [NO]
/live/end | [end broadcast](#broadcast_end) | [NO]
/live/enter | [enter a room](#broadcast_enter) | [NO]
/live/leave | [leave a room](#broadcast_leave) | [NO]

---


<div id="room_list"></div>

## room list

URL: /room/

AUTH: NO

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
             "roomID":"57e2267ec86ab45af3d14806",
             "name":"come on, baby",
             "cover":"http://example.com/img/fabd12adfe0cdf.jpg"
             "isPlaying": true,
             "audience": 110
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
    "roomID": "57e2267ec86ab45af3d14806"
  }
}
```
---


