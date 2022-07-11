package interfaces

import (
	"github.com/alura/go-api-rest/infra/implementation"
	"github.com/alura/go-api-rest/models"
)

type iDatabase interface {
	ConnectDB()
	GetPersonalities() ([]models.Personality, error)
	GetPersonality(id string) (models.Personality, error)
	CreatePersonality(personality models.Personality) error
	DeletePersonality(id string) error
	EditPersonality(personalityUpdate models.Personality, id string) error
}

var IDatabase iDatabase = &implementation.PostgreSql{}
