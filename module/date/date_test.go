// Package date
// @description
// @author      张盛钢
// @datetime    2023/3/20 11:07
package date

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Log(GetDiffDaysLunar(1, 1))
	t.Log(GetDiffDaysSolar("10-01"))
	t.Log(GetDiffDaysSolar("1-01"))

	t.Log(ImportDateFormatMsg())
}
