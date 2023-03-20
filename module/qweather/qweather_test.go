package qweather

import (
	"testing"
	"wechat_robot/global"
)

var cfg = &global.SysConfig{}

func TestGetLocationID(t *testing.T) {
	id, err := GetLocationID("北京")
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestGetQWeatherDetail(t *testing.T) {

	//detail, err := GetQWeatherDetail("101010100", "北京")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(detail)
}
