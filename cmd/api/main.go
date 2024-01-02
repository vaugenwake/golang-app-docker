package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

const ENV_FILE = ".env"

func main() {
	if err := godotenv.Load(ENV_FILE); err != nil {
		fmt.Printf("Could not load .env file, does it exist? Err: %v", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)

		resp["ping"] = "pong"

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte{})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	})

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Could not find port, Err: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatalf("Could not start server, Err: %v", err)
	}
}
