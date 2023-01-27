package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"github.com/jonghyeons/fantastic-cookie-show/services"
	"net/http"
)

type Seat struct {
	service services.ISeat
}

func NewSeatController(service services.ISeat) *Seat {
	return &Seat{service: service}
}

func (s *Seat) GetAll(c *gin.Context) {
	pid := c.Param("pid")
	err, list := s.service.GetAll(pid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
	return
}

func (s *Seat) SetOne(c *gin.Context) {
	var user models.User // From Auth
	var seat models.Seat

	err := c.ShouldBind(&seat)
	if err != nil || seat == (models.Seat{}) {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = s.service.Reserve(user, seat)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
	return
}

func (s *Seat) Payment(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}
