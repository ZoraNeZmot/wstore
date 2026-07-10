package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"wstore/internal/config"
	"wstore/internal/handlers"
	"wstore/internal/log"

	"github.com/sirupsen/logrus"
)

type cleanupFunc func() error

type App struct {
	Config *config.Config
	Server *http.Server

	cleanupFuncs []cleanupFunc // cleanup functions
	ctxStop      context.CancelFunc
}

func New() *App {
	conf := config.GetConfig()
	return &App{
		Config: conf,
		Server: &http.Server{
			Addr:      net.JoinHostPort(conf.ServerConfig.Host, conf.ServerConfig.Port),
			TLSConfig: conf.ServerConfig.TLSConfig,
		},
	}
}

func (a *App) Init() error {
	a.cleanupFuncs = append(a.cleanupFuncs, log.Init())

	logrus.Tracef("Initializing app with config: %v", a.Config)
	logrus.Infof("Current log level: %v", logrus.GetLevel())

	ctx, ctxStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	a.ctxStop = ctxStop
	go a.waitForShutdown(ctx)

	innerMux := http.NewServeMux()
	outerMux := http.NewServeMux()
	if err := handlers.RegisterSwagger(outerMux); err != nil {
		return err
	}
	if err := handlers.RegisterRoutes(innerMux); err != nil {
		return err
	}
	basePath := strings.TrimRight(a.Config.ServerConfig.BasePath, "/")
	outerMux.Handle(basePath+"/", http.StripPrefix(basePath, innerMux))
	a.Server.Handler = outerMux

	return nil
}

func (a *App) Run() error {
	defer a.ctxStop()
	if a.Server.TLSConfig != nil {
		logrus.Infof("Starting server on %s with TLS", a.Server.Addr)
		if err := a.Server.ListenAndServeTLS(a.Config.ServerConfig.CertFile, a.Config.ServerConfig.KeyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	} else {
		logrus.Infof("Starting server on %s without TLS", a.Server.Addr)
		if err := a.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}
	return nil
}

func (a *App) Stop() error {
	return a.cleanup()
}

func (a *App) cleanup() error {
	if a.Server != nil {
		if err := a.Server.Shutdown(context.Background()); err != nil {
			logrus.WithError(err).Errorf("Failed to shutdown server")
		}
	}
	for _, cleanup := range a.cleanupFuncs {
		if err := cleanup(); err != nil {
			logrus.WithError(err).Errorf("Failed to cleanup")
		}
	}
	return nil
}

func (a *App) waitForShutdown(ctx context.Context) {
	<-ctx.Done()
	logrus.Info("Shutting down...")
	if err := a.Stop(); err != nil {
		logrus.WithError(err).Errorf("Failed to stop app")
		os.Exit(1)
	}
	os.Exit(0)
}

func (a *App) registerHandlers(handlers handlers.Handlers) error {
	return handlers.RegisterRoutes(a.Server.Handler.(*http.ServeMux))
}
