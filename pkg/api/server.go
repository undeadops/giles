package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/valve"
	"github.com/undeadops/giles/pkg/storage"
)

// Server - Dependancy injection through interfaces
type Server struct {
	DB   storage.DataAccessLayer
	Port string
	Ctx  context.Context
}

// SetupServer - Setup base HTTP Server
func (s *Server) SetupServer() {

}

// Run - Start Http Server
func (s *Server) Run() {
	// Our graceful valve shut-off package to manage code preemption and
	// shutdown signaling.
	valv := valve.New()
	baseCtx := valv.Context()

	srv := http.Server{Addr: ":" + s.Port, Handler: chi.ServerBaseContext(baseCtx, s.Router())}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			valv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	srv.ListenAndServe()
}
