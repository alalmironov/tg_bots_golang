package main

type Update struct {
	Message Message
}

type Message struct {
	Text string
	Chat Chat
}

type Chat struct {
	Id int
}

type SendMessageRequest struct {
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}