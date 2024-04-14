package processing

import (
	"image"
	"image/color"
	"sync"
	"transformer/messages"
)

func ProcessImage(message *messages.ImageMessage) (image.Image, error) {
	// This could be extended in the future with different types of processing. For now, simply grayscaling the image should be enough
	out := make([][]color.RGBA, len(message.Data));
	for i:=0; i < len(message.Data); i++ {
		out[i] = make([]color.RGBA, len(message.Data[0].Data))
	}
	var group sync.WaitGroup;
	for x:=0; x<len(message.Data); x++ {
		group.Add(1);
		go func(x int) {
			defer group.Done();
			for y:=0; y<len(message.Data[0].Data); y++ {
				px := message.Data[x].Data[y];
				r := px.R;
				b := px.B;
				g := px.G;
				a := px.A;
				grey := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
				c:=uint8(grey+0.5)
				out[x][y] = color.RGBA{R: c, G: c, B: c, A: uint8(a)}
			}
		}(x)
	}
	group.Wait()
	rect := image.Rect(0,0,len(out),len(out[0]))
	outImg := image.NewRGBA(rect)
	for x:=0; x<len(out); x++ {
		for y:=0; y<len(out[x]); y++ {
			outImg.Set(x, y, out[x][y])

		}
	}
	return outImg, nil
}
