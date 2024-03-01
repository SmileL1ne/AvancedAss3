package main

import (
	"architecture_go/pkg/store/postgres"
	deliveryHTTP "architecture_go/services/contact/internal/delivery/http"
	pgRepo "architecture_go/services/contact/internal/repository/storage/postgres"
	ucContact "architecture_go/services/contact/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	dbInfo := &postgres.Database{
		Host:     "localhost",
		Port:     "7777",
		User:     "postgres",
		Password: "postgres",
		Name:     "ap_ass3",
	}

	db, err := postgres.OpenDB(dbInfo)
	if err != nil {
		panic("nah" + err.Error())
	}

	contactRepo := pgRepo.New(db)
	contactUC := ucContact.New(contactRepo)
	contactDelivery := deliveryHTTP.NewContactDelivery(contactUC)

	http.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			contactDelivery.CreateContact(w, r)
		case http.MethodGet:
			contactDelivery.ViewContact(w, r)
		case http.MethodPut:
			contactDelivery.UpdateContact(w, r)
		case http.MethodDelete:
			contactDelivery.DeleteContact(w, r)
		}
	})

	http.HandleFunc("/groups/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			contactDelivery.CreateGroup(w, r)
		case http.MethodGet:
			contactDelivery.GetGroupByID(w, r)
		case http.MethodPut:
			contactDelivery.UpdateGroup(w, r)
		case http.MethodDelete:
			contactDelivery.DeleteGroup(w, r)
		default:
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/groups/add-contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			contactDelivery.InsertContactToGroup(w, r)
		} else {
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/groups/remove-contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			contactDelivery.DeleteContactFromGroup(w, r)
		} else {
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})

	err = http.ListenAndServe("127.0.0.1:7000", nil)
	log.Fatal(err)
}
