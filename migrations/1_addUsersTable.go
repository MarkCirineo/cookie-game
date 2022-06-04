package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(d migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := d.Exec(`CREATE TABLE users(
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`)
		return err
	}, func(d migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := d.Exec(`DROP TABLE users`)
		return err
	})
}