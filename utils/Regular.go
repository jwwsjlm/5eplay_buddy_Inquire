package utils

import (
	"errors"
	"regexp"
	"strings"
)

func MatchArg1String(arg1 string) []string {
	// 定义正则表达式
	re := regexp.MustCompile(`arg1='(.*?)';`)

	// 使用正则表达式匹配字符串
	result := re.FindStringSubmatch(arg1)

	return result
}
func MatchGame(data string) []string {
	// 定义正则表达式
	re := regexp.MustCompile(`name tleft ban-bg(.*?)prize-con`)

	// 使用正则表达式匹配字符串
	result := re.FindStringSubmatch(data)

	return result
}
func ExtractTextBetween(str, start, end string) []string {
	var result []string

	for {
		startIndex := strings.Index(str, start)
		if startIndex == -1 {
			break
		}

		endIndex := strings.Index(str[startIndex+len(start):], end)
		if endIndex == -1 {
			break
		}

		result = append(result, str[startIndex+len(start):startIndex+len(start)+endIndex])
		str = str[startIndex+len(start)+endIndex+len(end):]
	}

	return result
}
func GetTextBetween(str, start, end string) (string, error) {
	startIndex := strings.Index(str, start)

	if startIndex == -1 {
		return "", errors.New("未找到文本")
	}

	endIndex := strings.Index(str[startIndex+len(start):], end)
	if endIndex == -1 {
		return "", errors.New("未找到文本")
	}

	return str[startIndex+len(start) : startIndex+len(start)+endIndex], nil
}
