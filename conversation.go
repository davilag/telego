package main

import "time"

// Abstraction to make send messages to an specific chat.
type Conversation struct {
	ChatID  int                    // Chat ID to which we are storing the session
	client  TelegramClient         // instance of the telegram client to make API calls
	Flow    Flow                   // Flow that is going to be executed in this instance of he conversation
	channel chan Update            // Channel from which we are going to receive new updates for the chat.
	exit    chan int               // Channel in which the conversation is going to send a message when the session is finished
	Context map[string]interface{} // Context which is going to allow to the user store data which is going to be accessible from every step
}

// Creates a conversation based on the update Chat Id
func NewConversation(chatId int, f Flow, channel chan Update, exit chan int) Conversation {
	return Conversation{
		ChatID:  chatId,
		client:  client,
		Flow:    f,
		channel: channel,
		exit:    exit,
		Context: make(map[string]interface{}),
	}
}

// Sends a message to the conversation
func (c *Conversation) SendMessage(m string) {
	client.SendMessageText(m, c.ChatID)
}

// Replies to a message in the conversation.
func (c *Conversation) ReplyToMessage(m string, messageId int) {
	client.ReplyToMessage(m, c.ChatID, messageId)
}

// This execution executes only one step, it doesn't create a session
func (c *Conversation) executeUpdate(u Update) FlowStep {
	return c.Flow.ActualStep(u, *c)
}

// Creates a session which is going to be listening to new updates in the c.channel.
// It's going to send a message with it chat id into the c.exit channel when it times out
// or the next step of the flow is nil
func (c *Conversation) createSession() {
	for {
		select {
		case u := <-c.channel:
			c.Flow.ActualStep = c.executeUpdate(u)
			if c.Flow.ActualStep == nil {
				c.endSession()
				return
			}
		case <-time.After(time.Duration(c.Flow.TimeToLive) * time.Second):
			c.endSession()
			return
		}
	}
}

// Ends the session sending a message with the chat Id to the exit channel
func (c *Conversation) endSession() {
	c.exit <- c.ChatID
}
