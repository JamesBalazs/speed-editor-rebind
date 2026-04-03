package main

import (
	"embed"
	_ "embed"
	"log"
	"strings"
	"time"

	speedEditor "github.com/JamesBalazs/speed-editor-client"
	"github.com/sstallion/go-hid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[Heartbeat]("heartbeat")
	application.RegisterEvent[[]uint16]("keyPress")
}

var connected bool

type Heartbeat struct {
	Connected bool
	Error     string
	Serial    string
}

// main function serves as the application's entry point. It initializes the application and creates a window.
// It subsequently runs the application and logs any error that might occur.
func main() {
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "speed-editor-rebind",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&SpeedEditorService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Speed Editor Rebind",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		MinWidth:         800,
		MinHeight:        900,
	})

	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}
	defer hid.Exit()

	go func() {
		for {
			client, err := speedEditor.NewClient()
			if err != nil && strings.Contains(err.Error(), "No HID devices with requested VID/PID found in the system.") {
				continue
			} else if err != nil {
				app.Event.Emit("heartbeat", Heartbeat{Connected: false, Error: err.Error()})
				time.Sleep(2 * time.Second)
				continue
			}

			err = client.Authenticate()
			if err != nil {
				app.Event.Emit("heartbeat", Heartbeat{Connected: false, Error: err.Error()})
				time.Sleep(2 * time.Second)
				continue
			}

			deviceInfo := client.GetDeviceInfo()

			for {
				app.Event.Emit("heartbeat", Heartbeat{Connected: true, Serial: deviceInfo.SerialNbr})
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
