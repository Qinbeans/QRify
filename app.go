package main

import (
	"context"
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) QRify(url string, size int, name string) (string, error) {
	//open dialog
	//get file path
	options := runtime.OpenDialogOptions{
		Title:                      "Select a directory",
		DefaultFilename:            "",
		DefaultDirectory:           "./",
		Filters:                    []runtime.FileFilter{},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: false,
	}
	path, err := runtime.OpenDirectoryDialog(a.ctx, options)
	if err != nil {
		return "", err
	}
	file := path + "/" + name
	err = qrcode.WriteFile(url, qrcode.Medium, size, file)
	if err != nil {
		return "", err
	}
	return "QR code created", nil
}
