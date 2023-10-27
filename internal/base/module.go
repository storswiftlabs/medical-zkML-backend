package base

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"medical-zkml-backend/internal/module"
	"reflect"
	"strings"
)

// GetDiseaseInfo Basic data loading
func GetDiseaseInfo(conf *viper.Viper) ([]module.Disease, module.DiseaseList) {
	diseaseList := make(module.DiseaseList)
	var diseases []module.Disease
	diseaseConf := conf.Sub("disease")
	if diseaseConf == nil {
		zap.L().Panic("Failed to obtain disease configuration")
	}

	list := sublist(diseaseConf)
	if list == nil {
		zap.L().Panic("Failed to read sublist: ", zap.Error(errors.New("sublist == nil")))
	}

	for _, disease := range list {
		diseases = append(diseases, module.Disease{
			Name:        disease,
			Description: interface2string(diseaseConf.Sub(disease).Get("description")),
		})
	}

	for _, key := range list {
		info := &module.DiseaseInfo{
			Name: key,
		}
		disease := diseaseConf.Sub(key)
		inject(disease, info)
		diseaseList[key] = info
	}
	zap.L().Info("Disease list, disease information loading completed")
	return diseases, diseaseList
}

func inAnArray(ele string, arr []string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func inject(disease *viper.Viper, info *module.DiseaseInfo) {

	info.Description = interface2string(disease.Get("description"))
	inputs := disease.Get("inputs").([]any)
	for _, elem := range inputs {
		param := elem.(map[string]any)
		info.Inputs = append(info.Inputs, module.Input{
			Name:        interface2string(param["name"]),
			Description: interface2string(param["description"]),
			Index:       param["index"].(int),
			Regular:     interface2string(param["regular"]),
			Warn:        interface2string(param["warn"]),
			InputMethod: interface2string(param["input method"]),
			Select:      interface2keyvalue(param["select"]),
		})
	}

	output := disease.Get("output").(map[string]any)
	info.Output.Description = interface2string(output["description"])
	info.Output.Result = interface2keyvalue(output["result"])
}

// sublist Obtain a secondary list
func sublist(meet *viper.Viper) (list []string) {
	allKeys := meet.AllKeys()

	if len(allKeys) == 0 {
		return nil
	}
	for _, key := range allKeys {
		disease := cases.Title(language.Dutch).String(strings.Split(key, ".")[0])
		if len(list) == 0 {
			list = append(list, disease)
			continue
		}
		if inAnArray(disease, list) {
			continue
		}
		list = append(list, disease)
	}

	zap.L().Info("list := sublist(diseaseConf): ", zap.Any("list", list))
	return
}

func interface2string(v any) string {
	if v == nil {
		return ""
	} else {
		return v.(string)
	}
}

func interface2keyvalue(v any) []module.KeyValue {
	if v == nil {
		panic("The selection option cannot be empty")
	}

	var keyvalues []module.KeyValue
	for _, kv := range v.([]any) {
		mapType := reflect.TypeOf(kv)
		if mapType.Kind() != reflect.Map {
			panic("The selection option must be a map")
		}
		keyType := mapType.Key()
		valueType := mapType.Elem()

		mapValue := reflect.ValueOf(kv)
		keys := mapValue.MapKeys()
		for _, key := range keys {
			keyvalues = append(keyvalues, module.KeyValue{
				Key:   key.Convert(keyType).Interface(),
				Value: mapValue.MapIndex(key).Convert(valueType).Interface(),
			})
		}

	}
	return keyvalues
}

func GetModuleList(conf *viper.Viper) []module.Module {
	var modules []module.Module
	moduleConf := conf.Get("module").([]any)
	fmt.Println(moduleConf)
	if len(moduleConf) == 0 {
		panic("Module list cannot be empty")
	}

	for _, value := range moduleConf {
		temp := value.(map[string]any)
		modules = append(modules, module.Module{
			Name:        interface2string(temp["name"]),
			Description: interface2string(temp["description"]),
		})
	}
	return modules
}
