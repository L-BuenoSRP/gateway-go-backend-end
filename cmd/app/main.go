package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/repository"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/service"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// String de conexão com o banco de dados
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	invoiceRepository := repository.NewInvoiceRepository(db)
	accountService := service.NewAccountService(accountRepository)
	invoiceService := service.NewInvoiceService(invoiceRepository, *accountService)
	port := getEnv("HTTP_PORT", "8080")

	srv := server.NewServer(accountService, invoiceService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
