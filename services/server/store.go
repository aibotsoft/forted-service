package server

import (
	"context"
	"database/sql"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
	"time"
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

func (s *Store) CheckInCache(ctx context.Context, key string) (int, bool) {
	if get, b := s.cache.Get(key); b {
		if value, ok := get.(int); ok {
			//s.log.Debugf("Got %q from cache ", key)
			return value, true
		}
	}
	return 0, false
}

func (s *Store) FormKey(name string, keys ...string) string {
	return name + ":" + strings.Join(keys[:], ":")
}

func (s *Store) CreateService(ctx context.Context, serviceName string) (int, error) {
	key := s.FormKey("Service", serviceName)
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateService", &serviceName).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateService error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateSport(ctx context.Context, sportName string, serviceId int) (int, error) {
	key := s.FormKey("Sport", sportName, strconv.Itoa(serviceId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateSport", &sportName, &serviceId).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateSport error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateLeague(ctx context.Context, leagueName string, sportId int) (int, error) {
	key := s.FormKey("League", leagueName, strconv.Itoa(sportId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateLeague", &leagueName, &sportId).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateLeague error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateTeam(ctx context.Context, name string, sportId int) (int, error) {
	key := s.FormKey("Team", name, strconv.Itoa(sportId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateTeam", &name, &sportId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateTeam error, TeamName=%v, sportId=%v", &name, &sportId)
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateFortedEvent(ctx context.Context, starts string, fortedHomeId int, fortedAwayId int) (int, error) {
	key := s.FormKey("FortedEvent", starts, strconv.Itoa(fortedHomeId), strconv.Itoa(fortedAwayId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateFortedEvent", starts, &fortedHomeId, &fortedAwayId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateFortedEvent error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateEvent(ctx context.Context, starts string, homeId int, awayId int, leagueId int) (int, error) {
	key := s.FormKey("Event", starts, strconv.Itoa(homeId), strconv.Itoa(awayId), strconv.Itoa(leagueId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateEvent", &starts, &homeId, &awayId, &leagueId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateEvent error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateMarket(ctx context.Context, marketName string, eventId int, url string, num int64) (int, error) {
	key := s.FormKey("Market", marketName, strconv.Itoa(eventId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateMarket", &marketName, &eventId, &url, &num).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateMarket error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreatePrice(ctx context.Context, price float64, marketId int, received string) (int, error) {
	var id int
	var createdAt *time.Time
	err := s.db.QueryRowContext(ctx, "uspCreatePrice",
		sql.Named("Price", &price),
		sql.Named("MarketId", &marketId),
		sql.Named("ReceivedAT", &received)).Scan(&id, &createdAt)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreatePrice error, price=%v, marketId=%v", price, marketId)
	}
	return id, nil
}

func (s *Store) CreateSurebet(ctx context.Context, FortedEventId int, AMarketId int, BMarketId int) (int, error) {
	key := s.FormKey("Surebet", strconv.Itoa(FortedEventId), strconv.Itoa(AMarketId), strconv.Itoa(BMarketId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateSurebet", &FortedEventId, &AMarketId, &BMarketId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateSurebet error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateLog(ctx context.Context, SurebetId int, FilterName string, FortedProfit string, InitiatorNum int, SkynetId int64, ReceivedAt string) (int, error) {
	var id int
	err := s.db.QueryRowContext(ctx, "uspCreateLog", &SurebetId, &FilterName, &FortedProfit, &InitiatorNum, &SkynetId, &ReceivedAt).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateLog error")
	}
	return id, nil
}

func (s *Store) CheckSkynetId(ctx context.Context, skynetId int64) (int, error) {
	var id int
	err := s.db.QueryRowContext(ctx, "uspCheckSkynetId", &skynetId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Store) LogAndReturnErr(err error, code codes.Code, msg string) error {
	s.log.Error(msg, err)
	return status.Errorf(code, "%v error: %v ", msg, err)
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

func (s *Store) InsertFullSurebet(ctx context.Context, sur *pb.Surebet) (int, error) {
	if sur.GetSkynetId() != 0 {
		logId, err := s.CheckSkynetId(ctx, sur.GetSkynetId())
		switch {
		case err == sql.ErrNoRows:
			break
		case err != nil:
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CheckSkynetId")
		case logId != 0:
			return 0, status.Errorf(codes.AlreadyExists, "skynetId %v already exists", sur.GetSkynetId())
		}
	}
	fortedServiceId, err := s.CreateService(ctx, "Forted")
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
	}
	fortedSport := "FortedUndefined"
	if sur.FortedSport != "" {
		fortedSport = sur.FortedSport
	}
	fortedSportId, err := s.CreateSport(ctx, fortedSport, fortedServiceId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
	}
	fortedLeague := "FortedUndefined"
	if sur.FortedLeague != "" {
		fortedLeague = sur.FortedLeague
	}
	fortedLeagueId, err := s.CreateLeague(ctx, fortedLeague, fortedSportId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
	}

	fortedHomeId, err := s.CreateTeam(ctx, sur.FortedHome, fortedSportId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
	}

	fortedAwayId, err := s.CreateTeam(ctx, sur.FortedAway, fortedSportId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
	}

	fortedEventId, err := s.CreateEvent(ctx, sur.Starts, fortedHomeId, fortedAwayId, fortedLeagueId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
	}

	var initiatorId int
	ids := make([]SurebetIds, 2)
	for i, ss := range sur.Members {
		ids[i].ServiceId, err = s.CreateService(ctx, ss.ServiceName)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateService")
		}
		ids[i].SportId, err = s.CreateSport(ctx, ss.SportName, ids[i].ServiceId)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateSport")
		}
		ids[i].LeagueId, err = s.CreateLeague(ctx, ss.LeagueName, ids[i].SportId)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateLeague")
		}
		ids[i].HomeId, err = s.CreateTeam(ctx, ss.Home, ids[i].SportId)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ids[i].AwayId, err = s.CreateTeam(ctx, ss.Away, ids[i].SportId)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateTeam")
		}
		ids[i].EventId, err = s.CreateEvent(ctx, sur.Starts, ids[i].HomeId, ids[i].AwayId, ids[i].LeagueId)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateEvent")
		}

		ids[i].MarketId, err = s.CreateMarket(ctx, ss.MarketName, ids[i].EventId, ss.Url, ss.Num)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateMarket")
		}
		if ss.Initiator {
			initiatorId = ids[i].MarketId
		}
		ids[i].PriceId, err = s.CreatePrice(ctx, ss.Price, ids[i].MarketId, sur.CreatedAt)
		if err != nil {
			return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreatePrice")
		}
	}
	surebetId, err := s.CreateSurebet(ctx, fortedEventId, ids[0].MarketId, ids[1].MarketId)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateSurebet")
	}
	logId, err := s.CreateLog(ctx, surebetId, sur.FilterName, sur.FortedProfit, initiatorId, sur.GetSkynetId(), sur.CreatedAt)
	if err != nil {
		return 0, s.LogAndReturnErr(err, codes.Internal, "store.CreateLog")

	}
	return logId, nil
}
