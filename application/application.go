package application

import (
	"example/vtr-mailer-service/actions/templating"
	"example/vtr-mailer-service/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Application struct {
	Cfg      *config.Config
	Template *templating.Templating
	Context  *gin.Context
}

func SetupApplication() (*Application, error) {
	cfg := config.Get()
	tmpl, _ := templating.Get()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.GetEnvironment() == "production" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	return &Application{
		Cfg:      cfg,
		Template: tmpl,
	}, nil
}
