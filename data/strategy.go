package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Strategy struct {
	Steps         map[string]bool `json:"steps"`
	RecallReasons []string        `json:"recall_reasons"`
}

func LoadStrategy(rc *RequestContext) bool {
	fileName := "dsl/strategy/" + rc.Strategy + ".json"
	// 打开 JSON 文件
	file, err := os.Open(fileName)
	if err != nil {
		logrus.Fatalf("cann't open %s ：%s", fileName, err)
		return false
	}
	defer file.Close()
	// 读取 JSON 文件内容
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cann't read %s ：%s", fileName, err)
		return false
	}
	// 解析 JSON 数据
	var strategy Strategy
	err = json.Unmarshal(bytes, &strategy)
	if err != nil {
		log.Fatalf("cann't unmarshal ：%s", err)
		return false
	}
	rc.Steps = strategy.Steps
	rc.StepsRule = make(map[string]string)
	rc.StepNum = 0
	for stepName, v := range rc.Steps {
		if v {
			rc.StepNum++
			fileName := "dsl/params/" + rc.Strategy + "/" + stepName + ".dsl"
			stepFile, err := os.Open(fileName)
			if err != nil {
				logrus.Fatalf("cann't open step file %s ：%s", fileName, err)
				return false
			}
			bytes, err := ioutil.ReadAll(stepFile)
			rc.StepsRule[stepName] = string(bytes)
			stepFile.Close()
		}
	}
	rc.RecallReasons = make(map[string][]Group)
	val, _ := json.Marshal(strategy.RecallReasons)
	logrus.Infof("lmx_test : %s", string(val))
	for _, reason := range strategy.RecallReasons {
		rc.RecallReasons[reason] = make([]Group, 0)
	}
	rc.RecallReasonNames = strategy.RecallReasons
	return true
}
