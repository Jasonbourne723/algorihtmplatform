package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"time"
)

type projectService struct {
}

var ProjectService = new(projectService)

// 项目 列表
func (p *projectService) List() []dto.ProjectDto {
	var projects []models.Project
	global.App.DB.Find(&projects)

	var projectDtos []dto.ProjectDto
	for _, project := range projects {
		projectDtos = append(projectDtos, mapToProjectDto(&project))
	}
	return projectDtos
}

// 新增 项目
func (p *projectService) Add(addProjectDto *dto.AddProjectDto) dto.ProjectDto {
	project := models.Project{
		Name:           addProjectDto.Name,
		Description:    addProjectDto.Description,
		CreateUserId:   addProjectDto.UserId,
		CreateUserName: addProjectDto.UserName,
		CreateDate:     time.Now(),
	}
	global.App.DB.Create(&project)
	return mapToProjectDto(&project)
}

// 更新 项目
func (p *projectService) Update(updateProjectDto *dto.UpdateProjectDto) (projectDto dto.ProjectDto, err error) {
	var project models.Project
	res := global.App.DB.First(&project, updateProjectDto.Id)

	if res.Error != nil {
		return projectDto, res.Error
	}
	project.Name = updateProjectDto.Name
	project.Description = updateProjectDto.Description

	global.App.DB.Save(&project)
	return mapToProjectDto(&project), err
}

// 删除 项目
func (p *projectService) Delete(id int64) {
	global.App.DB.Delete(&models.Project{}, id)
}

func mapToProjectDto(p *models.Project) dto.ProjectDto {
	return dto.ProjectDto{
		Id:             p.Id,
		Name:           p.Name,
		CreateUserId:   p.CreateUserId,
		CreateUserName: p.CreateUserName,
		CreateDate:     p.CreateDate,
		Description:    p.Description,
	}
}
