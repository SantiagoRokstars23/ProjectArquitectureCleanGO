package usecase

import "GO_PROJECT_ARQUITECTURE/internal/domain"

type ListCardsUseCase struct {
	repo domain.CardRepository
}

func NewListCardsUseCase(r domain.CardRepository) *ListCardsUseCase {
	return &ListCardsUseCase{repo: r}
}

func (uc *ListCardsUseCase) Execute() ([]domain.Card, error) {
	return uc.repo.FindAll()
}
