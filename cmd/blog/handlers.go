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
	Content  string
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
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю ответ:
	// 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:           "Escape",
		FeaturedPosts:   getFeaturedPosts(),
		MostRecentPosts: getMostRecentPosts(),
	}

	// Затем мы используем метод Execute() для записи содержимогошаблона в тело HTTP ответа.
	// Последний параметр в Execute() предоставляет возможность отправки динамических данных в шаблон.
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
		Content:  "",
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
			Tag:         "photography",
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
			Tag:         "nature",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_sunny-side-up.png",
			Title:       "Sunny Side Upl",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			Tag:         "Photography",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_water-falls.png",
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Tag:         "relaxation",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_through-the-mist.png",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Tag:         "adventure",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_awaken-early.png",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Tag:         "vacation",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "static/img/most-recent-posts__post_try-it-always.png",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Tag:         "travel",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
	}
}
