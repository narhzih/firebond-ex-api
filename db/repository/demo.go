package repository

import "firebond-ex-api.com/db/models"

type DemoRepository interface {
	CreateDemoData(demoData models.Demo) (models.Demo, error)
}
