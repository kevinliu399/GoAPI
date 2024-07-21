package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevinliu399/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)        // Enable filename and line number in log output
	var r *chi.Mux = chi.NewRouter() // struct with methods for routing HTTP requests
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
