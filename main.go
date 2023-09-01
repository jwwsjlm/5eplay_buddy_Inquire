package main

import (
	"5e/config"
	"5e/utils"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var url string
	var ti int
	flag.StringVar(&url, "url", "https://arena.5eplay.com/data/player/9548486whkirl", "个人简介地址")
	flag.IntVar(&ti, "time", 1, "每次获取延时单位为秒,默认一秒")
	flag.Parse()
	fmt.Println(url, ti)

	var info = config.PlayerInfo{}
	i := 0

	//url := "https://arena.5eplay.com/data/player/9548486whkirl"
	request, err := utils.SendGetRequest(url)

	err = utils.ParseToPlayerList(request, &info)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.Name, info.PlayerID, info.AvatarURL)

	url = "https://arena.5eplay.com/api/data/match_list/" + info.PlayerID + "?page=1&yyyy=2023&match_type="
	var array []config.AData
	//var acco []config.CurrentGamePlayersArray

	for {
		i++
		url = "https://arena.5eplay.com/api/data/match_list/" + info.PlayerID + "?page=" + strconv.Itoa(i) + "&yyyy=2023&match_type="
		request, err = utils.SendGetRequest(url)
		if err != nil {

			return
		}
		d, err := config.Deserialize(request)
		if err != nil {
			fmt.Println("json解析失败", err)
			return
		}
		if d.Success == false {
			break
		} else {
			array = append(array, d.Data...)
		}
	}
	acco := config.CurrentGamePlayersArray{}
	fmt.Printf("当前玩家总局数为:%d,玩家名字为:%s,玩家个人地址为:%s\n", len(array), info.Name, info.ProfileURL)
	for i, v := range array {

		url = "https://arena.5eplay.com/data/match/" + v.MatchCode
		a, _ := utils.SendGetRequest(url)
		var aa []config.CurrentGamePlayers
		cc := utils.ExtractTextBetween(string(a), "name tleft ban-bg", "icon_wrap")
		for i, vv := range cc {
			cgp := config.CurrentGamePlayers{Data: vv}
			utils.OrganizePlayerInfo(vv, &cgp)

			//aa = append(aa, config.CurrentGamePlayers{Data: vv})
			aa = append(aa, cgp)
			//aa[i].Match()

			//cgp.TeammateIds = append(cgp.TeammateIds)
			acco.Match = append(acco.Match, config.CurrentGamePlayers{Data: vv})
			if i < 5 {
				cgp.Up = 0
			} else {
				cgp.Up = 1
			}

			if cgp.Name == info.Name {
				acco.Location = i
				acco.Name = cgp.Name
				acco.Nameid = cgp.Nameid
				acco.Teammate = cgp.Teammate
				acco.RecordURL = url
			}
		}

		if acco.Location < 5 {
			for i, v := range aa {
				if i < 5 {
					if acco.Teammate == v.Teammate && v.Name != info.Name && acco.Teammate != "" {
						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
					}

				}
			}
		} else {
			for i, v := range aa {
				if i > 5 {
					if acco.Teammate == v.Teammate && v.Name != info.Name && acco.Teammate != "" {
						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
					}
				}
			}
		}
		time.Sleep(time.Second * time.Duration(ti))
		fmt.Printf("玩家姓名:%s,当前玩家的开黑好友为:%v,本次查询局数为%d,一共有%d\n", acco.Name, acco.TeammateIds, i, len(array))
	}
}
