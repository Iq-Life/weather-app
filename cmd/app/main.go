package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Config struct {
	OpenWeatherAPIKey string `env:"OPENWEATHER_API_KEY" env-required:"true"`
	Port              string `env:"PORT" env-default:"4200"`
}

type PageData struct {
	City    string
	Weather *WeatherData
	Error   string
}

var (
	cfg  Config
	tmpl *template.Template
)

func main() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		err = cleanenv.ReadConfig(".env", &cfg)
		if err != nil {
			fmt.Println("Error loading configuration:", err)
		}
	}

	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", homeHandler)
	router.Get("/{city}", weatherHandler)

	log.Printf("Server starting on http://localhost:%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	data := PageData{
		City: "Введите город в адресной строке (например: /moscow)",
	}

	if err := tmpl.Execute(writer, data); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func weatherHandler(writer http.ResponseWriter, request *http.Request) {
	city := strings.ToLower(chi.URLParam(request, "city"))

	city = strings.ReplaceAll(city, "-", " ")

	weather, err := getWeatherData(city, cfg.OpenWeatherAPIKey)

	caser := cases.Title(language.Russian)
	data := PageData{
		City: caser.String(city),
	}

	if err != nil {
		data.Error = fmt.Sprintf("Ошибка получения погоды для города %s: %v", city, err)
	} else {
		data.Weather = weather
	}

	if err := tmpl.Execute(writer, data); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
