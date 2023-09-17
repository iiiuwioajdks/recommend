package main

import (
	"encoding/json"
	data "recommend/data"
	dag_execute "recommend/execute"
	recall_group_exec "recommend/steps/recall"
	user_data_exec "recommend/steps/user_data"
	redis "recommend/tools/redis"

	"github.com/sirupsen/logrus"
)

func print(str string) {
	println(str)
}

const user_data_rule = `
rule "test_set_user"
begin
set_user(1)
end

rule "get_user_profile"
begin
get_user_profile(rc, 1)
end
`

const recall_rule = `
rule "test_set_group"
begin
set_group(1, "篮球")
set_group(2, "游戏")
set_group(3, "游泳")
set_group(4, "篮球")
set_group(5, "计算机")
end

rule "recall_i2i"
begin
recall_i2i(rc)
end

rule "merge"
begin
println("run recall merge")
end
`

const sort_rule = `
rule "predict"
begin
println("run sort predict")
end
`

func main() {
	// init
	redis.CreateRedisClient()

	rc := new(data.RequestContext)

	execute := new(dag_execute.DAG_EXECUTE)
	execute.Init(3)
	user_data_func := map[string]interface{}{
		"rc":               rc,
		"println":          print,
		"set_user":         user_data_exec.TestSetUser,
		"get_user_profile": user_data_exec.GetUserProfile,
	}
	execute.InitRuleBuilder(0, user_data_rule, user_data_func)
	recall_func := map[string]interface{}{
		"rc":         rc,
		"println":    print,
		"set_group":  recall_group_exec.TetsSetGroup,
		"recall_i2i": recall_group_exec.GetGroupI2I,
	}
	execute.InitRuleBuilder(1, recall_rule, recall_func)
	sort_func := map[string]interface{}{
		"println": print,
	}
	execute.InitRuleBuilder(2, sort_rule, sort_func)
	user_data_dag := [][]string{
		{},
		{"test_set_user"},
		{"get_user_profile"},
	}
	execute.InitRuleDag(0, user_data_dag)
	recall_dag := [][]string{
		{},
		{"test_set_group"},
		{"recall_i2i"},
		{"merge"},
	}
	execute.InitRuleDag(1, recall_dag)
	sort_dag := [][]string{
		{},
		{"predict"},
	}
	execute.InitRuleDag(2, sort_dag)

	execute.ExecuteDag()

	jsonData, _ := json.Marshal(rc.UserProfile)
	logrus.Infof("lmx_test user_profile : %s", jsonData)

	jsonData2, _ := json.Marshal(rc.Groups)
	logrus.Infof("lmx_test group_gids : %s", jsonData2)
}
