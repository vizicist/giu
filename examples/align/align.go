package main

import "github.com/vizicist/giu"

var text string

func loop() {
	giu.Window("window").Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("I'm a centered label"),
			giu.Button("I'm a centered button"),
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm a alined to right label"),
			giu.InputText(&text),
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm the label"),
			giu.Layout{
				giu.Label("I'm th e other label embeded in another layout"),
				giu.Label("I'm the next label"),
			},
			giu.Label("I'm the last label"),
		),
		giu.Label("Buttons in row:"),
		giu.Align(giu.AlignCenter).To(
			giu.Row(
				giu.Button("button 1"),
				giu.Button("button 2"),
			),
		),

		giu.Label("manual alignment"),
		giu.AlignManually(
			giu.AlignCenter,
			giu.Button("I'm button with 100 width").
				Size(100, 30),
			100, false,
		),
		giu.AlignManually(
			giu.AlignCenter,
			giu.InputText(&text),
			100, true,
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Alignment demo", 640, 480, 0)
	wnd.Run(loop)
}
