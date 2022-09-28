package main

import (
	g "github.com/vizicist/giu"
	"github.com/vizicist/imgui-go"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0)
	wnd.Run(loop)
}
