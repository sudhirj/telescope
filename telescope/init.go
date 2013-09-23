package telescope

import "github.com/gographics/imagick/imagick"

func Start() {
	imagick.Initialize()
}

func Stop() {
	imagick.Terminate()
}
