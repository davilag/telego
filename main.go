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
	telego.SetupMetrics()
	telego.Listen()
}

func firstStep(u Update, c Conversation) FlowStep {
	fmt.Println("I'm in the first step")
	c.Context["message"] = u.Message.Text
	response, _ := c.SendMessage("Hi there first step!")
	fmt.Println(response.ID)
	return secondStep
}

func secondStep(u Update, c Conversation) FlowStep {
	fmt.Println("I'm in the second step")
	response, _ := c.SendMessage(fmt.Sprintf("Your previous message was: %s", c.Context["message"]))
	fmt.Println(response.ID)
	return nil
}

func locationHandler(u Update, c Conversation) FlowStep {
	c.ReplyToMessage("That's a location huh?", u.Message.ID)
	fmt.Println("Latitude", u.Message.Location.Latitude)
	fmt.Println("Longitude", u.Message.Location.Longitude)
	return nil
}
