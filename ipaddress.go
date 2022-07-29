package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	//	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showIpadd(w fyne.Window) {
	//	a := app.New()
	//	w := a.NewWindow("ip address")

	//	w.Resize(fyne.NewSize(500, 500))
	myWindow.Resize(fyne.NewSize(500, 280))
	image1 := canvas.NewImageFromFile("new1.png")
	image1.FillMode = canvas.ImageFillOriginal
	labeltitle := widget.NewLabel("What is my county ip address?")
	labelip := widget.NewLabel("your ip address is..")
	labelvalue := widget.NewLabel("....")
	labelvalue2 := widget.NewLabel("....")
	labelvalue3 := widget.NewLabel("....")
	btn := widget.NewButton("Run", func() {
		labelvalue.Text = myIP()
		labelvalue.Refresh()
		labelvalue2.Text = myCity()
		labelvalue2.Refresh()
		labelvalue3.Text = myCountry()
		labelvalue3.Refresh()
	})
	ipContainer := container.NewVBox(
		image1,
		labeltitle,
		labelip,
		labelvalue,
		labelvalue2,
		labelvalue3,
		btn,
	)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, ipContainer))
	w.Show()
}
func myIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Query
}

func myCity() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.City
}
func myCountry() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Country
}

type IP struct {
	Query   string
	City    string
	Country string
}
