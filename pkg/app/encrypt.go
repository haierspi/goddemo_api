package app

import (
	"starfission_go_api/global"
	"starfission_go_api/pkg/util"
)

func Convert(str string) string {
	key := global.SecuritySetting.HtmlEncryptKey

	tokenByte := []rune(str)
	keyByte := []rune(key)
	tmpCode := util.XorEncodeStrRune(tokenByte, keyByte)

	return string(tmpCode)
}
