package persistence

import (
	"GO_PROJECT_ARQUITECTURE/internal/domain"
	"database/sql"
)

type PostgresCardRepository struct {
	db *sql.DB
}

func NewPostgresCardRepository(db *sql.DB) *PostgresCardRepository {
	return &PostgresCardRepository{db: db}
}

func (r *PostgresCardRepository) FindAll() ([]domain.Card, error) {

	rows, err := r.db.Query("SELECT id, number, name, limit FROM cards")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cards []domain.Card

	for rows.Next() {
		var c domain.Card
		err := rows.Scan(&c.ID, &c.Number, &c.Name, &c.Limit)
		if err != nil {
			return nil, err
		}

		cards = append(cards, c)
	}

	return cards, nil
}
