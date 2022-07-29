package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	//	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window) {
	//a := app.New()
	//w := a.NewWindow("weather app")
	//api using
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=72e18f3afb560d6316f327fb4bfd491b")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	Weather, err := UnmarshalWelcome(body)
	if err != nil {
		fmt.Println(err)
	}
	image := canvas.NewImageFromFile("weather-forecast.png")
	image.FillMode = canvas.ImageFillContain
	label1 := canvas.NewText("weather app", color.Black)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label2 := canvas.NewText(fmt.Sprintf("country %s", Weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("wind speed %.2f", Weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("temprature %.2f", Weather.Main.Temp), color.Black)
	label5 := canvas.NewText(fmt.Sprintln("Humidity ", Weather.Main.Humidity), color.Black)
	//label6 := canvas.NewText(fmt.Sprintln("clouds %s", Weather.Clouds), color.Black)
	weatherContainer := container.NewVBox(
		label1,
		image,
		label2,
		label3,
		label4,
		label5,
		container.NewGridWithColumns(1),
	)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))
	w.Show()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
