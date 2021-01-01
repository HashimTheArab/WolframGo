# WolframGo
A simple lib for WolframAlpha made in Go.

## Example
install it with go get github.com/Prim69/WolframGo
```go
package main

import (
  "fmt"
  "github.com/prim69/wolframgo"
)

func main() {
  wolframgo.appid = "your app id"
  
  res, err := wolframgo.GetSimpleResult("how many calories are in cereal")
  
  if err != nil {
    panic(err)
  }
  fmt.Println(res)
}
```
