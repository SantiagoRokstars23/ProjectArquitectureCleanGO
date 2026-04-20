package main

import (
	"database/sql"
	"log"
	"net/http"

	"GO_PROJECT_ARQUITECTURE/internal/delivery/handlers"
	"GO_PROJECT_ARQUITECTURE/internal/infrastructure/persistence"
	"GO_PROJECT_ARQUITECTURE/internal/usecase"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/cardsdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := persistence.NewPostgresCardRepository(db)
	useCase := usecase.NewListCardsUseCase(repo)
	handler := handlers.NewHandler(useCase)

	http.HandleFunc("/cards", handler.ListCards)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
