package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Host   string
	Port   string
	App    *fiber.App
	logger *zerolog.Logger
}

type Options func(*Server)

func WithHost(host string) Options {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port string) Options {
	return func(s *Server) {
		s.Port = port
	}
}

func WithApp(app *fiber.App) Options {
	return func(s *Server) {
		s.App = app
	}
}

func WithLogger(logger *zerolog.Logger) Options {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) Connect() error {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		s.logger.Info().
			Str("signal", sig.String()).Msg("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.App.ShutdownWithContext(ctx); err != nil {
			s.logger.Error().Err(err).Msg("server shutdown failed")
		}
	}()

	s.logger.Info().Str("addr", addr).Msg("starting fiber server")

	return s.App.Listen(addr)
}

func NewServer(opts ...Options) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
