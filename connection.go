package main

import (
	"fmt"
	"strings"
	"time"

	speedEditor "github.com/JamesBalazs/speed-editor-client"
	"github.com/JamesBalazs/speed-editor-client/input"
	"github.com/JamesBalazs/speed-editor-client/keys"
)

var (
	client   speedEditor.SpeedEditorInt
	keysById = keys.ById()
)

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
		app.Event.Emit(fmt.Sprintf("keyPress-%d", key.Id))

		if status, ok := speedEditorService.LedStatus[key.Id]; ok {
			if status.mode == "flash" {
				speedEditorService.SetKeyLit(key.Id, time.Now(), 250*time.Millisecond)
				consolidateLeds()

				go func() {
					time.Sleep(250 * time.Millisecond)

					consolidateLeds()
				}()
			}
		}
	}
}

func consolidateLeds() {
	leds := []uint32{}
	jogLeds := []uint8{}
	for keyId, status := range speedEditorService.LedStatus {
		if key, ok := keysById[keyId]; ok && key.Led != keys.LED_NONE {
			if time.Now().Before(status.litAt.Add(status.litFor)) {
				leds = append(leds, key.Led)
			} else {
				status.litAt = time.Time{}
				status.litFor = 0
			}
		} else if ok && key.JogLed != keys.LED_NONE {
			if time.Now().Before(status.litAt.Add(status.litFor)) {
				jogLeds = append(jogLeds, key.JogLed)
			} else {
				status.litAt = time.Time{}
				status.litFor = 0
			}
		}
	}

	client.SetLeds(leds)
	client.SetJogLeds(jogLeds)
}
