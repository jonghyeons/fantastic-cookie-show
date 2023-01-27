package services

import (
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"github.com/jonghyeons/fantastic-cookie-show/persistence"
)

type IPerformance interface {
	GetAll() (error, []models.Performance)
	GetOne(pid string) (error, models.Performance)

	SetOne(models.Performance) error
}

var _ IPerformance = (*Performance)(nil)

type Performance struct {
	dao persistence.IPerformanceRepository
}

func NewPerformanceService(dao persistence.IPerformanceRepository) *Performance {
	return &Performance{dao: dao}
}

func (p *Performance) GetAll() (error, []models.Performance) {
	err, list := p.dao.GetAll()
	if err != nil {
		return err, nil
	}
	return nil, list
}

func (p *Performance) GetOne(pid string) (error, models.Performance) {
	err, performance := p.dao.GetOne(pid)
	if err != nil {
		return err, *&models.Performance{}
	}
	return nil, performance
}

func (p *Performance) SetOne(performance models.Performance) error {
	err := p.dao.SetOne(performance)
	return err
}
