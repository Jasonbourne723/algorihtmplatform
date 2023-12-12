package bootstrap

import (
	"algorithmplatform/app/actuator"
	"algorithmplatform/app/common/dto"
	"algorithmplatform/global"
)

func InitializeActuator() {
	global.App.AlgoChan = make(chan *dto.RunAlgorithm, 10)
	go func() {
		actuator.Start()
	}()
}
