package main

import (
	"context"
	"log"

	"WebhookFlow/internal/models"
	"WebhookFlow/internal/repository"

	"github.com/google/uuid"
)

func seedOrder(repo *repository.OrderRepository) {

	order := &models.Order{
		ID:         uuid.New(),
		MerchantID: uuid.New(),
		Amount:     499.99,
		Status:     "created",
	}

	err := repo.Create(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("order inserted")
}
