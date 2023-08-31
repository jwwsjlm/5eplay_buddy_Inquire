package utils

import (
	"fmt"
	"strconv"
)

func hexXor(str1, str2 string) string {
	result := ""
	for i := 0; i < len(str1) && i < len(str2); i += 2 {
		num1, _ := strconv.ParseInt(str1[i:i+2], 16, 64)
		num2, _ := strconv.ParseInt(str2[i:i+2], 16, 64)
		xorResult := num1 ^ num2
		result += fmt.Sprintf("%02x", xorResult)
	}
	return result
}

func unsbox(str string) string {
	mapping := []int{0xf, 0x23, 0x1d, 0x18, 0x21, 0x10, 0x1, 0x26, 0xa, 0x9, 0x13, 0x1f, 0x28, 0x1b, 0x16, 0x17, 0x19, 0xd, 0x6, 0xb, 0x27, 0x12, 0x14, 0x8, 0xe, 0x15, 0x20, 0x1a, 0x2, 0x1e, 0x7, 0x4, 0x11, 0x5, 0x3, 0x1c, 0x22, 0x25, 0xc, 0x24}
	result := ""
	for i := 0; i < len(str); i++ {
		index := mapping[i] - 1
		result += string(str[index])
	}
	return result
}

func Getacw(x string) string {
	a := unsbox(x)
	return hexXor(a, "3000176000856006061501533003690027800375")
}
