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
	"time"
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

func (s *Server) CreateSurebetMany(ctx context.Context, req *pb.CreateSurebetManyRequest) (*pb.CreateSurebetManyResponse, error) {
	start := time.Now()
	var newSurebet int64 = 0
	for _, sur := range req.GetSurebet() {
		s.log.Info(sur)
		logId, err := s.store.InsertFullSurebet(ctx, sur)
		if err != nil {
			s.log.Info(err)
			continue
		}
		s.log.Debug(logId)
		newSurebet++
	}
	s.log.Info("time: ", time.Since(start))
	return &pb.CreateSurebetManyResponse{SurebetCount: newSurebet}, nil
}

//
//func (s *Server) CreateSurebet(ctx context.Context, request *pb.CreateSurebetRequest) (*pb.CreateSurebetResponse, error) {
//	sur := request.GetSurebet()
//	s.log.Debug(sur)
//	if sur.GetSkynetId() != 0 {
//		logId, err := s.store.CheckSkynetId(ctx, sur.GetSkynetId())
//		switch {
//		case err == sql.ErrNoRows:
//			break
//		case err != nil:
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CheckSkynetId")
//		case logId != 0:
//			return nil, status.Errorf(codes.AlreadyExists, "skynetId %v already exists", sur.GetSkynetId())
//		}
//	}
//
//	fortedServiceId, err := s.store.CreateService(ctx, "Forted")
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
//	}
//	fortedSport := "FortedUndefined"
//	if sur.FortedSport != "" {
//		fortedSport = sur.FortedSport
//	}
//	fortedSportId, err := s.store.CreateSport(ctx, fortedSport, fortedServiceId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
//	}
//	fortedLeague := "FortedUndefined"
//	if sur.FortedLeague != "" {
//		fortedLeague = sur.FortedLeague
//	}
//	fortedLeagueId, err := s.store.CreateLeague(ctx, fortedLeague, fortedSportId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
//	}
//
//	fortedHomeId, err := s.store.CreateTeam(ctx, sur.FortedHome, fortedSportId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
//	}
//
//	fortedAwayId, err := s.store.CreateTeam(ctx, sur.FortedAway, fortedSportId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
//	}
//
//	fortedEventId, err := s.store.CreateEvent(ctx, sur.Starts, fortedHomeId, fortedAwayId, fortedLeagueId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
//	}
//	var initiatorId int
//	ids := make([]SurebetIds, 2)
//	for i, ss := range sur.Members {
//		ids[i].ServiceId, err = s.store.CreateService(ctx, ss.ServiceName)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
//		}
//		ids[i].SportId, err = s.store.CreateSport(ctx, ss.SportName, ids[i].ServiceId)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
//		}
//		ids[i].LeagueId, err = s.store.CreateLeague(ctx, ss.LeagueName, ids[i].SportId)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
//		}
//		ids[i].HomeId, err = s.store.CreateTeam(ctx, ss.Home, ids[i].SportId)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
//		}
//		ids[i].AwayId, err = s.store.CreateTeam(ctx, ss.Away, ids[i].SportId)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
//		}
//		ids[i].EventId, err = s.store.CreateEvent(ctx, sur.Starts, ids[i].HomeId, ids[i].AwayId, ids[i].LeagueId)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
//		}
//
//		ids[i].MarketId, err = s.store.CreateMarket(ctx, ss.MarketName, ids[i].EventId, ss.Url, ss.Num)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateMarket")
//		}
//		if ss.Initiator {
//			initiatorId = ids[i].MarketId
//		}
//		ids[i].PriceId, err = s.store.CreatePrice(ctx, ss.Price, ids[i].MarketId, sur.CreatedAt)
//		if err != nil {
//			return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreatePrice")
//		}
//	}
//	surebetId, err := s.store.CreateSurebet(ctx, fortedEventId, ids[0].MarketId, ids[1].MarketId)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateSurebet")
//	}
//	_, err = s.store.CreateLog(ctx, surebetId, sur.FilterName, sur.FortedProfit, initiatorId, sur.GetSkynetId(), sur.CreatedAt)
//	if err != nil {
//		return nil, s.LogAndReturnErr(err, codes.Internal, "store.CreateLog")
//	}
//	s.log.Info(s.store.cache.Metrics)
//
//	return &pb.CreateSurebetResponse{SurebetId: int64(surebetId), SurebetHash: 0}, nil
//}
func (s *Server) LogAndReturnErr(err error, code codes.Code, msg string) error {
	s.log.Error(msg, err)
	return status.Errorf(code, "%v error: %v ", msg, err)
}
