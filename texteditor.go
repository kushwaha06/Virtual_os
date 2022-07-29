package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/widget"
)

var count int = 1

func showTextEditor() {
	a := app.New()
	w := a.NewWindow("text editor")
	//w.Resize(fyne.NewSize(600, 600))
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("text editor"),
		),
	)

	content.Add(widget.NewButton("add new file", func() {
		content.Add(widget.NewLabel("new file" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("enter  text")
	savebtn := widget.NewButton("Save", func() {
		savefiledialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textdata := []byte(input.Text)
				uc.Write(textdata)
			}, w)
		savefiledialog.SetFileName("new file" + strconv.Itoa(count-1) + ".txt")
		savefiledialog.Show()
	})
	openbtn := widget.NewButton("open", func() {
		openfiledailog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				Readdata, _ := ioutil.ReadAll(r)
				output := fyne.NewStaticResource("New file", Readdata)
				viewdata := widget.NewMultiLineEntry()
				viewdata.SetText(string(output.StaticContent))
				w := fyne.CurrentApp().NewWindow(string(output.StaticName))
				w.SetContent(container.NewScroll(viewdata))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
		openfiledailog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		openfiledailog.Show()
	})
	textcontainer := container.NewVBox(
		content,
		input,
		container.NewHBox(
			savebtn,
			openbtn,
		),
	)
	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, textcontainer),
	)
	w.ShowAndRun()
}
