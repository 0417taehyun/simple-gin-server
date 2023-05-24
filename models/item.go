package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ITEM_QUERY_BASE_PATH = "sql/item/"
	GET_ITEM             = "get_item.sql"
)

type ItemEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemModel struct {
	Database *sql.DB
}

func (itemModel ItemModel) GetItem(id string) (*ItemEntity, error) {
	database := itemModel.Database
	statement, err := database.Prepare("SELECT id, name FROM Item WHERE id = ?;")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)
	var item ItemEntity
	if err = row.Scan(&item.ID, &item.Name); err != nil {
		return nil, err
	}
	return &item, nil
}

func NewItemModel(database *sql.DB) *ItemModel {
	return &ItemModel{Database: database}
}
