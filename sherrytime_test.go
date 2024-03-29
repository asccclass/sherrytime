/*
    Test function) for test sherrytime.go 
    by C.H Liu
    date: 2019.03.21
*/
package sherrytime

import (
   "fmt"
   "testing"
   "time"
   "bytes"
   log "github.com/sirupsen/logrus" 
)

func TestNewSherryTime(t *testing.T) {
   // test st.Now() function is right or wrong
   st := NewSherryTime("Asia/Taipei", "-")
   tm := time.Now()
   var format bytes.Buffer
   format.WriteString("2006")
   format.WriteString(st.Delimiter)
   format.WriteString("01")
   format.WriteString(st.Delimiter)
   format.WriteString("02 15:04:05")
   expect := tm.Format(format.String())
   got := st.Now()
   if got != expect {
      t.Errorf("got [%s] expected [%s]", got, expect)
   } else {
      log.Printf("Function Now() passed.")
   }

   // 測試兩個時間差距
   diffTime := st.TimeDiff(tm, tm.Add(time.Second * 600), "04")
   if diffTime != "10" {
      t.Errorf("st.TimeDiff() has some problem.(should 10 but got %s)", diffTime)
   } else {
      log.Printf("Function TimeDiff(stime, etime time.Time, format string)(string) passed.")
   }

   // Is leapYear(year int) function checked.
   var leapYear int = 1880
   var IsLeap bool = true
   if st.leapYear(leapYear) != IsLeap {
      t.Errorf("st.leapYear() has some problem.(%d is leap year)", leapYear)
   } else {
      log.Printf("Function leapYear(int year) passed.")
   }

   // test lastMonthDay(yy, mm int) 
   var hasDays int = 31  // 2019.03 has 31 days
   if st.lastMonthDay(2019, 3) != hasDays {
      t.Errorf("st.lastMonthDay() has some problem.(Should be %d days)", hasDays)
   } else {
      log.Printf("Function lastMonthDays(yy, mm int) passed.")
   }

   // test chkDateW(year, month, day int) (bool) function
   var expectAns int = -1
   if st.chkDateW(2019, 2, 29) != expectAns  || st.chkDateW(2019, 3, 31) != 0 {
      t.Errorf("chkDateW(2019, 2, 29)(bool) has some problem.(Should be %d)", expectAns)
   } else { 
      log.Printf("Function chkDateW(year, month, day int)(bool) passed.")
   }

   // test toNumText(i int) function
   var zero = "零"
   var third = "三"
   if st.toNumText(0) != zero || st.toNumText(3) != third {
      t.Errorf("chkDateW(2019, 2, 29)(bool) has some problem. %v %v", st.toNumText(0), st.toNumText(4))
   } else { 
      log.Printf("Function toNumText(i int)(string) passed.")
   }

   // test toDayOrdW(yy, mm, dd int)
   if st.toDayOrdW(1, 1, 1) != 1  {
      t.Errorf("toDayOrdW(yy, mm, dd)(int) has some problem. value should be 1, but got %v", st.toDayOrdW(1,1,1))
   } else { 
      log.Printf("Function toDayOrdW(yy, mm, dd int)(int) passed.")
   }

   // test toDayOrdWs(yymmdd string)(int)
   if st.toDayOrdWs("1-1-1") != 1 {
      t.Errorf("toDayOrdWs(yymmdd string)(int) has some problem.")
   } else { 
      log.Printf("Function toDayOrdWs(yymmdd string)(int) passed.")
   }

   // test toDateW(dOrd int)
   yy, mm, dd, w := st.toDateW(1)
   if yy != 1 || mm != 1 || dd != 1 || w != 0 {
      t.Errorf("toDateW(dOrd int)(yy, mm, dd, w int) has some problem.")
   } else { 
      log.Printf("Function toDateW(dOrd int)(yy, mm, dd, w int) passed.")
   }

   // test toDateWs(dOrd int) (string)
   if st.toDateWs(1) != "1-1-1" {
      t.Errorf("toDateWs(dOrd int) (string) has some problem.")
   } else { 
      log.Printf("Function toDateWs(dOrd int) (string) passed.")
   }

   // test Today function
   if tm.Format("2006-01-02") != st.Today() {
      t.Errorf("Today() (string) has some problem.")
   } else { 
      log.Printf("Function Today() (string) passed.")
   }

   // test func (st *SherryTime) TransferFormat(d string)(string, error)
   transD, err := st.TransferFormat("10-17-19")
   if err != nil || transD != "2019-10-17" {
      if err != nil {
         if err == fmt.Errorf("%s 格式錯誤，僅容許日期", transD) {
            log.Printf("Function TransferFormat('10-17-19')(string, error) format error passed.")
         } else {
            t.Errorf("Function TransferFormat('10-17-19')(string, error) has some problem")
         }
      }  else {
         log.Printf("Function TransferFormat()(string, error) passed.")
      }
   } 
   transD, err = st.TransferFormat("2019-10-12")
   if err != nil {
      if transD != "2019-10-12" {
         t.Errorf("Function TransferFormat('2019-10-12')(string, error) has some problem." + err.Error())
      }  else {
         log.Printf("Function TransferFormat(d string)(string, error) passed.")
      }
   } else {
      log.Printf("Function TransferFormat('2019-10-12')(string, error) passed.")
   }
   log.Printf(st.NewUUID())

   // 測試DateDiff(st, ed string) 日期天數
   d := st.DateDiff("2020-03-21", "2020-03-25")
   if d != 4 {
         t.Errorf("Function DateDiff(st, ed string) has some problem.")
   } else {
         log.Printf("Function DateDiff(st, ed string) passed.")
   }
   d = st.DateDiff("2020-03-21 12:02:01", "2020-03-25 13:21:11")
   if d != 4 {
         t.Errorf("Function DateDiff(st, ed string) with time has some problem.")
   } else {
         log.Printf("Function DateDiff(st, ed string) with time passed.")
   }
   // 測試Year()
   yearx := st.Year()
   if yearx != "2020" {
         t.Errorf("Function st.Year() has some problem.")
   } else {
         log.Printf("Function Year() with time passed.")
   }
   yearx = st.DateAdd("2020-04-10", 5)
   if yearx != "2020-04-15" {
         t.Errorf("Function st.DateAdd(string, int) has some problem. should 2020-04-15, but got %s", yearx)
   } else {
         log.Printf("Function DateAdd(string, int) with time passed.")
   }

   // 天干地支
   hseb := st.HSEB(2020)
   if hseb != "庚子" {
      t.Errorf("Function st.DateAdd(string, int) has some problem. should 2020-04-15, but got %s", yearx)
   } else {
      log.Printf("HSEB passed.")
   }
}
