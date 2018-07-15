package main

import (
	"github.com/davilag/telego/kind"
	"github.com/davilag/telego/metrics"
)

type Telego struct {
	defaultHandler FlowStep           // Default handler which is going to be executed for those messages that don't have any flow assigned
	kindFlows      map[kind.Kind]Flow // Flows that are going to be executed based on the kind of the message
	commandFlows   map[string]Flow    // Flows that are goingto be executed based on the command that the message has
	updates        chan Update        // Channel on which we have to send the updates to be processed
}

var (
	client TelegramClient
	telego Telego
)

// Initialises the telegram instance with the telegram bot access token
// See https://core.telegram.org/bots/api#authorizing-your-bot
func Initialise(accessToken string) Telego {
	client = TelegramClient{
		AccessToken: accessToken,
	}
	updates, _ := NewSessionManager()
	telego = Telego{
		kindFlows:    map[kind.Kind]Flow{},
		commandFlows: map[string]Flow{},
		updates:      updates,
	}

	return telego
}

// Sets the default message handler for the telegram bot. It defines what we are going to do
// with messages that by default the bot doesn't understand (eg. send a description of the commands)
func (t *Telego) SetDefaultMessageHandler(f FlowStep) {
	t.defaultHandler = f
}

func (t *Telego) AddKindHandler(k kind.Kind, fs FlowStep) {
	t.AddKindHandlerSession(k, fs, 0)
}

func (t *Telego) AddKindHandlerSession(k kind.Kind, fs FlowStep, ttl int32) {
	f := Flow{
		ActualStep: fs,
	}
	t.kindFlows[k] = f
}

func (t *Telego) AddCommandHanlder(c string, fs FlowStep) {
	t.AddCommandHanlderSession(c, fs, 0)
}

func (t *Telego) AddCommandHanlderSession(c string, fs FlowStep, ttl int32) {
	f := Flow{
		ActualStep: fs,
		TimeToLive: ttl,
	}
	t.commandFlows[c] = f
}

// Main loop which is goint to be listening for updates.
func (t *Telego) Listen() {
	var offset int
	fetch := true
	for fetch {
		us := client.getUpdates(offset)
		for _, u := range us {
			metrics.MessageReceived()
			telego.updates <- u
			offset = u.UpdateID + 1
		}
	}
}

func (t *Telego) SetupMetrics() {
	go metrics.SetupMetrics()
}
