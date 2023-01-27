package persistence

import (
	"github.com/go-redis/redis"
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"gorm.io/gorm"
)

type IPerformanceRepository interface {
	GetAll() (error, []models.Performance)
	GetOne(pid string) (error, models.Performance)

	SetOne(models.Performance) error
}

var _ IPerformanceRepository = (*PerformanceRepository)(nil)

type PerformanceRepository struct {
	Engine *gorm.Dialector
	Cache  *redis.Client
}

func (p *PerformanceRepository) GetAll() (error, []models.Performance) {
	//TODO implement me
	panic("implement me")
}

func (p *PerformanceRepository) GetOne(pid string) (error, models.Performance) {
	//TODO implement me
	panic("implement me")
}

func (p *PerformanceRepository) SetOne(performance models.Performance) error {
	//TODO implement me
	panic("implement me")
}
