package test

import (
	"github.com/lastsweetop/gosip/sipserver"
	"testing"
)

func TestServerStart(t *testing.T) {
	sipsvr := sipserver.NewSipServer(5043)
	sipsvr.Start()
	defer sipsvr.Close()
}
