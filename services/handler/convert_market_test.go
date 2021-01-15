package handler

import (
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	side := &pb.SurebetSide{
		MarketName: "ИТ2М(0,5)",
	}
	err := Convert(side)
	if assert.NoError(t, err) {
		t.Log(side.Market)
	}
}
