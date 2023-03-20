package global

import (
	"encoding/json"
	"testing"
)

func TestConfig(t *testing.T) {

	k := &Keys{
		ChristmasHatURL:    "",
		BotName:            "",
		WeatherKey:         "",
		TianapiKey:         "",
		TianapiKey1:        "",
		HoneyLove:          "",
		LoverChName:        "",
		MasterAccount:      "",
		HouchangcunFans:    "",
		BanzhuanGroup:      "",
		BubeiGroup:         "",
		QweatherKey:        "83c3d4b6848f4abfa8981330bec334e5",
		BubeiStartDate:     "",
		WuZhuangShiMembers: "",
		RemindMsg:          "",
	}
	s := SysConfig{
		Keys: k,
	}
	ss, _ := json.Marshal(s)
	t.Log(string(ss))
}
