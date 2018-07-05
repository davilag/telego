package main

import "github.com/davilag/telego/kind"

func (m *Message) GetMessageKind() kind.MessageType {
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

	return kind.Text
}
