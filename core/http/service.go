package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/core/log"
)

// ServiceProvider provides a HTTP service.
type ServiceProvider interface {
	// Run runs a new HTTP service.
	Run()
}

// Service is a HTTP service.
type Service struct {
	name        string
	address     string
	handler     http.Handler
	logger      log.Logger
	diContainer di.ContainerProvider
}

// NewService returns a new HTTP service instance.
func NewService(
	name string,
	address string,
	handler http.Handler,
	logger log.Logger,
	diContainer di.ContainerProvider,
) *Service {

	return &Service{
		name:        name,
		logger:      logger,
		handler:     handler,
		address:     address,
		diContainer: diContainer,
	}
}

// Run runs a new HTTP service.
func (s *Service) Run() {

	s.logger.Info("%s service start listening address:%v", s.name, s.address)

	server := http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		s.logger.Info("%s service starting graceful shutdown...", s.name)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			s.logger.Error("%+v", err)
		}

		if err := s.diContainer.Shutdown(); err != nil {
			s.logger.Error("%+v", err)
		}

		s.logger.Info("%s service has been stopped", s.name)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(`%+v`, err)
	}
}
