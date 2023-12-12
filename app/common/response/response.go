package response

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32     `json:"statusCode"`
	Succeeded  bool      `json:"succeeded"`
	Errors     string    `json:"errors"`
	Timestamp  time.Time `json:"timestamp"`
}

type TResponse[T any] struct {
	Response
	Data T `json:"data"`
}

func SuccessWithData[T any](c *gin.Context, data T) {
	c.JSON(200, TResponse[T]{
		Response: Response{
			StatusCode: 200,
			Succeeded:  true,
			Errors:     "",
			Timestamp:  time.Now(),
		},
		Data: data,
	})
}

func Success(c *gin.Context) {
	c.JSON(200, Response{
		StatusCode: 200,
		Succeeded:  true,
		Errors:     "",
		Timestamp:  time.Now(),
	})
}

func FailError(c *gin.Context, err error) {
	var errString = err.Error()
	c.JSON(200, Response{
		StatusCode: 500,
		Succeeded:  false,
		Errors:     errString,
		Timestamp:  time.Now(),
	})
}

func FailWithMsg(c *gin.Context, msg string) {
	c.JSON(200, Response{
		StatusCode: 500,
		Succeeded:  false,
		Errors:     msg,
		Timestamp:  time.Now(),
	})
}

func SuccessWithFile(c *gin.Context, filePath string) {

	fileName := getFileName(filePath)
	c.Status(200)
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	fileName = url.QueryEscape(fileName)
	c.Header("Content-Disposition", "attachment; filename*=utf-8''"+fileName)
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)
}

func getFileName(path string) string {
	if path[len(path)-1] == '/' {
		return ""
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
