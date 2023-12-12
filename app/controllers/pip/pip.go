package pip

import (
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"

	"github.com/gin-gonic/gin"
)

// @Summary 查询python包列表 pip list
// @Schemes
// @Description
// @Tags pip
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Pip/List [get]
func List(c *gin.Context) {
	userId := c.GetInt64("userId")
	go services.PipService.List(userId)
	response.Success(c)
}

// @Summary 安装python包  pip install
// @Schemes
// @Description
// @Tags pip
// @Param package query string true "包名"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Pip/Install [get]
func Install(c *gin.Context) {

	packageName := c.Query("package")
	userId := c.GetInt64("userId")
	go services.PipService.Install(packageName, userId)

	response.Success(c)
}
