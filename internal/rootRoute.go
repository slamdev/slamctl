package internal

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"slamctl/internal/daemon"
	"syscall"
)

var router = mux.NewRouter()

var apiRoute = router.NewRoute().Name("api").PathPrefix("/api")

func StartDaemon() {
	httpServer := daemon.NewHttpServer(8181, router)

	box := packr.New("static", "../web/dist")
	router.NewRoute().Name("static").PathPrefix("/").Handler(http.FileServer(box))

	go func() {
		logrus.Info("starting http server")
		if err := httpServer.Start(); err != nil {
			logrus.WithField("error", fmt.Sprintf("%+v", err)).Fatal("http server failed to start")
		}
	}()

	waitForShutdown(func() {
		if err := httpServer.Stop(); err != nil {
			logrus.WithField("error", fmt.Sprintf("%+v", err)).Fatal("http server failed to stop")
		}
		logrus.Info("http server stopped")
	})
}

func waitForShutdown(shutdownHook func()) {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	sig := <-gracefulStop
	logrus.WithField("signal", sig).Info("shutdown signal received")
	shutdownHook()
}
