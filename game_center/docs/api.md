# Game Center

## 接口列表：

接口名称|接口描述|开发情况
---|---|--- 
types|[game type list](#types) | [NO]
list|[可玩游戏列表](#list)|[NO]
search|[搜索游戏](#search)|[NO]
start|[开始玩一个游戏](#start)|[NO]
end|[结束玩一个游戏](#end)|[NO]
preference|[获取一个游戏的配置信息](#preference)|[NO]
update_preference|[更新某个游戏的配置](#update_preference)|[NO] 

---
## game type list(types)

URL: /game/types

METHOD: GET 
Authorization: NO

RETURN:
```json
{
  "code": "200",
  "data": [
    {
      "id": 1,
      "name": "动作游戏"
    },
    {
      "id": 2,
      "name": "冒险游戏"
    },
    {
      "id": 3,
      "name": "街机游戏"
    },
    {
      "id": 4,
      "name": "桌面游戏"
    },
    {
      "id": 5,
      "name": "卡牌游戏"
    },
    {
      "id": 6,
      "name": "娱乐场游戏"
    },
    {
      "id": 7,
      "name": "休闲游戏"
    },
    {
      "id": 8,
      "name": "教育游戏"
    },
    {
      "id": 9,
      "name": "音乐游戏"
    },
    {
      "id": 10,
      "name": "解谜游戏"
    },
    {
      "id": 11,
      "name": "竞速游戏"
    },
    {
      "id": 12,
      "name": "角色扮演"
    },
    {
      "id": 13,
      "name": "模拟游戏"
    },
    {
      "id": 14,
      "name": "体育游戏"
    },
    {
      "id": 15,
      "name": "策略游戏"
    },
    {
      "id": 16,
      "name": "文字游戏"
    }
  ],
  "msg": "success"
}
```

---

<div id="list"></div>

## 获取可玩游戏列表 (list)

URL: /game

METHOD: GET

Authorization: NO

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
game_type | int | 是 | 游戏类型
page | int | 是 | 当前页, 默认第一页
page_num| int | 否 | 每页数目， 默认20

RETURN:

```json
{
  "code": "200",
  "data": {
    "list": [
      {
        "gameID": "123456789",
        "name": "world of warcraft",
        "gameTypeID": 12,
        "gameTypeName": "角色扮演",
        "description": "Wow is a grate game in history",
        "cover": "http://img4.imgtn.bdimg.com/it/u=2977877931,2204864369&fm=21&gp=0.jpg",
        "screenshots": [
          "http://img2.imgtn.bdimg.com/it/u=1803873670,2284693330&fm=11&gp=0.jpg",
          "http://img0.imgtn.bdimg.com/it/u=3951912182,2498656724&fm=11&gp=0.jpg"
        ],
        "playTimes": 10000,
        "playerNum": 1,
        "isFree": false,
        "payStatus": true
      },
      {
        "gameID": "123456788",
        "name": "league of legends",
        "gameTypeID": 1,
        "gameTypeName": "动作游戏",
        "description": "Lol is a grate game in history",
        "cover": "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcT9FXQ0uqQXlxFsd16A3YKLxCxJkDSqwLZWTBoh6psWjVR-KHkL",
        "screenshots": [
          "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJYt-xRPgp1JlWG0NzHHVDCetE-CWJfIAiW97NJ3WUamOT3QR2",
          "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTZq5K7qq9zyjvyabULU9gkREaXl85CHTA-qAmnlaWnbvHmaj6g"
        ],
        "playTimes": 20000,
        "playerNum": 10,
        "isFree": true,
        "payStatus": false
      }
    ],
    "pageNum": 1,
    "pageSize": 20,
    "totalPage": 5
  },
  "msg": "success"
}
```

字段名称|类型|描述
---|---|---|---
isFree| bool| 表示游戏是否免费
payStatus| bool|true表示已经购买， false表示未购买

---

## 搜索游戏（search)

URL: /game/search/:query

METHOD: GET

Authorization: NO

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
game_type | int | 否 | 游戏类型
page | int | 是 | 当前页, 默认第一页
page_num| int | 否 | 每页数目， 默认20


RETURN:
同 游戏列表

---

## 开始玩游戏 (start)

URL: /game/start

METHOD: POST

Authorization: YES

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
game_id | 游戏id | string | yes | 游戏id

<div id="start"></div>
RETURN:

```json
{
  "code": "200",
  "data": {
    "ip": "220.181.57.217",
    "port": 10001
  },
```

---

## 结束游戏 (end)

URL: /game/end

METHOD: POST

Authorization: YES

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
game_id | 游戏id | string | yes | 游戏id

<div id="start"></div>
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
