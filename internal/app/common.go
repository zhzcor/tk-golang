package app

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"strings"
	"time"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	app "tkserver/pkg/requestparam"
)

type Common struct {
}

type Calendar struct {
	date string
}

func (c Common) DoClientAgent(s string) string {
	userAgent := s
	mobileRe, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")

	data := mobileRe.FindString(userAgent)

	return data
}

// 获取指定时间所在月的开始 结束时间
func (c Common) GetMonthStartEnd(t time.Time) (time.Time, time.Time) {
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

func (c Common) GetCalendar(t time.Time) []*Calendar {
	year := t.Year()
	month := int(t.Month())
	days := c.MonthCount(year, month)

	var res []*Calendar
	for i := 0; i < days; i++ {
		va := &Calendar{}
		va.date = strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(i+1)
		res = append(res, va)
	}

	return res
}

//获取月的天数
func (c Common) MonthCount(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

func (c Common) GetDifficultyMap(id int) string {
	// 难易度，1：易，2：较易，3：较难，4：难，5：一般
	switch id {
	case 1:
		return "易"
	case 2:
		return "较易"
	case 3:
		return "较难"
	case 4:
		return "难"
	case 5:
		return "一般"
	default:
		return ""
	}
}

//获取指定开始,结束日期
func (c Common) GetDateStartEnd(t time.Time) (time.Time, time.Time) {
	timeStr := t.Format("2006-01-02")
	start, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	endNum := start.Unix() + 86400
	end := time.Unix(endNum, 0)
	return start, end
}

//数组去重复
func (c Common) RmDuplicate(list []int) []int {
	var x []int = []int{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

//添加操作日志
func (c Common) SetOperationLog(ctx *gin.Context, adminId int, record string) error {
	creation := store.WithContext(ctx).AdminOperationLog.Create().
		SetRecord(record).
		SetIP(app.RemoteIp(ctx.Request))
	if adminId > 0 {
		creation = creation.SetAdminID(adminId)
	}
	_, err := creation.Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//获取时间的周开始结束
func (c Common) FindWeekTimeByTime(date time.Time) (startDay, endDay time.Time) {
	weekDay := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}
	end := (time.Saturday - date.Weekday()).String()
	plus := int64(86400 * weekDay[end])
	owe := int64(86400 * (7 - 1 - weekDay[end]))
	endDay = time.Unix(date.Unix()+plus+86399, 0)
	startDay = time.Unix(date.Unix()-owe, 0)
	return
}

func (c Common) WeekDayDiff(t time.Time) (time.Time, time.Time) {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	offset2 := int(time.Sunday - t.Weekday())

	weekEnd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset2+7)

	return weekStart, weekEnd
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func (c Common) GetBetweenDates(date, date2 time.Time) []string {
	d := []string{}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl := "2006-01-02"
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

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func (c Common) GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return c.GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func (c Common) GetLastDateOfMonth(d time.Time) time.Time {
	return c.GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func (c Common) GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 计算日期相差多少天
// 返回值day>0, t1晚于t2; day<0, t1早于t2
func (c Common) SubDays(t1, t2 time.Time) (day int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t_ := t1
		t1 = t2
		t2 = t_
		swap = true
	}

	day = int(t1.Sub(t2).Hours() / 24)

	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		day += 1
	}

	if swap {
		day = -day
	}

	return
}

// 计算日期相差多少月
func (c Common) GetDiffDescDate(t1, t2 time.Time) (data string) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()
	h1 := t1.Hour()
	h2 := t2.Hour()
	Minute1 := t1.Minute()
	Minute2 := t2.Minute()

	yearInterval := y1 - y2
	monthInterval := m1 - m2
	dayInterval := d1 - d2
	hourInterval := h1 - h2
	minInterval := Minute1 - Minute2

	if yearInterval > 0 {
		return strconv.Itoa(yearInterval) + "年前"
	}

	if monthInterval > 0 {
		return strconv.Itoa(monthInterval) + "月前"
	}
	if dayInterval > 0 {
		return strconv.Itoa(dayInterval) + "天前"
	}
	if hourInterval > 0 {
		return strconv.Itoa(hourInterval) + "小时前"
	}
	if minInterval > 0 {
		return strconv.Itoa(minInterval) + "分钟前"
	}
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	/*if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval*/
	return
}

//判断是否登录
func (c Common) IsUserLogin(context *gin.Context)(uid int,err error){
	auth := context.GetHeader("Authorization")
	if auth != "" {
		authInfo := strings.Split(auth, " ")
		if len(authInfo) > 1 {
			authType := authInfo[0]
			tk := authInfo[1]
			switch {
			case authType == "Bearer":
				userService := UserCenter{}
				uid, err := userService.CheckToken(tk)
				if err != nil {
					return 0,err
				}

				return uid,nil
				// do something
			}
		}
	}

	return 0,nil
}
