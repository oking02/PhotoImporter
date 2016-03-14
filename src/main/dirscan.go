package main

import (
	"os"
	"log"
	"path/filepath"
	"image/jpeg"
	"github.com/nfnt/resize"
	"strings"
	"fmt"
)

func dirScan(root, dest string) []string {

	_, err := os.Stat(root)

	if err != nil {
		log.Fatal(err)
	}

	files := []string{}
	dirs := []string{}

	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() {

			name := strings.TrimPrefix(path,  root)
			name = strings.TrimPrefix(name,  "\\")

			p := dest  + name

			dirs = append(dirs, p)
		} else {
			files = append(files, path)
		}

		return nil
	})

	//Built all the directories at the destination first
	for _, dir := range dirs {
		os.Mkdir(dir, 'd')
	}

	if err != nil {
		log.Fatal(err)
	}

	return files

}

func copyAndResize(src, dest string) {

	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := resize.Resize(400, 400, img, resize.Lanczos3)

	out, err := os.Create(dest)

	defer out.Close()

	if err != nil {
		log.Fatal(err)
	}

	// write new image to file
	jpeg.Encode(out, m, nil)

}

func resizeAllPhotos(root, dest string){

	files := dirScan(root, dest)

	fmt.Println("Num Of Files: ", len(files))

	for _, file := range files {

		fmt.Println("-")
		fmt.Println(file)

		name := strings.TrimPrefix(file,  root)
		name = strings.TrimPrefix(name,  "\\")

		path := dest  + name

		copyAndResize(file, path)

	}

}
