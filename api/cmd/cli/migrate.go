package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	migrations "github.com/bjvanbemmel/benkyou"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	ErrNoArgumentProvided error = errors.New("No argument provided")
)

type CommandMigrate struct {
	db *sql.DB
}

func (c CommandMigrate) Execute(w io.ReadWriter, args ...string) error {
	if len(args) < 1 {
		return ErrNoArgumentProvided
	}

	command := args[0]

	godotenv.Load("../../../.env")
	conn := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	c.db, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer c.db.Close()

	switch command {
	case "status":
		return c.status()
	case "up":
		return c.up()
	case "down":
		return c.down()
	}

	return nil
}

func (c CommandMigrate) status() error {
	goose.SetBaseFS(migrations.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Status(c.db, "migrations"); err != nil {
		return err
	}

	return nil
}

func (c CommandMigrate) up() error {
	goose.SetBaseFS(migrations.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(c.db, "migrations"); err != nil {
		return err
	}

	return nil
}

func (c CommandMigrate) down() error {
	goose.SetBaseFS(migrations.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Down(c.db, "migrations"); err != nil {
		return err
	}

	return nil
}
