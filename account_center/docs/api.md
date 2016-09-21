# user操作接口列表：

## 接口列表：

接口名称|接口描述|开发情况
---|---|---
/user/register/sms | [获取注册码](#register_sms) | [YES]
/user/register/sms | [验证注册码](#varify_register_sms)| [YES]
/user/register/email| [获取注册码](#register_email) | [YES]
/user/register/email| [验证注册码](#varify_register_email)| [YES]
/user/register|[用户注册](#register)|[YES]
/user/login|[用户登录](#login)|[YES]
/user/auth|[根据token查询用户信息](#auth)|[ YES ]
/user/logout|[用户退出](#logout)|[YES]
/user/info|[用户信息查询](#nfo)|[ YES ]


---

<div id="register_sms"></div>

## 获取验证码

URL: /user/register/sms

AUTH: NO
> AUTH based on JWT

METHOD: POST

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
country|string| 是| 电话国际码, 如中国: +86
phone|string|是| 手机号

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "code": "692992"
  }
}
```
---

<div id="varify_register_sms"></div>

## 验证注册码

URL: /user/register/sms

AUTH: NO

METHOD: PUT

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
country|string| 是| 电话国际码, 如中国: +86
phone|string|是| 手机号
code|string|是| 用户输入的验证码


RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true
  },
}
```
---
<div id="register_email"></div>

## 获取验证码

URL: /user/register/email

AUTH: NO
> AUTH based on JWT

METHOD: POST

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
email|string|是| 邮箱

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "code": "692992"
  }
}
```
---

<div id="varify_register_email"></div>

## 验证注册码

URL: /user/register/email

AUTH: NO

METHOD: PUT

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
email|string|是| 邮箱
code|string|是| 用户输入的验证码


RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "success": true
  },
}
```

---

<div id="register"></div>

## 用户注册

URL: /user/register

AUTH: NO

METHOD: POST

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
account |string| 是| 手机号码或者邮箱
password|string|是|密码
password_confirm|string|是|密码重复
nickname|string|否|昵称

RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "userID": "57e226dac86ab45af3d14807",
    "phone": "13817782406",
    "nickname": "hello",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzcwMzA4NzQsImlkIjoiNTdlMjI2ZGFjODZhYjQ1YWYzZDE0ODA3In0.yDNkF_CL57gEmYfo5phqyzgTjYZmDZ7S_V0j_DNPqe8",
    "createdAt": 1474438874
  }
}
```
---

<div id="login"></div>

## 用户登录

URL: /user/login

AUTH: NO

METHOD: PUT

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
account|string|是|手机号或者邮箱
password|string|是|登录密码


RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "userID": "57e2267ec86ab45af3d14806",
    "phone": "13817782405",
    "nickname": "hello",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzcwMzExNTcsImlkIjoiNTdlMjI2N2VjODZhYjQ1YWYzZDE0ODA2In0.u1yhcg86Imd8yuGjcvqGwUN3yVJO8y4wWnSz9RNydTU",
    "createdAt": 1474438782
  }
}
```

----

<div id="info"></div>

### 获取用户信息
> 无法获取用户的手机, 邮箱等敏感信息

URL: /user/info/:user_id

AUTH: YES

METHOD: GET

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
user_id|string|是|用户id


RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "userID": "57e2267ec86ab45af3d14806",
    "nickname": "hello",
    "avatar": "http://www.example.com/avatar.jpg",
    "createdAt": 1474438782
  }
}
```

----

<div id="auth"></div>

### 根据token获取用户信息
> 优先通过header里的Authorization字段获取用户本人信息

URL: /user/auth/:token

AUTH: YES

METHOD: GET

PARAMETERS:

字段名称|类型|必须|描述
---|---|---|---
token|string|否|用户jwt


RETURN:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "userID": "57e226dac86ab45af3d14807",
    "phone": "13817782406",
    "nickname": "hello",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzcwMzA4NzQsImlkIjoiNTdlMjI2ZGFjODZhYjQ1YWYzZDE0ODA3In0.yDNkF_CL57gEmYfo5phqyzgTjYZmDZ7S_V0j_DNPqe8",
    "createdAt": 1474438874
  }
}
```

---

<div id="logout"></div>

### 退出登入

URL: /user/logout

AUTH: YES

METHOD: DELETE

RETURN:

字段名称|类型|必须|描述
---|---|---|---
Authorization | string(在header里)| 是 | bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzU4MTA5OTAsImlkIjoiNTdjZTk3NjJjODZhYjRlZjJiZjZiYjM1In0.OXaI9NqBo_y1xQMb71RcRGAyZj3OC4ouqrSWdQaTKSc

返回样例:
```json
{
  "code": "200",
  "msg": "success",
  "data": null
}
```

----

<div id="error_code"></div>

## 错误码

```json
ErrorCodeSuccess ErrorCode = "200"
```
