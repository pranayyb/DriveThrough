package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pranayyb/DriveThrough/driver"
	carHandler "github.com/pranayyb/DriveThrough/handler/car"
	engineHandler "github.com/pranayyb/DriveThrough/handler/engine"
	carService "github.com/pranayyb/DriveThrough/service/car"
	engineService "github.com/pranayyb/DriveThrough/service/engine"
	carStore "github.com/pranayyb/DriveThrough/store/car"
	engineStore "github.com/pranayyb/DriveThrough/store/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()

	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)
	carHandler := carHandler.NewCarHandler(carService)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		log.Fatal("error while executing schema file")
	}

	router.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("engine/{id}", engineHandler.GetEngineById).Methods("GET")
	router.HandleFunc("engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("Port: %s", port)
	log.Printf("Server listening on port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}

func executeSchemaFile(db *sql.DB, schemaFile string) error {
	sqlFile, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
