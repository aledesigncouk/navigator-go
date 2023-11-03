package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	_ "image/gif"
	_ "image/png"
)

//go:embed assets/quit-icon.png
var imgQuitBin []byte

//go:embed assets/back-icon.png
var imgBackBin []byte

//go:embed assets/camera-icon.png
var imgCameraBin []byte

//go:embed assets/navit-icon.png
var imgNavitBin []byte

//go:embed assets/sensors-icon.png
var imgSensorsBin []byte

//go:embed assets/main_bg.gif
var imgMainBgBin []byte

var resourceImgQuit = &fyne.StaticResource{
	StaticName:    "quit-icon.png",
	StaticContent: imgQuitBin,
}

var resourceImgBack = &fyne.StaticResource{
	StaticName:    "back-icon.png",
	StaticContent: imgBackBin,
}

var resourceImgCamera = &fyne.StaticResource{
	StaticName:    "camera-icon.png",
	StaticContent: imgCameraBin,
}

var resourceImgNavit = &fyne.StaticResource{
	StaticName:    "navit-icon.png",
	StaticContent: imgNavitBin,
}

var resourceImgSensors = &fyne.StaticResource{
	StaticName:    "sensors-icon.png",
	StaticContent: imgSensorsBin,
}

var resourceImgMainBg = &fyne.StaticResource{
	StaticName:    "main_bg.png",
	StaticContent: imgMainBgBin,
}
