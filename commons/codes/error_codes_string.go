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

		ErrorCodeGiftNotFound:          "gift not found",
		ErrorCodeSendGiftMsgApply:      "send gift message appply error",
		ErrorCodeSendGiftNotFound:      "send gift not found",
		ErrorCodeSendGiftListNotFound:  "send gift list not found",
		ErrorCodeSendGiftRank:          "send gift list rank error",
		ErrorCodeSendGiftNotify:        "send gift notify error",
		ErrorCodeSendGiftBroadcastRank: "send gift broadcast rank error",

		ErrorCodeProfileNotFound:    "profile not found",
		ErrorCodeNotEnoughSnowBall:  "not enough snow ball",
		ErrorCodeNotEnoughSnowFlake: "not enough  snow falke",
		ErrorCodeWithdraw:           "withdraw failed",
		ErrorCodeWithdrawRollback:   "withdraw rollback failed",
		ErrorCodeWithdrawCommit:     "withdraw commit failed",

		ErrorCodeRequestFriendAlreadySend: "request friend is already sent",
		ErrorCodeRequestFriendInvalidID:   "request friend invalid id",
		ErrorCodeRequestFriendSave:        "request friend save failed",
		ErrorCodeFriendSave:               "friend save failed",
		ErrorCodeRequestFriendRefuse:      "friend refuse failed",
		ErrorCodeFriendsNotFound:          "friends not found",

		ErrorCodeChatAlreadyExists: "chat already exists",
		ErrorCodeChatCreate:        "chat save failed",
		ErrorCodeChatNotExists:     "chat not exists",
		ErrorCodeChatMessageCreate: "chat message save failed",
		ErrorCodeChatMessageNotify: "chat message notfiy failed",
	}

}
