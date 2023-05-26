package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title           string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postData struct {
	Title    string `db:"title"`
	SubTitle string `db:"subtitle"`
	Content  string `db:"content"`
}

type featuredPostData struct {
	ID          string `db:"post_ID"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Tag         string `db:"tag"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}

type mostRecentPostData struct {
	ID          string `db:"post_ID"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Tag         string `db:"tag"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}

type createPostData struct {
	Title           string `json:"title"`
	SubTitle        string `json:"subtitle"`
	AuthorName      string `json:"authorname"`
	AuthorPhoto     string `json:"authorphoto"`
	AuthorPhotoName string `json:"authorphotoname"`
	Data            string `json:"data"`
	BigImage        string `json:"bigimage"`
	BigImageName    string `json:"bigimagename"`
	SmallImage      string `json:"smallimage"`
	SmallImageName  string `json:"smallimagename"`
	Content         string `json:"content"`
}

// Главная страница блога
func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featured_posts, err := getFeaturedPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		recently_posts, err := getMostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
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

// Страница просмотра постов

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] // Получаем orderID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку orderID в число

		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error1", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error2", 500)
			log.Println(err.Error())
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

// LOGIN
func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := postData{
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

// Admin
func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := postData{
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
	const query = `
		SELECT
			post_ID,
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

	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getMostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
			post_ID,
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

	var posts []mostRecentPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			content
		FROM
			post
		WHERE
			post_ID = ?
	`

	var order postData

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&order, query, postID)
	if err != nil {
		return postData{}, err
	}

	return order, nil
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "1Error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostData

		authorImg, err := base64.StdEncoding.DecodeString(req.AuthorPhoto)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileAuthor, err := os.Create("static/images/" + req.AuthorPhotoName)

		_, err = fileAuthor.Write(authorImg)

		bigImg, err := base64.StdEncoding.DecodeString(req.BigImage)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileBig, err := os.Create("static/images/" + req.BigImageName)

		_, err = fileBig.Write(bigImg)

		smallImg, err := base64.StdEncoding.DecodeString(req.AuthorPhoto)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileSmall, err := os.Create("static/images/" + req.AuthorPhotoName)

		_, err = fileSmall.Write(smallImg)

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "2Error", 500)
			log.Println(err.Error())
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "bd", 500)
			log.Println(err.Error())
			return
		}
		return
	}
}

func savePost(db *sqlx.DB, req createPostData) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			tag,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content
		)
		VALUES
		(
			?,
			"",
			?,
			?,
			CONCAT('../static/images/', ?),
			?,
			CONCAT('../static/images/', ?),
			?
		)
	`

	_, err := db.Exec(query, req.Title, req.SubTitle, req.AuthorName, req.AuthorPhotoName, req.Data, req.BigImageName, req.Content)
	return err
}
