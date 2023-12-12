package configset

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取全部配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.ConfigSetDto]
// @Router /ConfigSet/All [get]
func GetAll(c *gin.Context) {
	if res, err := services.ConfigSetService.GetAll(); err != nil {
		response.FailError(c, err)
		return
	} else {
		response.SuccessWithData(c, res)
	}
}

// @Summary 获取分页配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Param pageIndex query int true "pageindex"
// @Param pageSize query int true "pageSize"
// @Param name query string false "name"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.PageData[dto.ConfigSetDto]]
// @Router /ConfigSet [get]
func List(c *gin.Context) {
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
	pageData, err := services.ConfigSetService.GetPage(int(pageIndex), int(pageSize), name)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, pageData)
}

// @Summary 新增配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Param default body dto.AddConfigSetDto true "配置集"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.ConfigSetDto]
// @Router /ConfigSet [post]
func Add(c *gin.Context) {
	var addConfigSetDto dto.AddConfigSetDto
	if err := c.ShouldBindJSON(&addConfigSetDto); err != nil {
		response.FailError(c, err)
		return
	}
	configSet, err := services.ConfigSetService.Add(&addConfigSetDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, configSet)
}

// @Summary 更新配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Param default body dto.UpdateConfigSetDto true "配置集"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.ConfigSetDto]
// @Router /ConfigSet [put]
func Update(c *gin.Context) {
	var updateConfigSetDto dto.UpdateConfigSetDto
	if err := c.ShouldBindJSON(&updateConfigSetDto); err != nil {
		response.FailError(c, err)
		return
	}
	configSet, err := services.ConfigSetService.Update(&updateConfigSetDto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, configSet)
}

// @Summary 查询配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Param id path int true "配置集Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.ConfigSetDto]
// @Router /ConfigSet/{id} [get]
func GetOne(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	configSet, err := services.ConfigSetService.GetOne(id)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, configSet)
}

// @Summary 删除配置集
// @Schemes
// @Description
// @Tags ConfigSet
// @Param id path int true "配置集Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /ConfigSet/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	err = services.ConfigSetService.Delete(id)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.Success(c)
}
