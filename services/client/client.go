package client

import (
	"context"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type FortedClient struct {
	cfg  *config.Config
	log  *zap.SugaredLogger
	conn *grpc.ClientConn
	pb.FortedClient
	conf *config_client.ConfClient
}

func NewFortedClient(cfg *config.Config, log *zap.SugaredLogger, conf *config_client.ConfClient) *FortedClient {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Service.GrpcTimeout)
	defer cancel()
	addr, err := conf.GetGrpcAddr(ctx, "surebet-service")
	if err != nil {
		log.Panic("get grpc addr error for name: surebet-service")
	}
	log.Infow("connecting to surebet-service...", "addr", addr)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorw("NewFortedClient DialContext error", "addr", addr, "err", err)
	}
	cli := pb.NewFortedClient(conn)
	log.Infow("begin ping to surebet-service", "addr", addr)
	_, err = cli.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Errorw("ping to surebet-service error", "err", err)
	} else {
		log.Infow("ping done", "addr", addr)
	}
	return &FortedClient{cfg: cfg, log: log, conn: conn, FortedClient: cli, conf: conf}
}

func (s *FortedClient) Close() {
	err := s.conn.Close()
	if err != nil {
		s.log.Error(err)
	}
}
