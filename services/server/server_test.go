package server

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"testing"
)

func InitServerHelper(t *testing.T) (*Server, func()) {
	t.Helper()
	cfg := config.New()
	log := logger.New()
	//log.Infow("Begin service", "config", cfg)
	db := sqlserver.MustConnect(cfg)
	closer := func() {
		//log.Info("Closing db")
		if err := db.Close(); err != nil {
			log.Info(err)
		}
	}
	s := NewServer(cfg, log, db)
	return s, closer
}

//func TestServer_CreateSurebet(t *testing.T) {
//	s, closer := InitServerHelper(t)
//	defer closer()
//
//	var surebet []*pb.SurebetSide
//	surebet = append(surebet, &pb.SurebetSide{
//		ServiceName: "Pinnacle",
//	})
//
//	res, err := s.CreateSurebet(context.Background(), &pb.CreateSurebetRequest{
//		Surebet: &pb.Surebet{
//			SurebetSide: surebet,
//		},
//	})
//	assert.NoError(t, err, err)
//	assert.NotEmpty(t, res, res)
//	t.Log("response ", res.GetSurebetId())
//}
