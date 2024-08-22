package sherrytime
/**
 *  Note: 1.西曆4年羅馬奧古斯都帝停閏,故是年2月仍只有28天。
 *        2.教皇格勒哥里第13始改曆,自西曆1582年10月4日逕跳至15日或自1752年9月2日跳至14日
 *        3.西元前 n 年之值為 -n+1
   private $numText = array('零', '一', '二', '三', '四', '五', '六', '七', '八', '九', '十', '十一', '十二');
   private $_week = array('日','一','二','三','四','五','六');
   private $_weekx = array('一','二','三','四','五','六', '日');
*/

import (
   "fmt"
   "time"
   "sync"
   "strings"
   "strconv"
)

type Date struct {
   Year  int        // Year (e.g., 2014).
   Month time.Month // Month of the year (January = 1, ...).
   Day   int        // Day of the month, starting at 1.
}

type SherryTime struct {
   current time.Time
   numText [13]string
   location string
   Delimiter string   // 分隔符號
   dayOfMonths [12]int   // 每月天數
}

var lock = sync.RWMutex{}

// DateOf returns the Date in which a time occurs in that time's location.
func(st *SherryTime) DateOf(t time.Time)(Date) {
   var d Date
   d.Year, d.Month, d.Day = t.Date()
   return d
}

// SetTime(time time) 將UTC轉為Asia/Taipei UTC+8
func(st *SherryTime) SetCurrentTime(t time.Time) {
   location := time.FixedZone("UTCr+8", +8*60*60)
   st.current = t.In(location)
}

// toNumText(i int)  阿拉伯數字轉國字
func (st *SherryTime) toNumText(i int) (string) {  
   return st.numText[i] 
}

func (st *SherryTime) getTimeSince1582() uint64 {
   defer lock.Unlock()

   lock.Lock()
   t := time.Now()
   nanoSecond := t.UnixNano()

   nano := uint64(nanoSecond / 100)
   return nano + 122192928000000000
}

// <note> 西元1752年九月必須特別處理!
// <return> 0) success; !=0) fail
func (st *SherryTime) chkDateW(year, month, day int) (err int) {
   err = 0
   if year < 0 || month < 1 || month > 12 || day < 1 || day > st.lastMonthDay(year, month) {
       err--
   } else if year == 1582 && month == 10 && day >= 5 && day <= 14 {
       err--
   }
   return  err
}

// 判斷是否閏年，傳入參數：西元年
func (st *SherryTime) leapYear(year int) (bool) {
   if(year == 4)  { // 西曆4年羅馬奧古斯都帝停閏,故是年2月仍只有28天
      return  false 
   }  
   if(year < 0)  { // e.g.西元前1年(-1年) == 第-(-1+1)年
      year = -(year+1)
   }
   // 教皇格勒哥里第13始改曆,以西曆1582年10月5日為15日,中間銷去10日
   return year % 4 == 0 && (year <= 1582 || year % 100 != 0) ||  year % 400 == 0
}

// 取得某月(mm)最後一天的日期
func (st *SherryTime) lastMonthDay(yy, mm int) (int) {
   if mm != 0 {
      if mm == 2 && yy != 0 {
         if st.leapYear(yy) {
            return 29
         } else {
            return 28
         }
      } else {
         return st.dayOfMonths[mm-1] 
      }
   } else {
      return 31
   }
   // return (mm != 0 ? (mm == 2 && yy != 0 ? (st.leapYear(yy) ? 29 : 28) : st.dayOfMonths[mm-1]) : 31)
}

// 回傳每月的最後的天數
func(app *SherryTime) LastMonthDay(yy, mm int)(int) {
   return app.lastMonthDay(yy, mm)
}

