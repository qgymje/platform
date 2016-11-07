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
		ErrorCodeBroadcastTooShort:        "broadcast too short",
		ErrorCodeBroadcastNotify:          "broadcast notify failed",
		ErrorCodeBroadcastUpdate:          "broadcast update failed",
		ErrorCodeBroadcastCreate:          "broadcast create failed",
		ErrorCodeInvalidBroadcastringUser: "invalid broadcast user",
		ErrorCodeBroadcastNotFound:        "broadcast not found",
		ErrorCodeAgreement:                "agreement false",
		ErrorCodeBroadcastIsOn:            "broadcast is already playing",
		ErrorCodeBroadcastClosed:          "broadcast is closed",

		ErrorCodeDeleteChannel: "delete channel failed",
		ErrorCodeDeleteTopic:   "delte topic failed",

		ErrorCodeAudienceUpdate: "audience update failed",

		ErrorCodeBarrageCreate: "barrage create failed",
		ErrorCodeBarrageNotify: " barrage notify failed",
		ErrorCodeBarrageFind:   "barrage find error",

		// room errors
		ErrorCodeRoomAlreadyCreated: "room already created",
		ErrorCodeRoomCreate:         "room create failed",
		ErrorCodeRoomUpdate:         "room update failed",
		ErrorCodeRoomNotFound:       "room not found",
		ErrorCodeRoomFinder:         "room finder error",

		ErrorCodeFollow:   "follow room failed",
		ErrorCodeUnfollow: "unfollow room failed",

		// game errors
		ErrorCodeGameCreate:   "game create failed",
		ErrorCodeGameNotFound: "game not found",
		ErrorCodeGameFinder:   "game find error",

		ErrorCodeUserCouponNotFound:        "user coupon not found",
		ErrorCodeUserCouponFind:            "user coupon find failed",
		ErrorCodeSendCouponNumberNotEnough: "send coupon number not enought",

		ErrorCodeUserCouponUpdate: "user coupon update failed",
		ErrorCodeTakeCouponCreate: "take coupon create failed",

		ErrorCodeSendCouponNotFound:     "send coupon not found",
		ErrorCodeSendCouponClosed:       "send coupon is closed",
		ErrorCodeSendCouponAlreadyTaken: "send coupon already taken",

		ErrorCodeGiftNotFound: "gift not found",
	}
}
