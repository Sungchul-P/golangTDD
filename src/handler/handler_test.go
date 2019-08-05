package handler

import (
	"net/http"
	"testing"

	"io/ioutil"
	"net/http/httptest"
	"os"

	"github.com/gorilla/handlers"
	"github.com/stretchr/testify/assert"
)

var (
	server  *httptest.Server
	testURL string
)

type Response struct {
	Content string
	Code    int
}

func Test_Init(t *testing.T) {
	logfile, err := os.OpenFile("C:/Temp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// logfile, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	assert.Nil(t, err, "")
	h := Handler{}
	h.Init()

	// 웹 서버 실행
	server = httptest.NewServer(handlers.CombinedLoggingHandler(logfile, http.DefaultServeMux))
	testURL = server.URL // 웹 서버 접근 URL
}

// Ping API 테스트.
func Test_Ping(t *testing.T) {
	res, err := DoGet(testURL + "/ping")
	assert.Nil(t, err, "")
	assert.Equal(t, 200, res.Code, "PING API")           // 200 OK
	assert.Equal(t, "pong", res.Content, "PONG Message") // "pong"
}

func DoGet(url string) (*Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

// 존재하지 않는 페이지를 요청 할 경우 404 Page Not Found를 반환해야 한다. 이를 테스트하는 코드.
func Test_APINotFound(t *testing.T) {
	res, err := DoGet(testURL + "/myfunc")
	assert.Nil(t, err, "")
	assert.Equal(t, 404, res.Code, "Unknown API")
}

func Test_Div(t *testing.T) {
	res, err := DoGet(testURL + "/div/4/2")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "2", res.Content)

	// res, err = DoGet(testURL + "/div/a/4")
	// assert.Nil(t, err, "")
	// assert.Equal(t, http.StatusInternalServerError, res.Code, "Invalide argument")

	res, err = DoGet(testURL + "/div/4/a")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusInternalServerError, res.Code, "Invalide argument")

	res, err = DoGet(testURL + "/div/0/4")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusNotAcceptable, res.Code, "Invalide argument")

	res, err = DoGet(testURL + "/div/4/0")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusNotAcceptable, res.Code, "Invalide argument")
}
