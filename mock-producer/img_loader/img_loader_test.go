package imgloader_test

import (
	imgloader "mock-producer/img_loader"
	"testing"
)

func TestLoadImage(t *testing.T) {
	const IMAGE_PATH = "../data/cool-dog.jpg"
	const DOG_WIDTH = 4048;
	const DOG_HEIGHT = 3036;
	loaded, err := imgloader.LoadImage(IMAGE_PATH);
        if err != nil {
		t.Errorf("LoadImage returned an error: %s", err);
	}
	if len(loaded) != DOG_WIDTH || len(loaded[0]) != DOG_HEIGHT {
		t.Error("LoadImage returned zero-sized image");
	}
}
