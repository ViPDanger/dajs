package app

import (
	"DAJ/pkg/config"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg config.Config){
mux := http.NewServeMux()
ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
// Usecase

// 

go func() {
	log.Println("Listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}()
<-ctx.Done()
log.Println("Server is closed")
}

