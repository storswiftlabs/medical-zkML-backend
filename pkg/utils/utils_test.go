package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestSwap(t *testing.T) {
	strList := []string{"0", "0", "255", "0", "0", "0"}
	fmt.Println(strList)
	str := StringList2String(strList)
	fmt.Println(str)
	fmt.Println(String2StringList(str))
}

func TestHexToDecimal(t *testing.T) {
	//str := "0x0000000000000000000000000000000000000000000000000000000000000002"
	//t.Log(HexToDecimal(str))
	t.Log(strings.ReplaceAll("Acute Inflammations", " ", "_"))
}

//func TestCoverInput(t *testing.T) {
//	t.Log(CoverInput())
//}
