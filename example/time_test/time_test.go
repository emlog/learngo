package time_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 获取时间戳 int64
	timestamp := time.Now().Unix()
	t.Logf("unix timestamp: %d", timestamp)

	time.Now().Unix()

	// 格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	t.Log(tm.Format("2006-01-02 03:04:05 PM"))       // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	t.Log(time.Now().Format("2006/01/02/ 15:04:05")) // 15：04 是24小时制

	// 从字符串转为时间戳，layout:格式，value:要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	t.Log(tm2.Unix())

	// go语言并没有全局设置时区这么一个东西，每次输出时间都需要调用一个In()函数改变时区
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	t.Log("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
}
