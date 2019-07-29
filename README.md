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
}
```

### Go Test
```
make test
```

### Functions list
* toDayOrdW(yy, mm, dd int)(int)	取得某年月日之日序 
* toDayOrdWs(yymmdd string)(int) 	取得某年月日（字串）之日序
