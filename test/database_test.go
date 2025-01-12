package main

import (
	"blog/config"
	"testing"
)

func TestGetConnection(t *testing.T) {
	_ = config.GetConnection()

}
