package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	//	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {
	//	a := app.New()
	//	w := a.NewWindow("Gallery")
	root_src := "C:\\Users\\gyandeep kushwaha\\OneDrive\\Pictures"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	var picArr []string
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" {
				picArr = append(picArr)
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}
	//container.NewTabItem("image", canvas.NewImageFromFile(picArr[0])),)
	tabs.SetTabLocation(container.TabLocationTrailing)
	w.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(500, 280))
	//w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
