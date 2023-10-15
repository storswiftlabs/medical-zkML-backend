package utils

import (
	"fmt"
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
	var coverInput []string
	for _, input := range inputs {
		coverInput = append(coverInput, fmt.Sprintf("%s:%s", input.Name, input.Select))
	}
	return StringList2String(coverInput)
}
