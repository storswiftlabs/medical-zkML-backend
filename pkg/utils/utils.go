package utils

import (
	"encoding/json"
	"math/big"
	m "medical-zkml-backend/internal/module"
	"strings"
)

func Hexadecimal2Decimal(result string) int {
	bigInt := new(big.Int)
	bigInt.SetString(result, 0)

	return int(bigInt.Int64())
}

func StringList2String(list []string) string {
	return strings.Join(list, ", ")
}

func String2StringList(data string) []string {
	return strings.Split(data, ", ")
}

func CoverInput(premise *m.PredictionPremise) string {
	inputs := premise.Inputs

	type record struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var records []record
	//var coverInput []string
	for _, input := range inputs {
		records = append(records, record{
			Key:   input.Name,
			Value: input.SelectKey,
		})
	}
	jsonBytes, _ := json.Marshal(records)
	return string(jsonBytes)
}
