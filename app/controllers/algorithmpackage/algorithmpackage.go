package algorithmpackage

import (
	"algorithmplatform/app/common/dto"
	_ "algorithmplatform/app/common/dto"
	"algorithmplatform/app/common/response"
	_ "algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 算法包分页查询
// @Schemes
// @Description
// @Tags AlgorithmPackage
// @Param pageIndex query int64 true "pageIndex"
// @Param pageSize query int64 true "pageSize"
// @Param name query string false "name"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.PageData[dto.AlgorithmPackageDto]]
// @Router /AlgorithmPackage/Page [get]
func Page(c *gin.Context) {
	pageIndex, err := strconv.ParseInt(c.Query("pageIndex"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	name := c.Query("name")

	pages, err := services.AlgorithmPackageService.Page(int(pageIndex), int(pageSize), name)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, pages)
}

// @Summary 算法包查询
// @Schemes
// @Description
// @Tags AlgorithmPackage
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.AlgorithmPackageDto]
// @Router /AlgorithmPackage [get]
func List(c *gin.Context) {

	results, err := services.AlgorithmPackageService.List()
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, results)
}

// @Summary 新增算法包
// @Schemes
// @Description
// @Tags AlgorithmPackage
// @Param default body dto.AddAlgorithmPackageDto true "算法包"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.AlgorithmPackageDto]
// @Router /AlgorithmPackage [post]
func Add(c *gin.Context) {
	var addAlgorithmPackageDto dto.AddAlgorithmPackageDto
	if err := c.ShouldBindJSON(&addAlgorithmPackageDto); err != nil {
		response.FailError(c, err)
		return
	}
	if err := services.AlgorithmPackageService.Add(&addAlgorithmPackageDto); err != nil {
		response.FailError(c, err)
		return
	}
	response.Success(c)
}

// @Summary 下载算法包
// @Schemes
// @Description
// @Tags AlgorithmPackage
// @Param algorithmPackageId path int64 true "包Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[string]
// @Router /AlgorithmPackage/Down/{algorithmPackageId} [get]
func Down(c *gin.Context) {
	algorithmPackageId, err := strconv.ParseInt(c.Param("algorithmPackageId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}

	filePath, err := services.AlgorithmPackageService.DownLoad(algorithmPackageId)
	if err != nil {
		response.FailError(c, err)
		return
	}

	response.SuccessWithFile(c, filePath)
}
