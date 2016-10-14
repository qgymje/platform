package codes

var errorCodeMsg map[ErrorCode]string

func init() {
	errorCodeMsg = map[ErrorCode]string{
		ErrorCodeSuccess:        "success",
		ErrorCodeMissParameters: "missing parameters",

		// upload errors
		ErrorCodeUpload:       "upload failed",
		ErrorCodeUploadCreate: "upload create failed",
		ErrorCodeUploadSend:   "upload send failed",

		// token errors
		ErrorCodeInvalidToken:  "invalid token",
		ErrorCodeUnauthorized:  "unauthroized",
		ErrorCodeAuthFormat:    "auth format error",
		ErrorCodeGenerateToekn: "generate token error",

		// register errors
		ErrorCodePhoneAlreadyExist:    "phone already exist",
		ErrorCodeEmailAlreadyExist:    "email already exist",
		ErrorCodePasswordTooShort:     "password too short",
		ErrorCodeNickNameAlreadyExist: "nickname already exist",
		ErrorCodeCreateUserFail:       "create user fail",
		ErrorCodeRegisterNotify:       "register notify failed",
		ErrorCodeInvalidPhone:         "invalid phone number",
		ErrorCodeInvalidEmail:         "invalid email address",
		ErrorCodeSMSCodeNotify:        "sms code notify failed",
		ErrorCodeSMSCreate:            "sms code create failed",
		ErrorCodeEmailCodeNotify:      "email notify failed",
		ErrorCodeEmailCreate:          "email create failed",

		// login errors
		ErrorCodeLoginFailed:     "login failed",
		ErrorCodeUserNotFound:    "user not found",
		ErrorCodeUpdateTokenFail: "update token fail",
		ErrorCodeLoginNotify:     "login notify failed",

		// info errors
		ErrorCodeInvalidUserID: "invalid user id",

		// user list errors
		ErrorCodeUsersNotFound: "user list not found",
		ErrorCodeUserFinder:    "user finder error",

		// broadcasting  errors
		ErrorCodeBroadcastNotify:          "broadcasting notify failed",
		ErrorCodeBroadcastRoomUpdate:      "broadcast room update failed",
		ErrorCodeInvalidBroadcastringUser: "invalid broadcasting user",
		ErrorCodeAgreement:                "agreement false",

		// room errors
		ErrorCodeRoomAlreadyCreated: "room already created",
		ErrorCodeRoomCreate:         "room create failed",
		ErrorCodeRoomNotFound:       "room not found",
		ErrorCodeRoomFinder:         "room finder error",

		// game errors
		ErrorCodeGameCreate:   "game create failed",
		ErrorCodeGameNotFound: "game not found",
		ErrorCodeGameFinder:   "game find error",
	}
}
