# Game Center

## 接口列表：

接口名称|接口描述|开发情况
---|---|---
list|[可玩游戏列表](#list)|[YES]
start|[开始玩一个游戏](#start)|[ YES ]
preference|[获取一个游戏的配置信息](#preference)|[YES]
update_preference|[更新某个游戏的配置](#update_preference)|[YES]


### [错误码](#error_code)

---

<div id="list"></div>

## 获取可玩游戏列表 (list)

URL: /game

METHOD: GET

Authorization: YES

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---

<div id="list"></div>
RETURN:

```json
{
  "code": "200",
  "data": {
  },
  "msg": "success"
}
```

字段名称|类型|必须|描述
---|---|---|---
code|int|是|1,正确; 其它[错误码](#error_code)
msg|string|是|错误信息描述
data| any | 是 | null 或者 object or array


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
  },
  "msg": "success"
}
```

字段名称|类型|必须|描述
---|---|---|---
code|int|是|1,正确; 其它[错误码](#error_code)
msg|string|是|错误信息描述
data| any | 是 | null 或者 object or array



