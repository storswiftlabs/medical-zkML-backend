package db

import (
	"medical-zkml-backend/pkg/config"
	"testing"
)

func init() {
	InitMysql(config.NewConfig())
}

//func Test_getDBInfoFromConfig(t *testing.T) {
//	dbInfo := getDBInfoFromConfig(config.GetConfig(), "mongodb")
//	fmt.Println(dbInfo)
//}

func TestInitMysql(t *testing.T) {
	conf := config.NewConfig()
	InitMysql(conf)

}

//func TestInsertpropose(t *testing.T) {
//	conf := config.NewConfig()
//	InitMysql(conf)
//	diseases, _ := base.GetDiseaseInfo(conf)
//	var proposes []string
//	//db.Engine.Raw("SELECT hash FROM tbl_disease_propose WHERE name = ?", disease)
//	for _, disease := range diseases {
//		result := engine.Model(&module.DiseasePropose{}).Where("name = ?", disease.Name).Pluck("hash", &proposes)
//		if result.error != nil {
//			fmt.Println("error:", result.Error)
//		}
//		fmt.Println(proposes)
//	}
//
//}
/*
func TestGetPredicedList(t *testing.T) {
	conf := config.NewConfig()
	InitMysql(conf)
	list, err := GetPredicedList("user")
	if err != nil {
		return
	}
	t.Log(list)
}
*/
func TestUserQuery(t *testing.T) {
	query, err := UserQuery("0x01f5BB073e893d334FF9b0e239939982c124AF97")
	if err != nil {
		t.Error(err)
	}
	t.Log(query)
}

func TestGetArticle(t *testing.T) {
	articles, err := GetArticle([]string{"Acute_Inflammations"})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(articles)
}
