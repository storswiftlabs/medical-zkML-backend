package noir

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

//var (
//	disease   = "Acute_Inflammations"
//	module    = "kMeans"
//	data      = []string{"0", "0", "255", "0", "0", "0"}
//	scale     = "255"
//	zeroPoint = "0"
//)

func init() {
	InitNoir("nargo", fmt.Sprintf("%s/internal/plugin/circuit", rootPath))
}

//func TestCreateProject(t *testing.T) {
//	project, err := Noir.CreateProject(1, disease, module)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//	t.Log(project)
//}
//
//func TestNoir(t *testing.T) {
//
//	project, err := Noir.CreateProject(1, disease, module)
//	if err != nil {
//		t.Errorf(err.Error())
//		return
//	}
//	t.Log(project)
//
//	if err := Noir.DiseasePrediction(project, disease, module, data, scale, zeroPoint); err != nil {
//		t.Errorf(err.Error())
//		return
//	}
//}

func TestRunner(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	root := path.Dir(path.Dir(path.Dir(path.Dir(path.Dir(filename)))))
	fmt.Printf(root)
}
