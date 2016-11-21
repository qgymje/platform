# chat API list

## release notes
1. v0.0.1
    1. 2016.11.17
    2. initial API design

## 接口列表：

接口名称|接口描述|开发情况
---|---|---
/user/ | [user_list](#user_list) | [NO]
/profile/friend/ | [friend_list](#friend_list)| [NO]
/profile/friend/ | [add_friend](#add_friend)| [NO]
/chat/| [chat_list](#chat_list) | [NO]
/chat/| [create_chat](#create_chat) | [NO]
/chat/send| [send_message](#send_message) | [NO]
/notify/:token | [notify](#notify) | [NO]

---

<div id="user_list"></div>

## user list

URL: /user/

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

URL: /profile/friend

AUTH: YES

METHOD: POST

PARAMETERS:

name|type|must|desc
---|---|---|---
user_id| string| yes | friend's user_id

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
user_id| string| yes | friend's user_id

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
{"type": 11000, "data": {"chat_id": "123, "from_user_id":"57e226dac86ab45af3d1480","nickname":"somebody", "avatar":"http://www.example.com/avatar_of_somebody", content":"say hi", "created_at":147443887}}
```

> type id: 11000 means normal textual message in chat server
---
