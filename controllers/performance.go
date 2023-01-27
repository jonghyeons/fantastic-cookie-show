package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonghyeons/fantastic-cookie-show/models"
	"github.com/jonghyeons/fantastic-cookie-show/services"
	"net/http"
)

type Performance struct {
	service services.IPerformance
}

func NewPerformanceController(service services.IPerformance) *Performance {
	return &Performance{service: service}
}

func (pc *Performance) GetAll(c *gin.Context) {
	err, list := pc.service.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
	return
}

func (pc *Performance) GetOne(c *gin.Context) {
	pid := c.Param("pid")
	err, performance := pc.service.GetOne(pid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": performance,
	})
	return
}

func (pc *Performance) SetOne(c *gin.Context) {
	var performance models.Performance
	err := c.ShouldBind(&performance)
	if err != nil || performance == (models.Performance{}) {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = pc.service.SetOne(performance)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
	return
}
