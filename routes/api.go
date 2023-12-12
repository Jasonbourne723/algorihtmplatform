package routes

import (
	"algorithmplatform/app/controllers/algorithm"
	"algorithmplatform/app/controllers/algorithmpackage"
	"algorithmplatform/app/controllers/configset"
	"algorithmplatform/app/controllers/file"
	"algorithmplatform/app/controllers/operator"
	"algorithmplatform/app/controllers/pip"
	"algorithmplatform/app/controllers/project"
	"algorithmplatform/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("").Use(middleware.Recover()).Use(middleware.JWTAuth())
	{
		authGroup.GET("/Project", project.List)
		authGroup.POST("/Project", project.Add)
		authGroup.PUT("/Project", project.Update)
		authGroup.DELETE("/Project/:id", project.Delete)
	}

	{
		authGroup.GET("/Operator", operator.List)
		authGroup.POST("/Operator", operator.Add)
		authGroup.PUT("/Operator", operator.Update)
		authGroup.DELETE("/Operator/:id", operator.Delete)
		authGroup.GET("/Operator/:id", operator.GetOne)
		authGroup.POST("/Operator/UploadFiles", operator.UploadFiles)
		authGroup.POST("/Operator/UpdateFile", operator.UpdateFile)
	}

	{
		authGroup.GET("/File/GetFilesByOperatorId", file.GetFilesByOperatorId)
		authGroup.GET("/File/GetFileContent", file.GetFileContent)
		authGroup.GET("/File/GetFilesByAlgorithmId", file.GetFilesByAlgorithmId)
	}

	{
		authGroup.GET("/ConfigSet", configset.List)
		authGroup.GET("/ConfigSet/All", configset.GetAll)
		authGroup.GET("/ConfigSet/:id", configset.GetOne)
		authGroup.POST("/ConfigSet", configset.Add)
		authGroup.PUT("/ConfigSet", configset.Update)
		authGroup.DELETE("/ConfigSet/:id", configset.Delete)
	}

	{
		authGroup.POST("/Algorithm", algorithm.Post)
		authGroup.PUT("/Algorithm", algorithm.Put)
		authGroup.POST("/Algorithm/DataSourceConfig", algorithm.SetDataSourceConfig)
		authGroup.GET("/Algorithm/DataSourceConfig/:algorithmId", algorithm.GetDataSourceConfig)
		authGroup.GET("/Algorithm/GetByProjectId", algorithm.GetByProjectId)
		authGroup.POST("/Algorithm/UpdateAlgorithmFile", algorithm.UpdateAlgorithmFile)
		authGroup.GET("/Algorithm/Run/:algorithmId", algorithm.Run)
		authGroup.GET("/Algorithm/:algorithmId", algorithm.Get)
		authGroup.DELETE("/Algorithm/:algorithmId", algorithm.Delete)
	}

	{
		authGroup.POST("/AlgorithmPackage", algorithmpackage.Add)
		authGroup.GET("/AlgorithmPackage/Down/:algorithmPackageId", algorithmpackage.Down)
		authGroup.GET("/AlgorithmPackage/Page", algorithmpackage.Page)
		authGroup.GET("/AlgorithmPackage", algorithmpackage.List)
	}

	{
		authGroup.GET("/Pip/List", pip.List)
		authGroup.GET("/Pip/Install", pip.Install)
	}

}
