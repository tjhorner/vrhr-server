package ui

import (
	"bytes"
	"fmt"
	"image"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/skip2/go-qrcode"
)

func Error(err error) {
	fmt.Printf("Fatal error occurred: %v\n", err)

	fyneApp := app.New()

	titleLabel := widget.NewLabel("A critical error occurred")
	titleLabel.Wrapping = fyne.TextWrapOff
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle.Bold = true

	w := fyneApp.NewWindow("VRHR Error")

	w.SetFixedSize(true)

	w.SetContent(
		widget.NewVBox(
			titleLabel,
			widget.NewLabel(err.Error()),
			widget.NewButton("OK", func() {
				fyneApp.Quit()
			}),
		),
	)

	w.ShowAndRun()
}

func Run(pairingCode string) {
	fyneApp := app.New()

	qrPng, _ := qrcode.Encode(pairingCode, qrcode.Medium, 256)
	qrImg, _, _ := image.Decode(bytes.NewReader(qrPng))

	qr := canvas.NewImageFromImage(qrImg)
	qr.FillMode = canvas.ImageFillOriginal

	welcomeLabel := widget.NewLabel("The VRHR Server is now running and listening for connections.\n\nUse the QR code below to pair it with the app, or copy the pairing code manually.")
	welcomeLabel.Wrapping = fyne.TextWrapWord
	welcomeLabel.Alignment = fyne.TextAlignCenter

	w := fyneApp.NewWindow("VRHR Server")

	w.SetFixedSize(true)

	w.SetContent(
		widget.NewVBox(
			welcomeLabel,
			qr,
			widget.NewButton("Copy Pairing Code", func() {
				w.Clipboard().SetContent(pairingCode)
				fyneApp.SendNotification(&fyne.Notification{
					Title:   "VRHR Server",
					Content: "URL copied. Go to the app and paste it.",
				})
			}),
			widget.NewButton("OBS Setup", func() {
				ghURL, _ := url.Parse("https://github.com/tjhorner/vrhr-server/wiki/OBS-Setup")
				fyneApp.OpenURL(ghURL)
			}),
			widget.NewButton("Quit", func() {
				fyneApp.Quit()
			}),
		),
	)

	w.ShowAndRun()
}
