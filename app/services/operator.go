package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"database/sql"
	"strconv"
	"time"
)

type operatorService struct {
}

var OperatorService = new(operatorService)

func mapToOperatorDto(o *models.Operator) *dto.OperatorDto {
	return &dto.OperatorDto{
		Id:             o.Id,
		Name:           o.Name,
		ClassName:      o.ClassName,
		CreateUserName: o.CreateUserName,
		CreateUserId:   o.CreateUserId,
		Description:    o.Description,
		Type:           o.Type,
		FilePath:       o.FilePath,
		Version:        o.Version,
		CreateDate:     o.CreateDate,
		UpdateDate:     o.UpdateDate.Time,
	}
}

func (o *operatorService) List() []dto.OperatorDto {
	var operators []models.Operator
	global.App.DB.Find(&operators)

	var operatorDtos []dto.OperatorDto
	for _, ea := range operators {
		operatorDtos = append(operatorDtos, *mapToOperatorDto(&ea))
	}
	return operatorDtos
}

func (o *operatorService) Add(req *dto.AddOperatorDto) (*dto.OperatorDto, error) {
	operator := models.Operator{
		Name:           req.Name,
		Description:    req.Description,
		ClassName:      req.ClassName,
		Type:           req.Type,
		Version:        getVersion(),
		FilePath:       req.FilePath,
		CreateUserId:   req.UserId,
		CreateUserName: req.UserName,
		CreateDate:     time.Now(),
	}
	err := global.App.DB.Create(&operator).Error
	if err != nil {
		return nil, err
	}
	return mapToOperatorDto(&operator), nil
}

func (o *operatorService) Update(req *dto.UpdateOperatorDto) (*dto.OperatorDto, error) {
	var operator models.Operator
	if res := global.App.DB.First(&operator, req.Id); res.Error != nil {
		return nil, res.Error
	}
	operator.ClassName = req.ClassName
	operator.Description = req.Description
	operator.Name = req.Name
	operator.Type = req.Type
	operator.FilePath = req.FilePath
	operator.Version = getVersion()
	operator.UpdateDate = sql.NullTime{Time: time.Now(), Valid: true}
	global.App.DB.Save(&operator)
	FileService.CleanCacheByOperatorId(operator.Id)
	return mapToOperatorDto(&operator), nil
}

func (o *operatorService) Delete(id int64) {
	global.App.DB.Delete(&models.Operator{}, id)
}

func (o *operatorService) GetOne(id int64) (*dto.OperatorDto, error) {
	var operator models.Operator
	if res := global.App.DB.First(&operator, id); res.Error != nil {
		return nil, res.Error
	}
	return mapToOperatorDto(&operator), nil
}

func (o *operatorService) UpdateFile(updateFileDto dto.UpdateFileDto) error {
	var operator models.Operator
	if res := global.App.DB.First(&operator, updateFileDto.OperatorId); res.Error != nil {
		return res.Error
	}
	filePath, err := FileService.UpdateOperatorFile(updateFileDto.Path, updateFileDto.Content, updateFileDto.OperatorId)
	if err != nil {
		return err
	}
	operator.FilePath = filePath
	operator.Version = getVersion()
	global.App.DB.Save(&operator)
	return nil
}

func getVersion() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
