package data

import "sync"

type RequestContext struct {
	UserProfile UserProfile
	Groups      []Group
	Strategy    string
	Steps       map[string]bool
	StepNum     int
	StepsRule   map[string]string
	// 多路召回
	RecallReasons     map[string][]Group
	RecallReasonNames []string
	RecallMutex       sync.Mutex
}
