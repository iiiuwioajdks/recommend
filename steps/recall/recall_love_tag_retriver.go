package recall_group

import (
	data "recommend/data"
	redis "recommend/tools/redis"
	"sort"
	"strconv"
)

func RecallLoveTagRetriever(rc *data.RequestContext) {
	if len(rc.UserProfile.LoveTags) == 0 {
		return
	}
	groups := make(map[int64]int)
	for _, loveTag := range rc.UserProfile.LoveTags {
		gids_str, err := redis.RecallRedisClient.SMembers(redis.RecallRedisClient.Context(), loveTag+redis.RecallGroupKey).Result()
		if err != nil {
			return
		}
		for _, str := range gids_str {
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				continue
			}
			groups[num]++
		}
	}
	rc.RecallMutex.Lock()
	defer rc.RecallMutex.Unlock()
	postRecallLoveTag(rc, groups)
}

func postRecallLoveTag(rc *data.RequestContext, groups map[int64]int) {
	// 根据相似度分数来排序，目前先不限制 quota
	type sortHelp struct {
		Gid   int64
		Score int
	}
	var groupScoreSlice []sortHelp
	for key, value := range groups {
		groupScoreSlice = append(groupScoreSlice, sortHelp{Gid: key, Score: value})
	}
	sort.Slice(groupScoreSlice, func(i, j int) bool {
		return groupScoreSlice[i].Score > groupScoreSlice[j].Score
	})
	for _, kv := range groupScoreSlice {
		rc.Groups = append(rc.Groups, data.Group{
			Gid: kv.Gid,
		})
	}
}
