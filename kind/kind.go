package kind

// Kind is an enum to define all the different kind of
// Telegram messages that we can receive
type Kind int

const (
	// Text is the default message kind
	Text Kind = 0
	// Photo is the message kind that we have when we receive a photo
	Photo Kind = 1
	// Video is the message kind that we have when we receive a video
	Video Kind = 2
	// Location is the message kind that we have when we receive a location
	Location Kind = 3
	// Command is the message kind that we have when we receive a command
	Command Kind = 4
	// Contact is the message kind that we have when we receive a contact
	Contact Kind = 5
	// Sticker is the message kind that we have when we receive a sticker
	Sticker Kind = 6
	// VoiceMessage is the message kind that we have when we receive a voice message
	VoiceMessage Kind = 7
	// VideoNote is the message kind that we have when we receive a telegram video note
	VideoNote Kind = 8
)
