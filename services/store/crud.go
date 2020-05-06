package store

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func (s *Store) CreateService(ctx context.Context, serviceName string) (int64, error) {
	if len(serviceName) < 3 {
		return 0, errors.Errorf("cannot create service %q", serviceName)
	}
	key := s.FormKey("Service", serviceName)
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateService", &serviceName).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateService error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateSport(ctx context.Context, sportName string, serviceId int64) (int64, error) {
	key := s.FormKey("Sport", sportName, strconv.FormatInt(serviceId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateSport", &sportName, &serviceId).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateSport error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateLeague(ctx context.Context, leagueName string, sportId int64) (int64, error) {
	key := s.FormKey("League", leagueName, strconv.FormatInt(sportId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateLeague", &leagueName, &sportId).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "uspCreateLeague error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateTeam(ctx context.Context, name string, sportId int64) (int64, error) {
	key := s.FormKey("Team", name, strconv.FormatInt(sportId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateTeam", &name, &sportId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateTeam error, TeamName=%v, sportId=%v", &name, &sportId)
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateEvent(ctx context.Context, starts string, homeId int64, awayId int64, leagueId int64) (int64, error) {
	key := s.FormKey("Event", starts, strconv.FormatInt(homeId, 10), strconv.FormatInt(awayId, 10), strconv.FormatInt(leagueId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateEvent", &starts, &homeId, &awayId, &leagueId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateEvent error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateMarket(ctx context.Context, marketName string, eventId int64, url string, num int64) (int64, error) {
	key := s.FormKey("Market", marketName, strconv.FormatInt(eventId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateMarket", &marketName, &eventId, &url, &num).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateMarket error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreatePrice(ctx context.Context, price float64, marketId int64, received string) (int64, error) {
	var id int64
	var createdAt *time.Time
	err := s.db.QueryRowContext(ctx, "dbo.uspCreatePrice",
		sql.Named("Price", &price),
		sql.Named("MarketId", &marketId),
		sql.Named("ReceivedAT", &received)).Scan(&id, &createdAt)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreatePrice error, price=%v, marketId=%v", price, marketId)
	}
	return id, nil
}

func (s *Store) CreateSurebet(ctx context.Context, FortedEventId int64, AMarketId int64, BMarketId int64) (int64, error) {
	key := s.FormKey("Surebet", strconv.FormatInt(FortedEventId, 10), strconv.FormatInt(AMarketId, 10), strconv.FormatInt(BMarketId, 10))
	id, b := s.CheckInCache(ctx, key)
	if b {
		return id, nil
	}
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateSurebet", &FortedEventId, &AMarketId, &BMarketId).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateSurebet error")
	}
	s.cache.Set(key, id, 1)
	return id, nil
}

func (s *Store) CreateLog(ctx context.Context, SurebetId int64, FilterName string, FortedProfit float64, InitiatorNum int64, SkynetId int64, ReceivedAt string) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, "dbo.uspCreateLog", &SurebetId, &FilterName, &FortedProfit, &InitiatorNum, &SkynetId, &ReceivedAt).Scan(&id)
	if err != nil {
		return 0, errors.Wrapf(err, "uspCreateLog error")
	}
	return id, nil
}
