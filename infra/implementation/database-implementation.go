package implementation

import (
	"github.com/alura/go-api-rest/logger"
	"github.com/alura/go-api-rest/models"
)

type Database struct {
	personalities []models.Personality
}

func (d *Database) Put(personality models.Personality) bool {
	logger.Info("Implementation.Database.Put", "process started")
	d.personalities = append(d.personalities, personality)
	logger.Info("Implementation.Database.Put", "process finished")
	return true

}

func (d *Database) GetAll() []models.Personality {
	logger.Info("Implementation.Database.GetAll", "process started")
	logger.Info("Implementation.Database.GetAll", "process finished")
	return d.personalities
}
