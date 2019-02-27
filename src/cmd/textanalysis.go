package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gorilla/mux"

  "pkg/handlers"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/v1/analyze/text", handlers.AnalyzeText).Methods("POST")
  srv := &http.Server{
    Addr: ":7777",
    Handler: router,
  }
  log.Println("Starting textanalysis server")
  go func() {
    if err := srv.ListenAndServe(); err != nil {
      log.Println(err)
    }
  }()
  c := make(chan os.Signal, 1)
  signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
  <-c
  wait, _ := time.ParseDuration("15s")
  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()
  log.Println("Shutting down textanalysis server")
  srv.Shutdown(ctx)
  os.Exit(0)
}
