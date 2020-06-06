package middles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	got, err := Convert("1")
	assert.NoError(t, err)
	assert.Equal(t, -0.5, *got.Handicap)
}
