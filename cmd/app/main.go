package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Config struct {
	OpenWeatherAPIKey string `env:"OPENWEATHER_API_KEY" env-required:"true"`
}

func main() {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		// Пробуем загрузить из .env файла
		err = cleanenv.ReadConfig(".env", &cfg)
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			fmt.Println("Please create .env file with OPENWEATHER_API_KEY")
			os.Exit(1)
		}
	}

	var city string
	if len(os.Args) > 1 {
		city = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Print("Enter the city name: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		city = strings.TrimSpace(input)
	}

	weather, err := getWeatherData(city, cfg.OpenWeatherAPIKey)
	if err != nil {
		fmt.Println("Error feastching weather data:", err)
		return
	}

	displayWeatherData(weather)
}

func displayWeatherData(weather *WeatherData) {
	fmt.Printf("\nWeather for %s: \n", weather.Name)
	if len(weather.Weather) > 0 {
		caser := cases.Title(language.English)
		fmt.Printf("Description: %s\n", caser.String(weather.Weather[0].Description))
	}
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d%%\n\n", weather.Main.Humidity)
}

// cd40aa1acf82425aa46394937c11cfa6
