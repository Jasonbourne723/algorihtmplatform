package dto

type AlgorithmParam struct {
	Type       int32
	Inputs     []ItemParam
	Outputs    []ItemParam
	Algo       ItemParam
	CommonPara map[string]string
}

type ItemParam struct {
	Name string
	Para map[string]string
}
