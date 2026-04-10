package main

import (
	"fmt"
	"strings"
	"time"

	speedEditor "github.com/JamesBalazs/speed-editor-client"
	"github.com/JamesBalazs/speed-editor-client/input"
)

var client speedEditor.SpeedEditorInt

func connectSpeedEditor() {
	for {
		var err error
		client, err = speedEditor.NewClient()
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
		go func() {
			// TODO - this is a hack to get the connected string to work when the device connects before Vue is ready
			// Need to rework this to keep some BE state about the connection
			for {
				app.Event.Emit("heartbeat", Heartbeat{Connected: true, Serial: deviceInfo.SerialNbr})
				time.Sleep(2 * time.Second)
			}
		}()

		client.SetKeyPressHandler(handleKeyPress)
		client.Poll()
	}
}

func handleKeyPress(se speedEditor.SpeedEditorInt, report input.KeyPressReport) {
	for _, key := range report.Keys {
		app.Event.Emit(fmt.Sprintf("keyPress-%d", key.Id), map[string]string{"some": "data"})

		if mode, ok := speedEditorService.keyLedBehaviours[key.Id]; ok {
			if mode == "flash" {
				client.SetLeds([]uint32{key.Led})
				go func() {
					time.Sleep(250 * time.Millisecond)
					client.SetLeds([]uint32{})
				}()
			}
		}
	}
}
