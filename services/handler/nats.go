package handler

import (
	"github.com/nats-io/nats.go"
	"time"
)

func (h *Handler) NatsConnect() error {
	h.log.Info("begin_connect_nats")
	nc, err := nats.Connect("nats://192.168.1.10:30873")
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
	start := time.Now()
	if h.nats == nil {
		err = h.NatsConnect()
		if err != nil {
			return
		}
	}
	err = h.nats.Publish(subject, v)
	h.log.Infow("Publish_done", "time", time.Since(start))
	return err
}
