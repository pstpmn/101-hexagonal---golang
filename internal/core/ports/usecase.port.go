package ports

import (
	"101-hexagonal-golang/internal/core/domain"
)

type MembersUseCase interface {
	GetById(int) (*domain.Members, error)
}
