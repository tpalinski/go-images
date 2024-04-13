package imgloader

import (
	"image"
	"image/color"
	"image/jpeg"
	"mock-producer/messages"
	"os"

	"google.golang.org/protobuf/proto"
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

func DeserializeImage(image [][]color.Color, name string) ([]byte, error) {
	ser := &messages.ImageMessage {}
	ser.Name = name;
	for i:=0; i < len(image); i++ {
		col := &messages.ImageMessage_Col {}
		for j:=0; j < len(image[i]); j++ {
			r, g, b, a := image[i][j].RGBA();
			px := messages.Pixel{R: r, G: g, B: b, A: a}
			col.Data = append(col.Data, &px);
		}
		ser.Data = append(ser.Data, col)
	}
	return proto.Marshal(ser)
}
