package codes

type ErrorCoder interface {
	ErrorCode() ErrorCode
}

// ErrorCode 表示统一的错误代码
type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

func GetErrorMsgByCode(code ErrorCode) string {
	return errorCodeMsg[code]
}

const (
	ErrorCodeSuccess ErrorCode = "200"

	ErrorCodeMissParameters ErrorCode = "USR400101"

	// token 错误系列以1xx开始
	ErrorCodeTokenNotFound ErrorCode = "TOK403101"
	ErrorCodeInvalidToken  ErrorCode = "TOK403102"
	ErrorCodeUnauthorized  ErrorCode = "TOK403103"
	ErrorCodeAuthFormat    ErrorCode = "TOK403104"
	ErrorCodeGenerateToekn ErrorCode = "TOK500101"

	// register 注册系统以2xx开始
	ErrorCodePhoneAlreadyExist    ErrorCode = "REG400200"
	ErrorCodeEmailAlreadyExist    ErrorCode = "REG400201"
	ErrorCodePasswordTooShort     ErrorCode = "REG400202"
	ErrorCodeNickNameAlreadyExist ErrorCode = "REG400203"
	ErrorCodeInvalidPhone         ErrorCode = "REG400204"
	ErrorCodeInvalidEmail         ErrorCode = "REG400205"
	ErrorCodeCreateUserFail       ErrorCode = "REG500201"
	ErrorCodeRegisterNotify       ErrorCode = "REG500202"

	// SMS  errorcode
	ErrorCodeSMSCodeNotify ErrorCode = "SMS500302"
	ErrorCodeSMSCreate     ErrorCode = "SMS500303"

	// Email errorcode
	ErrorCodeEmailCodeNotify ErrorCode = "SMS500402"
	ErrorCodeEmailCreate     ErrorCode = "SMS500403"

	// login 登录系列以3xx开始
	ErrorCodeLoginFailed     ErrorCode = "LGN400301"
	ErrorCodeUserNotFound    ErrorCode = "LGN400302"
	ErrorCodeUpdateTokenFail ErrorCode = "LGN500303"
	ErrorCodeLoginNotify     ErrorCode = "LGN500304"

	// info 系列错误
	ErrorCodeInvalidUserID ErrorCode = "INF400402"

	// broadcasting 系列错误
	ErrorCodeBroadcastNotify          ErrorCode = "BRO500101"
	ErrorCodeBroadcastRoomUpdate      ErrorCode = "BRO500102"
	ErrorCodeInvalidBroadcastringUser ErrorCode = "BRO400202"
	ErrorCodeAgreement                ErrorCode = "BRO400301"
	ErrorCodeRoomAlreadyCreated       ErrorCode = "BRO400302"
	ErrorCodeRoomCreate               ErrorCode = "BRO500303"
)
