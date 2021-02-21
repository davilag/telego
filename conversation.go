package telego

import (
	"time"

	"github.com/davilag/telego/api"
)

// Conversation abstraction to make send messages to an specific chat.
type Conversation struct {
	ChatID  int                    // Chat ID to which we are storing the session
	client  TelegramClient         // instance of the telegram client to make API calls
	Flow    Flow                   // Flow that is going to be executed in this instance of he conversation
	channel chan api.Update        // Channel from which we are going to receive new updates for the chat.
	exit    chan int               // Channel in which the conversation is going to send a message when the session is finished
	Context map[string]interface{} // Context which is going to allow to the user store data which is going to be accessible from every step
}

// NewConversation creates a conversation based on the update Chat Id
func newConversation(chatID int, f Flow, channel chan api.Update, exit chan int, telego *Telego) Conversation {
	return Conversation{
		ChatID:  chatID,
		client:  telego.Client,
		Flow:    f,
		channel: channel,
		exit:    exit,
		Context: make(map[string]interface{}),
	}
}

// SendMessage sends a message to the conversation
func (c *Conversation) SendMessage(m string) (api.Message, error) {
	return c.client.SendMessageText(m, c.ChatID)
}

// ReplyToMessage replies to a message in the conversation.
func (c *Conversation) ReplyToMessage(m string, messageID int) (api.Message, error) {
	return c.client.ReplyToMessage(m, c.ChatID, messageID)
}

// SendMessageWithKeyboard sends a message showing a list of options to the user as a custom keyboard
func (c *Conversation) SendMessageWithKeyboard(m string, keyboardOptions []string) (api.Message, error) {
	return c.client.SendMessageWithKeyboard(m, c.ChatID, keyboardOptions)
}

func (c *Conversation) SendVideo(f string) (api.Message, error) {
	return c.client.SendVideo(f, c.ChatID)
}

// This execution executes only one step, it doesn't create a session
func (c *Conversation) executeUpdate(u api.Update) FlowStep {
	return c.Flow.ActualStep(u, *c)
}

// Creates a session which is going to be listening to new updates in the c.channel.
// It's going to send a message with it chat id into the c.exit channel when it times out
// or the next step of the flow is nil
func (c *Conversation) createSession(requeueChan chan api.Update) {
	addSessionMetric()
	requeue := false
	for {
		select {
		case u := <-c.channel:
			if u.UpdateID == 0 {
				return
			}
			if requeue {
				requeueChan <- u
				break
			}
			c.Flow.ActualStep = c.executeUpdate(u)
			if c.Flow.ActualStep == nil {
				c.endSession()
			}
		case <-time.After(time.Duration(c.Flow.TimeToLive) * time.Second):
			c.endSession()
			requeue = true
		}
	}
}

// Ends the session sending a message with the chat Id to the exit channel
func (c *Conversation) endSession() {
	c.exit <- c.ChatID
	finishSessionMetric()
}
