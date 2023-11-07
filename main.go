package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "image/png"
	"log"
)

func main() {
	navigator := app.New()
	mainWindow := navigator.NewWindow("Navigator")
	mainWindow.Resize(fyne.NewSize(1080, 720))

	background := canvas.NewImageFromResource(resourceImgMainBg)

	buttons := container.New(layout.NewGridLayout(4),
		createBtn(resourceImgNavit, runNavit),
		createBtn(resourceImgCamera, runCamera),
		createBtn(resourceImgSensors, runSensors),
		createBtn(resourceImgQuit, runQuit),
	)
	buttonArea := container.New(layout.NewCenterLayout(), buttons)

	content := container.New(layout.NewStackLayout(), background, buttonArea)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

func createBtn(resource *fyne.StaticResource, action func()) *fyne.Container {
	btn := widget.NewButton("", action)
	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillOriginal
	return container.New(layout.NewStackLayout(), btn, image)
}

// actions
func runNavit() {
	log.Println("Run Navit")
}

func runCamera() {
	log.Println("Run Camera")
}

func runSensors() {
	log.Println("Run Sensors")
}

func runQuit() {
	log.Println("Run Quit")
}
