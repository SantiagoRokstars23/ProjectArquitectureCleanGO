package usecase

import "GO_PROJECT_ARQUITECTURE/internal/domain"

type CardRepository interface {
	FindAll() ([]domain.Card, error)
	FindByID(id int) (*domain.Card, error)
	Save(card domain.Card) error
}

type CardUseCase struct {
	repo CardRepository
}

func NewCardUseCase(r CardRepository) *CardUseCase {
	return &CardUseCase{repo: r}
}

func (uc *CardUseCase) List() ([]domain.Card, error) {
	return uc.repo.FindAll()
}

func (uc *CardUseCase) Get(id int) (*domain.Card, error) {
	return uc.repo.FindByID(id)
}

func (uc *CardUseCase) Create(card domain.Card) error {
	return uc.repo.Save(card)
}