package graph

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	http      *http.Server
	resolver  generated.ResolverRoot
	runErr    error
	readiness bool
}

func CreateAndRun(port int, resolver generated.ResolverRoot) *Server {
	service := &Server{
		http: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
		resolver: resolver,
	}

	service.run()

	return service
}

func (s *Server) setupHandlers() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: s.resolver,
	}))

	h := http.NewServeMux()
	h.Handle("/query", correlationID(loggingMiddleware(srv)))
	s.http.Handler = h
}

func (s *Server) run() {
	log.Info("server is running")

	go func() {
		log.Debug("graphql service: addr=", s.http.Addr)
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.runErr = err
			log.WithError(err).Error("graphql service")
		}
	}()
	s.readiness = true
}

func (s *Server) Close(ctx context.Context) {
	if err := s.http.Shutdown(ctx); err != nil {
		log.WithError(err).Error("stopping graphql service")
	}
	log.Info("graphql service stopped")
}

func (s *Server) HealthCheck() error {
	if !s.readiness {
		return errors.New("http service isn't ready yet")
	}
	if s.runErr != nil {
		return errors.New("http service: run issue")
	}
	return nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.
			WithField("method", req.Method).
			WithField("uri", req.URL.RequestURI()).
			Trace("new graphql request")

		next.ServeHTTP(w, req)
	})
}

func correlationID(next http.Handler) http.Handler {
	const key = "X-Correlation-Id"

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		correlationID := uuid.NewV4().String()

		ctx = context.WithValue(ctx, key, correlationID) // nolint
		ctx = metadata.AppendToOutgoingContext(ctx, key, correlationID)

		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
