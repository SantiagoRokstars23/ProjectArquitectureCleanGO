package main

import (
	"database/sql"
	"log"

	"github.com/santiago/cards-service/internal/delivery/handlers"
	"github.com/santiago/cards-service/internal/infrastructure/persistence"
	"github.com/santiago/cards-service/internal/usecase"
	"github.com/santiago/cards-service/pkg/middleware"

		"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

/*func main() {
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
}*/

func main(){

/*	cfg := config.LoadConfig()

	db := db.NewPostgresConnection(cfg.DBUrl)*/

	db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/cardsdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := persistence.NewPostgresCardRepository(db)
	uc := usecase.NewCardUseCase(repo)
	handler := handlers.NewCardHandler(uc)



	r := gin.Default()

	public := r.Group("/cards")
	public.GET("", handler.List)

	protected := r.Group("/cards")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/:id", handler.Get)
		protected.POST("", handler.Create)
	}

	log.Println("Server running on :8080")
	r.Run(":8080")


}
