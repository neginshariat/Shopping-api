package repository

import (
	"database/sql"
	"fmt"
	"log"
	"shopping-app/config"
	"shopping-app/models"

	_ "github.com/lib/pq"
)

type Repository interface {
	ShowAllStock() ([]models.Store, error)
	ShowAllOrders() ([]models.Order, error)
	ShowOrderById(id int64) (models.Order, error)
	CreateOrder(order models.Order) int64
	EditOrder(order models.Order) int64
	DeleteOrder(id int64) int64
}

type stock struct {
	DB *sql.DB
}

func getDbUrl(conf *config.Config) string {
	//DB_URL=postgres://root:postgres@postgres:5432/shoppingdb?sslmode=disable
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName, conf.DatabaseSSLMode)
	return dbURL
}

func StockRepository(conf *config.Config) *stock {
	dbName := getDbUrl(conf)
	db, err := sql.Open("postgres", dbName)
	if err != nil {
		fmt.Println("is not able to connect to db", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err, "unable to ping and connect to db")
	}
	fmt.Println("connected successfully")

	return &stock{DB: db}
}

func (s *stock) ShowAllStock() ([]models.Store, error) {

	sqlStatement := `SELECT  pants, shoes, tshirts FROM store`

	var stock []models.Store
	row, err := s.DB.Query(sqlStatement)
	if err != nil {
		// TODO: It's better to return nil, err here.
		// log.Fatal should only be used from main() function
		// Same for the cases below where log.Fatal is used.
		log.Fatalf("Unable to execute the query %v", err)
	}
	defer row.Close()

	for row.Next() {
		var stockHolder models.Store
		err := row.Scan(&stockHolder.Pants, &stockHolder.Shoes, &stockHolder.TShirts)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		defer row.Close()
		stock = append(stock, stockHolder)

	}
	return stock, nil
}
func (s *stock) ShowAllOrders() ([]models.Order, error) {

	sqlStatement := `SELECT orid, pants,shoes, tshirts FROM order`

	var orders []models.Order
	row, err := s.DB.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	defer row.Close()

	for row.Next() {
		var orderHolder models.Order
		err := row.Scan(&orderHolder.OrID, &orderHolder.Pants, &orderHolder.Shoes, &orderHolder.TShirts)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		defer row.Close()
		orders = append(orders, orderHolder)

	}
	return orders, nil
}

func (s *stock) ShowOrderById(id int64) (models.Order, error) {

	var order models.Order

	sqlStatement := `SELECT orid, pants,shoes, tshirts FROM order WHERE orid = $1`

	row := s.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&order.OrID, &order.Pants, &order.Shoes, &order.TShirts)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("that is your order", order)
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return order, nil
}
func (s *stock) CreateOrder(order models.Order) int64 {


	/* 	sqlStatement := "INSERT INTO order (pants, shoes, tshirt ) VALUES ($1, $2, $3) RETURNING orid"

	   	var id int64

	   	err := s.db.QueryRow(sqlStatement, order.Pants, order.Shoes, order.TShirts).Scan(&id)
	   	if err != nil {
	   		log.Fatalf("Unable to execute the query. %v", err)
	   	} */
	insert, err := s.DB.Prepare("INSERT INTO order (orpants, orshoes, ortshirt ) VALUES ($1, $2, $3) RETURNING orid")
	if err != nil {
		log.Fatalf("invalid insert query")
	}




	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("unable to retreive id from last inserted record. %v", err)
	}
	fmt.Printf("inserted a single record %v", id)




func (s *stock) EditOrder(order models.Order) int64 {

	stmn, err := s.DB.Prepare("UPDATE order SET pants=$1, shoes=$2, tshirt=$3 WHERE orid=$4")
	if err != nil {
		log.Fatalf("invalid insert query")
	}




	res, err := stmn.Exec(order.Pants, order.Shoes, order.TShirts, order.OrID)
	if err != nil {
		log.Fatal("Unable to execute the query.", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Unable to checking the affected rows", err)
	}
	fmt.Printf("Total rows/records affected%v", rowsAffected)
	return rowsAffected
}
func (s *stock) DeleteOrder(id int64) int64 {

	sqlStatement := `DELETE FROM order WHERE orid = $1`
	res, err := s.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal("Unable to delete the query.", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Unable to checking the affected rows", err)
	}
	fmt.Printf("Total rows/records affected%v", rowsAffected)

	return rowsAffected
}
