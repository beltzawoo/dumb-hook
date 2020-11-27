package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Hello Widgets Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	label := widgets.NewQLabel(nil, 0)
	label.SetText("dumb-webhook")
	widget.Layout().AddWidget(label)

	window.Show()

	app.Exec()
}
