package usecase

import (
	"101-hexagonal-golang/internal/core/domain"
	"101-hexagonal-golang/internal/core/ports"
)

type MembersUseCase struct {
	MembersRepo ports.MembersRepository
}

func NewMembersUseCase(repo ports.MembersRepository) ports.MembersUseCase {
	return &MembersUseCase{
		MembersRepo: repo,
	}
}

func (repo MembersUseCase) GetById(id int) (*domain.Members, error) {
	member, err := repo.MembersRepo.GetById(id)
	return member, err
}
