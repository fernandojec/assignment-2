package main

import (
	"fmt"
	"log"

	"github.com/fernandojec/assignment-2/config"
	"github.com/fernandojec/assignment-2/domain/users"
	"github.com/fernandojec/assignment-2/pkg/dbconnect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	var err error
	config.AppConfig, err = config.LoadConfig()
	// err := godotenv.Load("../../.env")
	// err = godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	dbx, err := dbconnect.ConnectSqlx(dbconnect.DBConfig{
		Host:       config.AppConfig.Postgres.Host,
		Port:       config.AppConfig.Postgres.Port,
		Dbname:     config.AppConfig.Postgres.DbName,
		Dbuser:     config.AppConfig.Postgres.User,
		Dbpassword: config.AppConfig.Postgres.Password,
		Sslmode:    config.AppConfig.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("Cannot connect to DB:%v", err)
	}

	// store := sessions.NewCookieStore([]byte(config.AppConfig.Session.AuthSessionId))
	store := session.New(session.ConfigDefault)
	engine := html.New("views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/web",
		},
		StatusCode: 301,
	}))
	web := app.Group("web")
	users.RouterWebInit(web, dbx, store)

	log.Fatal(app.Listen(config.AppConfig.App.BaseWebPort))
}
