// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// GetApiDiseasesDiseaseInfoParams defines parameters for GetApiDiseasesDiseaseInfo.
type GetApiDiseasesDiseaseInfoParams struct {
	Disease string `form:"Disease" json:"Disease"`
}

// PostApiOperatorDeletePredictedRecordJSONBody defines parameters for PostApiOperatorDeletePredictedRecord.
type PostApiOperatorDeletePredictedRecordJSONBody struct {
	Ids  []int  `json:"ids"`
	User string `json:"user"`
}

// PostApiOperatorArticleJSONBody defines parameters for PostApiOperatorArticle.
type PostApiOperatorArticleJSONBody struct {
	Diseases []string `json:"diseases"`
	User     string   `json:"user"`
}

// PostApiOperatorDiseasePredictionJSONBody defines parameters for PostApiOperatorDiseasePrediction.
type PostApiOperatorDiseasePredictionJSONBody struct {
	Inputs []struct {
		Index  *int    `json:"index,omitempty"`
		Name   *string `json:"name,omitempty"`
		Select *string `json:"select,omitempty"`
	} `json:"inputs"`
	Module string `json:"module"`
	Name   string `json:"name"`
	User   string `json:"user"`
}

// PostApiOperatorPredictingOutcomesJSONBody defines parameters for PostApiOperatorPredictingOutcomes.
type PostApiOperatorPredictingOutcomesJSONBody struct {
	User string `json:"user"`
}

// PostApiOperatorRecommendJSONBody defines parameters for PostApiOperatorRecommend.
type PostApiOperatorRecommendJSONBody struct {
	Disease string `json:"disease"`
}

// PostApiOperatorVerifyPredictionResultsJSONBody defines parameters for PostApiOperatorVerifyPredictionResults.
type PostApiOperatorVerifyPredictionResultsJSONBody struct {
	Id   int    `json:"id"`
	User string `json:"user"`
}

// PostApiUserArticleCollectionJSONBody defines parameters for PostApiUserArticleCollection.
type PostApiUserArticleCollectionJSONBody struct {
	Url  string `json:"url"`
	User string `json:"user"`
}

// PostApiUserArticleCollectionCheckJSONBody defines parameters for PostApiUserArticleCollectionCheck.
type PostApiUserArticleCollectionCheckJSONBody struct {
	User string `json:"user"`
}

// PostApiUserGetUserJSONBody defines parameters for PostApiUserGetUser.
type PostApiUserGetUserJSONBody struct {
	User string `json:"user"`
}

// PostApiUserRegisterJSONBody defines parameters for PostApiUserRegister.
type PostApiUserRegisterJSONBody struct {
	Address string `json:"address"`
}

// PostApiOperatorDeletePredictedRecordJSONRequestBody defines body for PostApiOperatorDeletePredictedRecord for application/json ContentType.
type PostApiOperatorDeletePredictedRecordJSONRequestBody PostApiOperatorDeletePredictedRecordJSONBody

// PostApiOperatorArticleJSONRequestBody defines body for PostApiOperatorArticle for application/json ContentType.
type PostApiOperatorArticleJSONRequestBody PostApiOperatorArticleJSONBody

// PostApiOperatorDiseasePredictionJSONRequestBody defines body for PostApiOperatorDiseasePrediction for application/json ContentType.
type PostApiOperatorDiseasePredictionJSONRequestBody PostApiOperatorDiseasePredictionJSONBody

// PostApiOperatorPredictingOutcomesJSONRequestBody defines body for PostApiOperatorPredictingOutcomes for application/json ContentType.
type PostApiOperatorPredictingOutcomesJSONRequestBody PostApiOperatorPredictingOutcomesJSONBody

// PostApiOperatorRecommendJSONRequestBody defines body for PostApiOperatorRecommend for application/json ContentType.
type PostApiOperatorRecommendJSONRequestBody PostApiOperatorRecommendJSONBody

// PostApiOperatorVerifyPredictionResultsJSONRequestBody defines body for PostApiOperatorVerifyPredictionResults for application/json ContentType.
type PostApiOperatorVerifyPredictionResultsJSONRequestBody PostApiOperatorVerifyPredictionResultsJSONBody

// PostApiUserArticleCollectionJSONRequestBody defines body for PostApiUserArticleCollection for application/json ContentType.
type PostApiUserArticleCollectionJSONRequestBody PostApiUserArticleCollectionJSONBody

// PostApiUserArticleCollectionCheckJSONRequestBody defines body for PostApiUserArticleCollectionCheck for application/json ContentType.
type PostApiUserArticleCollectionCheckJSONRequestBody PostApiUserArticleCollectionCheckJSONBody

// PostApiUserGetUserJSONRequestBody defines body for PostApiUserGetUser for application/json ContentType.
type PostApiUserGetUserJSONRequestBody PostApiUserGetUserJSONBody

