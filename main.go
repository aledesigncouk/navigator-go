package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	navigator := app.New()
	mainWindow := navigator.NewWindow("Navigator")
	mainWindow.Resize(fyne.NewSize(1080, 720))

	// text := widget.NewLabel("text 1")
	buttons := container.New(layout.NewGridLayout(4), btnNavit(), btnCamera(), btnSensors(), btnQuit())
	content := container.New(layout.NewCenterLayout(), buttons)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

// buttons
func btnNavit() *fyne.Container {
	btn := widget.NewButton("", runNavit)
	iconNavit := canvas.NewImageFromFile("imgs/navit-icon.png")
	iconNavit.FillMode = canvas.ImageFillOriginal
	navitButton := container.New(layout.NewStackLayout(), btn, iconNavit)
	return navitButton
}

func btnCamera() *fyne.Container {
	btn := widget.NewButton("", runCamera)
	iconCamera := canvas.NewImageFromFile("imgs/camera-icon.png")
	iconCamera.FillMode = canvas.ImageFillOriginal
	cameraButton := container.New(layout.NewStackLayout(), btn, iconCamera)
	return cameraButton
}

func btnQuit() *fyne.Container {
	btn := widget.NewButton("", runQuit)
	iconQuit := canvas.NewImageFromFile("imgs/quit-icon.png")
	iconQuit.FillMode = canvas.ImageFillOriginal
	quitButton := container.New(layout.NewStackLayout(), btn, iconQuit)
	return quitButton
}

func btnSensors() *fyne.Container {
	btn := widget.NewButton("", runSensors)
	iconSensors := canvas.NewImageFromFile("imgs/sensors-icon.png")
	iconSensors.FillMode = canvas.ImageFillOriginal
	sensorsButton := container.New(layout.NewStackLayout(), btn, iconSensors)
	return sensorsButton
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
