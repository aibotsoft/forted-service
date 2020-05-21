package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_CreateSurebet(t *testing.T) {
	got, err := s.CreateSurebet(context.Background(), 3, []int64{1, 2, 4})
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	t.Log(got)
}

func TestSplitToString(t *testing.T) {
	a := []int64{1, 2}
	toString := SliceToString(a, "event")
	t.Log(toString)
}