// PostApiUserRegisterJSONRequestBody defines body for PostApiUserRegister for application/json ContentType.
type PostApiUserRegisterJSONRequestBody PostApiUserRegisterJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 疾病信息
	// (GET /api/diseases/disease_info)
	GetApiDiseasesDiseaseInfo(c *gin.Context, params GetApiDiseasesDiseaseInfoParams)
	// 疾病种类
	// (GET /api/diseases/diseases)
	GetApiDiseasesDiseases(c *gin.Context)
	// 删除预测记录
	// (POST /api/operator/DeletePredictedRecord)
	PostApiOperatorDeletePredictedRecord(c *gin.Context)
	// 文章
	// (POST /api/operator/article)
	PostApiOperatorArticle(c *gin.Context)
	// 疾病预测
	// (POST /api/operator/disease_prediction)
	PostApiOperatorDiseasePrediction(c *gin.Context)
	// 算子列表
	// (GET /api/operator/operator_list)
	GetApiOperatorOperatorList(c *gin.Context)
	// 预测结果
	// (POST /api/operator/predicting_outcomes)
	PostApiOperatorPredictingOutcomes(c *gin.Context)
	// 疾病建议
	// (POST /api/operator/recommend)
	PostApiOperatorRecommend(c *gin.Context)
	// 验证预测结果
	// (POST /api/operator/verify_prediction_results)
	PostApiOperatorVerifyPredictionResults(c *gin.Context)
	// 文章收藏
	// (POST /api/user/article_collection)
	PostApiUserArticleCollection(c *gin.Context)
	// 文章收藏检查
	// (POST /api/user/article_collection_check)
	PostApiUserArticleCollectionCheck(c *gin.Context)
	// 查询用户是否已注册
	// (POST /api/user/get_user)
	PostApiUserGetUser(c *gin.Context)
	// 用户注册
	// (POST /api/user/register)
	PostApiUserRegister(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetApiDiseasesDiseaseInfo operation middleware
func (siw *ServerInterfaceWrapper) GetApiDiseasesDiseaseInfo(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiDiseasesDiseaseInfoParams

	// ------------- Required query parameter "Disease" -------------
	if paramValue := c.Query("Disease"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument Disease is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "Disease", c.Request.URL.Query(), &params.Disease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter Disease: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiDiseasesDiseaseInfo(c, params)
}

// GetApiDiseasesDiseases operation middleware
func (siw *ServerInterfaceWrapper) GetApiDiseasesDiseases(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiDiseasesDiseases(c)
}

// PostApiOperatorDeletePredictedRecord operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorDeletePredictedRecord(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorDeletePredictedRecord(c)
}

// PostApiOperatorArticle operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorArticle(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorArticle(c)
}

// PostApiOperatorDiseasePrediction operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorDiseasePrediction(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorDiseasePrediction(c)
}

// GetApiOperatorOperatorList operation middleware
func (siw *ServerInterfaceWrapper) GetApiOperatorOperatorList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiOperatorOperatorList(c)
}

// PostApiOperatorPredictingOutcomes operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorPredictingOutcomes(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorPredictingOutcomes(c)
}

// PostApiOperatorRecommend operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorRecommend(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorRecommend(c)
}

// PostApiOperatorVerifyPredictionResults operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorVerifyPredictionResults(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorVerifyPredictionResults(c)
}

// PostApiUserArticleCollection operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserArticleCollection(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserArticleCollection(c)
}

// PostApiUserArticleCollectionCheck operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserArticleCollectionCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserArticleCollectionCheck(c)
}

// PostApiUserGetUser operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserGetUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserGetUser(c)
}

// PostApiUserRegister operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserRegister(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/api/diseases/disease_info", wrapper.GetApiDiseasesDiseaseInfo)

	router.GET(options.BaseURL+"/api/diseases/diseases", wrapper.GetApiDiseasesDiseases)

	router.POST(options.BaseURL+"/api/operator/DeletePredictedRecord", wrapper.PostApiOperatorDeletePredictedRecord)

	router.POST(options.BaseURL+"/api/operator/article", wrapper.PostApiOperatorArticle)

	router.POST(options.BaseURL+"/api/operator/disease_prediction", wrapper.PostApiOperatorDiseasePrediction)

	router.GET(options.BaseURL+"/api/operator/operator_list", wrapper.GetApiOperatorOperatorList)

	router.POST(options.BaseURL+"/api/operator/predicting_outcomes", wrapper.PostApiOperatorPredictingOutcomes)

	router.POST(options.BaseURL+"/api/operator/recommend", wrapper.PostApiOperatorRecommend)

	router.POST(options.BaseURL+"/api/operator/verify_prediction_results", wrapper.PostApiOperatorVerifyPredictionResults)

	router.POST(options.BaseURL+"/api/user/article_collection", wrapper.PostApiUserArticleCollection)

	router.POST(options.BaseURL+"/api/user/article_collection_check", wrapper.PostApiUserArticleCollectionCheck)

	router.POST(options.BaseURL+"/api/user/get_user", wrapper.PostApiUserGetUser)

	router.POST(options.BaseURL+"/api/user/register", wrapper.PostApiUserRegister)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"JQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
