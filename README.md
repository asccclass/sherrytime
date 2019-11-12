### Sherry Time Utility


#### file lists
```
sherrytime.go			// main function
sherrytime_test.go		// test function
```

### Installation
```
go get github.com/asccclass/sherrytime
```

### Usage
```
import(
   "github.com/asccclass/sherrytime"
)

func main() {
   st := NewSherryTime("Asia/Taipei", "-")  // Initial
   
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



### Reference
* [RFC 4122: Time-Based UUID](https://tools.ietf.org/html/rfc4122)
