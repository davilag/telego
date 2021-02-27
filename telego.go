package telego

import (
	"log"

	"github.com/davilag/telego/api"
	"github.com/davilag/telego/kind"
	"github.com/davilag/telego/metrics"
)

// Telego is the main struct on which we can define all the flows and handlers.
type Telego struct {
	defaultHandler FlowStep           // Default handler which is going to be executed for those messages that don't have any flow assigned
	kindFlows      map[kind.Kind]Flow // Flows that are going to be executed based on the kind of the message
	commandFlows   map[string]Flow    // Flows that are going to be executed based on the command that the message has
	updates        chan api.Update    // Channel on which we have to send the updates to be processed
	Client         TelegramClient     // Client allows send messages to Telegram chats
}

var (
	metricMessageSent     = "telego_message_sent"
	metricMessageReceived = "telego_message_received"
	metricSession         = "telego_sessions"
	telegramHost          = "https://api.telegram.org/bot"
	produceMetrics        = false
)

// Initialise inits the telegram instance with the telegram bot access token.
// See https://core.telegram.org/bots/api#authorizing-your-bot
func Initialise(accessToken string) *Telego {
	log.Println("Local telego")
	client := TelegramClient{
		AccessToken: accessToken,
	}
	telego := Telego{
		kindFlows:    map[kind.Kind]Flow{},
		commandFlows: map[string]Flow{},
		Client:       client,
	}
	updates, _ := newSessionManager(&telego)
	telego.updates = updates
	return &telego
}

// SetDefaultMessageHandler Sets the default message handler for the telegram bot. It defines
// what we are going to do with messages that by default the bot doesn't understand
// (eg. send a description of the commands)
func (t *Telego) SetDefaultMessageHandler(f FlowStep) {
	t.defaultHandler = f
}

// AddKindHandler adds the step that it is going to be executed when we receive a message
// of a certain kind
func (t *Telego) AddKindHandler(k kind.Kind, fs FlowStep) {
	t.AddKindHandlerSession(k, fs, 0)
}

// AddKindHandlerSession adds the step that it is going to be executed when we receive
// a message of a certain kind, keeping the information that the handler saves for the
// time defined in ttl
func (t *Telego) AddKindHandlerSession(k kind.Kind, fs FlowStep, ttl int32) {
	f := Flow{
		ActualStep: fs,
	}
	t.kindFlows[k] = f
}

// AddCommandHandler adds the step that it is going to be executed when we receive a certain command
func (t *Telego) AddCommandHandler(c string, fs FlowStep) {
	t.AddCommandHandlerSession(c, fs, 0)
}

// AddCommandHandlerSession adds the step that it is going to be executed when we receive
// a message certain command, keeping the information that the handler saves for the
// time defined in ttl
func (t *Telego) AddCommandHandlerSession(c string, fs FlowStep, ttl int32) {
	f := Flow{
		ActualStep: fs,
		TimeToLive: ttl,
	}
	t.commandFlows[c] = f
}

// Listen main loop which is going to be listening for updates.
func (t *Telego) Listen() {
	var offset int
	fetch := true
	for fetch {
		us := t.Client.getUpdates(offset)
		for _, u := range us {
			addMessageReceivedMetric()
			t.updates <- u
			offset = u.UpdateID + 1
		}
	}
}

// SetupMetrics sets up the metrics for a telegram bot. It exposes metrics when
// the bot sends a message, when the bot receives a message and the sessions that
// the bot is keeping with chat information.
func (t *Telego) SetupMetrics() {
	metrics.AddCounter(metricMessageSent, "Telego sending a message")
	metrics.AddCounter(metricMessageReceived, "Telego receiving a message")
	metrics.AddGauge(metricSession, "Sessions that telego is keeping waiting for messages")
	go metrics.ExposeMetrics()
}

// SetTelegramHost sets the telegram host where telegram is running
func (t *Telego) SetTelegramHost(tm string) {
	telegramHost = tm
}

func addMessageSentMetric() {
	if produceMetrics {
		m, ok := metrics.GetCounter(metricMessageSent)
		if ok {
			m.Inc()
		}
	}
}

func addMessageReceivedMetric() {
	if produceMetrics {
		m, ok := metrics.GetCounter(metricMessageReceived)
		if ok {
			m.Inc()
		}
	}
}

func addSessionMetric() {
	if produceMetrics {
		m, ok := metrics.GetGauge(metricSession)
		if ok {
			m.Inc()
		}
	}
}

func finishSessionMetric() {
	if produceMetrics {
		m, ok := metrics.GetGauge(metricSession)
		if ok {
			m.Dec()
		}
	}
}
