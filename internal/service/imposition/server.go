package imposition

import (
	"context"

	. "github.com/hkobayash/sample-grc/internal/app"
	"github.com/hkobayash/sample-grc/internal/service/sample"
)

type server struct {
	*App
}

func NewServer(app *App) sample.ImpositionServiceServer {
	return &server{
		App: app,
	}
}

func (s *server) CreateItem(ctx context.Context, req *sample.CreateItemRequest) (*sample.CreateItemResponse, error) {
	return &sample.CreateItemResponse{}, nil
}
