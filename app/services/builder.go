package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"strconv"
)

type algorithmBuilder struct {
}

var AlgorithmBuilder = new(algorithmBuilder)

func (a *algorithmBuilder) Build(algorithmId int64) (*dto.AlgorithmParam, error) {
	algorithm, err := AlgorithmService.GetOne(algorithmId)
	if err != nil {
		return nil, err
	}
	var inputs []dto.ItemParam
	var outputs []dto.ItemParam
	var algo dto.ItemParam

	for _, ea := range algorithm.InputAlgorithm {
		operator, err := OperatorService.GetOne(ea.OperatorId)
		if err != nil {
			continue
		}
		inputs = append(inputs, *mapToItemParam(&ea, operator))
	}

	algoOperator, err := OperatorService.GetOne(algorithm.AlgoAlgorithm.OperatorId)
	if err == nil {
		algo = *mapToItemParam(&algorithm.AlgoAlgorithm, algoOperator)
	}

	for _, ea := range algorithm.OutputAlgorithm {
		operator, err := OperatorService.GetOne(ea.OperatorId)
		if err != nil {
			continue
		}
		outputs = append(outputs, *mapToItemParam(&ea, operator))
	}

	commonPara, err := getCommonParaByAlgorithmId(algorithmId)
	if err != nil {
		return nil, err
	}

	return &dto.AlgorithmParam{
		Inputs:     inputs,
		Algo:       algo,
		Outputs:    outputs,
		Type:       algorithm.Type,
		CommonPara: commonPara,
	}, nil
}

func (a *algorithmBuilder) BuildForPacket(algorithmPackage *models.AlgorithmPackage) (*dto.AlgorithmParam, error) {

	var inputs []dto.ItemParam
	var outputs []dto.ItemParam
	var algo dto.ItemParam

	items := deserializeConfig(algorithmPackage.Config)

	for _, ea := range items.InputAlgorithm {
		operator, err := OperatorService.GetOne(ea.OperatorId)
		if err != nil {
			continue
		}
		inputs = append(inputs, *mapToItemParam(&ea, operator))
	}

	algoOperator, err := OperatorService.GetOne(items.AlgoAlgorithm.OperatorId)
	if err == nil {
		algo = *mapToItemParam(&items.AlgoAlgorithm, algoOperator)
	}

	for _, ea := range items.OutputAlgorithm {
		operator, err := OperatorService.GetOne(ea.OperatorId)
		if err != nil {
			continue
		}
		outputs = append(outputs, *mapToItemParam(&ea, operator))
	}

	commonPara := getCommonParaByCraneId(algorithmPackage.CraneId)
	return &dto.AlgorithmParam{
		Inputs:     inputs,
		Algo:       algo,
		Outputs:    outputs,
		Type:       0,
		CommonPara: commonPara,
	}, nil
}

func mapToItemParam(algorithmItem *dto.AlgorithmItemDto, operator *dto.OperatorDto) *dto.ItemParam {
	var paras map[string]string = make(map[string]string)

	for _, ea := range algorithmItem.ItemParams {
		paras[ea.Key] = ea.Value
	}
	itemParam := &dto.ItemParam{
		Name: operator.Name + "." + operator.ClassName,
		Para: paras,
	}
	return itemParam
}

func getCommonParaByCraneId(craneId int64) map[string]string {
	var exparams []models.ExParams
	var commonPara map[string]string = make(map[string]string, 10)
	if err := global.App.DB.Where("crane_Id = ?", craneId).Find(&exparams).Error; err != nil {
		return commonPara
	}

	for _, ea := range exparams {
		commonPara[ea.Name] = ea.Value
	}
	return commonPara
}

func getCommonParaByAlgorithmId(algorithmId int64) (map[string]string, error) {
	var dataSource models.DataSource
	if err := global.App.DB.Where("algorithm_id = ?", algorithmId).First(&dataSource).Error; err != nil {
		return nil, err
	}
	commonPara := getCommonParaByCraneId(dataSource.CraneId)
	commonPara["beginDate"] = dataSource.BeginDate.Format("2006-01-02 15:04:05")
	commonPara["endDate"] = dataSource.EndDate.Format("2006-01-02 15:04:05")
	commonPara["craneId"] = strconv.FormatInt(dataSource.CraneId, 10)
	commonPara["companyId"] = strconv.FormatInt(dataSource.CompanyId, 10)
	return commonPara, nil
}
