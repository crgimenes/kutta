package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

// The window icon ships inside the binary so a single executable carries its own
// icon on Windows and Linux (macOS uses the .app bundle icon instead). It is a
// 256px reduction of the assets/kutta.png master (the base for all icons).
//
//go:embed icon.png
var iconPNG []byte

// setWindowIcon decodes the embedded icon and installs it. A decode failure is
// non-fatal: the app just runs with the platform default icon.
func setWindowIcon() {
	img, err := png.Decode(bytes.NewReader(iconPNG))
	if err != nil {
		return
	}
	ebiten.SetWindowIcon([]image.Image{img})
}
