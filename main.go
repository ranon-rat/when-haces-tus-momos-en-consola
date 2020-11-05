package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"
	"time"
)

var letters = []string{}

func addthis() {
	pos := " .:!|10@"
	for _, x := range pos {
		for y := 0; y < 257/(len(pos)); y++ {
			letters = append(letters, string(x))
		}
	}
}
func frase(f string) {
	for _, i := range f {
		fmt.Print(string(i))
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}
func openThis(f io.Reader) {
	//this open the image and print the pixels
	img, _, _ := image.Decode(f)
	division := 5
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y/division; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X/division; x++ {
			r, g, b, _ := img.At(x*division, y*division).RGBA()
			color := ((r / 257) + (g / 257) + (b / 257)) / 3
			for int(color) > len(letters)-1 {
				color--
			}
			fmt.Print(letters[color])
		}
		fmt.Println("\033[32m")

	}
}

func main() {
	addthis()
	when, _ := os.Open("when.jpeg")
	defer when.Close()
	//// cuando you make your memes in videos
	frase("cuando haces tus momos en consola")
	fmt.Println()
	/// file open when
	openThis(when)
	futuro, _ := os.Open("futuro.jpg")
	defer futuro.Close()
	/////// the future is today you listen,old man?
	frase("el futuro es hoy ¿Oiste viejo?")
	openThis(futuro)
	baneando, _ := os.Open("baneando.jpg")
	defer baneando.Close()
	// but they band you
	frase("pero te terminan baneando")
	openThis(baneando)
	jaja, _ := os.Open("jajaja.jpeg")
	defer jaja.Close()
	//oh my contact eye ahahhahahahhah
	frase("oh! mi lente de contacto jajajajajajj")
	openThis(jaja)
}
