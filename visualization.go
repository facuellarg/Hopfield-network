package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	indexImg := 0
	width, height := 800, 800
	// The themes create the content. Currently only a dark theme is offered for GUI elements.
	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Image viewer")
	myImg, _ := VectorToImage(steps[0], 8, 8, width/8)
	texture := driver.CreateTexture(myImg, 1)

	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClick(func(m gxui.MouseEvent) {
		window.RemoveAll()
		indexImg++
		if indexImg < len(steps) {
			newImg := theme.CreateImage()
			myImg, _ = VectorToImage(steps[indexImg], 8, 8, width/8)
			texture = driver.CreateTexture(myImg, 1)
			newImg.SetTexture(texture)
			// img.SetTexture(texture)
			window.AddChild(newImg)

		}

	})

	window.OnClose(driver.Terminate)
}
