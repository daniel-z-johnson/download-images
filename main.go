package main

import (
	"flag"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"path"
	"regexp"

	"image"
	_ "image/gif"
	_ "image/jpeg"
)

func main() {
	fmt.Println("Program Start")
	folder := flag.String("f", "", "folder name")
	filename := flag.String("name", "page_%04d", "")
	downloadURL := flag.String("loc", "", "URL template for the downloads")
	start := flag.Int("start", 0, "start count")
	end := flag.Int("end", 0, "end count")
	flag.Parse()

	fmt.Println("Config being used")
	fmt.Printf("folder: '%s'\n", *folder)
	fmt.Printf("filename: '%s'\n", *filename)
	fmt.Printf("downlaodURL: '%s'\n", *downloadURL)
	fmt.Printf("start: '%d'\n", *start)
	fmt.Printf("end: '%d'\n", *end)

	if checkName(*folder) != nil {
		panic("folder name is invalid")
	}
	folderPath := path.Join("downloads", *folder)
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		panic(err)
	}

	for i := *start; i <= *end; i++ {
		err = downloadImage(folderPath, fmt.Sprintf(*filename, i), fmt.Sprintf(*downloadURL, i))
		if err != nil {
			fmt.Printf("Err downloading img: '%s'", err)
		}
	}

	fmt.Println("Finished")
}

func checkName(name string) error {
	pattern, err := regexp.Compile("^[a-z0-9]+$")
	if err != nil {
		return err
	}

	if !pattern.MatchString(name) {
		return fmt.Errorf("foldername is invlid")
	}

	return nil
}

func downloadImage(folderPath, filename, downloadURL string) error {
	filePath := path.Join(folderPath, filename)
	f1, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f1.Close()
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("did not get 200 status code but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}
	png.Encode(f1, img)
	return nil
}
