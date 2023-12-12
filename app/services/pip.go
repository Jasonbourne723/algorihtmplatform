package services

import (
	"algorithmplatform/app/hubs"
	"fmt"
	"os/exec"
)

type pipService struct{}

var PipService = new(pipService)

func (p *pipService) List(userId int64) {
	cmd := exec.Command("pip", "list")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("pip install error: " + err.Error())
		hubs.AlgorithmHub.SendPipList(err.Error(), userId)
	}
	cmd.Run()
	fmt.Println(string(out))
	hubs.AlgorithmHub.SendPipList(string(out), userId)
}

func (p *pipService) Install(packageName string, userId int64) {
	cmd := exec.Command("pip", "install", packageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("pip install error: " + err.Error())
		hubs.AlgorithmHub.SendPipInstall(err.Error(), userId)
	}
	cmd.Run()
	fmt.Println(string(out))
	hubs.AlgorithmHub.SendPipInstall(string(out), userId)
}
