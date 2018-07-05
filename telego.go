package main

type messageHandler func(Update, Conversation)

type Telego struct {
	client         TelegramClient
	defaultHandler messageHandler
}

// Initialises the telegram instance with the telegram bot access token
// See https://core.telegram.org/bots/api#authorizing-your-bot
func Initialise(accessToken string) Telego {
	return Telego{
		client: TelegramClient{
			AccessToken: accessToken,
		},
	}
}

// Sets the default message handler for the telegram bot. It defines what we are going to do
// with messages that by default the bot doesn't understand (eg. send a description of the commands)
func (t *Telego) SetDefaultMessageHandler(f messageHandler) {
	t.defaultHandler = f
}
