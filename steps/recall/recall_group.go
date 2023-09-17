package recall_group

import (
	"encoding/json"
	data "recommend/data"
	redis_client "recommend/tools/redis"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Print(str string) {
	println(str)
}

func TetsSetGroup(gid int64, tag string) {
	group := data.Group{
		Gid: gid,
		GroupInfoData: data.GroupInfo{
			AuthId: 2,
			Tag:    []string{tag},
		},
	}
	jsonData, err := json.Marshal(group)
	if err != nil {
		logrus.Errorf("TestSetGroup err %s", err.Error())
	}
	gid_str := strconv.FormatInt(gid, 10)
	redis_client.GroupRedisClient.Set(redis_client.GroupRedisClient.Context(), gid_str+redis_client.GroupInfoKey, string(jsonData), 0)

	for _, str := range group.GroupInfoData.Tag {
		redis_client.RecallRedisClient.SAdd(redis_client.RecallRedisClient.Context(), str+redis_client.RecallGroupKey, gid_str)
	}
}

// i2i 召回
func GetGroupI2I(rc *data.RequestContext) {
	groups := make([]data.Group, len(rc.UserProfile.UA.LoveGids))
	for i, love_gid := range rc.UserProfile.UA.LoveGids {
		love_gid_str := strconv.FormatInt(love_gid, 10)
		val, _ := redis_client.GroupRedisClient.Get(redis_client.GroupRedisClient.Context(), love_gid_str+redis_client.GroupInfoKey).Result()
		json.Unmarshal([]byte(val), &groups[i])
	}
	// i2i tag 召回
	for _, group := range groups {
		for _, tag := range group.GroupInfoData.Tag {
			gids_str, err := redis_client.RecallRedisClient.SMembers(redis_client.RecallRedisClient.Context(), tag+redis_client.RecallGroupKey).Result()
			if err != nil {
				return
			}
			for _, str := range gids_str {
				num, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					continue
				}
				rc.Groups = append(rc.Groups, data.Group{
					Gid: num,
				})
			}
		}
	}
}
