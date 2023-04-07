package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	dbDriverName = "mysql"
)

func main() {

	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе

	//createDatabase(dbx)

	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))
	mux.HandleFunc("/post", post)

	// Реализуем отдачу статики
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Start server")
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	// Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, "root:root@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}

func createDatabase(db *sqlx.DB) {
	const query = "CREATE SCHEMA blog;" +
		"USE blog;" +
		"CREATE TABLE blog.post" +
		"(" +
		"`post_id`      INT NOT NULL AUTO_INCREMENT," +
		"`title`        VARCHAR(255) NOT NULL," +
		"`subtitle`     VARCHAR(255) NOT NULL," +
		"`tag`          VARCHAR(255) NOT NULL," +
		"`image_url`    VARCHAR(255)," +
		"`author`       VARCHAR(255)," +
		"`author_url`   VARCHAR(255)," +
		"`publish_date` VARCHAR(255) NOT NULL," +
		"`featured`     TINYINT(1) DEFAULT 0," +
		"PRIMARY KEY (`post_id`)" +
		") ENGINE = InnoDB" +
		"CHARACTER SET = utf8mb4" +
		"COLLATE utf8mb4_unicode_ci;"
	// Составляем SQL-запрос для создания таблицы post

	request, err := db.Query(query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		panic(err)
	}
	defer request.Close()
}
