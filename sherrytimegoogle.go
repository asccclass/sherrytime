package sherrytime
// 轉換日期符合Google行事曆
import(
   "fmt"
   "strings"
   "strconv"
)


// daez)yyyy-mm-dd  timez)HH:MM return)西元年份月份日T時分秒
func (st *SherryTime) ToGoogleCalendarTime(datez, timez string, zone int)(string) {
   str := ""
   t := strings.Split(timez, ":")
   if len(t) != 2 {
      return str
   } 
   y := strings.Split(datez, "-")
   if len(y) != 3 {
      return str
   }
   h, err := strconv.Atoi(t[0])
   if err != nil {
      return str
   }
   h -= zone      // 切換時區
   ystr := ""
   if h < 0  { // 要先扣天數，然後取絕對值
      num := st.toDayOrdWs(y[0] + "-" + y[1] + "-" + y[2])
      num -= 1
      yy, mm, dd, _ := st.toDateW(num)
      ystr = fmt.Sprintf("%d%d%d", yy, mm, dd)
   } else {
      ystr =  strings.Join(y, "")
   }
   i := strconv.Itoa(h)
   if len(i) < 2  {
      i = "0" + i
   }
   str = ystr + "T" + i  + t[1] + "00"
   return str
}
