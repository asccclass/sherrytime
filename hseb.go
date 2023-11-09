/*
   算出某年的天干地支為何
*/
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
)

// 傳入參數：西元年，回傳天干地支
func (st *SherryTime) HSEB(year int) (string) {
   heavenlyStems := []string{"庚", "辛", "壬", "葵", "甲", "乙", "丙", "丁", "戊", "己"}
   earthlyBranches := []string{"申","酉","戌","亥", "子","丑","寅","卯","辰","巳","午","未"}
   hs := year % 10
   eb := year % 12
   return fmt.Sprintf("%s%s", heavenlyStems[hs], earthlyBranches[eb])
}
