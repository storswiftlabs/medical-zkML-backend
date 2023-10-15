package config

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func init() {
	NewConfig()
}
func TestConfig_GetDatabase(t *testing.T) {

	// fmt.Println(conf.Get("database.mysql.user"))
	fmt.Println(conf.Get("database.mongodb"))
	conf = GetConfig()
	fmt.Println(conf.Get("database.mongodb"))
}

func TestType(t *testing.T) {
	var variable string
	variable = "1"
	typeOfVariable := reflect.TypeOf(variable)
	fmt.Println("Type:", typeOfVariable.String())
	fmt.Println("Type name:", typeOfVariable.Name())
	fmt.Println("Kind:", typeOfVariable.Kind().String())
}

func TestRootPath(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(filename)))
	fmt.Println(rootPath)
}

func TestConfig(t *testing.T) {
	t.Log(conf.Get("disease"))
}
