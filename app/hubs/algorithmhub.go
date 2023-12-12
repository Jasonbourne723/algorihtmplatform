package hubs

import (
	"algorithmplatform/app/common"
	"algorithmplatform/global"
	"algorithmplatform/signalr"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

const (
	PIP_LIST         = "pipList"
	PIP_INSTALL      = "pipInstall"
	ALGORITHM_RESULT = "AlgorithmResult"
)

type algorithmHub struct {
	signalr.Hub
	initialized bool
}

var AlgorithmHub = new(algorithmHub)

func (h *algorithmHub) Initialize(ctx signalr.HubContext) {
	h.Hub.Initialize(ctx)
	// Initialize will be called on first connection to the hub
	// if position is sent before that, the HubContext is nil and application crashes
	// TODO: possible bug in singlar implementation?
	h.initialized = true
}

func (h *algorithmHub) OnConnected(connectionID string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	if h.initialized {
		token := h.AccessToken()
		userId, err := getUserIdByAccessToken(token)
		if err == nil {
			h.Groups().AddToGroup(userId, connectionID)
		}
	}
}

func (h *algorithmHub) OnDisconnected(connectionID string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	if h.initialized {
		token := h.AccessToken()
		userId, err := getUserIdByAccessToken(token)
		if err == nil {
			h.Groups().RemoveFromGroup(userId, connectionID)
		}
	}
}

func (h *algorithmHub) SendPipList(message string, userId int64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	if h.initialized {
		h.Clients().Group(strconv.Itoa(int(userId))).Send(PIP_LIST, message)
	}
}

func (h *algorithmHub) SendPipInstall(message string, userId int64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	if h.initialized {
		h.Clients().Group(strconv.Itoa(int(userId))).Send(PIP_INSTALL, message)
	}
}

func (h *algorithmHub) SendAlgorithmOutput(message string, userId int64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	if h.initialized {
		h.Clients().Group(strconv.Itoa(int(userId))).Send(ALGORITHM_RESULT, message)
	}
}

func getUserIdByAccessToken(accessToken string) (string, error) {
	accessToken = accessToken[len("Bearer "):]
	var claims common.CustomJwt
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.App.Config.Jwt.Secret), nil
	})
	if err != nil {
		return "", nil
	}
	return claims.UserId, nil
}
