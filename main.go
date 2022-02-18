package main

import (
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
	"log"
	"math/rand"
)

type Payment struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Cost	int    `json: "cost"`
	Date	time.Time    `json: "date"`
	Type	string    `json: "type"`
	Comment	string    `json: "comment"`
	Category int `json:"category"`
  }
  
  type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
  }

var payments []Payment
var categories []Category

func getPayments(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payments)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range payments {
       if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
        }
    }
   json.NewEncoder(w).Encode(&Payment{})
}


func createPayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var payment Payment
    _ = json.NewDecoder(r.Body).Decode(&payment)
    payment.ID = strconv.Itoa(rand.Intn(1000000))
    payments = append(payments, payment) 
    json.NewEncoder(w).Encode(payment)
}

func updatePayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range payments {
        if item.ID == params["id"] {
            payments = append(payments[:index], payments[index+1:]...)
            var payment Payment
            _ = json.NewDecoder(r.Body).Decode(&payment)
            payment.ID = params["id"]
            payments = append(payments, payment) 
            json.NewEncoder(w).Encode(payment)
            return
        }
    }
    json.NewEncoder(w).Encode(payments)
}

func deletePayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range payments {
        if item.ID == params["id"] {
            payments = append(payments[:index], payments[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(payments)
}



func getCategories(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categories)
}

func getCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range categories {
       if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
        }
    }
   json.NewEncoder(w).Encode(&Category{})
}


func createCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var category Category
    _ = json.NewDecoder(r.Body).Decode(&category)
    category.ID = strconv.Itoa(rand.Intn(1000000))
    categories = append(categories, category) 
    json.NewEncoder(w).Encode(category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range categories {
        if item.ID == params["id"] {
            categories = append(categories[:index], categories[index+1:]...)
            var category Payment
            _ = json.NewDecoder(r.Body).Decode(&category)
            category.ID = params["id"]
            categories = append(categories, category) 
            json.NewEncoder(w).Encode(category)
            return
        }
    }
    json.NewEncoder(w).Encode(categories)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range categories {
        if item.ID == params["id"] {
            categories = append(categories[:index], categories[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(categories)
}
// type Payment struct {
// 	ID      string `json:"id"`
// 	Title   string `json:"title"`
// 	Cost	int    `json: "cost"`
// 	Date	time.Time    `json: "date"`
// 	Type	string    `json: "type"`
// 	Comment	string    `json: "comment"`
// 	Category *Category `json:"category"`
//   }

func main() {
    r := mux.NewRouter()

	categories = append(categories, Category{id: 1, Title: "Category1"})

	categories = append(categories, Category{id: 2, Title: "Category2"})

    payments = append(payments, Payment{
		ID: "1", 
		Title: "Payment1",
		Cost: 1000,
		Date: time.Now(),
		Type: "income",
		Comment: "First income",
		Category: 1,
	})

	category2 := Category{Title: "Category2"}

    payments = append(payments, Payment{
		ID: "2", 
		Title: "Payment2",
		Cost: 2500,
		Date: time.Now(),
		Type: "outcome",
		Comment: "First outcome",
		Category: 2,
	})

    r.HandleFunc("/payments", getPayments).Methods("GET")
    r.HandleFunc("/payments/{id}", getPayment).Methods("GET")
    r.HandleFunc("/payments", createPayment).Methods("POST")
    r.HandleFunc("/payments/{id}", updatePayment).Methods("PUT")
    r.HandleFunc("/payments/{id}", deletePayment).Methods("DELETE")

	r.HandleFunc("/categories", getCategories).Methods("GET")
    r.HandleFunc("/categories/{id}", getCategory).Methods("GET")
    r.HandleFunc("/categories", createCategory).Methods("POST")
    r.HandleFunc("/categories/{id}", updateCategory).Methods("PUT")
    r.HandleFunc("/categories/{id}", deleteCategory).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}