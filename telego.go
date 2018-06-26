package main

type messageHandler func(Update, Conversation)

type Telego struct {
	client         TelegramClient
	defaultHandler messageHandler
}

func Initialise(accessToken string) Telego {
	return Telego{
		client: TelegramClient{
			AccessToken: accessToken,
		},
	}
}

func (t *Telego) SetDefaultMessageHandler(f messageHandler) {
	t.defaultHandler = f
}
