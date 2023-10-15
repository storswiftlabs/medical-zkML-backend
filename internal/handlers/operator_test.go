package handlers

import (
	"fmt"
	"medical-zkml-backend/pkg/config"
	"os"
	"testing"
)

func TestPredictionProcess(t *testing.T) {
	//
	//conf := config.NewConfig()
	//db.InitMysql(conf)
	//modules := base.GetModuleList(conf)
	//operator := Operator{
	//	Modules: modules,
	//}
	//user := "0x111"
	//disease := "Acute_Inflammations"
	//module := "kMeans"
	//params := []string{"36.4", "1", "1", "1", "1", "1"}
	//if err := operator.DiseasePrediction(user, disease, "", module, params); err != nil {
	//	t.Fatalf(err.Error())
	//	return
	//}
	//time.Sleep(50 * time.Second)
}

func TestReadJson(t *testing.T) {
	data, err := os.ReadFile("../plugin/abi/decision_tree/Lymphography.json")
	if err != nil {
		t.Errorf(err.Error())
	}
	content := string(data)
	t.Log(content)
}

func TestFunc(t *testing.T) {
	disease := "Lymphography"
	module := "decision_tree"
	config.NewConfig()
	t.Log(config.GetConfig().Get("contract"))
	t.Log(config.GetConfig().Sub("contract"))
	t.Log(config.GetConfig().Sub("contract").Sub(fmt.Sprintf("%s+%s", module, disease)).Get("address"))
}
