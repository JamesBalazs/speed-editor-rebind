# speed-editor-rebind
Cross platform GUI for rebinding DaVinci Resolve Speed Editor keys

*WORK IN PROGRESS* - I'll remove this notice when a V1 release is available

![Screenshot](img/ui-example.png)

## Dev notes

Fedora deps for wails3:

```sudo dnf install gtk3-devel webkit2gtk4.1-devel pkgconf gcc```

Generate keyboard layout:

```
go run cmd/generate-layout/main.go > frontend/partials/keyboard.html
```
(copy into index.html)

Install wails3 CLI:

```
go install github.com/wailsapp/wails/v3/cmd/wails3@latest
```
