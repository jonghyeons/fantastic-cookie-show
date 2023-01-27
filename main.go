package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonghyeons/fantastic-cookie-show/controllers"
	"github.com/jonghyeons/fantastic-cookie-show/persistence"
	"github.com/jonghyeons/fantastic-cookie-show/services"
	"net/http"
)

func main() {
	app := gin.New()
	pc := func() *controllers.Performance {
		dao := &persistence.PerformanceRepository{
			// TODO DB conn info
			Engine: nil,
			Cache:  nil,
		}
		service := services.NewPerformanceService(dao)
		controller := controllers.NewPerformanceController(service)
		return controller
	}()

	sc := func() *controllers.Seat {
		dao := &persistence.SeatRepository{
			// TODO DB conn info
			Engine: nil,
			Cache:  nil,
		}
		service := services.NewSeatService(dao)
		controller := controllers.NewSeatController(service)
		return controller
	}()

	app.POST("/performance", pc.SetOne)
	app.GET("/performance", pc.GetAll)

	app.GET("/performance/:pid", pc.GetOne)
	app.GET("/performance/:pid/seat", sc.GetAll)
	app.POST("/performance/:pid/seat", sc.SetOne)
	app.POST("/performance/:pid/seat/:sid/payment", sc.Payment)

	app.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusNotFound)
	})

	_ = app.Run()
}
