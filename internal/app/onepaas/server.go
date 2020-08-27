package onepaas

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/onepaas/onepaas/internal/app/onepaas/controller"
	_ "github.com/onepaas/onepaas/internal/pkg/validator"
	"github.com/rs/zerolog/log"
)

type ApiServer struct {
	*gin.Engine
	address string
	debug   bool
}

func NewApiServer(address string, debug bool) *ApiServer {
	e := gin.New()
	e.Use(logger.SetLogger())

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	return &ApiServer{
		Engine:  e,
		address: address,
		debug:   debug,
	}
}

func (as *ApiServer) setupRoutes() {
	healthz := new(controller.HealthzController)
	as.Engine.GET("/healthz", healthz.Healthz)

	v1 := as.Engine.Group("/v1")
	{
		users := new(controller.UsersController)
		v1.POST("/users", users.Add)
		//v1.GET("/users/:id", users.View)
	}
}

func (as *ApiServer) Run() error {
	as.setupRoutes()

	srv := &http.Server{
		Addr:    as.address,
		Handler: as.Engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().
				Err(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Err(err).Msg("Server forced to shutdown")

		return err
	}

	log.Info().Msg("Bye Bye")

	return nil
}