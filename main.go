package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var alto float32
var ancho float32

func main() {
	// pantalla
	alto = 650
	ancho = 300
	a := app.New()
	w := a.NewWindow("Inicio")
	w.Resize(fyne.NewSize(ancho, alto))

	titulo := canvas.NewText("Contable APP", color.White)
	titulo.TextStyle.Bold = true
	titulo.TextStyle.Monospace = true

	subtitulo := canvas.NewText("Bienvenido, [usuario]", color.White)
	subtitulo.TextStyle.Monospace = true

	// centrar el titulo y el subtitulo
	titulo.Alignment = fyne.TextAlignCenter
	subtitulo.Alignment = fyne.TextAlignCenter

	blank := canvas.NewRectangle(color.Transparent)
	blank.SetMinSize(fyne.NewSize(0, 30))

	// Creamos un contenedor para el titulo y subtitulo
	tituloContenedor := container.New(layout.NewVBoxLayout(), blank, titulo, subtitulo, blank)

	// botones
	ingreso := widget.NewButton("Ingreso", func() {
		fmt.Println("Ingreso")
		wIngreso := a.NewWindow("Ingreso")
		wIngreso.Resize(fyne.NewSize(ancho, alto))

		volver := widget.NewButton("Volver", func() {
			wIngreso.Close()
			w.Show()
		})

		// titulo
		titulo := canvas.NewText("Ingreso", color.White)
		titulo.TextStyle.Bold = true
		titulo.Alignment = fyne.TextAlignCenter
		subtitulo := widget.NewLabel("Ingresa el dinero que ingresó a tu saldo\nSu origen y motivo")
		subtitulo.Alignment = fyne.TextAlignCenter

		blank := canvas.NewRectangle(color.Transparent)
		blank.SetMinSize(fyne.NewSize(0, 30))

		cSuperior := container.New(layout.NewVBoxLayout(), volver, blank, titulo, subtitulo, blank)

		// seccion de inputs

		cantidad := widget.NewLabel("Cantidad:")
		in_cantidad := widget.NewEntry()

		origen := widget.NewLabel("Origen:")
		in_origen := widget.NewEntry()

		motivo := widget.NewLabel("Motivo:")
		in_motivo := widget.NewEntry()

		cDatos := container.New(layout.NewVBoxLayout(), cantidad, in_cantidad, origen, in_origen, motivo, in_motivo)

		text := widget.NewLabel("")

		// boton de guardar
		guardar := widget.NewButton("Guardar", func() {
			text.SetText("Guardado \n" + in_cantidad.Text + "\n" + in_origen.Text + "\n" + in_motivo.Text)

		})

		cIngreso := container.New(layout.NewVBoxLayout(), cSuperior, cDatos, guardar, text)

		wIngreso.SetContent(cIngreso)
		w.Hide()
		wIngreso.Show()
	})

	gasto := widget.NewButton("Gasto", func() {
		fmt.Println("Gasto")
		wGasto := a.NewWindow("Gasto")
		wGasto.Resize(fyne.NewSize(ancho, alto))

		volver := widget.NewButton("Volver", func() {
			wGasto.Close()
			w.Show()
		})

		// titulo
		titulo := canvas.NewText("Gasto", color.White)
		titulo.TextStyle.Bold = true
		titulo.Alignment = fyne.TextAlignCenter
		subtitulo := widget.NewLabel("Ingresa el dinero que gastaste de tu saldo\nDetalle y motivo")
		subtitulo.Alignment = fyne.TextAlignCenter

		blank := canvas.NewRectangle(color.Transparent)
		blank.SetMinSize(fyne.NewSize(0, 30))

		cSuperior := container.New(layout.NewVBoxLayout(), volver, blank, titulo, subtitulo, blank)

		// seccion de inputs

		cantidad := widget.NewLabel("Cantidad:")
		in_cantidad := widget.NewEntry()

		motivo := widget.NewLabel("Motivo:")
		in_motivo := widget.NewEntry()

		detalle := widget.NewLabel("Detalle:")
		in_detalle := widget.NewEntry()

		cDatos := container.New(layout.NewVBoxLayout(), cantidad, in_cantidad, motivo, in_motivo, detalle, in_detalle)

		text := widget.NewLabel("")

		// boton de guardar
		guardar := widget.NewButton("Guardar", func() {
			text.SetText("Guardado \n" + in_cantidad.Text + "\n" + in_motivo.Text + "\n" + in_detalle.Text)

		})

		cGasto := container.New(layout.NewVBoxLayout(), cSuperior, cDatos, guardar, text)

		wGasto.SetContent(cGasto)
		w.Hide()
		wGasto.Show()
	})

	saldo := widget.NewButton("Saldo", func() {
		fmt.Println("Saldo")
		wSaldo := a.NewWindow("Saldo")
		wSaldo.Resize(fyne.NewSize(ancho, alto))

		volver := widget.NewButton("Volver", func() {
			wSaldo.Close()
			w.Show()
		})

		// titulo
		titulo := canvas.NewText("Saldo", color.White)
		titulo.TextStyle.Bold = true
		titulo.Alignment = fyne.TextAlignCenter

		blank := canvas.NewRectangle(color.Transparent)
		blank.SetMinSize(fyne.NewSize(0, 30))

		cSaldo := container.New(layout.NewVBoxLayout(), volver, blank, titulo, blank)

		wSaldo.SetContent(cSaldo)
		w.Hide()
		wSaldo.Show()
	})

	historial := widget.NewButton("Historial", func() {
		fmt.Println("Historial")
		wHistorial := a.NewWindow("Historial")
		wHistorial.Resize(fyne.NewSize(ancho, alto))

		volver := widget.NewButton("Volver", func() {
			wHistorial.Close()
			w.Show()
		})

		// titulo
		titulo := canvas.NewText("Historial", color.White)
		titulo.TextStyle.Bold = true
		titulo.Alignment = fyne.TextAlignCenter

		blank := canvas.NewRectangle(color.Transparent)
		blank.SetMinSize(fyne.NewSize(0, 30))

		cHistorial := container.New(layout.NewVBoxLayout(), volver, blank, titulo, blank)

		wHistorial.SetContent(cHistorial)
		w.Hide()
		wHistorial.Show()
	})

	salir := widget.NewButton("Salir", func() {
		fmt.Println("Salir")
		w.Close()
	})

	// contenedor para los botones
	botonesContenedor := container.New(layout.NewVBoxLayout(), ingreso, blank, gasto, blank, saldo, blank, historial, blank, salir)

	// contenedor para el titulo y los botones
	inicio := container.New(layout.NewVBoxLayout(), tituloContenedor, botonesContenedor)

	w.SetContent(inicio)
	w.SetMaster()
	w.Show()

	// no mover esta ejecucion
	a.Run()
}
