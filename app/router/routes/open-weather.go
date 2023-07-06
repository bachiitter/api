package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type WeatherData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

func OpenWeatherRouter(app fiber.Router) {
	router := app.Group("/weather")

	router.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   5 * time.Minute,
		CacheControl: true,
	}))

	router.Get("/:location?", func(c *fiber.Ctx) error {

		apiKey := os.Getenv("OPEN_WEATHER_API_KEY")

		location := c.Query("location")

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", location, apiKey)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var data WeatherData

		if err := json.Unmarshal(body, &data); err != nil {
			return err
		}

		return c.JSON(data)
	})

}
