# SherryTime library

**SherryTime** is a time utility package written in Go. It provides a simple API for handling time zones, time formatting, UUID generation, and more.

## 📦 Installation

Install using `go get`:

```bash
go get github.com/asccclass/sherrytime
```

If you want to use the UUID feature, also install:

```bash
go get github.com/google/uuid
```

## 🚀 Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/asccclass/sherrytime"
)

func main() {
    // Initialize SherryTime with time zone Asia/Taipei and date-time separator "-"
    st := sherrytime.NewSherryTime("Asia/Taipei", "-")

    // Get current time
    log.Printf("Current Time: %v", st.Now())

    // Get today's date
    fmt.Println("Today's Date:", st.Today())

    // Generate a new UUID
    fmt.Println("Generated UUID:", st.NewUUID())
}
```

## 🧪 Testing

Run tests using Makefile:

```bash
make test
```

## 📚 Features

| Function Name | Return Type | Description           |
|---------------|-------------|-----------------------|
| `Today()`     | `string`    | Get today's date      |
| `Now()`       | `string`    | Get current time      |
| `NewUUID()`   | `string`    | Generate a new UUID   |

## 📁 File Structure

| File Name              | Description                        |
|------------------------|------------------------------------|
| `sherrytime.go`        | Main feature implementation        |
| `sherrytime_test.go`   | Unit tests                         |
| `sherrytimeformat.go`  | Time formatting functions          |
| `sherrytimegoogle.go`  | Google-related time features       |
| `sherrytimenanoid.go`  | Generate unique ID using NanoID    |
| `sherrytimeuuid.go`    | Generate unique ID using UUID      |
| `hseb.go`              | Additional helper functions        |
| `makefile`             | Automation scripts for build/test  |

## 📄 License

This project is licensed under the MIT License.

## 🙋‍♂️ Author

Developed and maintained by [LIU CHIH HAN (asccclass)](https://github.com/asccclass).

---

For more information, visit the [sherrytime GitHub page](https://github.com/asccclass/sherrytime).
