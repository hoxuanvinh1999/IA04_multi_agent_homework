package agt

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type Alternative int

type AgentID interface {
	Equal(ag AgentID) bool
	DeepEqual(ag AgentID) bool
	Clone() AgentID
	String() string
	Prefers(a Alternative, b Alternative)
	Start()
}
