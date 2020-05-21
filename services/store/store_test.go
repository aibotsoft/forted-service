package store

import (
	"context"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var s *Store

func TestMain(m *testing.M) {
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnectX(cfg)
	s = NewStore(cfg, log, db)
	m.Run()
	s.Close()
}

func surebetHelper(t *testing.T) *pb.Surebet {
	t.Helper()
	return &pb.Surebet{
		CreatedAt:    time.Time{}.Format(time.RFC3339Nano),
		Starts:       time.Time{}.Format(time.RFC3339Nano),
		FortedHome:   "FortedHome",
		FortedAway:   "FortedAway",
		FortedProfit: 6.66,
		Members: []*pb.SurebetSide{{
			Num:         1,
			ServiceName: "TestService_1",
			SportName:   "TestSport_1",
			LeagueName:  "TestLeague_1",
			Home:        "TestHome_1",
			Away:        "TestAway_1",
			MarketName:  "TestMarket_1",
			Price:       6.66,
			Url:         "http://testurl.loc",
			Initiator:   true,
		}, {Num: 2,
			ServiceName: "TestService_2",
			SportName:   "TestSport_2",
			LeagueName:  "TestLeague_2",
			Home:        "TestHome_2",
			Away:        "TestAway_2",
			MarketName:  "TestMarket_2",
			Price:       6.66,
			Url:         "http://testurl.loc",
			Initiator:   false,
		}}}
}
func IfErrorFail(t *testing.T, err error) {
	t.Helper()
	if assert.NoError(t, err) {
	} else {
		t.FailNow()
	}
}
func TestStore_CreateService(t *testing.T) {
	service := surebetHelper(t).Members[0].ServiceName
	got, err := s.CreateService(context.Background(), service)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
	_, err = s.CreateService(context.Background(), "")
	assert.Error(t, err)
}

func TestStore_CreateSport(t *testing.T) {
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	serviceId, err := s.CreateService(context.Background(), service)
	IfErrorFail(t, err)
	got, err := s.CreateSport(context.Background(), sport, serviceId)
	IfErrorFail(t, err)
	assert.NotEmpty(t, got, got)
}

func TestStore_CreateLeague(t *testing.T) {
	ctx := context.Background()
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	league := surebetHelper(t).Members[0].LeagueName
	serviceId, err := s.CreateService(ctx, service)
	IfErrorFail(t, err)
	sportId, err := s.CreateSport(ctx, sport, serviceId)
	IfErrorFail(t, err)
	got, err := s.CreateLeague(ctx, league, sportId)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
}

func TestStore_CreateTeam(t *testing.T) {
	ctx := context.Background()
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	home := surebetHelper(t).Members[0].Home
	serviceId, err := s.CreateService(ctx, service)
	IfErrorFail(t, err)
	sportId, err := s.CreateSport(ctx, sport, serviceId)
	IfErrorFail(t, err)
	got, err := s.CreateTeam(ctx, home, sportId)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
}

func TestStore_CreateEvent_Forted(t *testing.T) {
	ctx := context.Background()
	starts := surebetHelper(t).Starts
	serviceId, err := s.CreateService(ctx, "Forted")
	sportId, err := s.CreateSport(ctx, "FortedSport", serviceId)
	fortedHomeId, err := s.CreateTeam(ctx, surebetHelper(t).FortedHome, sportId)
	fortedAwayId, err := s.CreateTeam(ctx, surebetHelper(t).FortedAway, sportId)
	fortedLeagueId, err := s.CreateLeague(ctx, "FortedUndefined", sportId)
	got, err := s.CreateEvent(ctx, starts, fortedHomeId, fortedAwayId, fortedLeagueId)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got, got)
	}
}

