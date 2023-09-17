package dag_execute

import (
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

type DAG_EXECUTE struct {
	engines      []*engine.Gengine
	rule_builder []*builder.RuleBuilder
	rule_dag     [][][]string
}

type RuleAndDag struct {
	Rule     string
	Dag      [][]string
	Need_Dag bool
}

func (dag *DAG_EXECUTE) Init(step_num int) {
	dag.engines = make([]*engine.Gengine, step_num)
	for i := 0; i < step_num; i++ {
		dag.engines[i] = engine.NewGengine()
	}
	dag.rule_builder = make([]*builder.RuleBuilder, step_num)
	dag.rule_dag = make([][][]string, step_num)
}

func (dag *DAG_EXECUTE) InitRuleBuilder(engine_num int, rule string, funcs map[string]interface{}) {
	dataContext := context.NewDataContext()
	for key, val := range funcs {
		dataContext.Add(key, val)
	}
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(rule)
	if err != nil {
		panic(err)
	}
	dag.rule_builder[engine_num] = ruleBuilder
}

func (dag *DAG_EXECUTE) InitRuleDag(engine_num int, rule_dag [][]string) {
	dag.rule_dag[engine_num] = rule_dag
}

func (dag *DAG_EXECUTE) ExecuteDag() {
	for i, engine := range dag.engines {
		err := engine.ExecuteDAGModel(dag.rule_builder[i], dag.rule_dag[i])
		if err != nil {
			panic(err)
		}
	}
}
