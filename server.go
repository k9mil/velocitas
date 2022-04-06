package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		opts := badger.DefaultOptions("/tmp/badger")
		opts.Logger = nil
		db, err := badger.Open(opts)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Inserting 100,000 records.")
		start := time.Now()

		for n := 0; n < 100000; n++ {
			log.Printf("Item: %d", n)
			setTransaction(db)
		}

		log.Printf("Completed in %v", time.Since(start))

		defer db.Close()

		return c.Render("index", fiber.Map{
			"title": "Hello, World!",
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

func getTransaction(db *badger.DB) error {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("path"))

		if err != nil {
			log.Fatal(err)
		}

		err = item.Value(func(val []byte) error {
			fmt.Printf("%s\n", val)
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
