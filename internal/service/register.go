package service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	. "github.com/hkobayash/sample-grc/internal/app"
	"github.com/hkobayash/sample-grc/internal/service/imposition"
	"github.com/hkobayash/sample-grc/internal/service/sample"
)

func RegisterAll(server *grpc.Server, app *App) {
	sample.RegisterImpositionServiceServer(server, imposition.NewServer(app))

	healthpb.RegisterHealthServer(server, health.NewServer())
}
