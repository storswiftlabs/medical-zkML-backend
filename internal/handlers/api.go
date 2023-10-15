package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"medical-zkml-backend/api"
	"medical-zkml-backend/internal/module"
	"net/http"
	"sort"
)

type Medical struct {
	Disease
	Operator
	User
}

func (m Medical) GetApiDiseasesDiseaseInfo(c *gin.Context, params api.GetApiDiseasesDiseaseInfoParams) {
	if m.Disease.diseaseList == nil {
		returnError(http.StatusInternalServerError, "", errors.New("no disease list"), c)
		return
	}
	if !m.Disease.elemInArray(params.Disease) {
		message := "there are no such diseases in the database"
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	returnDataNoCount(http.StatusOK, "", m.Disease.diseaseInfo[params.Disease], c)
	return
}

func (m Medical) GetApiDiseasesDiseases(c *gin.Context) {
	count := int64(len(m.Disease.diseaseList))
	if count == 0 {
		returnError(http.StatusInternalServerError, "", errors.New("disease list loading failed"), c)
		return
	}
	returnData(http.StatusOK, count, "Query completed", m.Disease.diseaseList, c)
	return
}

func (m Medical) GetApiOperatorOperatorList(c *gin.Context) {
	count := int64(len(m.Operator.Modules))
	if count == 0 {
		returnError(http.StatusInternalServerError, "", errors.New("model list loading failed"), c)
		return
	}
	returnData(http.StatusOK, count, "ok", m.Operator.Modules, c)
	return
}

func (m Medical) PostApiOperatorDiseasePrediction(c *gin.Context) {
	var premise module.PredictionPremise

	// Obtain the input body parameter
	if err := c.BindJSON(&premise); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if premise.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(premise.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", premise.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	// Check the correctness of the data
	// Check if the disease name is in the disease list
	if !m.Disease.elemInArray(premise.Name) {
		returnError(http.StatusBadRequest, "", errors.New(fmt.Sprintf("this disease %s is not found in the lists of diseases", premise.Name)), c)
		return
	}

	// Check if the module is in the module list
	if !m.Operator.InModuleList(premise.Module) {
		returnError(http.StatusBadRequest, "", errors.New(fmt.Sprintf("this module %s is not found in the lists of modules", premise.Module)), c)
		return
	}

	// Check if the number of inputs parameters is equal
	if len(premise.Inputs) != m.Disease.CountInputsForName(premise.Name) {
		returnError(http.StatusBadRequest, fmt.Sprintf("%d%d", len(premise.Inputs), m.Disease.CountInputsForName(premise.Name)), errors.New("the number of input parameters entered does not match"), c)
		return
	}

	// Check if the index is continuous
	sort.Slice(premise.Inputs, func(i, j int) bool {
		return premise.Inputs[i].Index < premise.Inputs[j].Index
	})

	for index, input := range premise.Inputs {
		if index != input.Index {
			msg := "index error in input parameters"
			returnError(http.StatusBadRequest, msg, errors.New(msg), c)
			return
		}

	}

	// TODO: Check if the input parameters match
	if err := m.Disease.ParameterCheck(&premise); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	go m.Operator.DiseasePrediction(&premise)

	returnDataNoCount(http.StatusOK, "ok", nil, c)
	return
}

func (m Medical) PostApiOperatorPredictingOutcomes(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if req.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	fmt.Println("user: ", req.User)
	// Query database to obtain results
	list, err := m.Operator.GetPrediction(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	for i := 0; i < len(list); i++ {
		if list[i].Output == "-1" {
			continue
		}
		value, err := m.Disease.GetOutputValue(list[i].Disease, list[i].Output)
		if err != nil {
			returnError(http.StatusInternalServerError, err.Error(), err, c)
			return
		}
		list[i].Output = value
	}

	returnData(http.StatusOK, int64(len(list)), "ok", list, c)
	return
}

func (m Medical) PostApiOperatorDeletePredictedRecord(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		IDs  []int  `json:"ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if req.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	if err := m.Operator.DeletePredictedResults(req.User, req.IDs); err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "Record deleted successfully", nil, c)
	return
}

// PostApiOperatorRecommend TODO Accessing GPT for disease advice
func (m Medical) PostApiOperatorRecommend(c *gin.Context) {
	var req struct {
		Disease string `json:"disease"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	fmt.Println(req.Disease)
	hash, err := m.Operator.GetProposeByName(req.Disease)
	if err != nil {
		returnError(http.StatusInternalServerError, "", err, c)
		return
	}
	if hash == "" {
		returnError(http.StatusInternalServerError, "", errors.New("the suggestion list is empty"), c)
		return
	}
	returnDataNoCount(http.StatusOK, "ok", hash, c)
	return
}

func (m Medical) PostApiOperatorVerifyPredictionResults(c *gin.Context) {
	var req module.VerifyReq
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	information, err := m.Operator.VerifyInformation(&req)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "ok", information, c)
	return
}

func (m Medical) PostApiOperatorArticle(c *gin.Context) {
	var req struct {
		Diseases []string `json:"diseases"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}

	for _, disease := range req.Diseases {
		if !m.Disease.elemInArray(disease) {
			message := "there are no such diseases in the database"
			returnError(http.StatusBadRequest, message, errors.New(message), c)
			return
		}
	}

	articles, err := m.Operator.GetArticle(req.Diseases)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}
	returnData(http.StatusOK, int64(len(articles)), "ok", articles, c)
	return
}

func (m Medical) PostApiUserArticleCollection(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		Url  string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	collect, err := m.User.IsCollected(req.User, req.Url)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}
	if collect {
		if err := m.User.CancelArticleCollection(req.User, req.Url); err != nil {
			returnError(http.StatusInternalServerError, err.Error(), err, c)
			return
		}
		returnDataNoCount(http.StatusOK, "Successfully cancelled collecting articles", nil, c)
		return
	}

	if err := m.User.CollectArticles(req.User, req.Url); err != nil {
		message := fmt.Sprintf("Article collection failed")
		returnError(http.StatusInternalServerError, message, errors.New(fmt.Sprintf("%s: %s", message, err.Error())), c)
		return
	}
	returnDataNoCount(http.StatusOK, "Successfully collected articles", nil, c)
	return
}

func (m Medical) PostApiUserRegister(c *gin.Context) {
	var req struct {
		User string `json:"address"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	if req.User == "" {
		message := fmt.Sprintf("Address cannot be empty")
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	if isExists {
		returnDataNoCount(http.StatusBadRequest, "User already exists", nil, c)
		return
	}

	if err := m.User.UserRegistration(req.User); err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "login was successful", nil, c)
	return
}

func (m Medical) PostApiUserGetUser(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "ok", isExists, c)
	return
}

func (m Medical) PostApiUserArticleCollectionCheck(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}

	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	fmt.Println(req.User)
	list, err := m.User.FavoriteArticleList(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	returnData(http.StatusOK, int64(len(list)), "ok", list, c)
	return
}
