package handlers

import (
	"fmt"
	"medical-zkml-backend/internal/base"
	"medical-zkml-backend/pkg/config"
	"testing"
)

func TestNewDisease(t *testing.T) {
	conf := config.NewConfig()

	disease := NewDisease(base.GetDiseaseInfo(conf))
	fmt.Println(disease.GetDiseaseNameList())
}
