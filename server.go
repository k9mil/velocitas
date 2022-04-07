package main

import (
	"log"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
)

func main() {
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		opts := badger.DefaultOptions("/tmp/badger")
		opts.Logger = nil
		db, err := badger.Open(opts)

		if err != nil {
			log.Fatal(err)
		}

		start := time.Now()

		for n := 0; n < 10000; n++ {
			setTransaction(db)
		}

		calculatedTime := time.Since(start).Seconds()

		defer db.Close()
		return c.Render("index", fiber.Map{
			"time": calculatedTime,
		})
	})

	app.Static("/", "./views")
	log.Fatal(app.Listen(":3000"))
}

func setTransaction(db *badger.DB) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(randomdata.SillyName()), []byte(randomdata.SillyName()))

	if err != nil {
		log.Fatal(err)
	}

	if err := txn.Commit(); err != nil {
		log.Fatal(err)
	}

	return nil
}
