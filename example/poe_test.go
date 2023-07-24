package example

import (
	"fmt"
	"log"
	"time"

	"github.com/Calcium-Ion/poe-api"
)

func ExampleSendMessage() {
	c, err := poe_api.NewClient("", nil)
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
		panic(err)
	}
	for m := range poe_api.GetTextStream(res) {
		fmt.Println(m)
	}
	// output:
}
