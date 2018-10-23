package service

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/mafanr/vgo/util"
	"github.com/vmihailenco/msgpack"

	"github.com/labstack/echo"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/agent/protocol"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// SkyWalking ...
type SkyWalking struct {
	appRegisterC chan *appRegister
}

// NewSkyWalking ...
func NewSkyWalking() *SkyWalking {

	return &SkyWalking{
		appRegisterC: make(chan *appRegister, 10),
	}
}

// Start ...
func (sky *SkyWalking) Start() error {

	// start http server
	sky.httpSer()

	// init grpc
	sky.grpcSer()

	return nil
}

func (sky *SkyWalking) grpcSer() error {
	lis, err := net.Listen("tcp", misc.Conf.SkyWalking.RPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	protocol.RegisterApplicationRegisterServiceServer(s, &appRegister{})
	// protocol.RegisterInstanceDiscoveryServiceServer(s, &InstanceDiscoveryService{})
	// protocol.RegisterServiceNameDiscoveryServiceServer(s, &ServiceNameDiscoveryService{})
	// protocol.RegisterNetworkAddressRegisterServiceServer(s, &NetworkAddressRegisterService{})
	// protocol.RegisterJVMMetricsServiceServer(s, &JVMMetricsService{})
	// protocol.RegisterTraceSegmentServiceServer(s, &TraceSegmentService{})

	s.Serve(lis)
	return nil
}

func (sky *SkyWalking) httpSer() error {
	e := echo.New()
	e.GET("/agent/gRPC", rpcAddr)
	go e.Start(misc.Conf.SkyWalking.HTTPAddr)
	g.L.Info("Start:sky.httpSer", zap.String("httpSer", "ok"))
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

// appRegister ...
type appRegister struct {
}

// ApplicationCodeRegister ...
func (a *appRegister) ApplicationCodeRegister(ctx context.Context, al *protocol.Application) (*protocol.ApplicationMapping, error) {

	g.L.Info("ApplicationCodeRegister", zap.String("ApplicationCode", al.ApplicationCode))
	// AppRegister
	registerPacker := &util.AppRegister{
		Name: al.ApplicationCode,
	}

	buf, err := msgpack.Marshal(registerPacker)
	if err != nil {
		g.L.Warn("ApplicationCodeRegister:msgpack.Marshal", zap.String("error", err.Error()))
		return nil, err
	}

	skp := util.SkywalkingPacket{
		Type:    util.TypeOfAppRegister,
		Payload: buf,
	}

	payload, err := msgpack.Marshal(skp)
	if err != nil {
		g.L.Warn("ApplicationCodeRegister:msgpack.Marshal", zap.String("error", err.Error()))
		return nil, err
	}

	packet := &util.VgoPacket{
		Type:       util.TypeOfSkywalking,
		Version:    util.VersionOf01,
		IsSync:     util.TypeOfSyncYes,
		IsCompress: util.TypeOfCompressNo,
		ID:         gAgent.ID(),
		PayLoad:    payload,
	}

	if err := gAgent.client.WritePacket(packet); err != nil {
		g.L.Warn("ApplicationCodeRegister:gAgent.client.WritePacket", zap.String("error", err.Error()))
		return nil, err
	}

	// var id int32 = 1111
	// kv := &protocol.KeyWithIntegerValue{
	// 	Key:   al.ApplicationCode,
	// 	Value: id,
	// }

	// appMapping := &protocol.ApplicationMapping{
	// 	Application: kv,
	// }

	// return appMapping, nil
	return nil, nil
}
