package actuator

import (
	"algorithmplatform/app/hubs"
	"algorithmplatform/app/services"
	"algorithmplatform/global"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

// type AlgorithmActuator struct{}

func Run(algorithmId int64, userId int64) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()

	startupPath := services.FileService.GetStartupPath(algorithmId)
	req, err := services.AlgorithmBuilder.Build(algorithmId)
	if err != nil {
		fmt.Println(err.Error())
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	reqStr := string(bytes)

	//reqStr = "\"" + strings.ReplaceAll(reqStr, "\"", "\\\"") + "\""
	cmd := exec.Command("python", startupPath, reqStr)
	var done = make(chan int, 1)
	hubs.AlgorithmHub.SendAlgorithmOutput("开始运行", userId)
	go heartBeat(done, userId)
	out, _ := cmd.CombinedOutput()

	cmd.Run()
	done <- 0
	hubs.AlgorithmHub.SendAlgorithmOutput(string(out), userId)
	hubs.AlgorithmHub.SendAlgorithmOutput("运行结束", userId)
}

func RunTest(algorithmId int64, userId int64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()

	startupPath := services.FileService.GetStartupPath(algorithmId)
	req, err := services.AlgorithmBuilder.Build(algorithmId)
	if err != nil {
		fmt.Println(err.Error())
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	reqStr := string(bytes)

	//reqStr = "\"" + strings.ReplaceAll(reqStr, "\"", "\\\"") + "\""
	cmd := exec.Command("python", "-u", startupPath, reqStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			hubs.AlgorithmHub.SendAlgorithmOutput(readString, userId)
		}
	}()
	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()
}

func heartBeat(done chan int, userId int64) {
	second := 0
loop:
	for {
		select {
		case <-done:
			break loop
		case <-time.After(5 * time.Second):
			second += 5
			hubs.AlgorithmHub.SendAlgorithmOutput("已执行"+strconv.Itoa(second)+"秒", userId)
		}
	}
}

func Start() {
	workerNums := 5
	for i := 0; i < workerNums; i++ {
		go func() {
			for {
				c := <-global.App.AlgoChan
				Run(c.AlgorithmId, c.UserId)
			}
		}()
	}
}
