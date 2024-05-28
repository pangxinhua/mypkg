# my-pkg

# example
```go 
package main

import (
    "fmt"
    "github.com/pangxinhua/mypkg"
    "github.com/pangxinhua/mypkg/a"
    b "github.com/pangxinhua/mypkg/a/b"
)

func main() {
    var r string
    r = a.Say("tom")
    a.Hello("tom")
    a.Hello("tom1")
    b.World("jim")
    fmt.Println(r)
    r = mypkg.Read("tom")
    fmt.Println(r)
}

```
