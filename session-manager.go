package telego

type sessionManager struct {
	update   chan Update         // Update channel which expects new messages from Telegram
	exit     chan int            // Exit channel where it expects messages from the conversations to finish them
	requeue  chan Update         // Channel to receive requeue messages
	channels map[int]chan Update // Map from ChatID to channel which stores the channel to communicate with live sessions.
}

// newSessionManager initialises a session manager which is going to manage the sessions
// by chat id. It returns 3 channels, the first one is the channel where
// the session manager expects new updates from telegram, the second
// channel is the channel where the manager expects a message from the
// conversations to finish the session and the third channel is the channel where
// we are going to requeue messages that we were assigned to a session the first time
// but that they should be process as a new session.
func newSessionManager() (chan Update, chan int) {
	s := sessionManager{
		update:   make(chan Update, 100),
		exit:     make(chan int, 100),
		requeue:  make(chan Update, 100),
		channels: map[int]chan Update{},
	}
	go s.manageChannels()

	return s.update, s.exit
}

// Main loop for the session manager. It's listening for events in the exit and
// update channels. It gives priority to the exit and requeue channel.
// The reason why we're giving priority to the exit and the requeue channel is because
// we want to finish the session as soon as we've reached the timeout and if we have
// a requeued message, we want to give it higher priority.
func (s *sessionManager) manageChannels() {
	for {
		select {
		case cID := <-s.exit:
			s.doExit(cID)
			continue
		case u := <-s.requeue:
			s.manageUpdate(u)
		default:
		}
		select {
		case cID := <-s.exit:
			s.doExit(cID)
			continue
		case u := <-s.requeue:
			s.manageUpdate(u)
		case u := <-s.update:
			s.manageUpdate(u)
		}
	}
}

// Method to manage an update comming from the telegram API
func (s *sessionManager) manageUpdate(u Update) {
	chatID := u.Message.Chat.ID
	v, ok := s.channels[chatID]

	if !ok {
		v = s.startConversation(u)
	}
	if v != nil {
		v <- u
	}
}

func (s *sessionManager) doExit(cID int) {
	close(s.channels[cID])
	delete(s.channels, cID)
}

// Given a message, it checks if it contains any command
// and returns a flow based on that.
func getCommandFlows(m *Message) (Flow, bool) {
	command := m.GetCommand()
	if command == "" {
		return Flow{}, false
	}
	value, ok := telego.commandFlows[command]
	return value, ok
}

// Given a message, it checks its kind and returns a flow
// based on it.
func getKindFlows(m *Message) (Flow, bool) {
	k := m.GetKind()
	value, ok := telego.kindFlows[k]
	return value, ok
}

// Gets the flow to execute given a message, it gives priority
// to the command flows. Returns the default handler defined in
// the package if the message doesn't match any flow.
func getFlow(m *Message) Flow {
	if f, ok := getCommandFlows(m); ok {
		return f
	}
	if f, ok := getKindFlows(m); ok {
		return f
	}
	return Flow{
		ActualStep: telego.defaultHandler,
	}
}

// Initialises a conversation, retrieving the flow defined for each command/kind
// and executing the default handler if no flow has been defined for that message
func (s *sessionManager) startConversation(u Update) chan Update {
	f := getFlow(u.Message)
	if f.ActualStep == nil {
		return nil
	}
	var cu chan Update

	// We only create the channel if the flow has time to live
	// Otherwise, we don't create any channel as the flow will
	// execute just one handler.
	if f.TimeToLive != 0 {
		cu = make(chan Update)
		s.channels[u.Message.Chat.ID] = cu
	}
	c := NewConversation(u.Message.Chat.ID, f, cu, s.exit)

	// If the channel hasn't been created, we just execute the
	// handler for that update.
	if cu == nil {
		go c.executeUpdate(u)
	} else {
		go c.createSession(s.requeue)
	}
	return cu
}
