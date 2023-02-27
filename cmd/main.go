package main

import (
	"domain-driven-design/config"
	"domain-driven-design/db"
	"domain-driven-design/domain/repository"
	"domain-driven-design/domain/service"
	"domain-driven-design/infrastructure/persistence/datastore"
	"domain-driven-design/infrastructure/transport/http_transport"
	"domain-driven-design/router"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	engine := echo.New()
	cfg := config.LoadConfig("../config")

	dbCon, err := db.Connect(cfg.SqlitePath)
	if err != nil {
		log.Fatalln(err)
	}

	customerDatastore := datastore.NewCustomerDatastore(dbCon)
	customerRepo := repository.NewCustomerRepository(customerDatastore)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := http_transport.NewCustomerHandler(customerService)
	app := &router.RestServer{
		Engine:          engine,
		Config:          cfg,
		CustomerHandler: customerHandler,
	}
	err = router.InitRouter(app)
	if err != nil {
		log.Fatalln(err)
	}
}
