package server

import (
	"context"
	"database/sql"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	cfg *config.Config
	log *zap.SugaredLogger
	//store *Store
	gs *grpc.Server
	//pb.UnimplementedFortedServer
}

func NewServer(cfg *config.Config, log *zap.SugaredLogger, db *sql.DB) *Server {
	return &Server{
		cfg: cfg,
		log: log,
		//store: NewStore(cfg, log, db),
		gs: grpc.NewServer(),
	}
}
func (s *Server) Serve() error {
	addr := net.JoinHostPort("", "50051")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "net.Listen error")
	}
	pb.RegisterFortedServer(s.gs, s)
	s.log.Info("gRPC Proxy Server listens on addr ", addr)
	return s.gs.Serve(lis)
}
func (s *Server) GracefulStop() {
	s.log.Debug("begin gRPC server gracefulStop")
	s.gs.GracefulStop()
	s.log.Debug("end gRPC server gracefulStop")
}

func (s *Server) CreateSurebet(ctx context.Context, request *pb.CreateSurebetRequest) (*pb.CreateSurebetResponse, error) {
	surebet := request.GetSurebet()
	s.log.Info(surebet)
	return &pb.CreateSurebetResponse{}, nil
}

func (s *Server) HelloForted(ctx context.Context, req *pb.HelloFortedRequest) (*pb.HelloFortedResponse, error) {
	msg := req.GetMsg()
	s.log.Info(msg)
	return &pb.HelloFortedResponse{}, nil

	//return &pb.HelloFortedResponse{Msg: msg + " World"}, nil
}
