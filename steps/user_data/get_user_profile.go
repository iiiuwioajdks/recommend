package userdata

import (
	"encoding/json"
	data "recommend/data"
	redis_client "recommend/tools/redis"
	"strconv"

	"github.com/sirupsen/logrus"
)

func TestSetUser(uid int64) {
	user := data.UserProfile{
		Name:   "lmx",
		Addr:   "广东",
		Age:    20,
		Gender: 0,
		Uid:    uid,
		UA: data.UserAction{
			LoveGids: []int64{1, 5},
			HateGids: []int64{},
			LoveTags: []string{"篮球", "计算机"},
		},
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("TestSetUser err %s", err.Error())
	}
	uid_str := strconv.FormatInt(uid, 10)
	redis_client.UserRedisClient.Set(redis_client.UserRedisClient.Context(), uid_str+redis_client.UserProfileKey, string(jsonData), 0)
}

func GetUserProfile(rc *data.RequestContext, uid int64) {
	uid_str := strconv.FormatInt(uid, 10)
	val, err := redis_client.UserRedisClient.Get(redis_client.UserRedisClient.Context(), uid_str+redis_client.UserProfileKey).Result()
	if err != nil {
		logrus.Errorf("GetUserProfile err %s", err.Error())
	}
	json.Unmarshal([]byte(val), &rc.UserProfile)
}
