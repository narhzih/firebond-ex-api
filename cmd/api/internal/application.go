package internal

import (
	"firebond-ex-api.com/db/repository"
	"firebond-ex-api.com/services"
	"github.com/rs/zerolog"
)

type Application struct {
	Repositories repository.Repositories
	Logger       zerolog.Logger
	Services     services.Services
}
