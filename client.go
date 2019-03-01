package telego

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davilag/telego/api"
	"github.com/mitchellh/mapstructure"
)

const (
	getUpdates  = "/getUpdates?offset="
	sendMessage = "/sendMessage"
)

// TelegramClient manages the connection to the Telegram API
type TelegramClient struct {
	AccessToken string
}

// Returns a list of updates hitting the getUpdates method
// see https://core.telegram.org/bots/api#getupdates
func (c *TelegramClient) getUpdates(offset int) []api.Update {
	ep := fmt.Sprintf("%v%v%v%v", telegramHost, c.AccessToken, getUpdates, offset)
	r, e := http.Get(ep)

	if e != nil {
		panic(e)
	}
	defer r.Body.Close()

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		panic(readErr)
	}

	obj := api.TelegramResponse{Result: &[]api.Update{}}
	result := obj.Result.(*[]api.Update)

	jsonErr := json.Unmarshal(body, &obj)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return *result
}

// SendMessageText sends the given message to the given chat ID
func (c *TelegramClient) SendMessageText(message string, chatID int) (api.Message, error) {
	mo := api.MessageOut{
		Text:   message,
		ChatID: chatID,
	}
	return c.SendMessage(mo)
}

// ReplyToMessage sends a message to a chat replying to the indicated message.
func (c *TelegramClient) ReplyToMessage(message string, chatID int, messageID int) (api.Message, error) {
	mo := api.MessageOut{
		Text:             message,
		ChatID:           chatID,
		ReplyToMessageID: messageID,
	}

	return c.SendMessage(mo)
}

// SendMessageWithKeyboard sends a message showing a list of options to the user as a custom keyboard
func (c *TelegramClient) SendMessageWithKeyboard(message string, chatID int, keyboardOptions []string) (api.Message, error) {
	m := api.MessageOut{
		Text:   message,
		ChatID: chatID,
	}
	var rkm api.ReplyKeyboardMarkup

	rkm.Keyboard = make([][]api.KeyboardButton, len(keyboardOptions))
	for i, o := range keyboardOptions {
		rkm.Keyboard[i] = []api.KeyboardButton{
			api.KeyboardButton{
				Text: o,
			},
		}
	}
	m.ReplyMarkup = rkm

	return c.SendMessage(m)
}

// SendMessage sends a message with the filled MessageOut object.
func (c *TelegramClient) SendMessage(mo api.MessageOut) (api.Message, error) {
	b, e := json.Marshal(mo)

	if e != nil {
		return api.Message{}, e
	}

	ep := fmt.Sprintf("%v%v%v", telegramHost, c.AccessToken, sendMessage)
	resp, err := http.Post(ep, "application/json", bytes.NewReader(b))
	if err != nil {
		return api.Message{}, err
	}
	defer resp.Body.Close()

	var body api.TelegramResponse
	json.NewDecoder(resp.Body).Decode(&body)

	if !body.Ok {
		return api.Message{}, errors.New(body.Description)
	}
	addMessageSentMetric()
	var m api.Message

	err = mapstructure.Decode(body.Result, &m)
	if err != nil {
		return api.Message{}, err
	}
	return m, nil
}
