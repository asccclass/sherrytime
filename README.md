### Sherry Time Utility

世界協調時間（簡稱UTC）是最主要的世界時間標準，其以原子時秒長為基礎，在時刻上儘量接近於格林威治標準時間。


#### file lists
```
sherrytime.go			// main function
sherrytime_test.go		// test function
```

### Installation
```
go get github.com/google/uuid
go get github.com/asccclass/sherrytime
```

### Usage
```
import(
   "github.com/asccclass/sherrytime"
)

func main() {
   st := sherrytime.NewSherryTime("Asia/Taipei", "-")  // Initial
   
   log.Printf("%v", st.Now()) // current time: "2019-07-30 12:15:12"

   // 產生UUID
   fmt.Println(st.NewUUID()) // 2bf08894-47a3978-bbdd8c85-70d16847
}
```

### Go Test
```
make test
```

### Functions list
* toDayOrdW(yy, mm, dd int)(int)	取得某年月日之日序 
* toDayOrdWs(yymmdd string)(int) 	取得某年月日（字串）之日序
* UnixTime2Timestamp(second string)(string, error)	將Unix秒數時間，轉換為local的timestamp


### Reference
* [RFC 4122: Time-Based UUID](https://tools.ietf.org/html/rfc4122)
* [Comprehensive Guide to Dates and Times in Go](https://qvault.io/golang/golang-date-time/)
