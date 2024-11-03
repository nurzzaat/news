package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/news/internal/controller"
	"github.com/nurzzaat/news/internal/repository"
	"github.com/nurzzaat/news/pkg"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		webfinalapi.mobydev.kz
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	//defer app.CloseDBConnection()
	
	ginRouter := gin.Default()
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS ,HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-*, Cross-Origin-Resource-Policy , Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	controller.Setup(app, ginRouter)
	
	go ClearNews(app.Pql)

	ginRouter.Run(fmt.Sprintf(":%s", app.Env.PORT))
}

func ClearNews(db *pgxpool.Pool) {
	for {
		now := time.Now().UTC()
		log.Println("Daily routine to clear news is running...")

		nextDay := now.Add(24 * time.Hour)
		nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())

		durationUntilNextDay := nextDay.Sub(now)

		err := repository.ClearNews(context.Background(), db)
		if err != nil {
			fmt.Println("Error running daily clear task:", err)
		}
		time.Sleep(durationUntilNextDay)
	}
}
