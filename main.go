package main

import (
	"5e/config"
	"5e/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	i := 0

	url := "https://arena.5eplay.com/data/player/10304198uihisj"
	request, err := utils.SendGetRequest(url)
	n := utils.ExtractTextBetween(string(request), "\"username-tooltips\">", "</span>")
	ni := utils.ExtractTextBetween(string(request), "_g_player_domain ='", "'")
	if len(n) == 0 || len(ni) == 0 {
		return
	}
	fmt.Println(n, ni)

	url = "https://arena.5eplay.com/api/data/match_list/" + ni[0] + "?page=1&yyyy=2023&match_type="
	array := []config.AData{}

	acco := config.Acco{}
	for {
		i++
		url = "https://arena.5eplay.com/api/data/match_list/" + ni[0] + "?page=" + strconv.Itoa(i) + "&yyyy=2023&match_type="
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
	//utils.Setacw("")
	for i, v := range array {

		utils.Setacw("")
		//result := utils.Getacw("13D8E4E464E15558AFDE808F185EE64F7E4B5D93")
		//fmt.Println(result)
		url = "https://arena.5eplay.com/data/match/" + v.MatchCode

		//a, _ := utils.SendGetRequest(url)
		//b := utils.MatchArg1String(string(a))
		////fmt.Println(string(a))
		//acw := utils.Getacw(b[1])
		//utils.Setacw(acw)

		a, _ := utils.SendGetRequest(url)

		var aa []config.Match
		cc := utils.ExtractTextBetween(string(a), "name tleft ban-bg", "icon_wrap")
		fmt.Println(i, url, len(cc))
		for i, v := range cc {
			aa = append(aa, config.Match{Data: v})
			aa[i].Match()
			acco.Match = append(acco.Match, config.Match{Data: v})
			if i < 5 {
				aa[i].Up = 0
			} else {
				aa[i].Up = 1
			}
			if aa[i].Name == n[0] {
				acco.Location = i
				acco.Name = aa[i].Name
				acco.Nameid = aa[i].Nameid
				acco.Teammate = aa[i].Teammate
			}
			fmt.Println(aa[i].Name, aa[i].Avatar, aa[i].Teammate, aa[i].Nameid, aa[i].Up)
		}
		//fmt.Println(acco.Name, acco.Teammate, "本人id")
		if acco.Location < 5 {
			for i, v := range aa {
				if i < 5 {
					//fmt.Println(v.Name, "----", acco.Teammate, v.Teammate, "----", v.Name == ii, acco.Teammate == v.Teammate, acco.Teammate == v.Teammate && v.Name != ii)

					if acco.Teammate == v.Teammate && v.Name != n[0] && acco.Teammate != "" {
						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						//acco.TeammateIds = append(acco.TeammateIds, config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						//fmt.Println(acco.Name, v.Name)
					}

				}
			}
		} else {
			for i, v := range aa {
				if i > 5 {

					if acco.Teammate == v.Teammate && v.Name != n[0] && acco.Teammate != "" {
						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						//acco.TeammateIds = append(acco.TeammateIds, config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						//fmt.Println(acco.Name, acco.Teammate, "----", v.Name, v.Teammate)
						//fmt.Println(acco.Name, v.Name)
					}

				}
			}
		}
		time.Sleep(time.Second * 2)
		fmt.Println(acco.Location, acco.Nameid, acco.Name, acco.TeammateIds, n)
	}

	//fmt.Println(acco.Location, acco.Nameid, acco.Name, acco.TeammateIds, ii)

}
