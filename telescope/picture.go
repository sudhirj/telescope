package telescope

import "github.com/gographics/imagick/imagick"
import "net/http"
import "io/ioutil"

type picture struct {
	wand *imagick.MagickWand
}

func (p *picture) WriteTo(w http.ResponseWriter) {
	w.Write(p.wand.GetImageBlob())
}

func (p *picture) SizeToWidth(newWidth uint) {
	ratio := float64(p.wand.GetImageHeight()) / float64(p.wand.GetImageWidth())
	newHeight := ratio * float64(newWidth)
	p.wand.ResizeImage(newWidth, uint(newHeight), imagick.FILTER_LANCZOS, 1)
}

func (p *picture) Destroy() {
	p.wand.Destroy()
}

func LoadPicture(origin string) *picture {
	p := new(picture)
	originalImage, _ := http.Get(origin)
	defer originalImage.Body.Close()
	p.wand = imagick.NewMagickWand()
	img, _ := ioutil.ReadAll(originalImage.Body)
	p.wand.ReadImageBlob(img)
	return p
}
