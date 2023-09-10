package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
)

func main() {
	a := app.New()
	w := a.NewWindow("Settings")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
				systray.SetTooltip("hello")
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}

// package main

// import (
// 	"fmt"
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/getlantern/systray"
// )

// var myApp fyne.App
// var myWindow fyne.Window

// func main() {
// 	myApp := app.New()
// 	myWindow = myApp.NewWindow("System Tray Example")
// 	myWindow.Resize(fyne.NewSize(400, 300))

// 	myWindow.SetContent(container.NewVBox(widget.NewLabel("Hello from Fyne!"),
// 		widget.NewButton("Quit", func() {
// 			myApp.Quit()
// 		})),
// 	)
// 	systray.Run(onReady, nil)

// 	myWindow.ShowAndRun()
// }

// func onReady() {
// 	systray.SetTitle("My App")
// 	systray.SetTooltip("My App Tooltip")
// 	mReady := systray.AddMenuItem("Settings", "Setting the application")
// 	mQuit := systray.AddMenuItem("Quit", "Quit the application")
// 	go func() {
// 		select {
// 		case <-mReady.ClickedCh:
// 			fmt.Printf("hello ready")
// 			openWindow()
// 		case <-mQuit.ClickedCh:
// 			fmt.Printf("hello quit")
// 			myApp.Quit()
// 			systray.Quit()
// 		}
// 	}()
// }

// func openWindow() {
// 	myWindow.Show()
// 	fmt.Printf("hello openwindow")
// }
