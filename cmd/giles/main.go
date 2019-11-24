package main

import (
	"flag"
	"log"

	"github.com/undeadops/giles/pkg/api"
	"github.com/undeadops/giles/pkg/config"
	"github.com/undeadops/giles/pkg/storage/db"
)

const defaultPortVariable = "PORT"
const defaultPort = "3000"

// Debug - Enable debug logging
var Debug = flag.Bool("debug", false, "Enable Debug Logging")

// Setup is the function `main` calls in order to start the web server bootstrapped with the handler
// func Setup(config *config.Config) {

// 	// cache := cache.New(ds)

// 	frontend, err := handler.NewFrontend(
// 		ds,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	e, err := server.NewFrontend(
// 		frontend,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	port := defaultPort

// 	if os.Getenv(defaultPortVariable) != "" {
// 		port = os.Getenv(defaultPortVariable)
// 	}

// 	addr := fmt.Sprintf(":%s", port)

// 	e.Start(addr)
// }

func main() {
	flag.Parse()

	config := config.New()

	// Connect to database
	db, err := db.SetupDB(config.URI)
	if err != nil {
		// Implement better health checking/retry here or in lib
		log.Fatalf("Cannot set up Database: %v", err)
	}

	server := &api.Server{DB: db, Port: defaultPort}

	server.SetupServer()
	server.Run()
}