// <func> Get day ordinal for western calendar.
// The ordinal of 1/1/1 is 1.
// <return> >0) success; -1) fail.
func(st *SherryTime) toDayOrdW(yy, mm, dd int) (int) {
   var n int = 0
   if st.chkDateW(yy, mm, dd) == -1 {   
      n = -10 
   } else if yy > 0 {
      n = dd
      if mm > 2 && st.leapYear(yy) {
         n++
      }
      // 教皇格勒哥里第13改曆
      if yy > 1582 || yy == 1582 && (mm > 10 || mm == 10 && dd >= 15) {
         n -= 10
      }
      mm--;
      for i := 0; i < mm; i++  {
         n += st.dayOfMonths[i]
      }
      yy--
      n += yy * 365 + (yy/4);
      if yy >= 4  {  // 西曆4年羅馬奧古斯都帝停閏
         n--
      }
      if yy >= 1600  {   // 教皇格勒哥里第13改曆
         i := yy - 1600;
         n -= (i/100) - (i/400);
      }
   } else {
      n = -1
   }
   return n
}

// 開放給外面介面使用(for 相容性)
func(st *SherryTime) ToDayOrdWs(yymmdd string)(int) {
   return st.toDayOrdWs(yymmdd)
}

// 轉換成日序
func(st *SherryTime) toDayOrdWs(yymmdd string)(int)  {
   d := strings.Split(yymmdd, st.Delimiter)
   yy, _ := strconv.Atoi(d[0])
   mm, _ := strconv.Atoi(d[1])
   dd, _ := strconv.Atoi(d[2])
   return st.toDayOrdW(yy, mm, dd)
}

// 日序轉西元
//<func> From darily ordinal 'dOrd' to western date 'yy', 'mm' and 'dd'
//<return> >=0) week number(0 for Sunday); -1) fail
func (st *SherryTime) toDateW(dOrd int)(yy, mm, dd, week int) {
   var baseWeek int = 6

   yy = (int((float64(dOrd)/float64(365.25))) + 1)
   ord := st.toDayOrdW(yy, 1, 1)
   if ord < 0 {
      return -1, 0, 0, 0
   }
   dd = dOrd - ord + 1
   mm = 1
   n := st.lastMonthDay(yy, mm)
   yes := false

   for dd > n  {
      dd = dd - n
      mm = mm + 1
      if mm > 12  {
         yy++
         mm = 1
      }
      yes = (yy == 1582 && mm == 10)
      if yes {
         n = 21
      } else {
         n = st.lastMonthDay(yy, mm)
      }
   }
   if yes && dd > 4 {
      dd += 10
   }
   w := (dOrd + baseWeek) % 7
   return yy, mm, dd, w
}

// 日序轉西元(for 相容性）
func(st *SherryTime) ToDateWs(dayOrder int)(string) {
   return st.toDateWs(dayOrder)
}

// 日序轉西元
func (st *SherryTime) toDateWs(dOrd int) (string)  {
      var y, m, d int
      y = 0
      m = 0
      d = 0
      y, m , d, _ = st.toDateW(dOrd)
      return fmt.Sprintf("%d-%d-%d", y, m, d)
}

func( st *SherryTime) DateTimeBaseFormat(datetime bool)(string) {
   var format strings.Builder
   fmt.Fprint(&format, "2006")
   fmt.Fprint(&format, st.Delimiter)
   fmt.Fprint(&format, "01")
   fmt.Fprint(&format, st.Delimiter)
   fmt.Fprint(&format, "02")
   if(datetime) {
      fmt.Fprint(&format, " 15:04:05")
   }
   return format.String()
}

// 回傳年月日
func(app *SherryTime) TodayYMD()(int, int, int) {
   t := strings.Split(app.Today(), app.Delimiter)
   y, _ := strconv.Atoi(t[0])
   m, _ := strconv.Atoi(t[1])
   d, _ := strconv.Atoi(t[2])
   return y, m, d
}

// 目前日期
func (st *SherryTime) Today()(string) {
   st.current = time.Now()
   return st.current.Format(st.DateTimeBaseFormat(false))
}

// 取得目前系統日期時間
func (st *SherryTime) Now() (string) {
   st.current = time.Now()
   return st.current.Format(st.DateTimeBaseFormat(true))
}

