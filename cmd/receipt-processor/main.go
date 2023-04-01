package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AkankshaNichrelay/Receipt-Processor/internal/handler"
	"github.com/AkankshaNichrelay/Receipt-Processor/internal/receipts"
)

func main() {
	logger := log.Default()
	receipts := receipts.New(logger)
	handler := handler.New(logger, receipts)
	go http.ListenAndServe(":8080", handler.Router)
	log.Println("Listening on localhost:8080...")
	Stop()
}

// Stop to gracefully shutdown application upon receiving signal
func Stop() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Application stopped")
}
