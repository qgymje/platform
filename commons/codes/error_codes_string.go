package codes

var errorCodeMsg map[ErrorCode]string

func init() {
	errorCodeMsg = map[ErrorCode]string{
		ErrorCodeSuccess:        "success",
		ErrorCodeMissParameters: "missing parameters",

		// token 错误系列以1xx开始
		ErrorCodeInvalidToken:  "invalid token",
		ErrorCodeUnauthorized:  "unauthroized",
		ErrorCodeAuthFormat:    "auth format error",
		ErrorCodeGenerateToekn: "generate token error",

		// register 注册系统以2xx开始
		ErrorCodeNameAlreadyExist:     "name already exist",
		ErrorCodePasswordTooShort:     "password too short",
		ErrorCodeNickNameAlreadyExist: "nickname already exist",
		ErrorCodeCreateUserFail:       "create user fail",
		ErrorCodeRegisterNotify:       "register notify failed",
		ErrorCodeInvalidPhone:         "invalid phone number",
		ErrorCodeSMSCodeNotify:        "sms code notify failed",
		ErrorCodeSMSCreate:            "sms code create failed",

		// login 登录系列以3xx开始
		ErrorCodeLoginFailed:     "login failed",
		ErrorCodeUserNotFound:    "user not found",
		ErrorCodeUpdateTokenFail: "update token fail",
		ErrorCodeLoginNotify:     "login notify failed",

		// info 系列错误
		ErrorCodeInvalidUserID: "invalid user id",

		// broadcasting 系列错误
		ErrorCodeBroadcastNotify:          "broadcasting notify failed",
		ErrorCodeBroadcastRoomUpdate:      "broadcast room update failed",
		ErrorCodeInvalidBroadcastringUser: "invalid broadcasting user",
		ErrorCodeAgreement:                "agreement false",
		ErrorCodeRoomAlreadyCreated:       "broadcast room already created",
		ErrorCodeRoomCreate:               "broadcast room create failed",
	}
}
