package server

import (
	"context"
	"database/sql"
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
	cfg   *config.Config
	log   *zap.SugaredLogger
	store *Store
	gs    *grpc.Server
	pb.UnimplementedFortedServer
}

func NewServer(cfg *config.Config, log *zap.SugaredLogger, db *sql.DB) *Server {
	return &Server{
		cfg:   cfg,
		log:   log,
		store: NewStore(cfg, log, db),
		gs:    grpc.NewServer(),
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

type SurebetIds struct {
	ServiceId int
	SportId   int
	LeagueId  int
	HomeId    int
	AwayId    int
	EventId   int
	MarketId  int
	PriceId   int
}

func (s *Server) CreateSurebet(ctx context.Context, request *pb.CreateSurebetRequest) (*pb.CreateSurebetResponse, error) {
	sur := request.GetSurebet()
	s.log.Info(sur)

	fortedServiceId, err := s.store.CreateService(ctx, "Forted")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateService error: %v", err)
	}
	fortedSportId, err := s.store.CreateSport(ctx, "FortedSport", fortedServiceId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateSport error: %v", err)
	}

	fortedHomeId, err := s.store.CreateTeam(ctx, sur.FortedHome, fortedSportId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateTeam error: %v", err)
	}

	fortedAwayId, err := s.store.CreateTeam(ctx, sur.FortedAway, fortedSportId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateTeam error: %v", err)
	}

	fortedEventId, err := s.store.CreateFortedEvent(ctx, sur.Starts, fortedHomeId, fortedAwayId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateFortedEvent error: %v", err)
	}
	ids := make([]SurebetIds, 2)
	for i, ss := range sur.Members {
		ids[i].ServiceId, err = s.store.CreateService(ctx, ss.ServiceName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateService error: %v", err)
		}
		ids[i].SportId, err = s.store.CreateSport(ctx, ss.SportName, ids[i].ServiceId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateSport error: %v", err)
		}
		ids[i].LeagueId, err = s.store.CreateLeague(ctx, ss.LeagueName, ids[i].SportId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateLeague error: %v", err)
		}
		ids[i].HomeId, err = s.store.CreateTeam(ctx, ss.Home, ids[i].SportId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateTeam error: %v", err)
		}
		ids[i].AwayId, err = s.store.CreateTeam(ctx, ss.Away, ids[i].SportId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateTeam error: %v", err)
		}
		ids[i].EventId, err = s.store.CreateEvent(ctx, sur.Starts, ids[i].HomeId, ids[i].AwayId, ids[i].LeagueId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateEvent error: %v", err)
		}
		ids[i].MarketId, err = s.store.CreateMarket(ctx, ss.MarketName, ids[i].EventId, ss.Url)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreateMarket error: %v", err)
		}
		ids[i].PriceId, err = s.store.CreatePrice(ctx, ss.Price, ids[i].MarketId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "store.CreatePrice error: %v", err)
		}
		s.log.Debug("ServiceId", i, ids[i].ServiceId)
	}
	surebetId, err := s.store.CreateSurebet(ctx, fortedEventId, ids[0].MarketId, ids[1].MarketId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "store.CreateSurebet error: %v", err)
	}
	return &pb.CreateSurebetResponse{SurebetId: int64(surebetId), SurebetHash: 0}, nil
}
