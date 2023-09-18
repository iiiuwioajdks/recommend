package main

import (
	"encoding/json"
	"os"
	data "recommend/data"
	dag_execute "recommend/execute"
	recall_group_exec "recommend/steps/recall"
	user_data_exec "recommend/steps/user_data"
	mysql "recommend/tools/mysql"
	redis "recommend/tools/redis"
	"time"

	"github.com/sirupsen/logrus"
)

func print(str string) {
	println(str)
}

func main() {
	start := time.Now()
	// init
	redis.CreateRedisClient()
	mysql.CreateMysqlClient()

	rc := new(data.RequestContext)
	args := os.Args
	if len(args) == 2 {
		rc.Strategy = args[1]
	} else {
		rc.Strategy = "test"
	}
	// load json
	succ := data.LoadStrategy(rc)
	if !succ {
		logrus.Errorf("Load Strategy %s Error, Exit", rc.Strategy)
		return
	}

	execute := new(dag_execute.DAG_EXECUTE)
	run_step := 0
	execute.Init(rc.StepNum)
	if rc.Steps["user_data"] {
		user_data_func := map[string]interface{}{
			"rc":                    rc,
			"println":               print,
			"set_user":              user_data_exec.TestSetUser,
			"get_user_profile_ua":   user_data_exec.GetUserProfileUA,
			"get_user_profile_info": user_data_exec.GetUserProfileINFO,
		}
		execute.InitRuleBuilder(run_step, rc.StepsRule["user_data"], user_data_func)
		user_data_dag := [][]string{
			{},
			{"test_set_user"},
			{"get_user_profile_ua", "get_user_profile_info"},
		}
		execute.InitRuleDag(run_step, user_data_dag)
		run_step++
	}

	if rc.Steps["recall"] {
		recall_func := map[string]interface{}{
			"rc":         rc,
			"println":    print,
			"set_group":  recall_group_exec.TetsSetGroup,
			"recall_i2i": recall_group_exec.GetGroupI2I,
		}
		execute.InitRuleBuilder(run_step, rc.StepsRule["recall"], recall_func)
		recall_dag := [][]string{
			{},
			{"test_set_group"},
			{"recall_i2i"},
			{"merge"},
		}
		execute.InitRuleDag(run_step, recall_dag)
		run_step++
	}

	if rc.Steps["sort"] {
		sort_func := map[string]interface{}{
			"println": print,
		}
		execute.InitRuleBuilder(run_step, rc.StepsRule["sort"], sort_func)
		sort_dag := [][]string{
			{},
			{"predict"},
		}
		execute.InitRuleDag(run_step, sort_dag)
		run_step++
	}

	execute.ExecuteDag()

	jsonData, _ := json.Marshal(rc.UserProfile)
	logrus.Infof("lmx_test user_profile : %s", jsonData)

	jsonData2, _ := json.Marshal(rc.Groups)
	logrus.Infof("lmx_test group_gids : %s", jsonData2)
	logrus.Infof("ALL COST %d ms", time.Since(start).Milliseconds())
}
