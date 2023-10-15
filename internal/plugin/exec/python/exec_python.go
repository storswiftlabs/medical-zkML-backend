package python

import (
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/module"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var (
	python   string
	rootPath string
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(filename)

}

func init() {
	if runtime.GOOS == "windows" {
		python = "py.exe"
	} else {
		python = "python3"
	}
}

// QuantitativeData Calling local Python to perform data quantization
func QuantitativeData(disease string, data []string) (*module.Quantized, error) {
	cmd := exec.Command(python, append([]string{fmt.Sprintf("%s/quantification_data.py", rootPath)}, append([]string{disease}, data...)...)...)
	// Execute commands and capture output
	output, err := cmd.Output()
	if err != nil {
		zap.L().Error("Data processing: quantification data execution failed", zap.Error(err))
		return &module.Quantized{}, err
	}

	quantized := new(module.Quantized)

	lines := strings.Split(string(output), "\n")

	quantized.Data = string2stringList(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	quantized.Scale = strings.TrimSpace(strings.Split(lines[1], ":")[1])
	quantized.ZeroPoint = strings.TrimSpace(strings.Split(lines[2], ":")[1])
	zap.L().Info("Data processing: quantification data execution success")
	return quantized, nil
}

func NormalizedData(disease string, data []string) ([]string, error) {
	cmd := exec.Command(python, append([]string{fmt.Sprintf("%s/normalized_data.py", rootPath)}, append([]string{disease}, data...)...)...)
	output, err := cmd.Output()
	if err != nil {
		zap.L().Error("Data processing: normalization data execution failed", zap.Error(err))
		return nil, err
	}
	zap.L().Info("Data processing: normalization data execution success")
	return string2stringList(strings.TrimSpace(string(output))), nil
}

func string2stringList(data string) []string {
	var result []string
	data = strings.Trim(data, "[]")
	strList := strings.Split(data, ", ")
	for _, v := range strList {
		result = append(result, v)
	}
	return result
}
