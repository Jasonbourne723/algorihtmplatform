package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"time"
)

type configSetService struct{}

var ConfigSetService = new(configSetService)

func (c *configSetService) GetAll() ([]dto.ConfigSetDto, error) {
	var configSets = []models.ConfigSet{}
	if err := global.App.DB.Find(&configSets).Error; err != nil {
		return nil, err
	}
	var configSetDtos = []dto.ConfigSetDto{}
	for _, ea := range configSets {
		configSetDtos = append(configSetDtos, *mapToConfigSetDto(&ea))
	}
	return configSetDtos, nil
}

func (c *configSetService) GetPage(pageIndex int, pageSize int, name string) (*dto.PageData[dto.ConfigSetDto], error) {
	var total int64
	if err := global.App.DB.Model(models.ConfigSet{}).Where("name like ?", "%"+name+"%").Count(&total).Error; err != nil {
		return nil, err
	}
	var configSets []dto.ConfigSetDto
	if err := global.App.DB.Model(models.ConfigSet{}).Where("name like ?", "%"+name+"%").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&configSets).Error; err != nil {
		return nil, err
	}

	return &dto.PageData[dto.ConfigSetDto]{
		RecordCount: total,
		PageCount:   getPageCount(total, int64(pageSize)),
		TModel:      configSets,
	}, nil

}

func getPageCount(total int64, pageSize int64) (pageCount int64) {
	pageCount = total / pageSize
	if total%pageSize > 0 {
		pageCount += 1
	}
	return pageCount
}

func (c *configSetService) Add(req *dto.AddConfigSetDto) (*dto.ConfigSetDto, error) {
	configSet := models.ConfigSet{
		Name:          req.Name,
		Description:   req.Description,
		Tag:           req.Tag,
		KeyValuePairs: req.KeyValuePairs,
		CreateDate:    time.Now(),
		CreateUserId:  req.UserId,
	}
	if err := global.App.DB.Create(&configSet).Error; err != nil {
		return nil, err
	}
	return mapToConfigSetDto(&configSet), nil
}

func (c *configSetService) Update(req *dto.UpdateConfigSetDto) (*dto.ConfigSetDto, error) {
	var configSet models.ConfigSet
	if err := global.App.DB.First(&configSet, req.Id).Error; err != nil {
		return nil, err
	}
	configSet.Name = req.Name
	configSet.Description = req.Description
	configSet.Tag = req.Tag
	configSet.KeyValuePairs = req.KeyValuePairs

	if err := global.App.DB.Save(&configSet).Error; err != nil {
		return nil, err
	}
	return mapToConfigSetDto(&configSet), nil
}

func (c *configSetService) GetOne(id int64) (*dto.ConfigSetDto, error) {
	var configSet models.ConfigSet
	if err := global.App.DB.First(&configSet, id).Error; err != nil {
		return nil, err
	}
	return mapToConfigSetDto(&configSet), nil
}

func (o *configSetService) Delete(id int64) error {
	if err := global.App.DB.Delete(&models.ConfigSet{}, id).Error; err != nil {
		return err
	}
	return nil
}

func mapToConfigSetDto(c *models.ConfigSet) *dto.ConfigSetDto {
	return &dto.ConfigSetDto{
		Id:            c.Id,
		Name:          c.Name,
		Description:   c.Description,
		Tag:           c.Tag,
		KeyValuePairs: c.KeyValuePairs,
		CreateDate:    c.CreateDate,
	}
}
