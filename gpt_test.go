package gpt

import (
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriest(t *testing.T) {
	gone.Test(func(gpt *chatGPTClient) {
		assert.NotNil(t, gpt.Client)
	}, Priest)
}
