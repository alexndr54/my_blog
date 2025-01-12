package main

import (
	"blog/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddRedisSession(t *testing.T) {
	ctx, _ := config.GetFiberConfig()
	err := config.SetSession(ctx, "demo", "demo")
	assert.Nil(t, err)

}
