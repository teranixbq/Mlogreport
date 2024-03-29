package route

import (
	"mlogreport/app/middleware"
	userRepository "mlogreport/feature/user/repository"
	"mlogreport/feature/weekly/handler"
	"mlogreport/feature/weekly/repository"
	"mlogreport/feature/weekly/service"
	"mlogreport/utils/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteWeekly(c *gin.RouterGroup, db *gorm.DB) {
	userRepository := userRepository.NewUserRepository(db)
	weeklyRepository := repository.NewWeeklyRepository(db)
	weeklyService := service.NewWeeklyService(weeklyRepository, userRepository)
	weeklyHandler := handler.NewWeeklyHandler(weeklyService)

	admin := c.Group("periode", auth.JWTMiddleware(), middleware.IsRole("admin"))
	{
		admin.POST("", weeklyHandler.CreatePeriode)
		admin.GET("", weeklyHandler.GetAllPeriode)
		admin.PUT("/:id", weeklyHandler.UpdatePeriode)
	}

	advisor := c.Group("admin/weekly", auth.JWTMiddleware(), middleware.IsRole("advisor"))
	{
		advisor.GET("/:nim", weeklyHandler.GetAllWeeklyAdvisor)
		advisor.PATCH(":iduser/:id", weeklyHandler.UpdateStatus)
	}

	user := c.Group("weekly", auth.JWTMiddleware(), middleware.IsRole(""))
	{
		user.POST("", weeklyHandler.CreateWeekly)
		user.GET("", weeklyHandler.GetAllWeekly)
		user.PATCH("/:id", weeklyHandler.UpdateWeekly)
	}
}
