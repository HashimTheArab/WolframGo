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
  wolframgo.SetAppId("Your App ID")
  
  res, err := wolframgo.GetSimpleResult("How many calories are in cereal?")
  
  if err != nil {
    panic(err)
  }
  fmt.Println(res)
}
```
