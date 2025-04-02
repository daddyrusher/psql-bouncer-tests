package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS test_table (id UUID PRIMARY KEY)`)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	testUUID := uuid.New()
	_, err = db.Exec(ctx, "INSERT INTO test_table (id) VALUES ($1)", testUUID)
	if err != nil {
		log.Fatalf("Ошибка вставки UUID: %v", err)
	}

	fmt.Println("Тестовые данные успешно вставлены!")

	var retrievedUUID uuid.UUID
	err = db.QueryRow(ctx, "SELECT id FROM test_table LIMIT 1").Scan(&retrievedUUID)
	if err != nil {
		log.Fatalf("Ошибка чтения UUID: %v", err)
	}

	fmt.Printf("UUID из базы: %s\n", retrievedUUID)
}
