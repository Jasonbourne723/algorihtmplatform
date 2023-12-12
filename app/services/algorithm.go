package services

import (
	"algorithmplatform/app/common"
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"encoding/json"
	"strconv"
	"time"
)

type algorithmSerivce struct{}

var AlgorithmService = new(algorithmSerivce)

// 序列化配置
func generateConfig(addAlgorithmDto *dto.AlgorithmItemsDto) string {
	bytes, _ := json.Marshal(struct {
		InputAlgorithm  []dto.AlgorithmItemDto
		OutputAlgorithm []dto.AlgorithmItemDto
		AlgoAlgorithm   dto.AlgorithmItemDto
	}{
		addAlgorithmDto.InputAlgorithm,
		addAlgorithmDto.OutputAlgorithm,
		addAlgorithmDto.AlgoAlgorithm,
	})
	return string(bytes)
}

// 反序列化配置
func deserializeConfig(config string) *dto.AlgorithmItemsDto {
	var entity dto.AlgorithmItemsDto
	json.Unmarshal([]byte(config), &entity)
	return &entity
}

// 模型转换
func mapToAlgorithmDto(algorithm *models.Algorithm) *dto.AlgorithmDto {
	config := deserializeConfig(algorithm.Config)
	return &dto.AlgorithmDto{
		Id:              algorithm.Id,
		Name:            algorithm.Name,
		Description:     algorithm.Remark,
		Type:            algorithm.Type,
		ProjectId:       algorithm.ProjectId,
		CreateUserId:    algorithm.CreateUserId,
		CreateUserName:  algorithm.CreateUserName,
		CreateDate:      algorithm.CreateDate,
		InputAlgorithm:  config.InputAlgorithm,
		OutputAlgorithm: config.OutputAlgorithm,
		AlgoAlgorithm:   config.AlgoAlgorithm,
	}
}

// 模型转换
func mapToAlgorithmDtos(algorithms []models.Algorithm) []dto.AlgorithmDto {
	var algorithmDtos []dto.AlgorithmDto
	for _, a := range algorithms {
		algorithmDtos = append(algorithmDtos, *mapToAlgorithmDto(&a))
	}
	return algorithmDtos
}

// 根据projectId获取算法列表
func (a *algorithmSerivce) List(projectId int64) ([]dto.AlgorithmDto, error) {
	var algorithms []models.Algorithm
	if projectId == 0 {
		if err := global.App.DB.Find(&algorithms).Error; err != nil {
			return nil, err
		}
	} else {
		if err := global.App.DB.Where("project_id = ? ", projectId).Find(&algorithms).Error; err != nil {
			return nil, err
		}
	}
	return mapToAlgorithmDtos(algorithms), nil
}

// 查询单个算法
func (a *algorithmSerivce) GetOne(algorithmId int64) (*dto.AlgorithmDto, error) {
	var algorithm models.Algorithm
	if err := global.App.DB.First(&algorithm, algorithmId).Error; err != nil {
		return nil, err
	} else {
		return mapToAlgorithmDto(&algorithm), nil
	}
}

// 新增算法
func (a *algorithmSerivce) Add(addAlgorithmDto *dto.AddAlgorithmDto) (*dto.AlgorithmDto, error) {
	var config = generateConfig(&addAlgorithmDto.AlgorithmItemsDto)
	algorithm := models.Algorithm{
		Name:           addAlgorithmDto.Name,
		Remark:         addAlgorithmDto.Description,
		CreateUserId:   addAlgorithmDto.UserId,
		CreateUserName: addAlgorithmDto.UserName,
		ProjectId:      addAlgorithmDto.ProjectId,
		CreateDate:     time.Now(),
		Type:           addAlgorithmDto.Type,
		Config:         config,
		Version:        strconv.Itoa(int(time.Now().Unix())),
	}
	if err := global.App.DB.Create(&algorithm).Error; err != nil {
		return nil, err
	} else {
		return mapToAlgorithmDto(&algorithm), nil
	}
}

// 删除算法
func (a *algorithmSerivce) Delete(id int64) {
	global.App.DB.Delete(&models.Algorithm{}, id)
}

// 更新算法
func (a *algorithmSerivce) Update(updateAlgorithmDto *dto.UpdateAlgorithmDto) (*dto.AlgorithmDto, error) {
	var algorithm models.Algorithm
	err := global.App.DB.First(&algorithm, updateAlgorithmDto.Id).Error
	if err != nil {
		return nil, err
	}
	config := generateConfig(&updateAlgorithmDto.AlgorithmItemsDto)
	algorithm.Name = updateAlgorithmDto.Name
	algorithm.Remark = updateAlgorithmDto.Description
	algorithm.Type = updateAlgorithmDto.Type
	algorithm.Config = config
	if err = global.App.DB.Save(&algorithm).Error; err != nil {
		return nil, err
	} else {
		FileService.CleanCache(algorithm.Id)
		return mapToAlgorithmDto(&algorithm), nil
	}
}

// 获取数据源
func (a *algorithmSerivce) GetDataSource(algorithmId int64) (*dto.DataSourceDto, error) {
	var dataSource models.DataSource
	if err := global.App.DB.Where("algorithm_id = ?", algorithmId).First(&dataSource).Error; err != nil {
		return &dto.DataSourceDto{
			BeginDate: common.TimeStamp(time.Now()),
			EndDate:   common.TimeStamp(time.Now()),
		}, nil
	} else {
		return &dto.DataSourceDto{
			AlgorithmId: dataSource.AlgorithmId,
			BeginDate:   common.TimeStamp(dataSource.BeginDate),
			EndDate:     common.TimeStamp(dataSource.EndDate),
			CompanyId:   dataSource.CompanyId,
			CraneId:     dataSource.CraneId,
			CompanyName: dataSource.CompanyName,
			CraneName:   dataSource.CraneName,
		}, nil
	}
}

// 设置数据源
func (a *algorithmSerivce) SetDataSource(dataSourceDto dto.DataSourceDto) {
	global.App.DB.Where("algorithm_id = ?", dataSourceDto.AlgorithmId).Delete(&models.DataSource{})
	global.App.DB.Create(&models.DataSource{
		AlgorithmId: dataSourceDto.AlgorithmId,
		BeginDate:   dataSourceDto.BeginDate.ToTime(),
		EndDate:     dataSourceDto.EndDate.ToTime(),
		CompanyId:   dataSourceDto.CompanyId,
		CraneId:     dataSourceDto.CraneId,
		CompanyName: dataSourceDto.CompanyName,
		CraneName:   dataSourceDto.CraneName,
	})
}

// 运行算法
func (a *algorithmSerivce) Run(algorithmId int64, userId int64) {
	global.App.AlgoChan <- &dto.RunAlgorithm{
		AlgorithmId: algorithmId,
		UserId:      userId,
	}
}
