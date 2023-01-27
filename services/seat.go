package services

import (
	"errors"
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"github.com/jonghyeons/fantastic-cookie-show/persistence"
)

// Client - Server  - Cache - DB
type ISeat interface {
	GetAll(pid string) (error, []models.Seat)

	IsReservable(sid string) (error, bool)
	Reserve(user models.User, seat models.Seat) error

	Payment(user models.User, seat models.Seat) error
}

var _ ISeat = (*Seat)(nil)

type Seat struct {
	dao persistence.ISeatRepository
}

func NewSeatService(dao persistence.ISeatRepository) *Seat {
	return &Seat{dao: dao}
}

func (s *Seat) GetAll(pid string) (error, []models.Seat) {
	err, list := s.dao.GetAll(pid)
	if err != nil {
		return err, nil
	}
	return nil, list
}

func (s *Seat) IsReservable(sid string) (error, bool) {
	err, seatFromCache := s.dao.GetOneByCache(sid)
	if err != nil {
		return err, false
	}
	return nil, seatFromCache.Reservable
}

func (s *Seat) Reserve(user models.User, seat models.Seat) error {
	err, isReservable := s.IsReservable(seat.Id)
	if err != nil {
		return err
	}
	if isReservable {
		if err := s.dao.SetOneToCache(user, seat); err != nil {
			return err
		}
		if err := s.dao.SetOne(user, seat); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("can not reservable")
	}
}

func (s *Seat) Payment(user models.User, seat models.Seat) error {
	//TODO implement me
	panic("implement me")
}
