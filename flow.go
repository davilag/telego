package main

// Flow step, it accepts an update and a conversation and it returns the next step
// In order to no execute any more steps, it has to return nil
type FlowStep func(Update, Conversation) FlowStep

type Flow struct {
	ActualStep FlowStep // Step that has to be executed the next time that we receive an update
	TimeToLive int32    // Time to live of the session when executing this flow
}
