package codes

type ErrorCoder interface {
	ErrorCode() ErrorCode
}

// ErrorCode 表示统一的错误代码
type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

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

func GetErrorMsgByCode(code ErrorCode) string {
	return errorCodeMsg[code]
}

const (
	ErrorCodeSuccess ErrorCode = "200"

	ErrorCodeMissParameters ErrorCode = "USR400101"

	// token 错误系列以1xx开始
	ErrorCodeTokenNotFound ErrorCode = "USR403101"
	ErrorCodeInvalidToken  ErrorCode = "USR403102"
	ErrorCodeUnauthorized  ErrorCode = "USR403103"
	ErrorCodeAuthFormat    ErrorCode = "USR403104"
	ErrorCodeGenerateToekn ErrorCode = "USR500101"

	// register 注册系统以2xx开始
	ErrorCodeNameAlreadyExist     ErrorCode = "USR400201"
	ErrorCodePasswordTooShort     ErrorCode = "USR400202"
	ErrorCodeNickNameAlreadyExist ErrorCode = "USR400203"
	ErrorCodeCreateUserFail       ErrorCode = "USR500201"
	ErrorCodeRegisterNotify       ErrorCode = "USR500202"

	// login 登录系列以3xx开始
	ErrorCodeLoginFailed     ErrorCode = "USR400301"
	ErrorCodeUserNotFound    ErrorCode = "USR400302"
	ErrorCodeUpdateTokenFail ErrorCode = "USR500303"
	ErrorCodeLoginNotify     ErrorCode = "USR500304"

	// info 系列错误
	ErrorCodeInvalidUserID ErrorCode = "USR400402"

	// broadcasting 系列错误
	ErrorCodeBroadcastNotify          ErrorCode = "BRO500101"
	ErrorCodeBroadcastRoomUpdate      ErrorCode = "BRO500102"
	ErrorCodeInvalidBroadcastringUser ErrorCode = "BRO400202"
	ErrorCodeAgreement                ErrorCode = "BRO400301"
	ErrorCodeRoomAlreadyCreated       ErrorCode = "BRO400302"
	ErrorCodeRoomCreate               ErrorCode = "BRO500303"
)
