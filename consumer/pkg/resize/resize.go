package resize

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func Resize(img image.Image, id int) error {

	errCh := make(chan error)
	x, y, q := calcSize(uint(img.Bounds().Size().X), uint(img.Bounds().Size().Y))
	for i, _ := range x {

		go createImg(img, x[i], y[i], id, q[i], errCh)

		select {
		case err := <-errCh:
			return err
		default:
			continue
		}
	}

	return nil
}

func calcSize(x, y uint) (xU []uint, yU []uint, q []int) {

	xU = append(xU, x)
	yU = append(yU, y)
	q = append(q, 100)

	xU = append(xU, x-(x/4)) // 75%
	yU = append(yU, y-(y/4)) // 75%
	q = append(q, 75)

	xU = append(xU, x/2) // 50%
	yU = append(yU, y/2) // 50%
	q = append(q, 50)

	xU = append(xU, x/4) // 25%
	yU = append(yU, y/4) // 25%
	q = append(q, 25)

	return
}

func createImg(img image.Image, x, y uint, id, quality int, errCh chan error) {

	photo := resize.Resize(x, y, img, resize.Lanczos3)
	out, err := os.Create(fmt.Sprintf("%d_%d.jpg", id, quality))
	if err != nil {
		errCh <- err
		log.Println(err.Error())
	}

	defer out.Close()
	if err := jpeg.Encode(out, photo, nil); err != nil {
		errCh <- err
		log.Println(err.Error())
	}
}
