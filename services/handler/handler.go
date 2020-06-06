package handler

import (
	"context"
	"github.com/aibotsoft/forted-service/services/client"
	"github.com/aibotsoft/forted-service/services/middles"
	"github.com/aibotsoft/forted-service/services/store"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

type Handler struct {
	cfg    *config.Config
	log    *zap.SugaredLogger
	client *client.FortedClient
	store  *store.Store
	Conf   *config_client.ConfClient
	nats   *nats.EncodedConn
}

func (h *Handler) HandleSurebet(ctx context.Context, sur *pb.Surebet) error {
	start := time.Now()
	err := h.store.InsertFullSurebet(ctx, sur)
	if err != nil {
		h.log.Infow("InsertFullSurebet error", "err", err)
		return err
	}
	if len(sur.Members) == 2 {
		diff, err := middles.CalcMiddle(sur.Members[0].MarketName, sur.Members[1].MarketName)
		if err != nil {
			h.log.Error(err)
			return nil
		}
		if diff == 0 {
			err := h.Publish("surebet", sur)
			if err != nil {
				h.log.Info(err)
			}

			_, err = h.client.PlaceSurebet(ctx, &pb.PlaceSurebetRequest{Surebet: sur})
			if err != nil {
				h.log.Infow("client.PlaceSurebet error", "err", err)
			}
			h.log.Infow("got_surebet", "time", time.Since(start), "fid", sur.FortedSurebetId, "profit", sur.FortedProfit)
		} else {
			err := h.Publish("middle", sur)
			if err != nil {
				h.log.Info(err)
			}
			h.log.Infow("got_middle", "time", time.Since(start), "fid", sur.FortedSurebetId, "profit", sur.FortedProfit,
				"diff", diff, "m1", sur.Members[0].MarketName, "m2", sur.Members[1].MarketName, "sport", sur.FortedSport)
		}
	} else {
		h.log.Info("got_3way_surebet, not sended to surebet service", "time", time.Since(start), "fid", sur.FortedSurebetId, "profit", sur.FortedProfit)
	}
	return nil
}

func (h *Handler) Close() {
	h.store.Close()
	h.Conf.Close()
	h.client.Close()
	h.nats.Close()
}

func New(cfg *config.Config, log *zap.SugaredLogger, client *client.FortedClient, store *store.Store, conf *config_client.ConfClient) *Handler {
	return &Handler{cfg: cfg, log: log, client: client, store: store, Conf: conf}
}