func TestStore_CreateEvent(t *testing.T) {
	ctx := context.Background()
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	league := surebetHelper(t).Members[0].LeagueName
	home := surebetHelper(t).Members[0].Home
	away := surebetHelper(t).Members[0].Away
	starts := surebetHelper(t).Starts

	serviceId, err := s.CreateService(ctx, service)
	IfErrorFail(t, err)
	sportId, err := s.CreateSport(ctx, sport, serviceId)
	IfErrorFail(t, err)
	leagueId, err := s.CreateLeague(ctx, league, sportId)
	IfErrorFail(t, err)
	homeId, err := s.CreateTeam(ctx, home, sportId)
	IfErrorFail(t, err)
	awayId, err := s.CreateTeam(ctx, away, sportId)
	IfErrorFail(t, err)
	eventId, err := s.CreateEvent(ctx, starts, homeId, awayId, leagueId)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, eventId, eventId)
	}
}
func TestStore_CreateMarket(t *testing.T) {
	ctx := context.Background()
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	league := surebetHelper(t).Members[0].LeagueName
	home := surebetHelper(t).Members[0].Home
	away := surebetHelper(t).Members[0].Away
	market := surebetHelper(t).Members[0].MarketName
	url := surebetHelper(t).Members[0].Url
	num := surebetHelper(t).Members[0].Num
	starts := surebetHelper(t).Starts
	serviceId, err := s.CreateService(ctx, service)
	IfErrorFail(t, err)
	sportId, err := s.CreateSport(ctx, sport, serviceId)
	IfErrorFail(t, err)
	leagueId, err := s.CreateLeague(ctx, league, sportId)
	IfErrorFail(t, err)
	homeId, err := s.CreateTeam(ctx, home, sportId)
	IfErrorFail(t, err)
	awayId, err := s.CreateTeam(ctx, away, sportId)
	IfErrorFail(t, err)
	eventId, err := s.CreateEvent(ctx, starts, homeId, awayId, leagueId)
	IfErrorFail(t, err)
	marketId, err := s.CreateMarket(ctx, market, eventId, url, num)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, marketId, marketId)
	}
}

func TestStore_CreatePrice(t *testing.T) {
	ctx := context.Background()
	service := surebetHelper(t).Members[0].ServiceName
	sport := surebetHelper(t).Members[0].SportName
	league := surebetHelper(t).Members[0].LeagueName
	home := surebetHelper(t).Members[0].Home
	away := surebetHelper(t).Members[0].Away
	market := surebetHelper(t).Members[0].MarketName
	url := surebetHelper(t).Members[0].Url
	num := surebetHelper(t).Members[0].Num
	starts := surebetHelper(t).Starts
	createdAt := surebetHelper(t).CreatedAt
	serviceId, err := s.CreateService(ctx, service)
	IfErrorFail(t, err)
	sportId, err := s.CreateSport(ctx, sport, serviceId)
	IfErrorFail(t, err)
	leagueId, err := s.CreateLeague(ctx, league, sportId)
	IfErrorFail(t, err)
	homeId, err := s.CreateTeam(ctx, home, sportId)
	IfErrorFail(t, err)
	awayId, err := s.CreateTeam(ctx, away, sportId)
	IfErrorFail(t, err)
	eventId, err := s.CreateEvent(ctx, starts, homeId, awayId, leagueId)
	IfErrorFail(t, err)
	marketId, err := s.CreateMarket(ctx, market, eventId, url, num)
	IfErrorFail(t, err)
	// Проверяем CreatePrice, заодно проверим чтобы вставлялись дубли прайсов
	firstPriceId, err := s.CreatePrice(ctx, 6.66666, marketId, createdAt)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, firstPriceId, firstPriceId)
	}
	// Второй прайс на этот же маркет
	secondPriceId, err := s.CreatePrice(ctx, 9.99999, marketId, createdAt)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, secondPriceId, secondPriceId)
		assert.NotEqual(t, firstPriceId, secondPriceId)
	}
	// Повторяем тот же прайс, должен вернуть тот же Id
	thirdPriceId, err := s.CreatePrice(ctx, 9.99999, marketId, createdAt)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, thirdPriceId, thirdPriceId)
		assert.Equal(t, secondPriceId, thirdPriceId)
	}
	//А теперь четвертый прайс, который равен первом. Он должен быть вставлен (новый Id), так как он свежее
	fourthPriceId, err := s.CreatePrice(ctx, 6.66666, marketId, createdAt)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, fourthPriceId, fourthPriceId)
		assert.NotEqual(t, firstPriceId, fourthPriceId, "А теперь четвертый прайс, который равен первому. Он должен быть вставлен (новый Id), так как он свежее")
		assert.NotEqual(t, firstPriceId, thirdPriceId)
	}
}

func TestStore_InsertFullSurebet(t *testing.T) {
	ctx := context.Background()
	sur := surebetHelper(t)
	err := s.InsertFullSurebet(ctx, sur)
	IfErrorFail(t, err)
	t.Log(sur)
}
