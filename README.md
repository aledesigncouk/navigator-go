# Navigator (Flight of The Navigator)

Onboard computer system for campervan. **(work in progress)**

## Description

A single dashboard that provide multiple functions and connect different controls such Water Tank sensors, Reverse camera, GPS tracking, Satellite Navigation and Engine information.
Is the core of a larger system composed by the following projects:

- [Water Tank sensors](https://github.com/aledesigncouk/water-sensor)
- [GPS Tracker board](https://github.com/aledesigncouk/gps_tracker) <sup>1</sup>
- [GPS Tracker back end](https://github.com/aledesigncouk/gps-tracker-server) <sup>1</sup>
- [GPS Tracker front end](https://github.com/aledesigncouk/gps-tracker-frontend) <sup>1</sup>
- Canbus ECU integration <sup>2</sup>

<sup>1</sup> not public for now

<sup>2</sup> not started

### Hardware

- Raspberry Pi 4
- USB Video Capture device
- GPS Receiver (USB or GY-NEO6MV2 module)
- DC-DC Buck Step-down voltage converters

### Software

A custom DietPi is used as OS.
The main interface is created with Golang and Fyne library.
The navigation is based on the Open Source project Navit.

- [DietPi](https://dietpi.com/)
- [i3](https://i3wm.org/)
- [Go](https://go.dev/)
- [Fyne](https://fyne.io/)
- [Navit](https://github.com/navit-gps/navit)

### Schematics

### Blog updates

[Blog](https://www.yoroxid.com/category/navigator/)

### Notes
