package main

import (
	"flag"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"  

	"fyne.io/fyne/v2"
)

const (
	RES240      = "240p"
	RES360      = "360p"
	RES480      = "480p"
	RES720      = "720p"
	CONFIG_FILE = ".meeting-media"
	WM          = "WM"
	MM          = "MM"
)

func main() {	
	config := NewConfig()
	a := app.NewWithID("AttendantZoom")

	config.DebugMode = flag.Bool("d", false, "descarga falsa; imprimir información de depuración")
	flag.Parse()
	if *config.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("FUNCIONANDO EN MODO DE DEPURACIÓN")
	}

	progressBar := widget.NewProgressBar()
	config.Progress = &progress{0, "", progressBar}
	pbFormatter := func() string { return config.Progress.Title }
	config.Progress.ProgressBar.TextFormatter = pbFormatter

	mediaviewer := a.NewWindow("Visualizador")
	mediaviewer.Resize(fyne.NewSize(640, 360))
	initialLabel := widget.NewLabel("Selecciona una imagen")
	mediaviewer.SetContent(container.NewMax(
		initialLabel,
	))

	settingsTab := container.NewTabItem("", config.settingsGUI())
	settingsTab.Icon = theme.SettingsIcon()

	downloadedFilesTab := container.NewTabItem("Reunión", config.createDownloadedFilesView(mediaviewer))

	tabs := container.NewAppTabs(
		downloadedFilesTab,
		container.NewTabItem("Vida y Ministerio", config.mGUI(MM)),
		container.NewTabItem("Estudio de La Atalaya", config.mGUI(WM)),
		settingsTab,
	)

	w := a.NewWindow("Attendant Zoom")
	w.SetContent(container.NewVBox(tabs))

	mediaviewer.Show()
	w.ShowAndRun()
}