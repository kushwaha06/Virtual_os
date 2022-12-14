package main

import (
	//"fyne.io/fyne/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("vertual os")
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget
var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("C:\\Users\\gyandeep kushwaha\\OneDrive\\Documents\\myos\\desktop.png")
	btn1 = widget.NewButtonWithIcon("Weather app", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})
	btn2 = widget.NewButtonWithIcon("calculator", theme.ContentAddIcon(), func() {
		showcalci()
	})
	btn3 = widget.NewButtonWithIcon("Gallery", theme.StorageIcon(), func() {
		showGalleryApp(myWindow)
	})
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor()
	})
	btn5 = widget.NewButtonWithIcon("Bmi calculator", theme.ContentAddIcon(), func() {
		showBmicalc()
	})
	btn6 = widget.NewButtonWithIcon("Ip address", theme.InfoIcon(), func() {
		showIpadd(myWindow)
	})
	DeskBtn = widget.NewButtonWithIcon("home btn", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})
	panelContent = container.NewVBox((container.NewGridWithColumns(7, DeskBtn, btn1,
		btn2, btn3, btn4, btn5, btn6)))

	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.CenterOnScreen()
	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	myWindow.ShowAndRun()
}
