package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (r CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllcustomers := "select  customer_id, name, date_of_birth, city, zipcode, status from customers"

	insert := `INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES  (2005, 'Pappu Alawa', '2000-02-15', 'India', '454545', 1)`
	_, er := r.client.Query(insert)
	if er != nil {
		fmt.Println("insert error ")
	}
	rows, err := r.client.Query(findAllcustomers)
	if err != nil {
		fmt.Println("Error while fetching customer data.")
		return nil, err
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			panic("Error while mapping customer data.")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositiryDb() CustomerRepositoryDb {
	// Mysql connection
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open("mysql", dbString)
	if err != nil {
		// panic(err)
		fmt.Println("Error opening DB connection:", err)
		return CustomerRepositoryDb{}
	}

	// Ping the database to verify the connection
	if err := dbConn.Ping(); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return CustomerRepositoryDb{}
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	return CustomerRepositoryDb{dbConn}
}
