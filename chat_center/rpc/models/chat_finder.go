package models

import (
	"platform/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// ChatFinder chat finder
type ChatFinder struct {
	userID, chatID string
	offset, limit  int

	withMembers bool
	query       orm.QuerySeter

	chats   []*Chat
	members map[int64][]*Member
	chatIDs []string
}

// NewChatFinder new chat finder
func NewChatFinder() *ChatFinder {
	c := new(ChatFinder)
	c.chats = []*Chat{}
	c.members = make(map[int64][]*Member)
	c.chatIDs = []string{}
	c.query = GetDB().QueryTable(TableNameChat)
	return c
}

// Limit limit
func (c *ChatFinder) Limit(offset, limit int) *ChatFinder {
	c.offset = offset
	c.limit = limit

	c.query = c.query.Offset(int64(c.offset))
	c.query = c.query.Limit(c.limit)

	return c
}

// UserID set user id
func (c *ChatFinder) UserID(userID string) *ChatFinder {
	c.userID = userID
	c.query = c.query.Filter("user_id", userID)
	return c
}

// ChatID set chat id
func (c *ChatFinder) ChatID(chatID string) *ChatFinder {
	c.chatID = chatID
	id, _ := strconv.Atoi(chatID)
	c.query = c.query.Filter("id", id)
	return c
}

// WithMembers with member models
func (c *ChatFinder) WithMembers() *ChatFinder {
	c.withMembers = true
	return c
}

// Do the query
func (c *ChatFinder) Do() (err error) {
	defer func() {
		if err != nil {
			utils.Dump("models.ChatFinder.Do error: %+v", err)
		}
	}()

	n, err := c.query.All(&c.chats)
	if err != nil {
		return
	}
	if n == 0 {
		return ErrNotFound
	}

	if c.withMembers {
		if err = c.chatWithMembers(); err != nil {
			return
		}
	}

	return nil
}

func (c *ChatFinder) chatWithMembers() (err error) {
	if err = c.findMembers(); err != nil {
		return
	}

	for i, chat := range c.chats {
		c.chats[i].Members = c.members[chat.ID]
	}

	return
}

func (c *ChatFinder) getChatIDs() []string {
	if len(c.chats) > 0 {
		for i := range c.chats {
			c.chatIDs = append(c.chatIDs, c.chats[i].GetID())
		}
	}
	return c.chatIDs
}

func (c *ChatFinder) findMembers() (err error) {
	chatIDs := c.getChatIDs()
	members := []*Member{}
	if _, err = GetDB().QueryTable(TableNameMember).Filter("Chat__id__in", chatIDs).All(members); err != nil {
		return err
	}

	for _, member := range members {
		if _, exists := c.members[member.ID]; !exists {
			c.members[member.ID] = []*Member{}
		}
		c.members[member.ID] = append(c.members[member.ID], member)
	}

	return
}

// Result result
func (c *ChatFinder) Result() []*Chat {
	return c.chats
}

// One one
func (c *ChatFinder) One() *Chat {
	return c.chats[0]
}

// Count count
func (c *ChatFinder) Count() int64 {
	n, _ := c.query.Limit(-1).Count()
	return n
}
