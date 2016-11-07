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

	ErrorCodeMissParameters = "USR400101"

	// upload
	ErrorCodeUpload       = "UPL400101"
	ErrorCodeUploadCreate = "UPL500101"
	ErrorCodeUploadSend   = "UPL500102"

	// token 错误系列以1xx开始
	ErrorCodeTokenNotFound = "TOK403101"
	ErrorCodeInvalidToken  = "TOK403102"
	ErrorCodeUnauthorized  = "TOK403103"
	ErrorCodeAuthFormat    = "TOK403104"
	ErrorCodeGenerateToekn = "TOK500101"

	// register 注册系统以2xx开始
	ErrorCodePhoneAlreadyExist    = "REG400200"
	ErrorCodeEmailAlreadyExist    = "REG400201"
	ErrorCodePasswordTooShort     = "REG400202"
	ErrorCodeNickNameAlreadyExist = "REG400203"
	ErrorCodeInvalidPhone         = "REG400204"
	ErrorCodeInvalidEmail         = "REG400205"
	ErrorCodeCreateUserFail       = "REG500201"
	ErrorCodeRegisterNotify       = "REG500202"

	// SMS  errorcode
	ErrorCodeSMSCodeNotify = "SMS500302"
	ErrorCodeSMSCreate     = "SMS500303"

	// Email errorcode
	ErrorCodeEmailCodeNotify = "SMS500402"
	ErrorCodeEmailCreate     = "SMS500403"

	// login 登录系列以3xx开始
	ErrorCodeLoginFailed     = "LGN400301"
	ErrorCodeUserNotFound    = "LGN400302"
	ErrorCodeUpdateTokenFail = "LGN500303"
	ErrorCodeLoginNotify     = "LGN500304"

	// info 系列错误
	ErrorCodeInvalidUserID = "INF400402"

	// user list errors
	ErrorCodeUsersNotFound = "USL400101"
	ErrorCodeUserFinder    = "USL400102"

	// broadcasting 系列错误
	ErrorCodeBroadcastTooShort        = "BRO200101"
	ErrorCodeBroadcastNotify          = "BRO500101"
	ErrorCodeBroadcastUpdate          = "BRO500102"
	ErrorCodeBroadcastCreate          = "BRO500103"
	ErrorCodeInvalidBroadcastringUser = "BRO400202"
	ErrorCodeBroadcastNotFound        = "BRO400203"
	ErrorCodeAgreement                = "BRO400301"
	ErrorCodeBroadcastIsOn            = "BRO400302"
	ErrorCodeBroadcastClosed          = "BRO400303"

	ErrorCodeDeleteChannel = "CHN500101"
	ErrorCodeDeleteTopic   = "CHN500102"

	// Audience error
	ErrorCodeAudienceUpdate = "AUD500101"

	// Barrage error
	ErrorCodeBarrageCreate = "BAR500101"
	ErrorCodeBarrageNotify = "BAR500102"
	ErrorCodeBarrageFind   = "BAR400101"

	// Room errror
	ErrorCodeRoomAlreadyCreated = "ROM400302"
	ErrorCodeRoomCreate         = "ROM500303"
	ErrorCodeRoomUpdate         = "ROM500304"
	ErrorCodeRoomNotFound       = "ROM400101"
	ErrorCodeRoomFinder         = "ROM400102"

	ErrorCodeFollow   = "ROM500401"
	ErrorCodeUnfollow = "ROM500402"

	// game errors
	ErrorCodeGameCreate   = "GAM500101"
	ErrorCodeGameNotFound = "GAM400101"
	ErrorCodeGameFinder   = "GAM500102"

	// Coupon
	ErrorCodeUserCouponNotFound        = "COP404101"
	ErrorCodeUserCouponFind            = "COP500101"
	ErrorCodeSendCouponNumberNotEnough = "COP400102"

	ErrorCodeUserCouponUpdate ErrorCode = "COP500102"

	ErrorCodeTakeCouponCreate ErrorCode = "COP500103"

	ErrorCodeSendCouponNotFound = "COP400103"
	ErrorCodeSendCouponClosed   = "COP400104"

	ErrorCodeSendCouponAlreadyTaken = "COP400105"

	// Gift
	ErrorCodeGiftNotFound = "GIF400101"
)
