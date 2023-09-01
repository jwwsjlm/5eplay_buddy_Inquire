package utils

import (
	"fmt"
	"io"
	"net/http"
)

var acw_sc__v2 http.Cookie

func SendGetRequest(url string) ([]byte, error) {
	//url := "https://arena.5eplay.com/data/match/g151-n-20230830180505294466823"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&acw_sc__v2)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	b, err := io.ReadAll(res.Body)

	arg1 := MatchArg1String(string(b))
	//fmt.Println(string(b), arg1)
	if len(arg1) > 1 {
		Setacw(Getacw(arg1[1]))
		client := http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.AddCookie(&acw_sc__v2)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		b, err := io.ReadAll(res.Body)
		return b, nil
	}

	return b, nil

}
func Setacw(c string) {

	acw_sc__v2 = http.Cookie{
		Name:  "acw_sc__v2",
		Value: c,
	}
}
func ResetCookies() {
	acw_sc__v2 = http.Cookie{}
}
