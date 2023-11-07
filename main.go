package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "image/png"
	"os"
	"os/exec"
)

func main() {
	navigator := app.New()
	mainWindow := navigator.NewWindow("Navigator")
	mainWindow.Resize(fyne.NewSize(1080, 720))

	background := canvas.NewImageFromResource(resourceImgMainBg)

	// ******* CAMERA BUTTON, NEED TO BE EXTRACTED //
	camera := widget.NewButton("", func() {
		w3 := navigator.NewWindow("Camera")
		closeBtn := widget.NewButton("Close", func() { w3.Close() })
		cont := container.New(layout.NewGridLayout(1), closeBtn)
		w3.SetContent(cont)
		w3.Resize(fyne.NewSize(1080, 720))
		w3.Show()
	})
	image := canvas.NewImageFromResource(resourceImgCamera)
	image.FillMode = canvas.ImageFillOriginal
	// *********************** //

	buttons := container.New(layout.NewGridLayout(4),
		createBtn(resourceImgNavit, runNavit),
		camera,
		// createBtn(resourceImgCamera, runCamera),
		createBtn(resourceImgSensors, runSensors),
		createBtn(resourceImgQuit, runQuit),
	)
	buttonArea := container.New(layout.NewCenterLayout(), buttons)
	content := container.New(layout.NewStackLayout(), background, buttonArea)

	// ******** //
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
	cmd := exec.Command("code")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Running: ", string(out))
}

func runCamera(a *fyne.App) {
	//cameraWindow := a.NewWindow("Larger")
	//cameraWindow.SetContent(widget.NewLabel("More content"))
	//cameraWindow.Resize(fyne.NewSize(100, 100))
	//cameraWindow.Show()
}

func runSensors() {

}

func runQuit() {
	// close the app
	os.Exit(0)
}
