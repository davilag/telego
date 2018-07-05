package main

// Abstraction to make send messages to an specific chat.
type Conversation struct {
	ChatID int
	client TelegramClient
}

// Creates a conversation based on the update Chat Id
func NewConversation(u Update, client TelegramClient) Conversation {
	return Conversation{
		ChatID: u.Message.Chat.ID,
		client: client,
	}
}

// Sends a message to the conversation
func (c *Conversation) sendMessage(m string) {
	c.client.SendMessageText(m, c.ChatID)
}

// Replies to a message in the conversation.
func (c *Conversation) replyToMessage(m string, messageId int) {
	c.client.ReplyToMessage(m, c.ChatID, messageId)
}
