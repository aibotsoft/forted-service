package store

import (
	"context"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type Store struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	db    *sqlx.DB
	cache *ristretto.Cache
}

func NewStore(cfg *config.Config, log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{log: log, db: db, cache: cache.NewCache(cfg)}
}
func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		s.log.Info(err)
	}
	s.cache.Close()
}

//func (s *Store) CheckSkynetId(ctx context.Context, skynetId int64) (int, error) {
//	var id int
//	err := s.db.QueryRowContext(ctx, "dbo.uspCheckSkynetId", &skynetId).Scan(&id)
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}

func (s *Store) InsertFullSurebet(ctx context.Context, sur *pb.Surebet) error {
	//if sur.GetSkynetId() != 0 {
	//	logId, err := s.CheckSkynetId(ctx, sur.GetSkynetId())
	//	switch {
	//	case err == sql.ErrNoRows:
	//		break
	//	case err != nil:
	//		return s.LogAndReturnErr(err, codes.Internal, "store.CheckSkynetId")
	//	case logId != 0:
	//		return status.Errorf(codes.AlreadyExists, "skynetId %v already exists", sur.GetSkynetId())
	//	}
	//}
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

	//var initiatorId int64
	for _, ss := range sur.Members {
		ss.Forted.ServiceId, err = s.CreateService(ctx, ss.ServiceName)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
		}
		ss.Forted.SportId, err = s.CreateSport(ctx, ss.SportName, ss.Forted.ServiceId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
		}
		ss.Forted.LeagueId, err = s.CreateLeague(ctx, ss.LeagueName, ss.Forted.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
		}
		ss.Forted.HomeId, err = s.CreateTeam(ctx, ss.Home, ss.Forted.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ss.Forted.AwayId, err = s.CreateTeam(ctx, ss.Away, ss.Forted.SportId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ss.Forted.EventId, err = s.CreateEvent(ctx, sur.Starts, ss.Forted.HomeId, ss.Forted.AwayId, ss.Forted.LeagueId)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
		}

		ss.Forted.MarketId, err = s.CreateMarket(ctx, ss.MarketName, ss.Forted.EventId, ss.Url, ss.Num)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreateMarket")
		}
		//if ss.Initiator {
		//	initiatorId = ss.Forted.MarketId
		//}
		ss.Forted.PriceId, err = s.CreatePrice(ctx, ss.Price, ss.Forted.MarketId, sur.CreatedAt)
		if err != nil {
			return s.LogAndReturnErr(err, codes.Internal, "store.CreatePrice")
		}
	}
	//s.log.Info("hello")
	//sur.FortedSurebetId, err = s.CreateSurebet(ctx, fortedEventId, sur.Members[0].Forted.MarketId, sur.Members[1].Forted.MarketId)
	var MarketIdList []int64
	for i := range sur.Members {
		MarketIdList = append(MarketIdList, sur.Members[i].Forted.MarketId)
	}
	//s.log.Info(MarketIdList)

	sur.FortedSurebetId, err = s.CreateSurebet(ctx, fortedEventId, MarketIdList)
	if err != nil {
		return s.LogAndReturnErr(err, codes.Internal, "store.CreateSurebet")
	}
	//sur.LogId, err = s.CreateLog(ctx, sur.FortedSurebetId, sur.FilterName, sur.FortedProfit, initiatorId, sur.GetSkynetId(), sur.CreatedAt)
	//if err != nil {
	//	return s.LogAndReturnErr(err, codes.Internal, "store.CreateLog")
	//}
	return nil
}
