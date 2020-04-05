package server

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

func InitServerHelper(t *testing.T) *Server {
	t.Helper()
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnect(cfg)
	s := NewServer(cfg, log, db)
	return s
}

func TestServer_CreateSurebet(t *testing.T) {
	s := InitServerHelper(t)
	defer s.store.Close()

	//var surebet []*pb.SurebetSide
	req := &pb.CreateSurebetRequest{
		Surebet: &pb.Surebet{
			CreatedAt:    &time.Time{},
			SurebetHash:  "hash",
			Starts:       &time.Time{},
			FortedHome:   "FortedHome",
			FortedAway:   "FortedAway",
			FortedProfit: "6.66",
			Members: []*pb.SurebetSide{
				&pb.SurebetSide{
					Num:         "1",
					ServiceName: "TestService_1",
					SportName:   "TestSport_1",
					LeagueName:  "TestLeague_1",
					Home:        "TestHome_1",
					Away:        "TestAway_1",
					MarketName:  "TestMarket_1",
					Price:       6.66,
					Url:         "http://testurl.loc",
					Initiator:   "true",
				},
				&pb.SurebetSide{
					Num:         "2",
					ServiceName: "TestService_2",
					SportName:   "TestSport_2",
					LeagueName:  "TestLeague_2",
					Home:        "TestHome_2",
					Away:        "TestAway_2",
					MarketName:  "TestMarket_2",
					Price:       6.66,
					Url:         "http://testurl.loc",
					Initiator:   "",
				},
			},
		},
	}
	res, err := s.CreateSurebet(context.Background(), req)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, res, res)

	}
}
