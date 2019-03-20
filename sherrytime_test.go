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
   
}
