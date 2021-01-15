package handler

import (
	"context"
	"github.com/aibotsoft/forted-service/services/store"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

type Handler struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	store *store.Store
	nats  *nats.EncodedConn
}

func (h *Handler) HandleSurebet(ctx context.Context, sur *pb.Surebet) (err error) {
	start := time.Now()
	err = h.store.InsertFullSurebet(ctx, sur)
	if err != nil {
		h.log.Infow("InsertFullSurebet error", "err", err)
		return err
	}
	sur.Calc.SurebetType = "surebet"
	if len(sur.Members) == 2 {
		var isNotConverted bool
		for i := range sur.Members {
			err := Convert(sur.Members[i])
			if err != nil {
				h.log.Info("convert_error ", err, " market:", sur.Members[i].MarketName, " service:", sur.Members[i].ServiceName, " sport:", sur.FortedSport)
				isNotConverted = true
			}
		}
		CalcMiddle(sur)
		if sur.Calc.MiddleDiff > 0 && !isNotConverted {
			sur.Calc.SurebetType = "middle"
		}
	} else {
		sur.Calc.SurebetType = "3way_surebet"
	}
	err = h.Publish(sur.Calc.SurebetType, sur)
	if err != nil {
		h.log.Info(err)
	} else {
		h.log.Infow(sur.Calc.SurebetType, "t", time.Since(start), "p", sur.FortedProfit,
			"diff", sur.Calc.MiddleDiff, "m1", sur.Members[0].MarketName, "m2", sur.Members[1].MarketName,
			"sp", sur.FortedSport,
			"fid", sur.FortedSurebetId,
			"t1", sur.Members[0].ServiceName,
			"t2", sur.Members[1].ServiceName,
		)
	}
	return nil
}

func (h *Handler) Close() {
	h.store.Close()
	if h.nats != nil {
		h.nats.Close()
	}
}

func New(cfg *config.Config, log *zap.SugaredLogger, store *store.Store) *Handler {
	return &Handler{cfg: cfg, log: log, store: store}
}
