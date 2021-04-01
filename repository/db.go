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
	//ShowByCathegory(cat string) (models.Store, error)
	ShowOrdeById(id int64) (models.Order, error)
	CreateOrder(order models.Order) int64
	EditOrder(id int64, order models.Order) int64
	DeleteOrder(id int64) int64
}

type stock struct {
	db *sql.DB
}

func getDbUrl(conf *config.Config) string {
	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseName, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseSSLMode)
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

	return &stock{db: db}
}

func (s *stock) ShowAllStock() ([]models.Store, error) {

	sqlStatement := `SELECT  pants, shoes, tshirts FROM store`

	var stock []models.Store
	row, err := s.db.Query(sqlStatement)
	if err != nil {
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

/* func (s *stock) ShowByCathegory(cat models.Store) (models.Store, error) {

	var store models.Store

	sqlStatement := `SELECT ? FROM store`

	row := s.db.QueryRow(sqlStatement, cat)

	err := row.Scan(&store.Pants, &store.Shoes, &store.TShirts)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("that is your order", store)
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return store, nil

} */
func (s *stock) ShowOrderById(id int64) (models.Order, error) {

	var order models.Order

	sqlStatement := `SELECT orid, pants,shoes, tshirts FROM order WHERE orid = ?`

	row := s.db.QueryRow(sqlStatement, id)

	err := row.Scan(&order.OrID, &order.Pants, &order.Shoes, &order.TShirts)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		//fmt.Println(w, http.StatusBadRequest, "No entry found with the id="+id)
	case nil:
		fmt.Println("that is your order", order)
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return order, nil
}
func (s *stock) CreateOrder(order models.Order) int64 {

	insert, err := s.db.Prepare("INSERT INTO order (orpants, orshoes, ortshirt ) VALUES ( $1, $2, $3 ) RETURNING orid")
	if err != nil {
		log.Fatalf("invalid insert query")
	}
	//"INSERT INTO order \\(orpants, orshoes, ortshirt\\) VALUES \\(\\?, \\?, \\?\\) RETURNING orid"

	res, err := insert.Exec(order.Pants, order.Shoes, order.TShirts)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("unable to retreive id from last inserted record. %v", err)
	}
	fmt.Printf("inserted a single record %v", id)
	return id
}
func (s *stock) EditOrder(id int64, order models.Order) int64 {

	sqlStatement := `UPDATE order SET orpants=$2, orshoes=$3, ortshirt=$4 WHERE orid=$1`
	res, err := s.db.Exec(sqlStatement, id, order.Pants, order.Shoes, order.TShirts)
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

	sqlStatement := `DELETE FROM order WHERE orid = ?`
	res, err := s.db.Exec(sqlStatement, id)
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
