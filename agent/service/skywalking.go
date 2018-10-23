package service

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mafanr/vgo/agent/misc"
)

// SkyWalking ...
type SkyWalking struct {
}

// NewSkyWalking ...
func NewSkyWalking() *SkyWalking {

	return &SkyWalking{}
}

// Start ...
func (sky *SkyWalking) Start() error {
	e := echo.New()
	e.GET("/agent/gRPC", rpcAddr)
	go e.Start(misc.Conf.SkyWalking.HTTPAddr)

	return nil
}

// Close ...
func (sky *SkyWalking) Close() error {
	return nil
}

// rpcAddr ...
func rpcAddr(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{misc.Conf.SkyWalking.RPCAddr})
}
