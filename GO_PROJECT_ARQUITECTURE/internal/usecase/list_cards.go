package usecase

import "github.com/santiago/cards-service/internal/domain"

type ListCardsUseCase struct {
	repo domain.CardRepository
}

func NewListCardsUseCase(r domain.CardRepository) *ListCardsUseCase {
	return &ListCardsUseCase{repo: r}
}

func (uc *ListCardsUseCase) Execute() ([]domain.Card, error) {
	return uc.repo.FindAll()
}
