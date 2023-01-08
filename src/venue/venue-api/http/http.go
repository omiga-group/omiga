package http

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type HttpServer interface {
	ListenAndServe() error
}

type httpServer struct {
	logger        *zap.SugaredLogger
	appConfig     configuration.AppConfig
	graphQLServer *handler.Server
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	graphQLServer *handler.Server) (HttpServer, error) {
	return &httpServer{
		appConfig:     appConfig,
		graphQLServer: graphQLServer,
		logger:        logger,
	}, nil
}

func (hs *httpServer) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("Omiga - Venue", "/graphql"))
	mux.Handle("/graphql", hs.graphQLServer)
	mux.HandleFunc("/health", hs.healthHandler)

	handler := cors.AllowAll().Handler(mux)

	hs.logger.Infof("Listening on: %s", hs.appConfig.ListeningInterface)

	return http.ListenAndServe(hs.appConfig.ListeningInterface, handler)
}

func (hs *httpServer) healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy\n")
}
