package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"sync"
)

func getImageDimensions(filename string) (width, height int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

func worker(files <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range files {
		width, height, err := getImageDimensions(file)
		if err != nil {
			continue
		}
		fmt.Printf(`<img src="%s" width="%d" height="%d" />\n`, filepath.Base(file), width, height)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: программа <файл1> <файл2> ...")
		return
	}

	files := make(chan string, len(os.Args)-1)
	var wg sync.WaitGroup

	// Запуск рабочих горутин
	for i := 0; i < 4; i++ { // 4 рабочих горутины
		wg.Add(1)
		go worker(files, &wg)
	}

	// Отправка файлов в канал
	for _, filename := range os.Args[1:] {
		files <- filename
	}
	close(files)

	wg.Wait()
}