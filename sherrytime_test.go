/*
    Test function) for test sherrytime.go 
    by C.H Liu
    date: 2019.03.21
*/
package sherrytime

import (
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
   format.WriteString(st.delimiter)
   format.WriteString("01")
   format.WriteString(st.delimiter)
   format.WriteString("02 15:04:05")
   expect := tm.Format(format.String())
   got := st.Now()
   if got != expect {
      t.Errorf("got [%s] expected [%s]", got, expect)
   } else {
      log.Printf("Function Now() passed.")
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
   
}
