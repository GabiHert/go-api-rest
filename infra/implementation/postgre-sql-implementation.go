package implementation

import (
	"github.com/alura/go-api-rest/logger"
	"github.com/alura/go-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
	DB  *gorm.DB
	err error
}

func (pg *PostgreSql) ConnectDB() {
	logger.Info("implementation.PostgreSql.ConnectDB", "Process started")

	dsn := "host=localhost user=postgres password=1809 dbname=alura_api_rest_register port=5433 sslmode=disable"
	pg.DB, pg.err = gorm.Open(postgres.Open(dsn))

	if pg.err != nil {
		logger.Error("implementation.PostgreSql.connectDB", "Process error", "Error connecting to database")
	}
	logger.Info("implementation.PostgreSql.ConnectDB", "Process finished")

}

func (pg *PostgreSql) GetPersonalities() ([]models.Personality, error) {
	logger.Info("implementation.PostgreSql.GetPersonalities", "Process started")

	var personalities []models.Personality
	err := pg.DB.Find(&personalities).Error

	if err != nil {
		logger.Error("implementation.PostgreSql.connectDB", "Process error", "Error getting personalities")
		return personalities, err
	}

	logger.Info("implementation.PostgreSql.GetPersonalities", "Process finished")

	return personalities, nil
}

func (pg *PostgreSql) GetPersonality(id string) (models.Personality, error) {
	logger.Info("implementation.PostgreSql.GetPersonality", "Process started", "id:", id)

	var personality models.Personality
	err := pg.DB.First(&personality, id).Error
	if err != nil {
		logger.Error("implementation.PostgreSql.connectDB", "Process error", "Error getting personality", "id:", id)
		return personality, err
	}
	logger.Info("implementation.PostgreSql.GetPersonality", "Process finished")

	return personality, nil
}

func (pg *PostgreSql) CreatePersonality(personality models.Personality) error {

	logger.Info("implementation.PostgreSql.GetPersonalities", "Process started", "personality:", personality)

	err := pg.DB.Create(&personality).Error
	if err != nil {
		logger.Error("implementation.PostgreSql.GetPersonalities", "Process error", "Error creating personality", "personality:", personality)
		return err
	}

	return nil
}

func (pg *PostgreSql) DeletePersonality(id string) error {

	logger.Info("implementation.PostgreSql.DeletePersonality", "Process started", "id:", id)
	var personality models.Personality

	err := pg.DB.Delete(personality, id).Error
	if err != nil {
		logger.Error("implementation.PostgreSql.DeletePersonality", "Process error", "Error deleting personality", "id:", id)
		return err
	}

	logger.Info("implementation.PostgreSql.DeletePersonality", "Process started", "id:", id)

	return nil
}

func (pg *PostgreSql) EditPersonality(personalityUpdate models.Personality, id string) error {

	logger.Info("implementation.PostgreSql.EditPersonality", "Process started", "id:", id)
	var personality models.Personality

	errFirst := pg.DB.First(&personality, id).Error

	personality.Story = personalityUpdate.Story
	personality.Name_ = personalityUpdate.Name_

	if errFirst != nil {
		logger.Error("implementation.PostgreSql.EditPersonality", "Process error", "Error editing personality", "id:", id)
		return errFirst
	}

	errSave := pg.DB.Save(&personality).Error

	if errSave != nil {
		logger.Error("implementation.PostgreSql.EditPersonality", "Process error", "Error editing personality", "id:", id)
		return errSave
	}

	logger.Info("implementation.PostgreSql.EditPersonality", "Process started", "id:", id)

	return nil
}
