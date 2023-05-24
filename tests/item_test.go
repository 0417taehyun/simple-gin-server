package tests

import (
	"database/sql"
	"regexp"
	"simple-gin-server/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ItemTestSuite struct {
	suite.Suite
	database *sql.DB
	mock     sqlmock.Sqlmock
	item     *models.ItemEntity
	queries  map[string]string
}

func (suite *ItemTestSuite) SetupTest() {
	var err error
	suite.database, suite.mock, err = sqlmock.New()
	if err != nil {
		suite.T().Fatalf("error invoked with MySQL: %s", err.Error())
	}
	suite.item = &models.ItemEntity{ID: uuid.New().String(), Name: "Mac"}
}

func (suite *ItemTestSuite) TeardownTest() {
	suite.item = &models.ItemEntity{}

}

func (suite *ItemTestSuite) TestGetItem() {
	model := models.NewItemModel(suite.database)
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(suite.item.ID, suite.item.Name)

	testCases := []struct {
		name     string
		id       string
		expected *models.ItemEntity
	}{
		{
			name:     "NotFound",
			id:       "test",
			expected: nil,
		},
		{
			name:     "Success",
			id:       suite.item.ID,
			expected: suite.item,
		},
	}
	for _, testCase := range testCases {
		suite.T().Run(testCase.name, func(t *testing.T) {
			prepare := suite.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, name FROM Item WHERE id = ?;`))
			prepare.ExpectQuery().WithArgs(suite.item.ID).WillReturnRows(rows)

			item, err := model.GetItem(testCase.id)
			if err != nil {
				suite.T().Error(err.Error())
			}
			suite.Equal(testCase.expected, item)
		})
	}

}

func TestItemTestSuite(t *testing.T) {
	gin.SetMode(gin.TestMode)
	suite.Run(t, new(ItemTestSuite))
}
