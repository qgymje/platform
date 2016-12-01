# chat API list

## release notes
1. v0.0.1
    1. 2016.11.17
    2. initial API design


2. v0.0.2
    1. 2016.11.25
    2. update API design

2. v1.0.0
    1. 2016.12.1

## list:

name|desc|dev status
---|---|---
/user/recommend | [user_recommend](#user_recommend) | [NO]
/friend/ | [friend_list](#friend_list)| [YES]
/friend/| [add_friend](#add_friend)| [YES]
/chat/| [chat_list](#chat_list) | [YES]
/chat/| [create_chat](#create_chat) | [YES]
/chat/send| [send_message](#send_message) | [YES]
/notify/:token | [notify](#notify) | [YES]

---

<div id="user_list"></div>

## user list

URL: /user/recommend

AUTH: YES

METHOD: GET 

PARAMETERS:

name|type|must|desc
---|---|---|---
page | int | yes | current page
page_size| int | no | size per page, default value is 2

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
      "list": [
            {
                "userID": "57e2267ec86ab45af3d14806",
                "nickname": "hello",
                "avatar": "http://www.example.com/avatar.jpg",
                "createdAt": 1474438782
            }, 
            {
                "userID": "57e2267ec86ab45af3d14807",
                "nickname": "hello2",
                "avatar": "http://www.example.com/avatar.jpg",
                "createdAt": 1474438782
            }
      ],
      "page":1,
      "pageSize":7,
      "totalPage":2
  }
}
```

---

<div id="add_friend"></div>

## add friend

URL: /friend

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|must|desc
---|---|---|---
user_id| string| yes | friend's user_id
message| string| no | request message, could be empty

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
      "success": true,
  }
}
```
---

<div id="create_chat"></div>

## create chat

URL: /chat/

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|must|desc
---|---|---|---
members[]| array| yes | friend's user_id, at least one id
name| string| yes | chat name, default is target's username

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
      "success": true,
      "chatID": "12",
  }
}
```
---

<div id="send_message"></div>

## send _message

URL: /chat/send

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|must|desc
---|---|---|---
chat_id | string| yes | chat_id
content | string | yes | message that user typed

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
      "success": true,
      "messageID": "123",
  }
}
```
---

<div id="notify"></div>

## notify

URL: /notify/:token

AUTH: YES

METHOD: Websocket

PARAMETERS:

name|type|must|desc
---|---|---|---
token| string| yes | user token

RETURN:

RECEIVER:

```json
{"type":11000,"data":{"message_id":"36","chat_id":"1","user_id":"57e2267ec86ab45af3d14806","content":"你好","send_time":1480569427}}
```

> in this case: type id: 11000 means normal textual message in chat server


ATTENTION:
> this link should be connected when APP start, it's a global connection that used to send notifications and chats.
---
