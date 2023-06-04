package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Device Scanner")

	// IP 输入框和执行按钮
	startIP := widget.NewEntry()
	endIP := widget.NewEntry()
	executeButton := widget.NewButton("Execute", func() {
		// 执行按钮点击事件处理逻辑
		// 在这里可以调用扫描函数并更新界面
	})

	// 进度条
	progressBar := widget.NewProgressBar()

	// 在线终端区域
	onlineTerminals := container.NewVBox()

	// 主界面布局
	content := container.New(layout.NewBorderLayout(nil, nil, nil, onlineTerminals),
		container.NewVBox(
			widget.NewLabel("Start IP:"),
			startIP,
			widget.NewLabel("End IP:"),
			endIP,
			executeButton,
			progressBar,
		),
		onlineTerminals,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
