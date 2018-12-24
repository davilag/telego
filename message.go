package telego

import "github.com/davilag/telego/kind"

// GetKind Returns the message kind based on the parameters that it has
func (m *Message) GetKind() kind.Kind {

	if m.Entities != nil {
		if m.getBotCommandEntity() != nil {
			return kind.Command
		}
	}

	if m.Location != nil {
		return kind.Location
	}
	if m.Contact != nil {
		return kind.Contact
	}
	if m.Photo != nil {
		return kind.Photo
	}
	if m.Sticker != nil {
		return kind.Sticker
	}
	if m.Video != nil {
		return kind.Video
	}
	if m.Voice != nil {
		return kind.VoiceMessage
	}
	if m.VideoNote != nil {
		return kind.VideoNote
	}

	// If we haven't identified the type of the message, we treat it as
	// a plain message.
	return kind.Text
}

// Returns the Message Entity which type is bot_command
func (m *Message) getBotCommandEntity() *MessageEntity {
	if m.Entities == nil {
		return nil
	}

	for _, e := range *m.Entities {
		if e.Type == "bot_command" {
			return &e
		}
	}

	return nil
}

// GetCommand retrieves the command from the Message.Text field
func (m *Message) GetCommand() string {
	me := m.getBotCommandEntity()

	if me == nil {
		return ""
	}

	runes := []rune(m.Text)
	return string(runes[me.Offset+1 : me.Offset+me.Length])
}
