package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
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
	label.SetText("dumb-hook")
	widget.Layout().AddWidget(label)

	contentText := widgets.NewQTextEdit(nil)
	contentText.SetPlaceholderText("Message Contents")
	widget.Layout().AddWidget(contentText)

	usernameLine := widgets.NewQLineEdit(nil)
	usernameLine.SetPlaceholderText("Username")
	widget.Layout().AddWidget(usernameLine)

	avatarLine := widgets.NewQLineEdit(nil)
	avatarLine.SetPlaceholderText("Avatar Url")
	widget.Layout().AddWidget(avatarLine)

	urlLine := widgets.NewQLineEdit(nil)
	urlLine.SetPlaceholderText("Webhook Url")
	widget.Layout().AddWidget(urlLine)

	ttsBox := widgets.NewQCheckBox(nil)
	ttsBox.SetText("TTS")
	widget.Layout().AddWidget(ttsBox)

	button := widgets.NewQPushButton2("Send", nil)
	button.ConnectClicked(func(bool) {
		m := webhookMessage{contentText.ToPlainText(), usernameLine.Text(), avatarLine.Text(), ttsBox.IsChecked()}
		sendWebhook(m, urlLine.Text())
	})
	widget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
