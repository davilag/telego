package main

type Conversation struct {
	ChatID int
	client TelegramClient
}

func NewConversation(u Update, client TelegramClient) Conversation {
	return Conversation{
		ChatID: u.Message.Chat.ID,
		client: client,
	}
}

func (c *Conversation) sendMessage(m string) {
	c.client.SendMessageText(m, c.ChatID)
}

func (c *Conversation) replyToMessage(m string, messageId int) {
	c.client.ReplyToMessage(m, c.ChatID, messageId)
}
