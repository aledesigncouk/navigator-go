package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	_ "image/gif"
	_ "image/png"
)

//go:embed imgs/quit-icon.png
var imgQuitBin []byte

//go:embed imgs/back-icon.png
var imgBackBin []byte

//go:embed imgs/camera-icon.png
var imgCameraBin []byte

//go:embed imgs/navit-icon.png
var imgNavitBin []byte

//go:embed imgs/back-icon.png
var imgSensorsBin []byte

//go:embed imgs/main_bg.gif
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
