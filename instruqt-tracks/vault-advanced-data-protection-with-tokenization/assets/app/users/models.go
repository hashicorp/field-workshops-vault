package users

import (
	"log"

	"github.com/Andrew-Klaas/vault-go-demo-tokenization/config"
)

// //Data ...
// type Data struct {
// 	dbuser DBuser
// 	user   User
// }

//User holds example user PII data
type User struct {
	Cust_no string
	First   string
	Last    string
	Ssn     string
	Addr    string
	Bday    string
	Salary  float32
}

//GetRecords returns all customer records from the database
func GetRecords() ([]User, error) {
	var err error

	rows, err := config.DB.Query("SELECT * FROM vault_go_demo;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Cust_no, &u.First, &u.Last, &u.Ssn, &u.Addr, &u.Bday, &u.Salary)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("users: %v\n", users)
	return users, nil
}

//GetUsers returns all customer records from the database
func GetUsers() ([]string, error) {
	var err error

	rows, err := config.DB.Query("SELECT u.usename FROM pg_catalog.pg_user u;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]string, 0)
	for rows.Next() {
		var un string
		err := rows.Scan(&un)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, un)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("users: %v\n", users)
	return users, nil
}
