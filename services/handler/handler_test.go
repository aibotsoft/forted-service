package handler

import (
	"github.com/aibotsoft/forted-service/services/store"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var h *Handler

func TestMain(m *testing.M) {
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnectX(cfg)
	sto := store.NewStore(cfg, log, db)
	conf := config_client.New(cfg, log)
	h := New(cfg, log, sto, conf)
	m.Run()
	h.Close()
}

func TestHandler_SendNats(t *testing.T) {
	nc, err := nats.Connect("nats://192.168.1.10:3087")
	assert.NoError(t, err)
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	assert.NoError(t, err)
	defer c.Close()
	err = c.Publish("foo", "Hello World")
	assert.NoError(t, err)
	//t.Log(nc)
}
