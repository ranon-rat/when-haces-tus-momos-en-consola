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
	pos := " .:!i1|0"
	for _, x := range pos {
		for y := 0; y <= 257/len(pos); y++ {
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
func main() {
	addthis()
	when, err := os.Open("when.jpeg")
	if err != nil {

		fmt.Println("the file didnt exist, retard")
		panic(err)
	}
	defer when.Close()
	//// cuando you make your memes in videos
	frase("cuando haces tus momos en videos")

	fmt.Println()

	/// file open when

	openThis(when)
	futuro, err := os.Open("futuro.jpg")
	if err != nil {
		fmt.Println("shit this didnt exist")
		panic(err)
	}
	defer futuro.Close()
	/////// the future is today you listen,old man?
	frase("el futuro es hoy ¿Oiste viejo?")

	openThis(futuro)
	baneando, err := os.Open("baneando.jpg")
	if err != nil {
		fmt.Println("fuck you")
	}
	defer baneando.Close()
	frase("pero te terminan baneando")
	openThis(baneando)

	jaja, err := os.Open("jajaja.jpeg")
	if err != nil {
		fmt.Println("fuck you")
		panic(err)
	}
	defer jaja.Close()
	frase("oh! mi lente de contacto jajajajajajj")
	openThis(jaja)

}
func openThis(f io.Reader) {
	img, _, _ := image.Decode(f)

	division := 5

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y/division; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X/division; x++ {
			_, _, a, _ := img.At(x*division, y*division).RGBA()
			fmt.Print(letters[(a/257)%uint32(len(letters))])
		}
		fmt.Println()

	}
}
