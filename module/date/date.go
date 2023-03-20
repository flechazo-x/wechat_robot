package date

import (
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/flechazo-x/go_common/util/timex"
	"github.com/golang-module/carbon"
	"wechat_robot/global"
)

var (
	yuanDan      = "01-01" // 元旦
	valentineDay = "02-14" // 情人节
	anniversary  = "09-17" // 纪念日
	day520       = "05-20"
)

// ImportDateFormatMsg 导入日期格式消息
func ImportDateFormatMsg() string {
	return fmt.Sprintf("%s.\n距离元旦[01-01]还有 %d 天\n距离春节还有 %d 天\n距离[02-14]情人节还有 %d 天\n距离[520]还有 %d 天\n距离七夕,七月初七还有 %d 天\n"+
		"距离纪念日还有 %d 天\n距离%s的生日还有 %d 天",
		global.GetCfg().Keys.RemindMsg,
		GetDiffDaysSolar(yuanDan),
		GetDiffDaysLunar(1, 1),
		GetDiffDaysSolar(valentineDay),
		GetDiffDaysSolar(day520),
		GetDiffDaysLunar(7, 7),
		GetDiffDaysSolar(anniversary), global.GetCfg().Keys.LoverChName,
		GetDiffDaysSolar(global.GetCfg().Keys.LoverBirthday))
}

func getYearDay(year int, date string) string {
	return fmt.Sprintf(`%d-%s`, year, date)
}

// GetDiffDaysSolar 获取两个时间相差的天数阳历
func GetDiffDaysSolar(futureDate string) int {
	furD := carbon.Parse(getYearDay(carbon.Parse(timex.LocalNow().String()).Year(), futureDate)) //未来农历日期
	curD := carbon.Parse(timex.Today(timex.DateTPL))                                             //当前年月日
	day := int(furD.ToStdTime().Sub(curD.ToStdTime()).Hours() / 24)
	if day > 0 {
		return day
	}
	furDay := furD.ToStdTime().AddDate(1, 0, 0)
	return int(furDay.Sub(curD.ToStdTime()).Hours() / 24)
}

// GetDiffDaysLunar 获取两个时间相差的天数农历
func GetDiffDaysLunar(solarMonth, solarDay int64) int {
	year := int64(timex.LocalNow().Year())
	hour := int64(timex.LocalNow().Hour())
	minute := int64(timex.LocalNow().Minute())
	second := int64(timex.LocalNow().Second())
	furD := carbon.Parse(getLunar2SolarDate(year, solarMonth, solarDay, hour, minute, second)) //未来日期
	curD := carbon.Parse(timex.Today(timex.DateTPL))                                           //当前年月日
	day := int(furD.ToStdTime().Sub(curD.ToStdTime()).Hours() / 24)
	if day > 0 {
		return day
	}
	furDay := carbon.Parse(getLunar2SolarDate(year+1, solarMonth, solarDay, hour, minute, second)) //未来日期
	return int(furDay.ToStdTime().Sub(curD.ToStdTime()).Hours() / 24)
}

// 农历转阳历
func getLunar2SolarDate(year, month, day, hour, minute, second int64) string {
	c := calendar.ByLunar(year, month, day, hour, minute, second, timex.IsLeapYear(timex.CST8Now().Year())).Solar
	return fmt.Sprintf("%d-%d-%d", c.GetYear(), c.GetMonth(), c.GetDay())
}
