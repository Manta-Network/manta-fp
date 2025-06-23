package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/Manta-Network/manta-fp/metrics"
	fpcfg "github.com/Manta-Network/manta-fp/symbiotic-fp/config"
	"github.com/Manta-Network/manta-fp/symbiotic-fp/router"

	"github.com/gin-gonic/gin"
	"github.com/lightningnetwork/lnd/kvdb"
	"github.com/lightningnetwork/lnd/signal"
	"go.uber.org/zap"
)

// Server is the main daemon construct for the Finality Provider server. It handles
// spinning up the RPC sever, the database, and any other components that the
// Taproot Asset server needs to function.
type Server struct {
	started int32

	cfg    *fpcfg.Config
	logger *zap.Logger

	db          kvdb.Backend
	interceptor signal.Interceptor

	metricsServer *metrics.Server
	httpServer    *http.Server

	quit chan struct{}
}

// NewFinalityproviderServer creates a new server with the given config.
func NewFinalityProviderServer(cfg *fpcfg.Config, l *zap.Logger, db kvdb.Backend, sig signal.Interceptor) *Server {
	return &Server{
		cfg:         cfg,
		logger:      l,
		db:          db,
		interceptor: sig,
		quit:        make(chan struct{}, 1),
	}
}

func (s *Server) StartFinalityProviderServer() error {
	if atomic.AddInt32(&s.started, 1) != 1 {
		return nil
	}

	// Start the metrics server.
	promAddr, err := s.cfg.Metrics.Address()
	if err != nil {
		return fmt.Errorf("failed to get prometheus address: %w", err)
	}
	s.metricsServer = metrics.Start(promAddr, s.logger)

	// Start the http server
	registry := router.NewRegistry()
	r := gin.Default()
	registry.Register(r)
	apiAddress, err := s.cfg.Api.Address()
	if err != nil {
		return fmt.Errorf("failed to get http address: %w", err)
	}

	var httpServer *http.Server
	httpServer = &http.Server{
		Addr:    apiAddress,
		Handler: r,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("api server starts failed", zap.String("err", err.Error()))
		}
	}()
	s.httpServer = httpServer

	s.logger.Info("success to start finality Provider server!")

	return nil
}

// RunUntilShutdown runs the main EOTS manager server loop until a signal is
// received to shut down the process.
func (s *Server) RunUntilShutdown() error {

	defer func() {
		s.logger.Info("Shutdown complete")
	}()

	defer func() {
		s.logger.Info("Closing database...")
		if err := s.db.Close(); err != nil {
			s.logger.Error(fmt.Sprintf("Failed to close database: %v", err)) // Log the error
		} else {
			s.logger.Info("Database closed")
		}
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			s.logger.Error(fmt.Sprintf("Failed to close http server: %v", err)) // Log the error
		}
		s.metricsServer.Stop(context.Background())
		s.logger.Info("Metrics server stopped")
	}()

	// Wait for shutdown signal from either a graceful server stop or from
	// the interrupt handler.
	<-s.interceptor.ShutdownChannel()

	return nil
}
