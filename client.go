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

func (c *TelegramClient) SendMessageText(message string, chatID int) error {
	mo := MessageOut{
		Text:   message,
		ChatID: chatID,
	}
	return c.SendMessage(mo)
}

func (c *TelegramClient) SendMessage(m MessageOut) error {
	b, e := json.Marshal(m)

	if e != nil {
		panic(e)
	}

	ep := fmt.Sprintf("%v%v%v", telegramAPI, c.AccessToken, sendMessage)
	resp, err := http.Post(ep, "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println(e)
	}
	defer resp.Body.Close()

	var body response
	json.NewDecoder(resp.Body).Decode(&body)

	if !body.Ok {
		return errors.New(body.Description)
	}

	return nil
}
