package domain

type CardRepository interface {
	FindAll() ([]Card, error)
}