// 回傳目前系統時間
func(st *SherryTime) CurrentTime()([]string) {
   t := strings.Split(st.Now(), " ")
   return strings.Split(t[1], ":")
}

// 回傳純日期 "2022-09-21"
func(st *SherryTime) PureDate(d string)(string) {
   s := d
   t := strings.Split(d, " ")
   if len(t) >= 2 {
      s = t[0]
   }
   return s
}

// 日期基準點＋n天
func(st *SherryTime) DateAdd(stdate string, n int)(string) {
   ord := st.toDayOrdWs(stdate)
   ord = ord + n
   return st.toDateWs(ord)
   
}

// 計算時間差 分)return
func(st *SherryTime) TimeDiffStr(std, ed string)(int)  {
   x := strings.Split(std, ":")
   e := strings.Split(ed, ":")

   h, _ := strconv.Atoi(x[0])
   m, _ := strconv.Atoi(x[1])
   d, _ := strconv.Atoi(x[2])
   hh, _ := strconv.Atoi(e[0])
   mm, _ := strconv.Atoi(e[1])
   dd, _ := strconv.Atoi(e[2])
   
   if h > hh || h < hh {
      return (h - hh) * 3600
   } else if hh == hh {
      if m > mm || m < mm {
         return (m - mm) * 60
      } else {
         if d > dd || d < dd {
            return (d - dd) 
         }
      }
   } 
   return 0
}

// 計算兩個時間差距，回傳 "04")分
func(st *SherryTime) TimeDiff(stime, etime time.Time, format string)(string) {
   diff := etime.Sub(stime)
   out := time.Time{}.Add(diff)
   return out.Format(format)
}

// 計算兩個日期差距天數
func(st *SherryTime) DateDiff(stdate, ed string)(int) {
   return st.toDayOrdWs(st.PureDate(ed)) - st.toDayOrdWs(st.PureDate(stdate))
}

func(st *SherryTime) BaseX(i int)(string) {
   t := st.PureDate(st.Now());
   s := strings.Split(t, st.Delimiter)
   return s[i]
}
 
// 回傳本年度年分
func(st *SherryTime) Year()(string) {
   return st.BaseX(0)
}
// 回傳本年度月份
func(st *SherryTime) Month()(string) {
   return st.BaseX(1)
}
// 回傳目前日期
func(st *SherryTime) Day()(string) {
   return st.BaseX(2)
}

// 回傳大寫之年月日星期 return)年月日星期幾
func(st *SherryTime) SepDay(d string)(string, string, string, string) {
   source := strings.Split(d, st.Delimiter)
   dayOrder := st.toDayOrdWs(d)
   yy, mm, _, ww := st.toDateW(dayOrder)  // int
   s := st.toNumText(ww)
   if ww == 0 {
      s = "日"
   }
   return strconv.Itoa(yy), st.toNumText(mm), source[2], s
}

// 回傳大寫之年月日星期
func(st *SherryTime) SepToday()(string, string, string, string) {
   return st.SepDay(st.Today())
}

// 回傳特定日期的星期幾
func(app *SherryTime) WeekDay(yy, mm, dd int)(int, error) {
   m := ""
   d := ""
   if mm < 10 {
      m = fmt.Sprintf("0%d", mm)
   } else {
      m = fmt.Sprintf("%d", mm)
   }
   if dd < 10 {
      d = fmt.Sprintf("0%d", dd)
   } else {
      d = fmt.Sprintf("%d", dd)
   }
   dt := fmt.Sprintf("%d-%s-%s", yy, m, d)
   t, err := time.Parse("2006-01-02", dt)
   if err != nil {
      return -1, err
   }
   return int(t.Weekday()), nil
}

func NewSherryTime(locate, del string) (*SherryTime) {
   return &SherryTime {
      current: time.Now(),
      location: locate,
      numText: [13]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"},
      Delimiter: del,
      dayOfMonths: [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
   }
}
