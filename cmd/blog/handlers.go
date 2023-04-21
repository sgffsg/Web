package main

import (
	"html/template"
	"log"
	"net/http"
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
	Title       string
	Subtitle    string
	Tag         string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

type mostRecentPostData struct {
	Title       string
	Subtitle    string
	Tag         string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

// Главная страница блога
func index(w http.ResponseWriter, r *http.Request) {
	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := indexPage{
		Title:           "Escape",
		FeaturedPosts:   getFeaturedPosts(),
		MostRecentPosts: getMostRecentPosts(),
	}

	// Затем мы используем метод Execute() для записи содержимогошаблона в тело HTTP ответа.
	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

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

func getFeaturedPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			Tag:         "",
			ImgModifier: "static/img/featured-posts__post_the-road-ahead.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you've never been before.",
			Tag:         "adventure",
			ImgModifier: "static/img/featured-posts__post_from-top-down.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
	}
}

func getMostRecentPosts() []mostRecentPostData {
	return []mostRecentPostData{
		{
			ImgModifier: "static/img/most-recent-posts__post_still-standing-tall.png",
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Tag:         "",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_sunny-side-up.png",
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			Tag:         "",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_water-falls.png",
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Tag:         "",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_through-the-mist.png",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Tag:         "",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_awaken-early.png",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Tag:         "",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_try-it-always.png",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Tag:         "",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
	}
}
