package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/kosa3/pexels-go"
)

func getWord(path string) ([]string, error) {
	verbFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer verbFile.Close()

	var lines []string
	scanner := bufio.NewScanner(verbFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func getCaption() string {
	verbLines, err := getWord("sexVerbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	partLines, err := getWord("bodyParts.txt")
	if err != nil {
		log.Fatal(err)
	}
	objectLines, err := getWord("householdObjects.txt")
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	min := 1
	verbMax := len(verbLines) - 1
	partMax := len(partLines) - 1
	objectMax := len(objectLines) - 1

	verb := verbLines[rand.Intn(verbMax-min)+min]
	part := partLines[rand.Intn(partMax-min)+min]
	object := objectLines[rand.Intn(objectMax-min)+min]

	caption := verb + " his " + part + " with a " + object
	return caption

}

func getPhotoURL() string {
	cli := pexels.NewClient("123-abc")
	ctx := context.Background()
	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query:   "men",
		Page:    1,
		PerPage: 80,
	})
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 80
	sourceURL := ps.Photos[rand.Intn(max-min)+min].Src.Medium
	return sourceURL
}

func getPhoto() int64 {
	resp, err := http.Get(getPhotoURL())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("dude.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	photo, _ := io.Copy(out, resp.Body)
	return photo
}

func captionPhoto() {
	im, err := gg.LoadImage("dude.jpg")
	size := im.Bounds().Size()
	if err != nil {
		log.Fatal(err)
	}
	dc := gg.NewContext(size.X, size.Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("Inspiration/Inspiration-Regular.ttf", 32); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)
	dc.DrawString(getCaption(), 50, 20)
	dc.SavePNG("out.png")
}

func main() {
	getPhoto()
	captionPhoto()

}
