package main

import (
	"log"
	"net/http"

	"github.com/Andrew-klaas/vault-go-demo-tokenization/users"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", users.Index)
	http.HandleFunc("/records", users.Records)
	http.HandleFunc("/dbview", users.DbView)
	http.HandleFunc("/addrecord", users.Addrecord)
	http.HandleFunc("/updaterecord", users.UpdateRecord)
	http.HandleFunc("/dbusers", users.DbUserView)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening on port 9090...")
	http.ListenAndServe(":9090", nil)
}
