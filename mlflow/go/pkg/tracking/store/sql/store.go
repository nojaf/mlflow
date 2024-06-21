package sql

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/mlflow/mlflow/mlflow/go/pkg/config"
	"github.com/mlflow/mlflow/mlflow/go/pkg/sql"

	_ "github.com/ncruces/go-sqlite3/embed" // embed sqlite3 driver
)

type TrackingSQLStore struct {
	config *config.Config
	db     *gorm.DB
}

func NewTrackingSQLStore(logger *logrus.Logger, config *config.Config) (*TrackingSQLStore, error) {
	database, err := sql.NewDatabase(logger, config.StoreURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %q: %w", config.StoreURL, err)
	}

	return &TrackingSQLStore{config: config, db: database}, nil
}
