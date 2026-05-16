package main

import (
	"crypto-portfolio/internal/portfolio"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	portfolioHandler := setupPortfolioHandler()
	portfolioHandler.RegisterRoutes(mux)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Fatal(err)
		}
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func setupPortfolioHandler() *portfolio.Handler {
	repo := portfolio.NewRepository()
	service := portfolio.NewService(repo)
	return portfolio.NewHandler(service)
}
