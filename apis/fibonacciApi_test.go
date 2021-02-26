package apis

import (
	"bytes"
	"fibunacci/daos"
	"fibunacci/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type apiTestCase struct {
	tag              string
	method           string
	urlToServe       string
	urlToHit         string
	body             string
	function         gin.HandlerFunc
	status           int
	responseFilePath string
	expectedResult   string
}

func init() {
	util.LoadEnvFromFileForTests()
	os.Setenv("POSTGRES_HOST", "localhost")
}

// Creates new router in testing mode
func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

// Used to run single API test case. It makes HTTP request and returns its response
func testAPI(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.urlToServe, test.urlToHit, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if http.StatusOK == res.Code {
			s1 := strings.Split(res.Body.String(), ":")[0]
			assert.Equal(t, test.expectedResult, s1, test.tag)
		}
	}
}

func init() {
	//util.LoadEnvFromFileForTests()
	daos.NewSqlDao("../config/db.yaml")
}

func TestGetFibonacciSeries(t *testing.T) {
	runAPITests(t, []apiTestCase{
		{"t1 - get a Fibonacci series", "GET", "/fseries/:count", "/fseries/10", "", GetFibonacciSeries, http.StatusOK, "", "\" [0 1 1 2 3 5 8 13 21 34] "},
		{"t2 - get a Invalid count", "GET", "/fseries/:count", "/fseries/-1", "", GetFibonacciSeries, http.StatusBadRequest, "", ""},
		{"t1 - get a Fibonacci number", "GET", "/fnumber/:ordinal", "/fnumber/11", "", GetFibonacciNumberForOrdinal, http.StatusOK, "", "\" 89 "},
		{"t1 - get a Fibonacci number with invalid ordinal", "GET", "/fnumber/:ordinal", "/fnumber/-1", "", GetFibonacciNumberForOrdinal, http.StatusBadRequest, "", ""},
	})
}
