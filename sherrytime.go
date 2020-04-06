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

type SherryTime struct {
   now time.Time
   numText [13]string
   location string
   delimiter string	// 分隔符號
   dayOfMonths [12]int	// 每月天數
}

var lock = sync.RWMutex{}

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
      n += yy * 365 + yy/4;
      if yy >= 4  {  // 西曆4年羅馬奧古斯都帝停閏
         n--
      }
      if yy >= 1600  {   // 教皇格勒哥里第13改曆
         i := yy - 1600;
         n -= i/100 - i/400;
      }
   } else {
      n = -1
   }
   return n
}

// 轉換成日序
func(st *SherryTime) toDayOrdWs(yymmdd string)(int)  {
   d := strings.Split(yymmdd, st.delimiter)
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

   yy = (int((float64(dOrd)/365.25)) + 1)
   ord := st.toDayOrdW(yy, 1, 1)
   if ord < 0 {
      return -1, 0, 0, 0
   }
   dd = dOrd - ord + 1
   mm = 1
   yes := yy == 1582 && mm == 10
   var n int = 0
   if yes {
      n = 21
   } else {
      n = st.lastMonthDay(yy, mm)
   }
   for dd > n  {
      dd -= n
      if mm++; mm > 12  {  
         yy++
         mm = 1
      }
   }
   if yes && dd > 4 {
      dd += 10
   }
   w := (dOrd + baseWeek) % 7
   return yy, mm, dd, w
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
   fmt.Fprint(&format, st.delimiter)
   fmt.Fprint(&format, "01")
   fmt.Fprint(&format, st.delimiter)
   fmt.Fprint(&format, "02")
   if(datetime) {
      fmt.Fprint(&format, " 15:04:05")
   }
   return format.String()
}

// 目前日期
func (st *SherryTime) Today()(string) {
   return st.now.Format(st.DateTimeBaseFormat(false))
}

// 取得目前日期時間
func (st *SherryTime) Now() (string) {
   return st.now.Format(st.DateTimeBaseFormat(true))
}

// 回傳純日期
func(st *SherryTime) PureDate(d string)(string) {
   s := d
   t := strings.Split(d, " ")
   if len(t) >= 2 {
      s = t[0]
   }
   return s
}

// 計算兩個日期差距天數
func(st *SherryTime) DateDiff(stdate, ed string)(int) {
   return st.toDayOrdWs(st.PureDate(ed)) - st.toDayOrdWs(st.PureDate(stdate))
}

// 回傳本年度年分
func(st *SherryTime) Year()(string) {
   t := st.PureDate(st.Now());
   s := strings.Split(t, st.delimiter)
   return s[0]
}

func NewSherryTime(locate, del string) (*SherryTime) {
   return &SherryTime {
      now: time.Now(),
      location: locate,
      numText: [13]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"},
      delimiter: del,
      dayOfMonths: [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
   }
}
