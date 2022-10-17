package repositories

import (
	"101-hexagonal-golang/internal/core/domain"
	"101-hexagonal-golang/internal/core/ports"

	"gorm.io/gorm"
)

type members struct {
	db *gorm.DB
}

func NewMembersRepo(db *gorm.DB) ports.MembersRepository {
	return &members{
		db: db,
	}
}

func (db members) Create(m *domain.Members) (*domain.Members, error) {
	return m, nil
}

func (m members) GetById(id int) (*domain.Members, error) {
	result := domain.Members{}
	err := m.db.Table("members").First(&result, id).Error
	return &result, err
}
