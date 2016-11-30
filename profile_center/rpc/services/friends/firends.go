package friends

import (
	"platform/commons/codes"
	"platform/profile_center/rpc/models"
	"platform/utils"
)

// Config config of friends list
type Config struct {
	UserID            string
	PageNum, PageSize int
}

// Friends friends
type Friends struct {
	config       *Config
	friendList   []*FriendInfo
	friendFinder *models.FriendFinder
	errorCode    codes.ErrorCode
}

// NewFriends create a new friends
func NewFriends(c *Config) *Friends {
	f := new(Friends)
	f.config = c
	f.friendList = []*FriendInfo{}
	f.friendFinder = models.NewFriendFinder().Limit(f.config.PageNum, f.config.PageSize)
	return f
}

// ErrorCode error cdoe
func (f *Friends) ErrorCode() codes.ErrorCode {
	return f.errorCode
}

// Do do the main job
func (f *Friends) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("friends.Friends.Do error:%+v", err)
		}
	}()

	if err = f.findFriends(); err != nil {
		f.errorCode = codes.ErrorCodeFriendsNotFound
		return
	}
	return
}

func (f *Friends) findFriends() (err error) {
	return f.friendFinder.Do()
}

// Result the friends result
func (f *Friends) Result() []*FriendInfo {
	modelFriends := f.friendFinder.Result()
	for i := range modelFriends {
		friendInfo := modelFriendToSrvFriend(f.config.UserID, modelFriends[i])
		f.friendList = append(f.friendList, friendInfo)
	}
	return f.friendList
}

// Count count
func (f *Friends) Count() int64 {
	return f.friendFinder.Count()
}
