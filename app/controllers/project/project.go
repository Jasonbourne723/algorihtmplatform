package project

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 工程列表
// @Schemes
// @Description
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.ProjectDto]
// @Router /Project [get]
func List(c *gin.Context) {
	projects := services.ProjectService.List()
	response.SuccessWithData(c, projects)
}

// @Summary 新增工程
// @Schemes
// @Description
// @Tags Project
// @Param default body dto.AddProjectDto true "参数"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.ProjectDto]
// @Router /Project [post]
func Add(c *gin.Context) {
	var dto dto.AddProjectDto
	err := c.BindJSON(&dto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	dto.UserId = c.GetInt64("userId")
	dto.UserName = c.GetString("name")

	project := services.ProjectService.Add(&dto)
	response.SuccessWithData(c, project)
}

// @Summary 修改工程
// @Schemes
// @Description
// @Tags Project
// @Param default body dto.UpdateProjectDto true "参数"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.ProjectDto]
// @Router /Project [put]
func Update(c *gin.Context) {
	var dto dto.UpdateProjectDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		response.FailError(c, err)
		return
	}

	project, err := services.ProjectService.Update(&dto)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, project)
}

// @Summary 删除工程
// @Schemes
// @Description
// @Tags Project
// @Param id path int64 true "项目Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Project/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.FailError(c, err)
	}
	services.ProjectService.Delete(id)
	response.Success(c)
}
