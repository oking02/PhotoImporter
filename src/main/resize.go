package main

import (

	"os"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
)

func compress() {

	hres := "C:\\Users\\ollyking\\Documents\\GoTests\\src\\github.com\\oking02\\PhotoImporter\\testdata\\hres.jpg"
	lres := "C:\\Users\\ollyking\\Documents\\GoTests\\src\\github.com\\oking02\\PhotoImporter\\testdata\\lres.jpg"

	file, err := os.Open(hres)

	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := resize.Resize(400, 400, img, resize.Lanczos3)

	out, err := os.Create(lres)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
