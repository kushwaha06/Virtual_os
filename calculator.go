package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showcalci() {
	// newa app
	//	a := app.New()
	//	w := a.NewWindow("Calculator")
	// resiging
	ishistory := false
	output := ""
	input := widget.NewLabel(output)
	historystr := ""
	history := widget.NewLabel(historystr)

	var prevArr []string

	prev := widget.NewButton("prev", func() {

		if ishistory {
			historystr = " "
			//	myWindow.Resize(fyne.NewSize(500, 280))
		} else {
			for i := len(prevArr) - 1; i >= 0; i-- {
				historystr = historystr + prevArr[i]
				historystr += "\n"
			}
		}
		ishistory = !ishistory
		history.SetText(historystr)
	})
	clear := widget.NewButton("clear", func() {
		output = ""
		input.SetText(output)
	})
	openbr := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	closebr := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	multiply := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	plus := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	minus := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	dot := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	backbtn := widget.NewButton("back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	divide := widget.NewButton("%", func() {
		output = output + "%"
		input.SetText(output)
	})
	equal := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + "=" + ans
				prevArr = append(prevArr, strToAppend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		input.SetText(output)
	})
	onebtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twobtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threebtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	fourbtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fivebtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixbtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	sevenbtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightbtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	ninebtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	zerobtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	//input := widget.NewLabel("")
	// created a buttton
	calcContainer := container.NewVBox(
		input, history, container.NewGridWithColumns(1, container.NewGridWithColumns(2,
			prev, backbtn),
			container.NewGridWithColumns(4, clear, openbr, closebr, divide),
			container.NewGridWithColumns(4, ninebtn, eightbtn, sevenbtn, multiply),
			container.NewGridWithColumns(4, fourbtn, fivebtn, sixbtn, plus),
			container.NewGridWithColumns(4, onebtn, twobtn, threebtn, minus),
			container.NewGridWithColumns(3, zerobtn, dot, equal),
		),
	)
	w := myApp.NewWindow("calci")
	w.Resize(fyne.NewSize(500, 280))
	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, calcContainer),
	)

	// running
	w.Show()
}
