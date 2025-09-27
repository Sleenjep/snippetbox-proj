package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Sleenjep/snippetbox-proj/snippetbox/pkg/models/postgresql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgresql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")

	passBytes, err := os.ReadFile("../secrets/pg_pass.txt")
	if err != nil {
		log.Fatal(err)
	}
	password := strings.TrimSpace(string(passBytes))

	dsn := flag.String("dsn",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			"user_pg", password, "localhost", 5432, "snippetbox-pg-db"),
		"PostgreSQL DSN",
	)

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgresql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера (http://localhost:4000/)")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// var version string
	// if err := db.QueryRow("SELECT version()").Scan(&version); err != nil {
	// 	return nil, fmt.Errorf("ошибка тестового запроса: %w", err)
	// }
	// log.Printf("Подключено к PostgreSQL: %s\n", version)

	return db, nil
}
