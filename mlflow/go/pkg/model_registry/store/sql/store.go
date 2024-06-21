package sql

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/mlflow/mlflow/mlflow/go/pkg/config"
	"github.com/mlflow/mlflow/mlflow/go/pkg/sql"
)

type ModelRegistrySQLStore struct {
	config *config.Config
	db     *gorm.DB
}

func NewModelRegistrySQLStore(logger *logrus.Logger, config *config.Config) (*ModelRegistrySQLStore, error) {
	database, err := sql.NewDatabase(logger, config.RegistryStoreURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %q: %w", config.RegistryStoreURL, err)
	}

	return &ModelRegistrySQLStore{config: config, db: database}, nil
}
