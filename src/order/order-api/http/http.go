package http

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/rs/cors"
)

type HttpServer interface {
	ListenAndServe() error
}

type httpServer struct {
	appSettings   configuration.AppSettings
	graphQLServer *handler.Server
}

func NewHttpServer(
	appSettings configuration.AppSettings,
	graphQLServer *handler.Server) (HttpServer, error) {
	return &httpServer{
		appSettings:   appSettings,
		graphQLServer: graphQLServer,
	}, nil
}

func (hs *httpServer) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("Omiga - Order", "/graphql"))
	mux.Handle("/graphql", hs.graphQLServer)
	mux.HandleFunc("/health", hs.healthHandler)

	handler := cors.AllowAll().Handler(mux)

	return http.ListenAndServe(hs.appSettings.ListeningInterface, handler)
}

func (hs *httpServer) healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy\n")
}
