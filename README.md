# **GOENV**

It is simple librari for working with env.

### `go go get github.com/youpps/genv`

## **Examples:**

With Load function:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/youpps/genv"
)

func main() {
    err := genv.Load(".env-custom")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(os.Getenv("KEY")) // 0987654321
}
```

Or autoload:

```go
package main

import (
    "fmt"
    "os"

    _ "github.com/youpps/genv/autoload"
)

func main() {
    fmt.Println(os.Getenv("KEY")) // 1234567890
}
```
