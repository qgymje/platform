package chats

import "platform/chat_center/rpc/models"

// ChatInfo chat info
type ChatInfo struct {
	ChatID    string
	Name      string
	Avatar    string
	UserID    string
	Members   []string
	CreatedAt int64
}

func modelChatToSrvChat(m *models.Chat) *ChatInfo {
	return &ChatInfo{
		ChatID:    m.GetID(),
		Name:      m.Name,
		Avatar:    m.Avatar,
		UserID:    m.UserID,
		Members:   m.MembersToStrings(),
		CreatedAt: m.CreatedAt.Unix(),
	}
}
