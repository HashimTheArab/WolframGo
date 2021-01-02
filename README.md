# WolframGo
A simple lib for WolframAlpha made in Go.
install it with go get github.com/Prim69/WolframGo

## Example
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

## Example with a discord bot
```go

package main

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
        "github.com/prim69/wolframgo"
	"strings"
        "fmt"
)

type Ask struct {
	
}

func (p Ask) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
  args := strings.Split(ctx.GetMessage().Content, " ")
  
  if len(args) < 2 {
    _, _ = session.ChannelMessageSend(ctx.GetChannel().ID, "Ask me something.")
  } else {
    
    wolframgo.SetAppId("Your App ID")
    args := args[1:]
    question := strings.Join(args, " ")
    
    res, err := wolframgo.GetSimpleResult("How many calories are in cereal?")
  
    if err != nil {
       _, _ := session.ChannelMessageSend(ctx.GetChannel().ID, "An error occured: " + err.Error())
    }
   
	  content := discordgo.MessageSend{
			Content: res,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}
    
		_, err = session.ChannelMessageSendComplex(ctx.GetChannel().ID, &content)
	 return err
   }
   return nil
}

```
