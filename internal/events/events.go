package events
import "time"

//AgentID identifies one agent, for example "a1"
type AgentID string 

//Event is anything agent can report 
type Event interface{
	AgentID() AgentID
	Time() time.Time
}

//Meta holds the fields every event shares
type Meta struct{
	Agent AgentID
	At time.Time
}

//NewMeta stamps an event for agent id with the current time
func NewMeta(id AgentID) Meta{
	return Meta{Agent: id, At: time.Now()}
}

func (m Meta) AgentID() AgentID {return m.Agent}
func (m Meta) Time() time.Time {return m.At}

//TextDelta is a chunk of output the agent printed
type TextDelta struct{
	Meta
	Text string
}

//Status Changed reports a move from one status to another
type StatusChanged struct {
	Meta
	From Status
	To Status
}

//Exit reports that the agent's process ended
type Exit struct {
	Meta
	Code int
}
