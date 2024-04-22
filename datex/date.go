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

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(startDate, endDate string) []string {
	d := []string{}
	timeFormatTpl := FmtDateTime
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = FmtDate
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func TodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
