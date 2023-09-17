package data

type UserProfile struct {
	Uid    int64
	Name   string
	Addr   string
	Age    int
	Gender int
	UA     UserAction
}

type UserAction struct {
	LoveGids []int64
	HateGids []int64
	LoveTags []string
}
