// Package ticker
// @description 定时任务组
// @author		Edwards.Z
// @datetime    2023/3/15 17:11
package ticker

func Ticker() {
	go MasterTicker()
}
