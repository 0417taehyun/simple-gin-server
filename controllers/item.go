package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"simple-gin-server/models"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ItemController struct {
	Model *models.ItemModel
}

func (itemController ItemController) GetItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindUri(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": err.Error()})
		return
	}

	result, err := itemController.Model.GetItem(item.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			message := fmt.Sprintf("item %s not found", item.ID)
			c.JSON(http.StatusNotFound, gin.H{"detail": message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func NewItemController(database *sql.DB) *ItemController {
	model := models.NewItemModel(database)
	return &ItemController{Model: model}
}
