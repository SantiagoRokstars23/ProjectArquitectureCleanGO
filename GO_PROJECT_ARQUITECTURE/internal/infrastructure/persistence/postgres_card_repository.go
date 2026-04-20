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

func (r *PostgresRepo) FindByID(id int) (*domain.Card, error) {
	row := r.db.QueryRow("SELECT id, number, name, limit FROM cards WHERE id=$1", id)

	var c domain.Card
	err := row.Scan(&c.ID, &c.Number, &c.Name, &c.Limit)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PostgresRepo) Save(card domain.Card) error {
	_, err := r.db.Exec(
		"INSERT INTO cards (number, name, limit) VALUES ($1,$2,$3)",
		card.Number, card.Name, card.Limit,
	)
	return err
}
