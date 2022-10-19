package books_test

import (
	assertion "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"kurdi-go/domain/domain_entities"
	"kurdi-go/infrastructure/infrastructure_database"
	"kurdi-go/infrastructure/infrastructure_repositories"
	"testing"
)

func TestListOnlyDiscounted(test *testing.T) {
	assert := assertion.New(test)
	infrastructure_database.Connect()
	// preparation
	booksRepository := infrastructure_repositories.NewStockRepository()
	books, err := booksRepository.ListAllDiscounted()
	if err != nil {
		test.Error(err)
	}
	for _, book := range books {
		actual := book.IsDiscounted
		assert.True(actual, "The book should be discounted.")
	}
}

func TestListWithWriteSellingPrice(test *testing.T) {
	assert := assertion.New(test)
	infrastructure_database.Connect()
	// preparation
	booksRepository := infrastructure_repositories.NewStockRepository()
	books, err := booksRepository.ListAllDiscountedModel()
	if err != nil {
		test.Error(err)
	}
	for _, book := range books {
		var actualBook domain_entities.Book
		// scan neglect gorm hooks
		infrastructure_database.PostgresDB.Session(&gorm.Session{SkipHooks: true}).Model(domain_entities.Book{}).Where("id = ?", book.ID).First(&actualBook)
		if book.IsDiscounted {
			actual := book.SellingPrice
			expected := actualBook.SellingPrice - actualBook.Discount
			assert.Equal(expected, actual, "discounted price selling price should be calculated.")
		}

	}
}
