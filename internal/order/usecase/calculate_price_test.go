package usecase

import (
	"database/sql"
	"testing"

	"github.com/ab-costa/imersao-go/internal/order/entity"
	"github.com/ab-costa/imersao-go/internal/order/infra/database"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CalculateFinalPriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculateFinalPriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")

	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price floate NOT NULL, PRIMARY KEY (id))")

	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TearDown() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceUseCaseTestSuite))
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TestCaseCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10.0, 2.0)

	suite.NoError(err)

	order.CalculateFinalPrice()

	calculateFinalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	CalculateFinalPriceUsecase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := CalculateFinalPriceUsecase.Execute(calculateFinalPriceInput)

	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
