package server

import (
	"context"
	"github.com/aibotsoft/forted-service/services/handler"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	cfg     *config.Config
	log     *zap.SugaredLogger
	handler *handler.Handler
	gs      *grpc.Server
	pb.UnimplementedFortedServer
}

func (s *Server) CreateSurebet(ctx context.Context, request *pb.CreateSurebetRequest) (*pb.CreateSurebetResponse, error) {
	sur := request.GetSurebet()
	err := s.handler.HandleSurebet(ctx, sur)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "InsertFullSurebet error")
	}
	return &pb.CreateSurebetResponse{SurebetId: sur.SurebetId}, nil
}

func NewServer(cfg *config.Config, log *zap.SugaredLogger, handler *handler.Handler) *Server {
	return &Server{cfg: cfg, log: log, handler: handler, gs: grpc.NewServer()}
}
func (s *Server) Serve() error {
	addr, err := s.handler.Conf.GetGrpcAddr(context.Background(), s.cfg.Service.Name)
	if err != nil {
		s.log.Panic(err)
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "net.Listen error")
	}
	pb.RegisterFortedServer(s.gs, s)
	s.log.Info("gRPC Server listens on addr ", addr)
	return s.gs.Serve(lis)
}

func (s *Server) Close() {
	s.log.Debug("begin gRPC server gracefulStop")
	s.gs.GracefulStop()
	s.handler.Close()
	s.log.Debug("end gRPC server gracefulStop")
}

//func (s *Server) CreateSurebetMany(ctx context.Context, req *pb.CreateSurebetManyRequest) (*pb.CreateSurebetManyResponse, error) {
//	start := time.Now()
//	var surebetCount int64 = 0
//	for _, sur := range req.GetSurebet() {
//		err := s.store.InsertFullSurebet(ctx, sur)
//		if err != nil {
//			s.log.Info(err)
//			continue
//		}
//		s.log.Debug(sur.LogId)
//		surebetCount++
//	}
//	s.log.Info("time: ", time.Since(start))
//	return &pb.CreateSurebetManyResponse{SurebetCount: surebetCount}, nil
//}
