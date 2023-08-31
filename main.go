package main

import (
	"5e/config"
	"5e/utils"
	"fmt"
	"strconv"
)

func main() {
	i := 0
	url := "https://arena.5eplay.com/api/data/match_list/13800525utgrhp?page=1&yyyy=2023&match_type="
	array := []config.AData{}
	ii := "技巧少女不会受孕"
	acco := config.Acco{}
	for {
		i++
		url = "https://arena.5eplay.com/api/data/match_list/13800525utgrhp?page=" + strconv.Itoa(i) + "&yyyy=2023&match_type="
		request, err := utils.SendGetRequest(url)
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
	for i, v := range array {
		fmt.Println(i, v.MatchCode)
		utils.Setacw("")
		//result := utils.Getacw("13D8E4E464E15558AFDE808F185EE64F7E4B5D93")
		//fmt.Println(result)
		url = "https://arena.5eplay.com/data/match/" + v.MatchCode
		a, _ := utils.SendGetRequest(url)
		b := utils.MatchArg1String(string(a))
		//fmt.Println(string(a))
		acw := utils.Getacw(b[1])
		utils.Setacw(acw)

		a, _ = utils.SendGetRequest(url)

		var aa []config.Match
		cc := utils.ExtractTextBetween(string(a), "name tleft ban-bg", "icon_wrap")
		for i, v := range cc {
			aa = append(aa, config.Match{Data: v})
			aa[i].Match()
			acco.Match = append(acco.Match, config.Match{Data: v})
			if i < 5 {
				aa[i].Up = 0
			} else {
				aa[i].Up = 1
			}
			if aa[i].Name == ii {
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

					if acco.Teammate == v.Teammate && v.Name != ii && acco.Teammate != "" {

						acco.Teammate_ids = append(acco.Teammate_ids, v.Name)
						//fmt.Println(acco.Name, v.Name)
					}

				}
			}
		} else {
			for i, v := range aa {
				if i > 5 {

					if acco.Teammate == v.Teammate && v.Name != ii && acco.Teammate != "" {
						acco.Teammate_ids = append(acco.Teammate_ids, v.Name)
						//fmt.Println(acco.Name, acco.Teammate, "----", v.Name, v.Teammate)
						//fmt.Println(acco.Name, v.Name)
					}

				}
			}
		}
	}

	fmt.Println(acco.Location, acco.Nameid, acco.Name, acco.Teammate_ids, ii)

}
