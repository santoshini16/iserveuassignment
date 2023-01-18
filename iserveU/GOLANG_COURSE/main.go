package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var db *gorm.DB
var err error

type NewEmployees struct {
	ID   string `json:"id"`
	Name string `json:"empName"`
}

func createNewEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newEmployee NewEmployees
	json.Unmarshal(reqBody, &newEmployee)
	db.Create(&newEmployee)
	fmt.Println("Create New Person")
	json.NewEncoder(w).Encode(newEmployee)
}

func main() {
	db, err = gorm.Open("mysql", "root:root@123@(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open!")
	} else {
		fmt.Println("Connection Established!")
	}
	defer db.Close()
	db.AutoMigrate(&NewEmployees{})
	router := mux.NewRouter()
	router.HandleFunc("/createEmployee", createNewEmployee).Methods("POST")
	http.ListenAndServe(":8000", router)

}
