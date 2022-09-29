package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/natnapat/simple_helpdesk/handler"
	"github.com/natnapat/simple_helpdesk/repository"
	"github.com/natnapat/simple_helpdesk/service"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	//initiation
	initTimeZone()
	initConfig()
	db := initDatabase()

	//hexagonal architecture setup
	ticketRepository := repository.NewTicketRepositoryDB(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	//router
	router := mux.NewRouter()
	router.HandleFunc("/tickets", ticketHandler.GetTickets).Methods(http.MethodGet)
	router.HandleFunc("/tickets/{status}", ticketHandler.GetTicketsByStatus).Methods(http.MethodGet)
	router.HandleFunc("/tickets", ticketHandler.CreateTicket).Methods(http.MethodPost)
	router.HandleFunc("/tickets/{id:[0-9]+}", ticketHandler.UpdateTicket).Methods(http.MethodPut)

	//cors setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(router)

	//http listening
	fmt.Println("listening...")
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), handler)
	http.Handle("/", router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = bkk
}

func initDatabase() *sqlx.DB {
	connStr := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v",
		viper.GetString("db.driver"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
		viper.GetString("db.sslmode"))

	db, err := sqlx.Open(viper.GetString("db.driver"), connStr)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
