package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func gasto() {
	fmt.Println("gasto")
}

func main() {
	a := app.New()
	w := a.NewWindow("CONTABLE APP")
	w.Resize(fyne.NewSize(300, 600)) // ancho y alto

	// titulo centrado arriba
	titulo := widget.NewLabel("CONTABLE APP \nHola, [nombre de usuario]")
	titulo.Alignment = fyne.TextAlignCenter

	// boton ingreso
	botonIngreso := widget.NewButton("Ingreso", func() {
		fmt.Println("ingreso")
		// boton para vovler a la pantalla anterior
		volver := widget.NewButton("Volver", func() {
			fmt.Println("volver")
		})

		// titulo centrado arriba
		titulo := widget.NewLabel("INGRESO")
		titulo.Alignment = fyne.TextAlignCenter

		// NewVBox agrupa los elementos en un contenedor vertical
		ingreso := container.NewVBox(titulo, volver)

		// setup content
		w.SetContent(ingreso)
	})

	// boton gasto
	botonGasto := widget.NewButton("Gasto", gasto)

	// centrar botones verticalmente
	botones := container.NewVBox(botonIngreso, botonGasto)

	// NewVBox agrupa los elementos en un contenedor vertical
	inicio := container.NewVBox(titulo, botones)

	// setup content
	w.SetContent(inicio)

	w.ShowAndRun()
}
