package client

import (
	"context"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type FortedClient struct {
	log  *zap.SugaredLogger
	conn *grpc.ClientConn
	pb.FortedClient
}

func NewFortedClient(cfg *config.Config, log *zap.SugaredLogger) *FortedClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Infow("connecting to surebet-service...", "addr", cfg.SurebetService.GrpcPort)
	conn, err := grpc.DialContext(ctx, net.JoinHostPort("", cfg.SurebetService.GrpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicw("NewFortedClient DialContext error", "addr", cfg.SurebetService.GrpcPort, "err", err)
	}
	return &FortedClient{
		log:          log,
		conn:         conn,
		FortedClient: pb.NewFortedClient(conn),
	}
}

func (s *FortedClient) Close() {
	err := s.conn.Close()
	if err != nil {
		s.log.Error(err)
	}
}
