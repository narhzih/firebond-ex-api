package internal

import (
	"firebond-ex-api.com/db/repository"
	"github.com/rs/zerolog"
)

type Application struct {
	Repositories repository.Repositories
	Logger       zerolog.Logger
}
