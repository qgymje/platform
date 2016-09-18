# user操作接口列表：

## 接口列表：

接口名称|接口描述|开发情况
---|---|---
register|[用户注册](#register)|[YES]
login|[用户登录](#login)|[YES]
auth|[根据token查询用户信息](#auth)|[ YES ]
logout|[用户退出](#logout)|[YES]
info|[用户信息查询](#nfo)|[ YES ]


### [错误码](#error_code)

---

<div id="register"></div>

## 用户注册(register)

URL: /user/register

METHOD: POST

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
name|string| 是| 用户名
password|string|是|密码
nickname|string|否|昵称

<div id="user_info"></div>
RETURN:

```json
{
  "code": "200",
  "data": {
    "userID": "57cf8925c86ab4291ad33f3a",
    "name": "helloworld5",
    "nickname": "hello",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzU4MTA4NTMsImlkIjoiNTdjZjg5MjVjODZhYjQyOTFhZDMzZjNhIn0.9GIy6J8-gIUf-qggS1ICqM6UiO4qYSgkCbNtXqkMkp8",
    "regTime": 1473218853
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

<div id="login"></div>

## 用户登录(login)

URL: /user/login 

METHOD: PUT

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
name|string|是|用户名
password|string|是|登录密码


RETURN:

[同用户登录](#user_info)

----

<div id="auth"></div>

### 根据token获取用户信息(auth)

URL: /user/auth/:token

METHOD: GET

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
token|string|是|用户jwt


RETURN:

[同用户登录](#user_info)


<div id="info"></div>

### 获取个人资料(info)

URL: /user/info/:user_id

METHOD: GET

PARAMETER:

字段名称|类型|必须|描述
---|---|---|---
user_id|string|是|用户id


RETURN:

[同用户登录](#user_info)

----

<div id="logout"></div>

### 退出登入(logout)

URL: /user/logout

METHOD: DELETE

RETURN:

字段名称|类型|必须|描述
---|---|---|---
Authorization | string(在header里)| 是 | bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzU4MTA5OTAsImlkIjoiNTdjZTk3NjJjODZhYjRlZjJiZjZiYjM1In0.OXaI9NqBo_y1xQMb71RcRGAyZj3OC4ouqrSWdQaTKSc

返回样例:
```json
{
  "code": 1,
  "msg": "success",
  "data": null
}
```

----

<div id="error_code"></div>

## 错误码

```
ErrorCodeSuccess ErrorCode = "200"

ErrorCodeMissParameters ErrorCode = "USR400101"

// token 错误系列以1xx开始
ErrorCodeTokenNotFound ErrorCode = "USR403101"
ErrorCodeInvalidToken  ErrorCode = "USR403102"
ErrorCodeUnauthorized  ErrorCode = "USR403103"
ErrorCodeAuthFormat    ErrorCode = "USR403104"

// register 注册系统以2xx开始
ErrorCodeNameAlreadyExist     ErrorCode = "USR400201"
ErrorCodePasswordTooShort     ErrorCode = "USR400202"
ErrorCodeNickNameAlreadyExist ErrorCode = "USR400203"
ErrorCodeCreateUserFail       ErrorCode = "USR500201"

// login 登录系列以3xx开始
ErrorCodeLoginFailed     ErrorCode = "USR400301"
ErrorCodeUserNotFound    ErrorCode = "USR400302"
ErrorCodeUpdateTokenFail ErrorCode = "USR500303"

// info 系列错误
ErrorCodeInvalidUserID ErrorCode = "USR400402"
```
