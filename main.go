package main

import (
	"fmt"
	"os"

	"github.com/davilag/telego/kind"
)

func main() {
	telego := Initialise(os.Getenv("APIToken"))
	telego.AddCommandHanlderSession("start", firstStep, 30)
	telego.AddKindHandler(kind.Location, locationHandler)
	telego.Listen()
}

func firstStep(u Update, c Conversation) FlowStep {
	fmt.Println("I'm in the first step")
	c.Context["message"] = u.Message.Text
	c.SendMessage("Hi there first step!")
	return secondStep
}

func secondStep(u Update, c Conversation) FlowStep {
	fmt.Println("I'm in the second step")
	c.SendMessage(fmt.Sprintf("Your previous message was: %s", c.Context["message"]))
	return nil
}

func locationHandler(u Update, c Conversation) FlowStep {
	c.ReplyToMessage("That's a location huh?", u.Message.ID)
	return nil
}
