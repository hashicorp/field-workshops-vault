package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Andrew-Klaas/vault-go-demo-tokenization/config"
)

//Index ...
func Index(w http.ResponseWriter, req *http.Request) {

	err := config.TPL.ExecuteTemplate(w, "index.gohtml", config.AppDBuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Dbview ...
func DbView(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	cRecords, err := GetRecords()
	if err != nil {
		log.Fatal(err)
	}

	err = config.TPL.ExecuteTemplate(w, "dbview.gohtml", cRecords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Records ...
func Records(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	cRecords, err := GetRecords()
	if err != nil {
		log.Fatal(err)
	}

	for i := 3; i < len(cRecords); i++ {
		u := cRecords[i]

		data := map[string]interface{}{
			"value":          u.Ssn,
			"transformation": "ssn",
		}

		response, err := config.Vclient.Logical().Write("transform/decode/vault_go_demo", data)
		if err != nil {
			log.Fatal(err)
		}
		decval := response.Data["decoded_value"].(string)
		ssn := decval

		if err != nil {
			log.Fatal(err)
		}
		cRecords[i].Ssn = string(ssn)
	}

	err = config.TPL.ExecuteTemplate(w, "records.gohtml", cRecords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//DbUserView ...
func DbUserView(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	dbUsers, err := GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	err = config.TPL.ExecuteTemplate(w, "dbusers.gohtml", dbUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Addrecord ...
func Addrecord(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		f := req.FormValue("first")
		l := req.FormValue("last")
		ssn := req.FormValue("ssn")
		adr := req.FormValue("address")
		bd := req.FormValue("birthday")
		slry := req.FormValue("salary")

		// convert form values
		f64, err := strconv.ParseFloat(slry, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		conSlry := float32(f64)

		u := User{
			Cust_no: "",
			First:   f,
			Last:    l,
			Ssn:     ssn,
			Addr:    adr,
			Bday:    bd,
			Salary:  conSlry,
		}

		data := map[string]interface{}{
			"value":          u.Ssn,
			"transformation": "ssn",
		}

		response, err := config.Vclient.Logical().Write("transform/encode/vault_go_demo", data)
		if err != nil {
			log.Fatal(err)
		}
		encval := response.Data["encoded_value"].(string)

		u.Ssn = encval

		_, err = config.DB.Exec("INSERT INTO vault_go_demo (FIRST, LAST, SSN, ADDR, BDAY, SALARY) VALUES ($1, $2, $3, $4, $5, $6)", u.First, u.Last, u.Ssn, u.Addr, u.Bday, u.Salary)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, req, "/records", http.StatusSeeOther)
	}
	err := config.TPL.ExecuteTemplate(w, "addrecord.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//UpdateRecord ...
func UpdateRecord(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		cn := req.FormValue("cust_no")
		f := req.FormValue("first")
		l := req.FormValue("last")
		ssn := req.FormValue("ssn")
		adr := req.FormValue("address")
		bd := req.FormValue("birthday")
		slry := req.FormValue("salary")

		// convert form values
		f64, err := strconv.ParseFloat(slry, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		conSlry := float32(f64)

		u := User{
			Cust_no: cn,
			First:   f,
			Last:    l,
			Ssn:     ssn,
			Addr:    adr,
			Bday:    bd,
			Salary:  conSlry,
		}

		data := map[string]interface{}{
			"value":          u.Ssn,
			"transformation": "ssn",
		}

		response, err := config.Vclient.Logical().Write("transform/encode/vault_go_demo", data)
		if err != nil {
			log.Fatal(err)
		}
		encval := response.Data["encoded_value"].(string)

		u.Ssn = encval

		convcn, err := strconv.Atoi(u.Cust_no)
		if err != nil {
			log.Fatal(err)
		}
		_, err = config.DB.Exec("UPDATE vault_go_demo SET CUST_NO=$1, FIRST=$2, LAST=$3, SSN=$4, ADDR=$5, BDAY=$6, SALARY=$7 WHERE CUST_NO=$1;", convcn, u.First, u.Last, u.Ssn, u.Addr, u.Bday, u.Salary)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, req, "/records", http.StatusSeeOther)
	}
	err := config.TPL.ExecuteTemplate(w, "updaterecord.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
