package application

import (
	"example/vtr-mailer-service/actions/templating"
	"example/vtr-mailer-service/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Application struct {
	DB       *config.DB
	Cfg      *config.Config
	Template *templating.Templating
	Context  *gin.Context
}

func SetupApplication() (*Application, error) {
	cfg := config.Get()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.GetEnvironment() == "production" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	tmpl, _ := templating.Get()
	db, errDatabase := config.BootDatabase(cfg.GetDBConfig())
	if errDatabase != nil {
		log.Fatal(errDatabase.Error())
	}

	return &Application{
		DB:       db,
		Cfg:      cfg,
		Template: tmpl,
	}, nil
}
