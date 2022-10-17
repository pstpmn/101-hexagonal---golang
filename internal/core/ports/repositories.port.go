package ports

import (
	"101-hexagonal-golang/internal/core/domain"
)

type MembersRepository interface {
	GetById(id int) (*domain.Members, error)
	Create(member *domain.Members) (*domain.Members, error)
}

type PostRepository interface {
	GetById(id int) (*domain.Members, error)
	Create(member *domain.Members) (*domain.Members, error)
}
