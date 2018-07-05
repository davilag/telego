package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	telego := Initialise("282459302:AAGFKKE_pOPIBXbiVeR8CocXUYt2HShtJig")
	telego.SetDefaultMessageHandler(handler)
	telego.Listen()
}

func handler(u Update, c Conversation) {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
