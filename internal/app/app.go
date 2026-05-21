package app

import (
	"crypto-portfolio/internal/portfolio"
	"crypto-portfolio/internal/prices"
	"net/http"
)

func NewHandler() *portfolio.Handler {
	repo := portfolio.NewRepository()
	priceClient := prices.NewClient(&http.Client{})
	service := portfolio.NewService(repo, priceClient)
	return portfolio.NewHandler(service)
}
