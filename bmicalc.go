package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showBmicalc() {
	a := app.New()
	w := a.NewWindow("Bmi calculator")
	bmiimg := canvas.NewImageFromFile("bmi.png")
	bmiimg.FillMode = canvas.ImageFillOriginal
	label1 := canvas.NewText("", color.Black)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextSize = 15
	//result
	result := canvas.NewText("Bmi calculator", color.Black)
	result.Alignment = fyne.TextAlignCenter
	result.TextSize = 15
	//taking user input
	inputh := widget.NewEntry()
	inputh.SetPlaceHolder("please Enter your height in cm:")

	inputw := widget.NewEntry()
	inputw.SetPlaceHolder("please Enter your weight in kg:")
	//btn
	btn := widget.NewButton("Bmi", func() {
		h, _ := strconv.ParseFloat(inputh.Text, 64)
		w, _ := strconv.ParseFloat(inputw.Text, 64)
		result.Text = bmicalculator(h/100, w)
		result.Refresh()

	})
	bmicalci := container.NewVBox(
		label1,
		result,
		inputh,
		inputw,
		btn,
		bmiimg,
	)
	//	w := myApp.NewWindow("bmi")
	w.Resize(fyne.NewSize(500, 280))
	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, bmicalci),
	)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()

}

func bmicalculator(height, weight float64) string {
	//	height := float64(177.0 / 100)
	//	weight := float64(80.0)
	var bmi float64 = weight / math.Pow(height, 2)
	if bmi <= 18.5 {
		fmt.Println("you are underweight and bmi is :" + strconv.Itoa(int(bmi)))
		return "you are underweight and your bmi is: " + strconv.Itoa(int(bmi))
	} else if bmi > 18.5 && bmi < 24.9 {
		fmt.Println("you are normal and bmi is :" + strconv.Itoa(int(bmi)))
		return "you are normal and your bmi is : " + strconv.Itoa(int(bmi))
	} else if bmi > 25 && bmi < 29.9 {
		fmt.Println("you are overweight and bmi is " + strconv.Itoa(int(bmi)))
		return "you are overweight and bmi is: " + strconv.Itoa(int(bmi))
	} else if bmi > 30 && bmi < 34.9 {
		fmt.Println("you are in obesity class1  and bmi is :" + strconv.Itoa(int(bmi)))
		return "you are in obesity class 1 and bmi is : " + strconv.Itoa(int(bmi))
	} else if bmi > 35 && bmi < 39.9 {
		fmt.Println("you are in obesity class 2 and bmi is :" + strconv.Itoa(int(bmi)))
		return "you are in obesity class 2 and bmi is : " + strconv.Itoa(int(bmi))
	} else if bmi > 40 {
		fmt.Println("you are in obesity class 3 and bmi :" + strconv.Itoa(int(bmi)))
		return "you are in obesity class 3 and bmi is: " + strconv.Itoa(int(bmi))
	} else {
		fmt.Println("error")
		return "error"
	}
}
