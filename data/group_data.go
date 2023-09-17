package data

type Group struct {
	Gid           int64
	GroupInfoData GroupInfo
}

type GroupInfo struct {
	AuthId int64
	Tag    []string
}
