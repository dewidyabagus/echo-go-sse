package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-sse/api"
	messageHandler "go-sse/api/v1/message"
	messageService "go-sse/business/message"
	"go-sse/config"
	messageRepository "go-sse/modules/message"
	"go-sse/modules/migration"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.PgHost, config.PgPort, config.PgUsername, config.PgPassword, config.PgDbname,
	)

	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic("Error database connection")
	}

	migration.MigrationTables(db)

	return db
}

func main() {
	// Konfigurasi App
	config := config.GetConfiguration()

	// Koneksi Database
	db := newDatabaseConnection(config)

	// Inisialisasi message repository
	msgRepo := messageRepository.NewRepository(db)

	// Inisialisasi message service
	msgSvc := messageService.NewService(msgRepo)

	// Inisialisasi controller
	msgHandle := messageHandler.NewController(msgSvc)

	// ====== WEB SERVER ======
	e := echo.New()
	api.RegisterRoutes(e, msgHandle)

	go func() {
		address := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)

		if err := e.Start(address); err != nil {
			log.Println("shutting down the server")
		}
	}()

	// tunggu sinyal interupsi untuk mematikan server dengan
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// batas waktu 10 detik untuk mematikan server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	// ====== END WEB SERVER ====== //
}
