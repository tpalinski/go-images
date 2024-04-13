package imgloader

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func LoadImage(path string) ([][]color.Color, error) {
	img, err := os.Open(path);
	if err != nil {
		return nil, err
	}
	defer img.Close();
	image, err := jpeg.Decode(img);
	if err != nil {
		return nil, err
	}
	pixels := imageToArray(image);
	return pixels, nil
}

func imageToArray(img image.Image) [][]color.Color {
        dims := img.Bounds().Size();
	var pixels [][]color.Color;
	for i:=0; i < dims.X; i++ {
		var col []color.Color
		for j:=0; j < dims.Y; j++ {
			col = append(col, img.At(i, j));
		}
		pixels = append(pixels, col);
	}
	return pixels
}
