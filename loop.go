package main

func (t *Telego) Listen() {
	var offset int
	for {
		us := t.client.getUpdates(offset)

		for _, u := range us {
			conv := NewConversation(u, t.client)
			if t.defaultHandler != nil {
				go t.defaultHandler(u, conv)
			}
			offset = u.UpdateID + 1
		}
	}
}
