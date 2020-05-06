package handler

import (
	"context"
	"github.com/aibotsoft/forted-service/services/client"
	"github.com/aibotsoft/forted-service/services/store"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"go.uber.org/zap"
)

type Handler struct {
	cfg    *config.Config
	log    *zap.SugaredLogger
	client *client.FortedClient
	store  *store.Store
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

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, client *client.FortedClient, store *store.Store) *Handler {
	return &Handler{cfg: cfg, log: log, client: client, store: store}
}
