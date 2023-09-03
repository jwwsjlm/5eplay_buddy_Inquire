package main

import (
	"5e/config"
	"5e/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type InputData struct {
	InputText string `json:"inputText"`
}

type OutputData struct {
	Result string `json:"result"`
}

var ti int

func FriendQuery(url string) string {

	var info = config.PlayerInfo{}
	i := 0

	//url := "https://arena.5eplay.com/data/player/9548486whkirl"
	request, err := utils.SendGetRequest(url)

	err = utils.ParseToPlayerList(request, &info)
	if err != nil {
		fmt.Println(err)
		return ""
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

			return ""
		}
		d, err := config.Deserialize(request)
		if err != nil {
			fmt.Println("json解析失败", err)
			return ""
		}
		if d.Success == false {
			break
		} else {
			array = append(array, d.Data...)
		}
	}
	acco := config.CurrentGamePlayersArray{}
	fmt.Printf("当前玩家总局数为:%d,玩家名字为:%s,玩家个人地址为:%s\n", len(array), info.Name, info.ProfileURL)
	for _, v := range array {

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

						//acco.AddMatchRecord(config.FriendInformation{RecordURL: url})
						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						acco.AddMatchRecords(config.FriendInformation{Name: v.Name, RecordURL: []string{url}})
					}

				}
			}
		} else {
			for i, v := range aa {
				if i > 5 {
					if acco.Teammate == v.Teammate && v.Name != info.Name && acco.Teammate != "" {
						//acco.TeammateIds = append(acco.TeammateIds, config.FriendInformation{Name: v.Name, RecordURL: url})
						//acco.AddMatchRecord(config.FriendInformation{RecordURL: url})

						acco.Append(config.FriendInformation{Name: v.Name, Match: "https://arena.5eplay.com/data/player/" + v.Nameid})
						acco.AddMatchRecords(config.FriendInformation{Name: v.Name, RecordURL: []string{url}})
					}
				}
			}
		}
		time.Sleep(time.Second * time.Duration(ti))

		//fmt.Printf("玩家姓名:%s,当前玩家的开黑好友为:%v,本次查询局数为%d,一共有%d\n", acco.Name, acco.TeammateIds, i, len(array), string(mjson))
	}
	mjson, _ := json.Marshal(acco)
	fmt.Println(i, string(mjson))
	return string(mjson)
}
func GinFriendQuery(c *gin.Context) {
	//body, _ := io.ReadAll(c.Request.Body)
	//fmt.Println(string(body))
	var req config.GinRequestParams
	err := c.Bind(&req)
	fmt.Println("data为", req.Data)

	if err != nil {
		c.Error(errors.New("传递的数据错误2"))

		return
	}

	data := FriendQuery(req.Data)

	c.String(200, data)
}
func main() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "http://127.0.0.1:8199/")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", "http://127.0.0.1:8199/")
	} else {
		cmd = exec.Command("xdg-open", "http://127.0.0.1:8199/")
	}
	//var url string
	//
	//url = "https://arena.5eplay.com/data/player/14101939b0wwhm"
	//flag.StringVar(&url, "url", url, "个人简介地址")
	//flag.IntVar(&ti, "time", 1, "每次获取延时单位为秒,默认一秒")
	//flag.Parse()
	//fmt.Println(url, ti)
	//
	//FriendQuery(url)
	//return
	router := gin.Default()

	// 静态文件路由，用于加载 HTML 文件
	router.StaticFile("/", "./html/index.html")

	// 处理函数，接收 POST 请求并处理数据
	router.POST("/process", GinFriendQuery)
	err := cmd.Run()
	if err != nil {
		fmt.Println("打开网页时发生错误:", err)
	}
	router.Run(":8199")

}
