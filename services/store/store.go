package store

import (
	"context"
	"database/sql"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Store struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	db    *sql.DB
	cache *ristretto.Cache
}

func NewStore(cfg *config.Config, log *zap.SugaredLogger, db *sql.DB) *Store {
	return &Store{log: log, db: db, cache: cache.NewCache(cfg)}
}
func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		s.log.Fatal(err)
	}
	s.cache.Close()
}

func (s *Store) CheckSkynetId(ctx context.Context, skynetId int64) (int, error) {
	var id int
	err := s.db.QueryRowContext(ctx, "dbo.uspCheckSkynetId", &skynetId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) InsertFullSurebet(ctx context.Context, sur *pb.Surebet) error {
	if sur.GetSkynetId() != 0 {
		logId, err := s.CheckSkynetId(ctx, sur.GetSkynetId())
		switch {
		case err == sql.ErrNoRows:
			break
		case err != nil:
			return s.LogAndReturnErr(err, codes.Internal, "store.CheckSkynetId")
		case logId != 0:
			return status.Errorf(codes.AlreadyExists, "skynetId %v already exists", sur.GetSkynetId())
		}
	}
	fortedServiceId, err := s.CreateService(ctx, "Forted")
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
	}
	fortedSport := "FortedUndefined"
	if sur.FortedSport != "" {
		fortedSport = sur.FortedSport
	}
	fortedSportId, err := s.CreateSport(ctx, fortedSport, fortedServiceId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
	}
	fortedLeague := "FortedUndefined"
	if sur.FortedLeague != "" {
		fortedLeague = sur.FortedLeague
	}
	fortedLeagueId, err := s.CreateLeague(ctx, fortedLeague, fortedSportId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
	}

	fortedHomeId, err := s.CreateTeam(ctx, sur.FortedHome, fortedSportId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
	}

	fortedAwayId, err := s.CreateTeam(ctx, sur.FortedAway, fortedSportId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
	}

	fortedEventId, err := s.CreateEvent(ctx, sur.Starts, fortedHomeId, fortedAwayId, fortedLeagueId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
	}

	var initiatorId int64
	for _, ss := range sur.Members {
		ss.ServiceId, err = s.CreateService(ctx, ss.ServiceName)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
		}
		ss.SportId, err = s.CreateSport(ctx, ss.SportName, ss.ServiceId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
		}
		ss.LeagueId, err = s.CreateLeague(ctx, ss.LeagueName, ss.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
		}
		ss.HomeId, err = s.CreateTeam(ctx, ss.Home, ss.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ss.AwayId, err = s.CreateTeam(ctx, ss.Away, ss.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ss.EventId, err = s.CreateEvent(ctx, sur.Starts, ss.HomeId, ss.AwayId, ss.LeagueId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
		}

		ss.MarketId, err = s.CreateMarket(ctx, ss.MarketName, ss.EventId, ss.Url, ss.Num)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateMarket")
		}
		if ss.Initiator {
			initiatorId = ss.MarketId
		}
		ss.PriceId, err = s.CreatePrice(ctx, ss.Price, ss.MarketId, sur.CreatedAt)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreatePrice")
		}
	}
	sur.FortedSurebetId, err = s.CreateSurebet(ctx, fortedEventId, sur.Members[0].MarketId, sur.Members[1].MarketId)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateSurebet")
	}
	sur.LogId, err = s.CreateLog(ctx, sur.FortedSurebetId, sur.FilterName, sur.FortedProfit, initiatorId, sur.GetSkynetId(), sur.CreatedAt)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateLog")

	}
	return nil
}
