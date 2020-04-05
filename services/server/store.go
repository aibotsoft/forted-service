package server

import (
	"database/sql"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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
}

func (s *Store) CheckInCache(ctx context.Context, name string, keys ...string) (int, bool) {
	if len(keys) == 0 || name == "" {
		return 0, false
	}
	keysJoin := strings.Join(keys[:], ":")
	key := name + ":" + keysJoin
	if get, b := s.cache.Get(key); b {
		if value, ok := get.(int); ok {
			s.log.Debugf("Got %q from cache ", key)
			return value, true
		}
	}
	return 0, false
}

func (s *Store) FormKey(name string, keys ...string) string {
	return name + ":" + strings.Join(keys[:], ":")
}

func (s *Store) SaveInCache(ctx context.Context, key string) (int, bool) {
	if get, b := s.cache.Get(key); b {
		if value, ok := get.(int); ok {
			s.log.Debugf("Got %q from cache ", key)
			return value, true
		}
	}
	return 0, false
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

func (s *Store) CreateFortedEvent(ctx context.Context, starts time.Time, fortedHomeId int, fortedAwayId int) (int, error) {
	key := s.FormKey("FortedEvent", starts.String(), strconv.Itoa(fortedHomeId), strconv.Itoa(fortedAwayId))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateFortedEvent", &starts, &fortedHomeId, &fortedAwayId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateFortedEvent error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateEvent(ctx context.Context, starts time.Time, homeId int, awayId int, leagueId int) (int, error) {
	key := s.FormKey("Event", starts.String(), strconv.Itoa(homeId), strconv.Itoa(awayId), strconv.Itoa(leagueId))
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

func (s *Store) CreateMarket(ctx context.Context, eventId int, market string) (int, error) {
	key := s.FormKey("Market", strconv.Itoa(eventId), market)
	s.log.Debug(key)
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "uspCreateMarket", &eventId, &market).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateMarket error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}
