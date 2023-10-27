package module

import (
	"gorm.io/gorm"
)

type DiseaseInfo struct {
	Name        string
	Description string
	Inputs      []Input
	Output      struct {
		Description string
		Result      []KeyValue
	}
}

type Input struct {
	Name        string
	Description string
	Index       int
	Regular     string
	Warn        string
	InputMethod string
	Select      []KeyValue
}

type KeyValue struct {
	Key   any
	Value any
}

type Disease struct {
	Name        string
	Description string
}

type DiseaseList map[string]*DiseaseInfo

type Module struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserMedicalInformation struct {
	Name   string `json:"name"`
	Inputs []any  `json:"inputs"`
}

type PredictionPremise struct {
	User   string `json:"user"`
	Name   string `json:"name"`
	Module string `json:"module"`
	Inputs []struct {
		Name        string `json:"name"`
		Index       int    `json:"index"`
		SelectKey   string `json:"select-key"`
		SelectValue string `json:"select-value"`
	} `json:"inputs"`
}

type DiseasePropose struct {
	gorm.Model
	Name string `json:"name" gorm:"notnull, index:idx_name"`
	Hash string `json:"hash" gorm:"notnull"`
}

type Quantized struct {
	Data      []string
	Scale     string
	ZeroPoint string
}

type RecordPrediction struct {
	gorm.Model
	Status     Status
	Message    string
	User       string
	Disease    string
	Module     string
	Inputs     string `gorm:"type:text"`
	Normalized string
	Quantized  string
	Scale      string
	ZeroPoint  string
	StartTime  int64 `gorm:"start_time"`
	EndTime    int64 `gorm:"end_time"`
	Output     int
}

type Status string

var (
	Start                Status = "start"
	NormalizationFailed  Status = "normalization failed"
	NormalizationSuccess Status = "normalization success"
	QuantizedFailed      Status = "quantization failed"
	QuantizedSuccess     Status = "quantization success"
	CreateSuccess        Status = "create project success"
	CreateFailed         Status = "create project failed"
	PredictionComplete   Status = "prediction complete"
	PredictionFailed     Status = "prediction failed"
	GetResultSuccess     Status = "get result success"
	GetResultFailed      Status = "get result failed"
	Complete             Status = "complete"
)

type GetPredictedResult struct {
	User string
	Id   uint
}

type PredictedResult struct {
	ID        uint
	Disease   string
	Module    string
	StartTime int64 `gorm:"start_time"`
	EndTime   int64 `gorm:"end_time"`
	Status    Status
	Inputs    string
	Message   string
	Output    string
}

type UserPredictionValidation struct {
	gorm.Model
	User               string `gorm:"user"`
	RecordPredictionID uint   `gorm:"record_prediction_id"`
	Result             string `gorm:"result"`
	Proof              string `gorm:"type:text"`
	Disease            string `json:"disease"`
	Module             string `json:"module"`
	IsVerified         bool   `json:"is_verified"`
}

type User struct {
	gorm.Model
	Name string `gorm:"name,not null"`
}

type VerifyReq struct {
	User string `json:"user"`
	ID   uint   `json:"id"`
}

type VerifyResultNoModel struct {
	ContractAddress  string `json:"contract_address"`
	ContractFunction string `json:"contract_function"`
	Proof            string `json:"proof"`
	Result           string `json:"result"`
	Disease          string `json:"disease"`
	Module           string `json:"module"`
	ABI              string `json:"abi"`
}

type Article struct {
	gorm.Model
	Disease string `gorm:"disease"`
	URL     string `gorm:"url"`
	Time    int64  `gorm:"time"`
	Icon    string `gorm:"type:text"`
	Title   string `gorm:"title"`
}

type ArticleResult struct {
	Disease string `json:"disease" gorm:"disease"`
	URL     string `json:"url" gorm:"url"`
	Time    int64  `json:"time" gorm:"time"`
	Icon    string `json:"icon" gorm:"type:text"`
	Title   string `json:"title" gorm:"title"`
}

type ArticleCollection struct {
	gorm.Model
	User string
	Url  string
}
