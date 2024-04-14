package main_test

import (
	"fmt"
	"github.com/mathieuhays/go-specs-greet/adapters"
	"github.com/mathieuhays/go-specs-greet/adapters/grpcserver"
	"github.com/mathieuhays/go-specs-greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
