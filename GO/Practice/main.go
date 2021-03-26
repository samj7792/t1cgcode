package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "samuel:password@tcp(127.0.0.1:3306)/practice")
	db.Ping()
	if err != nil {
		log.Fatal("sql.Open: ", err)
	}
	defer db.Close()

	// var (
	// 	Customerid  int
	// 	FirstName   string
	// 	LastName    string
	// 	City        string
	// 	OrderId     int
	// 	OrderNumber int
	// )
	// stmt, err := db.Prepare("SELECT Customerid, FirstName, LastName FROM Customers WHERE Customerid=?")
	// if err != nil {
	// 	log.Fatal("db.Prepare: ", err)
	// }
	// defer stmt.Close()
	// rows, err := stmt.Query(2) // Parameter here fills ? in statement above
	// if err != nil {
	// 	log.Fatal("stmt.Query: ", err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&Customerid, &FirstName, &LastName)
	// 	if err != nil {
	// 		log.Fatal("rows.Scan: ", err)
	// 	}
	// 	log.Println(Customerid, FirstName, LastName)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal("rows.Err: ", err)
	// }

	// tbl, err := db.Query("SELECT Customers.LastName, Customers.FirstName, Customers.City, Orders.OrderId, Orders.OrderNumber, Orders.Customerid FROM Customers JOIN Orders WHERE Customers.Customerid = Orders.Customerid")
	// if err != nil {
	// 	log.Fatal("db.Query: ", err)
	// }
	// defer tbl.Close()
	// for tbl.Next() {
	// 	err := tbl.Scan(&LastName, &FirstName, &City, &OrderId, &OrderNumber, &Customerid)
	// 	if err != nil {
	// 		log.Fatal("rows.Scan: ", err)
	// 	}
	// 	log.Println(LastName, FirstName, City, OrderId, OrderNumber, Customerid)
	// }
	// err = tbl.Err()
	// if err != nil {
	// 	log.Fatal("rows.Err: ", err)
	// }

	// stmt, err := db.Prepare("INSERT INTO Customers(LastName, FirstName, City) Values(?, ?, ?)")
	stmt, err := db.Prepare("INSERT INTO Orders(OrderId, OrderNumber, Customerid) Values(?, ?, ?)")
	if err != nil {
		log.Fatal("db.Prepare: ", err)
	}
	// res, err := stmt.Exec("Builder", "Bob", "New York")
	res, err := stmt.Exec(1234, 5, 4)
	if err != nil {
		log.Fatal("stmt.Exec", err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("res.LastInsertId: ", err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal("res.RowsAffected: ", err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
