package data

type RequestContext struct {
	UserProfile UserProfile
	Groups      []Group
	Strategy    string
	Steps       map[string]bool
	StepNum     int
	StepsRule   map[string]string
}
