package friends

import "platform/profile_center/rpc/models"

// FriendInfo friend info
type FriendInfo struct {
	UserID string
}

func modelFriendToSrvFriend(selfID string, m *models.Friend) *FriendInfo {
	f := &FriendInfo{}
	if m.FromUserID == selfID {
		f.UserID = m.ToUserID
	} else {
		f.UserID = m.FromUserID
	}
	return f
}
