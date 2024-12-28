package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"tgBotEcho/internal/config"
)

var Conn *sql.DB

func init() {
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.File.DBPort,
		config.File.DBUser,
		config.File.DBPass,
		config.File.DBName,
		config.File.DBSSLMode,
	)
	// Открываем соединение
	var err error
	Conn, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("error opening database: %v", err))
	}

	// Проверяем соединение
	if err = Conn.Ping(); err != nil {
		panic(fmt.Errorf("error opening database: %v", err))
	}
}
