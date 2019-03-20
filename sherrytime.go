package sherrytime
/**
 *  Note: 1.西曆4年羅馬奧古斯都帝停閏,故是年2月仍只有28天。
 *        2.教皇格勒哥里第13始改曆,自西曆1582年10月4日逕跳至15日或自1752年9月2日跳至14日
 *        3.西元前 n 年之值為 -n+1
*/


import (
   "time"
   "bytes"
   // log "github.com/sirupsen/logrus"
)

/*
   private $numText = array('零', '一', '二', '三', '四', '五', '六', '七', '八', '九', '十', '十一', '十二');
   private $_week = array('日','一','二','三','四','五','六');
   private $_weekx = array('一','二','三','四','五','六', '日');
*/

type SherryTime struct {
   now time.Time
   numText [12]string
   location string
   delimiter string	// 分隔符號
   dayOfMonths [12]int	// 每月天數
}

// 阿拉伯數字轉國字
func (st *SherryTime) toNumText(i int) (string) {  
   return st.numText[i] 
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

// 取得目前時間
func (st *SherryTime) Now() (string) {
   var format bytes.Buffer
   format.WriteString("2006")
   format.WriteString(st.delimiter)
   format.WriteString("01")
   format.WriteString(st.delimiter)
   format.WriteString("02 15:04:05")
   return st.now.Format(format.String())
}

func NewSherryTime(locate, del string) (*SherryTime) {
   return &SherryTime {
      now: time.Now(),
      location: locate,
      numText: ['�~[�', '�~@', '�~L', '�~I', '�~[~[', '�~T', '�~E�', '�~C', '�~E�', '�~]', '�~M~A',, '�~M~A�~@', '�]M~
      delimiter: del,
      dayOfMonths: [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
   }
}

/*
func main() {
   st := NewSherryTime("Asia/Taipei", "-")

   log.Printf("%v", st.Now()) // local time
}
*/
