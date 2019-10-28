package sherrytime
/*
   將不同日期格式轉化成正確日期格式 yyyy-mm-dd H:i:s
*/

import(
   "fmt"
   "strings"
   "strconv"
)

// Year2West  將年份轉換為西元年
func(st *SherryTime) Year2West(d string, isWestYear bool)(string) {
   if(isWestYear) {  // 西元年
      if len(d) == 2 {  // 19 格式
         today := strings.Split(st.Today(), st.delimiter) // 取得今日日期
         d = today[0][0:2] + d
      } else if len(d) == 4 {
         return d
      }
   }  else {  // 民國年
      year, err := strconv.Atoi(d)
      if err != nil {
         return ""
      }
      year += 1911   // 民國年轉為西元年
      return strconv.Itoa(year)
   }
   return d
}

// 轉換日期格式（不含時分秒）
func (st *SherryTime) TransferFormat(d string)(string, error) {
   if d == "" {
      return "", fmt.Errorf("未傳入日期資料")
   }
   var dt strings.Builder
   fmt.Fprint(&dt, "")

   // 先判斷是否有時分秒格式
   x := strings.Split(d, " ")
   if len(x)  == 2 {
      return "", fmt.Errorf("%s 格式錯誤，僅容許日期", d)
   }

   // 取得今日日期
   today := strings.Split(st.Today(), st.delimiter)
   str := strings.Split(d, st.delimiter)
   if len(str) != 3 {
      return "", fmt.Errorf("date format error")
   }
   if len(str[0]) != 4 {  // not 2019
      if(str[2] == today[0][2:4])  {   // 月-日-年
         _, err := fmt.Fprintf(&dt, "%s%s%s%s%s", st.Year2West(str[2], true), st.delimiter, str[0], st.delimiter, str[1])
         if err != nil {
            return "", err
         }
      } 
   }  else {   // 格式正常 2019-10-28
      fmt.Fprintf(&dt, d)
   }
   return dt.String(), nil
}
