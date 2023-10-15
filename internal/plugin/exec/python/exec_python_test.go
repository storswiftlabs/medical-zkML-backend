package python

import (
	"fmt"
	"testing"
)

func TestQuantized(t *testing.T) {
	disease := "Acute_Inflammations"
	data := []string{"0.0", "0.0", "1.0", "0.0", "0.0", "0.0"}
	//fmt.Println(append([]string{"123"}, data...))
	quantized, err := QuantitativeData(disease, data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(quantized)
}

func TestNormalizedData(t *testing.T) {
	disease := "Acute_Inflammations"
	data := []string{"35.5", "0", "1", "0", "0"}
	fmt.Println(data)
	normailzed, err := NormalizedData(disease, data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(normailzed)
	//quantized, err := QuantitativeData(normailzed)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	//fmt.Println(quantized)
}

func TestRootPath(t *testing.T) {
	fmt.Println(rootPath)
}
