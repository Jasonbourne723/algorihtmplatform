package file

import (
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取算子文件目录结构
// @Schemes
// @Description
// @Tags File
// @Param operatorId query int64 true "算子Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.FileDto]
// @Router /File/GetFilesByOperatorId [get]
func GetFilesByOperatorId(c *gin.Context) {
	operatorId, err := strconv.ParseInt(c.Query("operatorId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	fileDtos, err := services.FileService.GetFilesByOperatorId(operatorId)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, fileDtos)
}

// @Summary 获取算法文件目录结构
// @Schemes
// @Description
// @Tags File
// @Param algorithmId query int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.FileDto]
// @Router /File/GetFilesByAlgorithmId [get]
func GetFilesByAlgorithmId(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Query("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	fileDtos, err := services.FileService.GetFilesByAlgorithmId(algorithmId)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, fileDtos)
}

// @Summary 获取文件文本内容
// @Schemes
// @Description
// @Tags File
// @Param path query string true "文件路径"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[string]
// @Router /File/GetFileContent [get]
func GetFileContent(c *gin.Context) {
	path := c.Query("path")
	bytes, err := os.ReadFile(path)
	if err != nil && err != io.EOF {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, string(bytes))
}

// @Summary 清空缓存
// @Schemes
// @Description
// @Tags File
// @Param algorithmId path int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /File/CleanCache [get]
func CleanCache(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Param("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	services.FileService.CleanCache(algorithmId)
	response.Success(c)
}
