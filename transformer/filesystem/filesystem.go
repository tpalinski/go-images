package filesystem

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"sync"
)

var lock sync.Mutex;

func SaveImage(img image.Image, name string) {
	path := fmt.Sprintf("./data/%s.jpg", name);
	if _, err := os.Stat(path); err == nil {
		return
	}
	lock.Lock()
	outFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// Encode the image as JPEG and write to the output file
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		panic(err)
	}
	outFile.Close()
	lock.Unlock()
}
