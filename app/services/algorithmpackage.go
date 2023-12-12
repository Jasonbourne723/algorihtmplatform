package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"algorithmplatform/utils"
	"os"
	"path"
	"time"
)

type algorithmPackageService struct{}

var AlgorithmPackageService = new(algorithmPackageService)

func mapToAlgorithmPackageDtos(algorithmPackages []models.AlgorithmPackage) []dto.AlgorithmPackageDto {
	var packageDtos []dto.AlgorithmPackageDto
	for _, ea := range algorithmPackages {

		items := deserializeConfig(ea.Config)
		packageDtos = append(packageDtos, dto.AlgorithmPackageDto{
			Id:              ea.Id,
			Name:            ea.Name,
			AlgorithmId:     ea.AlgorithmId,
			AlgorithmName:   ea.AlgorithmName,
			CompanyId:       ea.CompanyId,
			CraneId:         ea.CraneId,
			Period:          ea.Period,
			Cron:            ea.Cron,
			CreateDate:      ea.CreateDate,
			InputAlgorithm:  items.InputAlgorithm,
			OutputAlgorithm: items.OutputAlgorithm,
			AlgoAlgorithm:   items.AlgoAlgorithm,
		})
	}
	return packageDtos
}

func (a *algorithmPackageService) Page(pageIndex int, pageSize int, name string) (*dto.PageData[dto.AlgorithmPackageDto], error) {

	var total int64
	if err := global.App.DB.Model(models.AlgorithmPackage{}).Where("name like ?", "%"+name+"%").Count(&total).Error; err != nil {
		return nil, err
	}

	var packages []models.AlgorithmPackage
	if err := global.App.DB.Where("name like ?", "%"+name+"%").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&packages).Error; err != nil {
		return nil, err
	}

	packageDtos := mapToAlgorithmPackageDtos(packages)
	return &dto.PageData[dto.AlgorithmPackageDto]{
		RecordCount: total,
		PageCount:   getPageCount(total, int64(pageSize)),
		TModel:      packageDtos,
	}, nil
}

func (a *algorithmPackageService) List() ([]dto.AlgorithmPackageDto, error) {

	var packages []models.AlgorithmPackage
	if err := global.App.DB.Find(&packages).Error; err != nil {
		return nil, err
	}
	return mapToAlgorithmPackageDtos(packages), nil
}

func (a *algorithmPackageService) Add(addAlgorithmPackageDto *dto.AddAlgorithmPackageDto) error {

	config := generateConfig(&dto.AlgorithmItemsDto{
		InputAlgorithm:  addAlgorithmPackageDto.InputAlgorithm,
		OutputAlgorithm: addAlgorithmPackageDto.OutputAlgorithm,
		AlgoAlgorithm:   addAlgorithmPackageDto.AlgoAlgorithm,
	})

	var algorithm models.Algorithm
	if err := global.App.DB.First(&algorithm, addAlgorithmPackageDto.AlgorithmId).Error; err != nil {
		return err
	}
	algorithmPackage := &models.AlgorithmPackage{
		Name:             addAlgorithmPackageDto.Name,
		AlgorithmId:      addAlgorithmPackageDto.AlgorithmId,
		AlgorithmName:    addAlgorithmPackageDto.AlgorithmName,
		CompanyId:        addAlgorithmPackageDto.CompanyId,
		CraneId:          addAlgorithmPackageDto.CraneId,
		Cron:             addAlgorithmPackageDto.Cron,
		Period:           addAlgorithmPackageDto.Period,
		AlgorithmVersion: algorithm.Version,
		Config:           config,
		ModelFileId:      0,
		Status:           0,
		FilePath:         "",
		CreateDate:       time.Now(),
		IsDelete:         0,
	}
	filePath, err := FileService.GetPackageFilePath(algorithmPackage)
	if err != nil {
		return err
	}
	algorithmPackage.FilePath = filePath
	if err = global.App.DB.Create(algorithmPackage).Error; err != nil {
		return err
	}
	return nil
}

func (a *algorithmPackageService) DownLoad(algorithmPackageId int64) (string, error) {

	var algorithmPackage models.AlgorithmPackage
	if err := global.App.DB.First(&algorithmPackage, algorithmPackageId).Error; err != nil {
		return "", err
	}

	dir := getAlgorithmPackageDir(algorithmPackageId)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}
	zipPath := path.Join(dir, algorithmPackage.FilePath)

	if err := utils.NewEosClient().DownLoad(algorithmPackage.FilePath, zipPath); err != nil {
		return "", err
	}
	return zipPath, nil
}
