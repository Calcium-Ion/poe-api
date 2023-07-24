# Golang Poe API
https://github.com/ading2210/poe-api 的golang版本

# Notice
Fork from https://github.com/lwydyby/poe-api, 在此基础上进行了一些修改，并修复了一些bug

# Instructions

## install

```bash
go get github.com/Calcium-Ion/poe-api
```

## use

```golang

import (
    "fmt"
    "log"
    "time"
    
    "github.com/Calcium-Ion/poe-api"
)


func ExampleSendMessage() {
    c, err := poe_api.NewClient("your p-b", "your formkey", nil)
    if err != nil {
    log.Printf("failed to create client: %v", err)
    }
	
    res, err := c.SendMessage("ChatGPT", "一句话描述golang的channel", true, 30*time.Second)
    if err != nil {
        log.Printf("failed to send message: %v", err)
    }
    fmt.Println(poe_api.GetFinalResponse(res))
	
    res, err = c.SendMessage("ChatGPT", "channel是并发安全的吗", false, 30*time.Second)
    if err != nil {
        log.Printf("failed to send message: %v", err)
    }
    // 流式返回 每次返回新增的数据
    for m := range poe_api.GetTextStream(res) {
        fmt.Println(m)
    }
}
```
