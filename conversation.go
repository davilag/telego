package main

type Conversation struct {
	ChatId int
	client TelegramClient
}

func NewConversation(u Update, client TelegramClient) Conversation {
	return Conversation{
		ChatId: u.Message.Chat.ID,
		client: client,
	}
}

func (c *Conversation) sendMessage(m string) {
	c.client.SendMessageText(m, c.ChatId)
}

func (c *Conversation) replyToMessage(m string, messageId int) {
	c.client.ReplyToMessage(m, c.ChatId, messageId)
}
