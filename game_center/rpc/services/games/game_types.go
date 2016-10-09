package games

// GameType game type
type GameType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GameTypes predefined game types
var GameTypes = []GameType{
	{1, "动作游戏"},
	{2, "冒险游戏"},
	{3, "街机游戏"},
	{4, "桌面游戏"},
	{5, "卡牌游戏"},
	{6, "娱乐场游戏"},
	{7, "休闲游戏"},
	{8, "教育游戏"},
	{9, "音乐游戏"},
	{10, "解谜游戏"},
	{11, "竞速游戏"},
	{12, "角色扮演"},
	{13, "模拟游戏"},
	{14, "体育游戏"},
	{15, "策略游戏"},
	{16, "文字游戏"},
}

func gameTypeByID(id int) string {
	for _, t := range GameTypes {
		if id == t.ID {
			return t.Name
		}
	}
	return ""
}
