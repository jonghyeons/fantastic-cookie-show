package persistence

import (
	"github.com/go-redis/redis"
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"gorm.io/gorm"
)

// Client - Server  - Cache - DB
type ISeatRepository interface {
	GetAll(pid string) (error, []models.Seat)
	GetOneByCache(sid string) (error, models.SeatCache)

	SetOne(user models.User, seat models.Seat) error
	SetOneToCache(user models.User, seat models.Seat) error
}

var _ ISeatRepository = (*SeatRepository)(nil)

type SeatRepository struct {
	Engine *gorm.Dialector
	Cache  *redis.Client
}

func (s *SeatRepository) GetAll(pid string) (error, []models.Seat) {
	//TODO implement me
	panic("implement me")
}

func (s *SeatRepository) GetOneByCache(sid string) (error, models.SeatCache) {
	//TODO implement me
	panic("implement me")
}

func (s *SeatRepository) SetOne(user models.User, seat models.Seat) error {
	//TODO implement me
	panic("implement me")
}

func (s *SeatRepository) SetOneToCache(user models.User, seat models.Seat) error {
	//TODO implement me
	panic("implement me")
}
