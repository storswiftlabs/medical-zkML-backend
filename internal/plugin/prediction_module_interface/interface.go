package prediction_module_interface

import (
	m "medical-zkml-backend/internal/module"
	"medical-zkml-backend/internal/plugin/exec/noir"
)

type DiseasePrediction interface {
	DiseasePrediction(user string, recordPrediction *m.RecordPrediction, disease, module string, quantized *m.Quantized)
}

var dp DiseasePrediction

func Register(name string) {
	switch name {
	case "Local Noir":
		dp = DiseasePrediction(noir.InitNoir("nargo", "internal/plugin/circuit"))
	}
}

func GetPredictionModule() DiseasePrediction {
	return dp
}
