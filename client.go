package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	telegramAPI = "https://api.telegram.org/bot"
	getUpdates  = "/getUpdates?offset="
	sendMessage = "/sendMessage"
)

type response struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

type TelegramClient struct {
	AccessToken string
}

// Returns a list of updates hitting the getUpdates method
// see https://core.telegram.org/bots/api#getupdates
func (c *TelegramClient) getUpdates(offset int) []Update {
	ep := fmt.Sprintf("%v%v%v%v", telegramAPI, c.AccessToken, getUpdates, offset)
	r, e := http.Get(ep)

	if e != nil {
		panic(e)
	}
	defer r.Body.Close()

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		panic(readErr)
	}

	obj := TelegramResponse{Result: &[]Update{}}
	result := obj.Result.(*[]Update)

	jsonErr := json.Unmarshal(body, &obj)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return *result
}

// It sends the given message to the given chat ID
func (c *TelegramClient) SendMessageText(message string, chatID int) error {
	mo := MessageOut{
		Text:   message,
		ChatID: chatID,
	}
	return c.SendMessage(mo)
}

// Send a message to a chat replying to the indicated message.
func (c *TelegramClient) ReplyToMessage(message string, chatID int, messageID int) error {
	mo := MessageOut{
		Text:             message,
		ChatID:           chatID,
		ReplyToMessageID: messageID,
	}

	return c.SendMessage(mo)
}

// Sends a message with the filled MessageOut object.
func (c *TelegramClient) SendMessage(m MessageOut) error {
	b, e := json.Marshal(m)

	if e != nil {
		return e
	}

	ep := fmt.Sprintf("%v%v%v", telegramAPI, c.AccessToken, sendMessage)
	resp, err := http.Post(ep, "application/json", bytes.NewReader(b))
	if err != nil {
		return e
	}
	defer resp.Body.Close()

	var body response
	json.NewDecoder(resp.Body).Decode(&body)

	if !body.Ok {
		return errors.New(body.Description)
	}

	return nil
}
