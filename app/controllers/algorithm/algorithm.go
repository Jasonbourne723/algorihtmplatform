package algorithm

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 查询算法列表
// @Schemes
// @Description
// @Tags Algorithm
// @Param projectId query int64 true "项目Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.AlgorithmDto]
// @Router /Algorithm/GetByProjectId [get]
func GetByProjectId(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	algorithms, err := services.AlgorithmService.List(projectId)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, algorithms)
}

// @Summary 查询单个算法
// @Schemes
// @Description
// @Tags Algorithm
// @Param algorithmId path int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.AlgorithmDto]
// @Router /Algorithm/{algorithmId} [get]
func Get(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Param("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	algorithmDto, err := services.AlgorithmService.GetOne(algorithmId)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, algorithmDto)
}

// @Summary 新增算法
// @Schemes
// @Description
// @Tags Algorithm
// @Param default body dto.AddAlgorithmDto true "算法"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.AlgorithmDto]
// @Router /Algorithm [post]
func Post(c *gin.Context) {
	var addAlgorithmDto dto.AddAlgorithmDto
	if err := c.ShouldBindJSON(&addAlgorithmDto); err != nil {
		response.FailError(c, err)
		return
	}
	algorithmDto, err := services.AlgorithmService.Add(&addAlgorithmDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, algorithmDto)
}

// @Summary 更新算法
// @Schemes
// @Description
// @Tags Algorithm
// @Param default body dto.UpdateAlgorithmDto true "算法"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.AlgorithmDto]
// @Router /Algorithm [put]
func Put(c *gin.Context) {
	var updateAlgorithmDto dto.UpdateAlgorithmDto
	if err := c.ShouldBindJSON(&updateAlgorithmDto); err != nil {
		response.FailError(c, err)
		return
	}
	algorithmDto, err := services.AlgorithmService.Update(&updateAlgorithmDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, algorithmDto)
}

// @Summary 删除算法
// @Schemes
// @Description
// @Tags Algorithm
// @Param algorithmId path int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Algorithm/{algorithmId} [delete]
func Delete(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Param("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	services.AlgorithmService.Delete(algorithmId)
	response.Success(c)
}

// @Summary 运行算法
// @Schemes
// @Description
// @Tags Algorithm
// @Param algorithmId path int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Algorithm/run/{algorithmId} [get]
func Run(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Param("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	userId := c.GetInt64("userId")
	services.AlgorithmService.Run(algorithmId, userId)
	response.Success(c)
}

// @Summary 更新算法文件
// @Schemes
// @Description
// @Tags Algorithm
// @Param default body dto.UpdateAlgorithmFileDto true "算法文件更新"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Algorithm/UpdateAlgorithmFile [post]
func UpdateAlgorithmFile(c *gin.Context) {
	var updateAlgorithmFileDto dto.UpdateAlgorithmFileDto
	err := c.ShouldBindJSON(&updateAlgorithmFileDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	file, err := os.OpenFile(updateAlgorithmFileDto.Path, os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		response.FailError(c, err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(updateAlgorithmFileDto.Content))
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.Success(c)
}

// @Summary 获取数据源配置
// @Schemes
// @Description
// @Tags Algorithm
// @Param algorithmId path int64 true "算法Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response[dto.DataSourceDto]
// @Router /Algorithm/DataSourceConfig/{algorithmId} [get]
func GetDataSourceConfig(c *gin.Context) {
	algorithmId, err := strconv.ParseInt(c.Param("algorithmId"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	dataSourceDto, err := services.AlgorithmService.GetDataSource(algorithmId)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, dataSourceDto)
}

// @Summary 设置数据源配置
// @Schemes
// @Description
// @Tags Algorithm
// @Param default body dto.DataSourceDto true "数据源"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Algorithm/DataSourceConfig [post]
func SetDataSourceConfig(c *gin.Context) {
	var dataSourceDto dto.DataSourceDto
	err := c.ShouldBindJSON(&dataSourceDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	services.AlgorithmService.SetDataSource(dataSourceDto)
	response.Success(c)
}
