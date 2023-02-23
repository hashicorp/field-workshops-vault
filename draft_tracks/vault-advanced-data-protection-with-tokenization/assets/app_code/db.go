package config

import (

	// "database/sql"
	// "net/http"
	// "github.com/vault/api"

	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
	_ "github.com/lib/pq"
)

//DBuser holds DB user info
type DBuser struct {
	Username string
	Password string
}

//DB Connection
var DB *sql.DB
var AppDBuser DBuser
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

//Vclient holds our HashiCorp Vault Client
var Vclient, _ = api.NewClient(&api.Config{Address: "http://vault-ui.default.svc:8200", HttpClient: httpClient})
var tokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
var K8sAuthRole = "vault_go_demo"
var K8sAuthPath = "auth/kubernetes/login"

func init() {
	fmt.Printf("Vault client init....\n")

	buf, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		log.Fatal(err)
	}
	jwt := string(buf)
	fmt.Printf("K8s Service Account JWT: %v", jwt)

	config := map[string]interface{}{
		"jwt":  jwt,
		"role": K8sAuthRole,
	}

	//Login
	secret, err1 := Vclient.Logical().Write(K8sAuthPath, config)
	if err1 != nil {
		log.Fatal(err)
	}
	Vclient.SetToken(secret.Auth.ClientToken)

	//Pull dynamic database credentials
	data, err := Vclient.Logical().Read("database/creds/vault_go_demo")
	if err != nil {
		log.Fatal(err)
	}
	username := data.Data["username"]
	password := data.Data["password"]
	SQLQuery := "postgres://" + username.(string) + ":" + password.(string) + "@pq-postgresql-headless.default.svc:5432/vault_go_demo?sslmode=disable"

	AppDBuser.Username = username.(string)
	AppDBuser.Password = password.(string)

	//Don't do this in production!
	fmt.Printf("\nDB Username: %v\n", AppDBuser.Username)
	fmt.Printf("DB Password: %v\n\n", AppDBuser.Password)

	//DB setup
	DB, err = sql.Open("postgres", SQLQuery)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database\n")

	SQLQuery = "DROP TABLE vault_go_demo;"
	DB.Exec(SQLQuery)
	SQLQuery = "CREATE TABLE vault_go_demo (CUST_NO SERIAL PRIMARY KEY, FIRST TEXT NOT NULL, LAST TEXT NOT NULL, SSN TEXT NOT NULL, ADDR CHAR(50), BDAY DATE DEFAULT '1900-01-01', SALARY REAL DEFAULT 25500.00);"
	DB.Exec(SQLQuery)
	SQLQuery = "INSERT INTO vault_go_demo (FIRST, LAST, SSN, ADDR, BDAY, SALARY) VALUES('John', 'Doe', '435-59-5123', '456 Main Street', '1980-01-01', 60000.00);"
	DB.Exec(SQLQuery)
	SQLQuery = "INSERT INTO vault_go_demo (FIRST, LAST, SSN, ADDR, BDAY, SALARY) VALUES('Jane', 'Smith', '765-24-2083', '331 Johnson Street', '1985-02-02', 120000.00);"
	DB.Exec(SQLQuery)
	SQLQuery = "INSERT INTO vault_go_demo (FIRST, LAST, SSN, ADDR, BDAY, SALARY) VALUES('Ben', 'Franklin', '111-22-8084', '222 Chicago Street', '1985-02-02', 180000.00);"
	DB.Exec(SQLQuery)

}

// Create Table vault-go-demo (
// 	CUST_NO SERIAL PRIMARY KEY,
// 	FIRST               TEXT NOT NULL,
// 	LAST                TEXT NOT NULL,
// 	SSN                 TEXT NOT NULL,
// 	ADDR                CHAR(50),
// 	BDAY			    DATE DEFAULT '1900-01-01',
// 	SALARY              REAL DEFAULT 25500.00
// );
