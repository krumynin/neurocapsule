package server

import (
	"context"
	"go.uber.org/zap"
	"net/http/httptest"
	serverRPC "stash.lamoda.ru/gotools/rpc/server"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/config"
	serverGenerated "stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/generated/server"
)

type Server struct {
	*serverGenerated.Server
}

func NewServer(cfg *config.Config, version string) *Server {
	srv := serverGenerated.New(version, cfg.LogIndex)

	srv.AddOnError(func(ctx context.Context, err error) {
		md := serverRPC.GetSafeMetadataFromContext(ctx)
		if md == nil || md.Logger == nil {
			return
		}

		md.Logger.Debug("Catch error hook", zap.Error(err))
	})

	return &Server{
		Server: srv,
	}
}
func NewTestServer(cfg *config.Config) *httptest.Server {
	srv := NewServer(cfg, "test")
	return httptest.NewServer(srv)
}

func (s Server) Run() {
	s.Server.Run()
}
