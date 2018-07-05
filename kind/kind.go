package kind

type MessageType int

const (
	Text     MessageType = 0
	Photo    MessageType = 1
	Video    MessageType = 2
	Location MessageType = 3
	Command  MessageType = 4
	Contact  MessageType = 5
	Sticker  MessageType = 6
)
