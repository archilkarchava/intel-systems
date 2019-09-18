package main

import (
	"image"
	"log"
	"os"

	_ "golang.org/x/image/bmp"
)

/* func convertBmpToBmp(r io.Reader) (Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return img, err
} */

// Figure represents geometrical figure
type Figure struct {
	image.Image
}

func (f Figure) width() int {
	return f.Bounds().Max.X
}

func (f Figure) height() int {
	return f.Bounds().Max.Y
}

func (f Figure) isRectangle() bool {
	for i := 0; i < f.width(); i++ {
		r, g, b, a := f.At(i, f.height()-1).RGBA()
		// log.Printf("i: %d // r: %d, g: %d, b: %d, a: %d\n", i, r, g, b, a)
		if !((r == 0) && (g == 0) && (b == 0) && (a == 65535)) {
			return false
		}
	}
	for i := 0; i < f.height(); i++ {
		r, g, b, a := f.At(f.width()-1, i).RGBA()
		// log.Printf("i: %d // r: %d, g: %d, b: %d, a: %d\n", i, r, g, b, a)
		if !((r == 0) && (g == 0) && (b == 0) && (a == 65535)) {
			return false
		}
	}
	for i := f.width() - 1; i >= 0; i-- {
		r, g, b, a := f.At(i, 0).RGBA()
		// log.Printf("i: %d // r: %d, g: %d, b: %d, a: %d\n", i, r, g, b, a)
		if !((r == 0) && (g == 0) && (b == 0) && (a == 65535)) {
			return false
		}
	}
	for i := f.height() - 1; i >= 0; i-- {
		r, g, b, a := f.At(0, i).RGBA()
		// log.Printf("i: %d // r: %d, g: %d, b: %d, a: %d\n", i, r, g, b, a)
		if !((r == 0) && (g == 0) && (b == 0) && (a == 65535)) {
			return false
		}
	}
	return true
}

func (f Figure) isSquare() bool {
	return (f.width() == f.height()) && (f.isRectangle())
}

func (f Figure) isTriangle() bool {
	return (f.width() == f.height()) && (f.isRectangle())
}

// func checkAngle

func main() {
	reader, err := os.Open("data/triangle.bmp")
	if err != nil {
		log.Fatalf("Ошибка открытия файла. %s", err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Ошибка считывания изображения. %s", err)
	}
	figure := Figure{img}
	log.Println("Is rectangle:", figure.isRectangle())
	log.Println("Is square:", figure.isSquare())
	log.Println("Is triangle:", figure.isTriangle())

}
