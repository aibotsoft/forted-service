package server

import (
	"context"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func InitStore(t *testing.T) *Store {
	t.Helper()
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnect(cfg)
	store := NewStore(cfg, log, db)
	return store
}
func TestStore_CheckInCache(t *testing.T) {
	s := InitStore(t)
	defer s.Close()
	var dig int
	s.cache.Set("Test:one:0", 66, 1)
	time.Sleep(time.Millisecond * 10)
	got, b := s.CheckInCache(context.Background(), "Test")
	assert.False(t, b)
	got, b = s.CheckInCache(context.Background(), "Test", "one", strconv.Itoa(dig))
	if assert.True(t, b, b) {
		assert.NotEmpty(t, 66, got)
	}
}

func TestStore_CreateService(t *testing.T) {
	s := InitStore(t)
	defer s.Close()
	got, err := s.CreateService(context.Background(), "TestService")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
	got, err = s.CreateService(context.Background(), "")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
}

func TestStore_CreateSport(t *testing.T) {
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(context.Background(), "TestService")
	if assert.NoError(t, err, err) {
		got, err := s.CreateSport(context.Background(), "TestSport", serviceId)
		if assert.NoError(t, err, err) {
			assert.NotEmpty(t, got, got)
		}
	}
}

func TestStore_CreateLeague(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "TestService")
	if assert.NoError(t, err, err) {
		sportId, err := s.CreateSport(ctx, "TestSport", serviceId)
		if assert.NoError(t, err) {
			got, err := s.CreateLeague(ctx, "TestLeague", sportId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, got, got)
			}
		}
	}
}

func TestStore_CreateTeam(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "TestService")
	if assert.NoError(t, err, err) {
		sportId, err := s.CreateSport(ctx, "TestSport", serviceId)
		if assert.NoError(t, err) {
			got, err := s.CreateTeam(ctx, "TestTeam", sportId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, got, got)
			}
		}
	}
}

func TestStore_CreateFortedEvent(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "Forted")
	sportId, err := s.CreateSport(ctx, "FortedSport", serviceId)
	fortedHomeId, err := s.CreateTeam(ctx, "FortedHome", sportId)
	fortedAwayId, err := s.CreateTeam(ctx, "FortedAway", sportId)

	got, err := s.CreateFortedEvent(ctx, &time.Time{}, fortedHomeId, fortedAwayId)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
}

func TestStore_CreateEvent(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "TestService")
	if assert.NoError(t, err) {
		sportId, err := s.CreateSport(ctx, "TestSport", serviceId)
		if assert.NoError(t, err) {
			leagueId, err := s.CreateLeague(ctx, "TestLeague", sportId)
			if !assert.NoError(t, err) {
				return
			}
			homeId, err := s.CreateTeam(ctx, "TestHome", sportId)
			if !assert.NoError(t, err) {
				return
			}
			awayId, err := s.CreateTeam(ctx, "TestAway", sportId)
			if !assert.NoError(t, err) {
				return
			}
			eventId, err := s.CreateEvent(ctx, &time.Time{}, homeId, awayId, leagueId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, eventId, eventId)
			}
		}
	}
}

func TestStore_CreateMarket(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "TestService")
	if assert.NoError(t, err) {
		sportId, err := s.CreateSport(ctx, "TestSport", serviceId)
		if assert.NoError(t, err) {
			leagueId, err := s.CreateLeague(ctx, "TestLeague", sportId)
			if !assert.NoError(t, err) {
				return
			}
			homeId, err := s.CreateTeam(ctx, "TestHome", sportId)
			if !assert.NoError(t, err) {
				return
			}
			awayId, err := s.CreateTeam(ctx, "TestAway", sportId)
			if !assert.NoError(t, err) {
				return
			}
			eventId, err := s.CreateEvent(ctx, &time.Time{}, homeId, awayId, leagueId)
			if !assert.NoError(t, err) {
				return
			}
			marketId, err := s.CreateMarket(ctx, "TestMarket", eventId, "http://testurl.loc")
			if assert.NoError(t, err) {
				assert.NotEmpty(t, marketId, marketId)
			}
		}
	}
}

func TestStore_CreatePrice(t *testing.T) {
	ctx := context.Background()
	s := InitStore(t)
	defer s.Close()
	serviceId, err := s.CreateService(ctx, "TestService2")
	if assert.NoError(t, err) {
		sportId, err := s.CreateSport(ctx, "TestSport2", serviceId)
		if assert.NoError(t, err) {
			leagueId, err := s.CreateLeague(ctx, "TestLeague2", sportId)
			if !assert.NoError(t, err) {
				return
			}
			homeId, err := s.CreateTeam(ctx, "TestHome2", sportId)
			if !assert.NoError(t, err) {
				return
			}
			awayId, err := s.CreateTeam(ctx, "TestAway2", sportId)
			if !assert.NoError(t, err) {
				return
			}
			eventId, err := s.CreateEvent(ctx, &time.Time{}, homeId, awayId, leagueId)
			if !assert.NoError(t, err) {
				return
			}
			marketId, err := s.CreateMarket(ctx, "TestMarket2", eventId, "")
			if !assert.NoError(t, err) {
				return
			}
			// Проверяем CreatePrice, заодно проверим чтобы вставлялись дубли прайсов
			firstPriceId, err := s.CreatePrice(ctx, 6.66666, marketId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, firstPriceId, firstPriceId)
			}
			// Второй прайс на этот же маркет
			secondPriceId, err := s.CreatePrice(ctx, 9.99999, marketId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, secondPriceId, secondPriceId)
				assert.NotEqual(t, firstPriceId, secondPriceId)
			}
			// Повторяем тот же прайс, должен вернуть тот же Id
			thirdPriceId, err := s.CreatePrice(ctx, 9.99999, marketId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, thirdPriceId, thirdPriceId)
				assert.Equal(t, thirdPriceId, thirdPriceId)
			}
			// А теперь четвертый прайс, который равен первом. Он должен быть вставлен (новый Id), так как он свежее
			fourthPriceId, err := s.CreatePrice(ctx, 6.66666, marketId)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, fourthPriceId, fourthPriceId)
				assert.NotEqual(t, firstPriceId, fourthPriceId, "А теперь четвертый прайс, который равен первому. Он должен быть вставлен (новый Id), так как он свежее")
				assert.NotEqual(t, firstPriceId, thirdPriceId)
			}
		}
	}
}

//func TestStore_CreateSurebet(t *testing.T) {
//
//}
