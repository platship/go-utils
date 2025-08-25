package datex

import (
	"strconv"
	"time"
)

const (
	FmtDate              = "2006-01-02"
	FmtTime              = "15:04:05"
	FmtDateTime          = "2006-01-02 15:04:05"
	FmtDateTimeNoSeconds = "2006-01-02 15:04"
	FmtDateTimeNoMinutes = "2006-01-02 15"
)

/**
 * @desc: 获取当前时间戳（秒级）
 * @return {*}
 */
func NowUnix() int64 {
	return time.Now().Unix()
}

/**
 * @desc: 秒时间戳转时间
 * @param undefined
 * @return {*}
 */
func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}

/**
 * @desc: 毫秒时间戳转时间
 * @param undefined
 * @return {*}
 */
func UnixMilliToTime(unix int64) time.Time {
	return time.UnixMilli(unix)
}

/**
 * @desc: 当前毫秒时间戳
 * @return {*}
 */
func NowUnixMilli() int64 {
	return TimeToUnixMilli(time.Now())
}

/**
 * @desc: 时间转毫秒时间戳
 * @param undefined
 * @return {*}
 */
func TimeToUnixMilli(t time.Time) int64 {
	return t.UnixMilli()
}

/**
 * @desc: 格式化毫秒时间戳
 * @param undefined
 * @return {*}
 */
func FormatUnixMilli(timestamp int64, format string) string {
	if format == "" {
		format = FmtDateTime
	}
	return TimeToString(time.Unix(0, timestamp*int64(time.Millisecond)), format)
}

/**
 * @desc: 时间格式化
 * @return {*}
 */
func TimeToString(time time.Time, layout string) string {
	if time.IsZero() {
		return ""
	}
	return time.Format(layout)
}

/**
 * @desc: 字符串时间转时间类型
 * @param undefined
 * @param undefined
 * @return {*}
 */
func StringToTime(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

/**
 * @desc: 获取今天
 * @return yyyy-MM-dd
 */
func GetToday() string {
	return time.Now().Format(FmtDate)
}

/**
 * @desc: 返回指定时间当天的开始时间
 * @param undefined
 * @return {*}
 */
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

/**
 * 将时间格式换成 xx秒前，xx分钟前...
 * 规则：
 * 59秒--->刚刚
 * 1-59分钟--->x分钟前（23分钟前）
 * 1-24小时--->x小时前（5小时前）
 * 昨天--->昨天 hh:mm（昨天 16:15）
 * 前天--->前天 hh:mm（前天 16:15）
 * 前天以后--->mm-dd（2月18日）
 */
func PrettyTime(milliseconds int64) string {
	t := UnixMilliToTime(milliseconds)
	duration := (NowUnixMilli() - milliseconds) / 1000
	if duration < 60 {
		return "刚刚"
	} else if duration < 3600 {
		return strconv.FormatInt(duration/60, 10) + "分钟前"
	} else if duration < 86400 {
		return strconv.FormatInt(duration/3600, 10) + "小时前"
	} else if TimeToUnixMilli(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24))) <= milliseconds {
		return "昨天 " + TimeToString(t, FmtTime)
	} else if TimeToUnixMilli(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24*2))) <= milliseconds {
		return "前天 " + TimeToString(t, FmtTime)
	} else {
		return TimeToString(t, FmtDate)
	}
}

/**
 * @desc: 获取本周周一的日期
 * @return {*}
 */
func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format(FmtDate)
	return
}

/**
 * @desc: 获取上周的周一日期
 * @return {*}
 */
func GetLastWeekFirstDate() (weekMonday string) {
	thisWeekMonday := GetFirstDateOfWeek()
	TimeMonday, _ := time.Parse(FmtDate, thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format(FmtDate)
	return
}

/**
 * @desc: 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
 * @return {*}
 */
func GetFirstDateOfMonth() (monthMonday string) {
	now := time.Now()
	offset := -30
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	monthMonday = weekStartDate.Format(FmtDate)
	return
}

/**
 * @desc: 获取两个时间的天数差
 * @return {*}
 */
func GetDaysByToday(start, end time.Time) int {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	addTime.Unix()
	if addTime.Unix() <= start.Unix() {
		return 0
	} else if (addTime.Unix() - 86400) <= start.Unix() {
		return 1
	}
	//计算相差天数
	return int(end.Sub(start).Hours() / 24)
}

// GetBetweenTimes 返回从 startTime 到 endTime（含）的所有时间戳字符串，步长单位可以取 "day"、"hour"、"minute" 或 "month"，
// stepVal 表示步长数值，如果未提供则默认步长为 1 天
func GetBetweenTimes(startTime, endTime string, unit string, stepVal ...int) []string {
	d := []string{}
	timeFormatTpl := FmtDateTime
	// 如果传入的 startTime 格式与默认格式长度不一致，则截取默认格式对应长度
	if len(timeFormatTpl) != len(startTime) {
		timeFormatTpl = timeFormatTpl[0:len(startTime)]
	}

	startT, err := time.Parse(timeFormatTpl, startTime)
	if err != nil {
		return d
	}
	endT, err := time.Parse(timeFormatTpl, endTime)
	if err != nil {
		return d
	}
	if endT.Before(startT) {
		return d
	}
	// 默认步长为 1
	step := 1
	if len(stepVal) > 0 {
		step = stepVal[0]
	}

	// 添加初始时间
	curr := startT
	d = append(d, curr.Format(timeFormatTpl))

	for i := 0; ; i++ {
		var next time.Time
		switch unit {
		case "day":
			next = curr.AddDate(0, 0, step)
		case "hour":
			next = curr.Add(time.Duration(step) * time.Hour)
		case "minute":
			next = curr.Add(time.Duration(step) * time.Minute)
		case "month":
			next = curr.AddDate(0, step, 0)
		default:
			// 单位不合法时直接返回当前结果
			return d
		}
		// 如果下一个时间不早于结束时间，则退出循环，不将 endT 加入结果
		if !next.Before(endT) {
			break
		}
		d = append(d, next.Format(timeFormatTpl))
		curr = next

		// 防止无限循环，外层写个上限 (可根据实际需求调整)
		if i >= 99 {
			break
		}
	}
	return d
}

// 获取当天0点时间
func GetTodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
