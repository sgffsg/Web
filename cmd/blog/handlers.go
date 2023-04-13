package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title           string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postPage struct {
	Title    string
	SubTitle string
}

type featuredPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Tag         string `db:"tag"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

type mostRecentPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Tag         string `db:"tag"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

// Главная страница блога
func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featured_posts, err := getFeaturedPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		recently_posts, err := getMostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			Title:           "Escape",
			FeaturedPosts:   featured_posts,
			MostRecentPosts: recently_posts,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

// Используем функцию template.ParseFiles() для чтения файла шаблона.
// Если возникла ошибка, мы запишем детальное сообщение ошибки и
// используя функцию http.Error() мы отправим пользователю ответ:
// 500 Internal Server Error (Внутренняя ошибка на сервере)

// В случае ошибки парсинга - возвращаем 500
// Используем стандартный логгер для вывода ошбики в консоль
// Не забываем завершить выполнение ф-ии

// Затем мы используем метод Execute() для записи содержимогошаблона в тело HTTP ответа.
// Последний параметр в Execute() предоставляет возможность отправки динамических данных в шаблон.

// Страница просмотра постов
func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := postPage{
		Title:    "The Road Ahead",
		SubTitle: "The road ahead might be paved - it might not be.",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func getFeaturedPosts(db *sqlx.DB) ([]featuredPostData, error) {
	// Составляем SQL-запрос для получения записей для секции featured-posts
	const query = `
		SELECT
			title,
			subtitle,
			tag,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 1
		LIMIT 2
	`

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func getMostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	// Составляем SQL-запрос для получения записей для секции featured-posts
	const query = `
		SELECT
			title,
			subtitle,
			tag,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
		LIMIT 6
	`

	var posts []mostRecentPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}
