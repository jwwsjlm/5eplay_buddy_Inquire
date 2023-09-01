package utils

import (
	"5e/config"
	"errors"
)

func ParseToPlayerList(s []byte, c *config.PlayerInfo) error {
	str := string(s)
	n, e := GetTextBetween(str, "\"username-tooltips\">", "</span>") //游戏名字获取
	//fmt.Println(str)
	if e != nil {
		return errors.New("未找到name")
	}
	c.Name = n
	n, e = GetTextBetween(str, "_g_player_domain ='", "'") //游戏id获取
	if e != nil {
		return errors.New("未找到游戏id")
	}
	c.PlayerID = n

	n, e = GetTextBetween(str, "',link:'", "'") //个人简介地址
	if e != nil {
		return errors.New("未找到个人简介地址")
	}
	c.ProfileURL = n
	n, e = GetTextBetween(str, "avatar-img position-center\" src=\"", "\" width") //头像获取
	if e != nil {
		return errors.New("未找到头像地址")
	}
	c.AvatarURL = n
	return nil

}
func OrganizePlayerInfo(str string, m *config.CurrentGamePlayers) {

	m.Name, _ = GetTextBetween(str, "username-tooltips\">", "</span>")

	m.Teammate, _ = GetTextBetween(str, "class=\"formation", "icon-titletip")

	m.Nameid, _ = GetTextBetween(str, "https://arena.5eplay.com/data/player/", "\"")
	//fmt.Println(m.Name, m.Teammate, m.Nameid)
}
