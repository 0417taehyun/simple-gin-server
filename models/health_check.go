package models

import (
	"database/sql"
)

const (
	HEALTH_CHECK_QUERY_BASE_PATH = "sql/health_check/"
	CHECK_DATABASE               = "check_database.sql"
)

type HealthCheckModel struct {
	Database *sql.DB
}

func (healthCheckModel HealthCheckModel) CheckDatabase() error {
	database := healthCheckModel.Database
	statement, err := database.Prepare("SELECT 1;")
	if err != nil {
		return err
	}
	defer statement.Close()

	var ignore interface{}
	err = statement.QueryRow().Scan(&ignore)
	if err != nil {
		return err
	}
	return nil
}

func NewHealthCheckModel(database *sql.DB) *HealthCheckModel {
	return &HealthCheckModel{Database: database}
}
