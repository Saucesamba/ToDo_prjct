// Слой работы с БД
//Реализация операций CRUD для Task

package db

//db.Query(): выполняет SQL-запрос, который возвращает набор строк (sql.Rows).
//rows.Next(): переходит к следующей строке результатов.
//rows.Scan(): читает значения из текущей строки и записывает их в переменные.
//db.QueryRow(): выполняет запрос, который возвращает одну строку.
//db.Exec(): выполняет SQL-запрос, который не возвращает строк (например, INSERT, UPDATE, DELETE).

import (
	"MyProject/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {

	Source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := sql.Open("postgres", Source)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}
	return db, nil
}
