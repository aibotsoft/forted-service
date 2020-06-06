package handler

import (
	"github.com/nats-io/nats.go"
)

func (h *Handler) NatsConnect() error {
	h.log.Infow("begin_connect_nats", "url", h.cfg.Broker.Url)
	nc, err := nats.Connect(h.cfg.Broker.Url)
	if err != nil {
		return err
	}
	h.nats, err = nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		return err
	}
	h.log.Info("end_connect_nats")
	return nil
}
func (h *Handler) Publish(subject string, v interface{}) (err error) {
	if h.nats == nil {
		err = h.NatsConnect()
		if err != nil {
			return
		}
	}
	err = h.nats.Publish(subject, v)
	return err
}
