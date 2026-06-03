package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

type Player struct {
	ID    int
	Name  string
	Email string
}

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	var data Request
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid2 request", http.StatusBadRequest)
		return
	}

	if data.Message == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}

	dsn := "host=localhost port=5432 user=postgres password=111 dbname=test1 sslmode=disable"

	// Открываем подключение
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Подключено к PostgreSQL!")

	// Читаем данные
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Обработка результатов
	var players []Player = make([]Player, 0)
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.ID, &p.Name, &p.Email); err != nil {
			http.Error(w, fmt.Sprintf("scan error: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Printf("ID: %d, Имя: %s, Email: %s\n", p.ID, p.Name, p.Email)
		players = append(players, p)
	}
	if len(players) == 0 {
		http.Error(w, "no players found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func main() {

	http.HandleFunc("/api", handlerAPI)
	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}
