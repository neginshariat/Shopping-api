package repository

import (
	"database/sql"
	"log"
	"shopping-app/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var or = &models.Order{
	OrID:    1,
	Pants:   "short",
	Shoes:   "sport",
	TShirts: "green",
}

var st = &models.Store{
	Pants:   "lang",
	Shoes:   "highheels",
	TShirts: "red",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func TestShowAllStock(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer db.Close()
	query := `SELECT  pants, shoes, tshirts FROM store`

	rows := sqlmock.NewRows([]string{"pants", "shoes", "tshirts"}).AddRow(st.Pants, st.Shoes, st.TShirts)
	mock.ExpectQuery(query).WillReturnRows(rows)

	order, err := repo.ShowAllStock()
	assert.NotEmpty(t, order)
	assert.NoError(t, err)
	assert.Len(t, order, 1)
	// TODO: Here is would be good to also check that the expected order
	// content is returned.
}

func TestShowAllOrders(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer db.Close()
	query := `SELECT orid, pants,shoes, tshirts FROM order`

	rows := sqlmock.NewRows([]string{"orid", "pants", "shoes", "tshirts"}).AddRow(or.OrID, or.Pants, or.Shoes, or.TShirts)
	mock.ExpectQuery(query).WillReturnRows(rows)

	order, err := repo.ShowAllOrders()
	assert.NotEmpty(t, order)
	assert.NoError(t, err)
	assert.Len(t, order, 1)
}
func TestShowOrderById(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer func() {
		db.Close()
	}()

	query := "SELECT orid, pants,shoes, tshirts FROM  order WHERE orid = \\$1"

	rows := sqlmock.NewRows([]string{"id", "pants", "shoes", "tshirts"}).AddRow(or.OrID, or.Pants, or.Shoes, or.TShirts)

	mock.ExpectQuery(query).WithArgs(or.OrID).WillReturnRows(rows)

	order, err := repo.ShowOrderById(or.OrID)
	assert.NotNil(t, order)
	assert.NoError(t, err)

	assert.Equal(t, *or, order)

}

func TestCreateOrder(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer func() {
		db.Close()
	}()


	//query := "INSERT INTO order (orpants, orshoes, ortshirt ) VALUES ( ?, ?, ? ) RETURNING orid"
	query := "INSERT INTO order*"

	prep := mock.ExpectPrepare(query)

	prep.ExpectExec().WithArgs("short", "sport", "green").WillReturnResult(sqlmock.NewResult(1, 1))



	order := repo.CreateOrder(*or)
	assert.NotEmpty(t, order)
	assert.NotNil(t, order)
	// TODO: Same here, we should validate that the expected content of the
	// order is returned.
}

func TestEditOrder(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer func() {
		db.Close()
	}()


	//query := "UPDATE order SET pants= \\$2, shoes= \\$3, tshirt= \\$4 WHERE orid= \\$1"
	query := "UPDATE order*"

	prep := mock.ExpectPrepare(query)

	prep.ExpectExec().WithArgs("short", "sport", "green", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	order := repo.EditOrder(*or)
	assert.NotEmpty(t, order)

}
func TestDeleteOrder(t *testing.T) {
	db, mock := NewMock()
	repo := &stock{db}
	defer func() {
		db.Close()
	}()

	query := "DELETE FROM order WHERE orid = \\$1"


	mock.ExpectExec(query).WithArgs(or.OrID).WillReturnResult(sqlmock.NewResult(0, 1))

	order := repo.DeleteOrder(or.OrID)
	assert.NotEmpty(t, order)

}
