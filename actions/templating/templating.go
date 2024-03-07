package templating

import (
	"embed"
	"errors"
	"io/fs"
	"strings"

	"github.com/rs/zerolog/log"
)

type Templating struct {
	Templates *Templates
}

func Get() (*Templating, error) {
	templates, err := LoadTemplates()
	if err != nil {
		return nil, err
	}

	return &Templating{
		Templates: templates,
	}, nil
}

const dir = "templates"

type Templates map[string]string

func (s *Templates) Get(script string) string {
	return (*s)[script]
}

//go:embed templates/*
var scriptsFS embed.FS

// LOAD FILES
func LoadTemplates() (*Templates, error) {
	scripts := make(Templates)

	files, _ := fs.ReadDir(scriptsFS, dir)
	for _, script := range files {
		name := script.Name()
		cleanName := strings.Split(name, ".")[0]

		if len(scripts[cleanName]) > 0 {
			log.Error().Msgf("Script name collision")
			return nil, errors.New("script name collision")
		}

		file, _ := scriptsFS.ReadFile(dir + "/" + name)
		scripts[cleanName] = string(file)
	}

	return &scripts, nil
}
