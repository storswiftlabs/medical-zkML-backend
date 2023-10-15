package handlers

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/module"
	"sort"
	"strconv"
)

type Disease struct {
	diseaseList []module.Disease
	diseaseInfo module.DiseaseList
}

func NewDisease(diseaseList []module.Disease, diseaseInfo module.DiseaseList) Disease {
	return Disease{
		diseaseList: diseaseList,
		diseaseInfo: diseaseInfo,
	}
}

func (d *Disease) GetDiseaseNameList() []string {
	var names []string
	list := d.diseaseList
	for _, disease := range list {
		names = append(names, disease.Name)
	}
	return names
}

func (d *Disease) elemInArray(elem string) bool {
	fmt.Println(d.diseaseList)
	list := d.diseaseList
	for _, disease := range list {
		if disease.Name == elem {
			return true
		}
	}
	return false
}

func (d *Disease) CountInputsForName(name string) int {
	return len(d.diseaseInfo[name].Inputs)
}

func (d *Disease) GetDescriptionByName(disease string) string {
	list := d.diseaseList
	for _, v := range list {
		if disease == v.Name {
			return v.Description
		}
	}
	return ""
}

func (d *Disease) GetOutputValue(disease, key string) (string, error) {
	output := d.diseaseInfo[disease].Output.Result
	for _, kv := range output {
		if kv.Key == key {
			return kv.Value.(string), nil
		}
	}
	zap.L().Error(fmt.Sprintf("%s disease not found the key %s", disease, key))
	return "", errors.New("disease mapping search failed")
}

func (d *Disease) ParameterCheck(info *module.PredictionPremise) error {
	inputs := d.diseaseInfo[info.Name].Inputs

	for index, input := range inputs {
		scope := len(input.Select)
		if scope == 1 {
			// If the scope is equal to 1, it means manual input, otherwise it means selection
			// TODO: Perform more detailed verification for manual input
			continue
		}
		//If the range is not equal to 1, it indicates selection, and the input parameters must be within the range

		sort.Slice(input.Select, func(i, j int) bool {
			return input.Select[i].Value.(int) < input.Select[j].Value.(int)
		})

		choose, err := strconv.Atoi(info.Inputs[index].Select)
		if err != nil {
			return err
		}

		if !(input.Select[0].Value.(int) <= choose && input.Select[len(input.Select)-1].Value.(int) >= choose) {
			return errors.New("the selected parameter is out of range")
		}
	}
	return nil
}
