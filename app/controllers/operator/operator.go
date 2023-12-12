package operator

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/common/response"
	"algorithmplatform/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 算子列表
// @Schemes
// @Description
// @Tags Operator
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[[]dto.OperatorDto]
// @Router /Operator [get]
func List(c *gin.Context) {
	operators := services.OperatorService.List()
	response.SuccessWithData(c, operators)
}

// @Summary 新增算子
// @Schemes
// @Description
// @Tags Operator
// @Param default body dto.AddOperatorDto true "参数"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.OperatorDto]
// @Router /Operator [post]
func Add(c *gin.Context) {
	req := dto.AddOperatorDto{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailError(c, err)
		return
	}
	operator, err := services.OperatorService.Add(&req)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, operator)
}

// @Summary 修改算子
// @Schemes
// @Description
// @Tags Operator
// @Param default body dto.UpdateOperatorDto true "参数"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.OperatorDto]
// @Router /Operator [put]
func Update(c *gin.Context) {
	req := dto.UpdateOperatorDto{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailError(c, err)
		return
	}
	if operator, err := services.OperatorService.Update(&req); err != nil {
		response.FailError(c, err)
		return
	} else {
		response.SuccessWithData(c, operator)
	}
}

// @Summary 删除算子
// @Schemes
// @Description
// @Tags Operator
// @Param id path int64 true "算子Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Operator/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	services.OperatorService.Delete(id)
	response.Success(c)
}

// @Summary 查询算子
// @Schemes
// @Description
// @Tags Operator
// @Param id path int64 true "算子Id"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[dto.OperatorDto]
// @Router /Operator/{id} [get]
func GetOne(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.FailError(c, err)
		return
	}
	operator, err := services.OperatorService.GetOne(id)
	if err != nil {
		response.FailError(c, err)
		return
	}
	response.SuccessWithData(c, operator)
}

// @Summary 上传算子文件
// @Schemes
// @Description
// @Tags Operator
// @Param file formData  file true "文件"
// @Accept json
// @Produce json
// @Success 200 {object} response.TResponse[string]
// @Router /Operator/UploadFiles [post]
func UploadFiles(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		response.FailError(c, err)
		return
	}
	if fileKey, err := services.FileService.UploadFile(file); err != nil {
		response.FailError(c, err)
	} else {
		response.SuccessWithData(c, fileKey)
	}

}

// @Summary 更新算子文件
// @Schemes
// @Description
// @Tags Operator
// @Param default body   dto.UpdateFileDto true "更新文件"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /Operator/UpdateFile [post]
func UpdateFile(c *gin.Context) {
	var updateFileDto dto.UpdateFileDto
	if err := c.ShouldBindJSON(&updateFileDto); err != nil {
		response.FailError(c, err)
		return
	}
	if err := services.OperatorService.UpdateFile(updateFileDto); err != nil {
		response.FailError(c, err)
		return
	}
	response.Success(c)
}
