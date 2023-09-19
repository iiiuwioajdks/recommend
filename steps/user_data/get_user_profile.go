package userdata

import (
	"encoding/json"
	data "recommend/data"
	"recommend/tools/mysql"
	redis_client "recommend/tools/redis"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func TestSetUser(uid int64) {
	user := data.UserProfile{
		Name:     "lmx",
		Addr:     "广东",
		Age:      20,
		Gender:   0,
		Uid:      uid,
		LoveTags: []string{"后端", "推荐系统"},
		UA: data.UserAction{
			LoveGids: []int64{1, 5},
			HateGids: []int64{},
		},
	}
	// ua set redis
	jsonData, err := json.Marshal(user.UA)
	if err != nil {
		logrus.Errorf("TestSetUser err %s", err.Error())
	}
	uid_str := strconv.FormatInt(uid, 10)
	redis_client.UserRedisClient.Set(redis_client.UserRedisClient.Context(), uid_str+redis_client.UserProfileKey, string(jsonData), 0)
	// other info set mysql
	// 创建表
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS user_profiles (
			uid INT PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			addr VARCHAR(100) NOT NULL,
			age INT NOT NULL,
			gender INT NOT NULL,
			love_tags VARCHAR(255)
		)
	`

	_, err = mysql.UserMysqlClient.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
	insertQuery := `
		INSERT INTO user_profiles (uid, name, addr, age, gender, love_tags)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	loveTagStr := strings.Join(user.LoveTags, ",")
	mysql.UserMysqlClient.Exec(insertQuery, user.Uid, user.Name, user.Addr, user.Age, user.Gender, loveTagStr)
}

func GetUserProfileUA(rc *data.RequestContext, uid int64) {
	uid_str := strconv.FormatInt(uid, 10)
	// get user_profile ua
	val, err := redis_client.UserRedisClient.Get(redis_client.UserRedisClient.Context(), uid_str+redis_client.UserProfileKey).Result()
	if err != nil {
		logrus.Errorf("GetUserProfile redis err %s", err.Error())
	}
	json.Unmarshal([]byte(val), &rc.UserProfile.UA)
}

func GetUserProfileINFO(rc *data.RequestContext, uid int64) {
	query := "SELECT uid, name, addr, age, gender,love_tags FROM user_profiles where uid=?"
	love_tags := ""
	err := mysql.UserMysqlClient.QueryRow(query, uid).Scan(&rc.UserProfile.Uid, &rc.UserProfile.Name, &rc.UserProfile.Addr, &rc.UserProfile.Age, &rc.UserProfile.Gender, &love_tags)
	rc.UserProfile.LoveTags = strings.Split(love_tags, ",")
	if err != nil {
		logrus.Errorf("GetUserProfile mysql query err %s", err.Error())
	}
}
