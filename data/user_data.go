package data

type UserProfile struct {
	// 这些信息一般不会发生变化，存在 mysql 中
	Uid    int64
	Name   string
	Addr   string
	Age    int
	Gender int
	// UA 存 redis
	UA UserAction
}

type UserAction struct {
	LoveGids []int64
	HateGids []int64
	LoveTags []string
}
