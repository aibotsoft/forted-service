package handler

import (
	"context"
	"github.com/aibotsoft/forted-service/services/client"
	"github.com/aibotsoft/forted-service/services/store"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
	"go.uber.org/zap"
)

type Handler struct {
	cfg    *config.Config
	log    *zap.SugaredLogger
	client *client.FortedClient
	store  *store.Store
	Conf   *config_client.ConfClient
}

func (h *Handler) HandleSurebet(ctx context.Context, sur *pb.Surebet) error {
	err := h.store.InsertFullSurebet(ctx, sur)
	if err != nil {
		h.log.Infow("InsertFullSurebet error", "err", err)
		return err
	}

	_, err = h.client.PlaceSurebet(ctx, &pb.PlaceSurebetRequest{Surebet: sur})
	if err != nil {
		h.log.Infow("client.PlaceSurebet error", "err", err)
	}
	return nil
}

func (h *Handler) Close() {
	h.store.Close()
	h.Conf.Close()
	h.client.Close()
}

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, client *client.FortedClient, store *store.Store, conf *config_client.ConfClient) *Handler {
	return &Handler{cfg: cfg, log: log, client: client, store: store, Conf: conf}
}
