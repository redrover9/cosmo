package main

import (
	"bufio"
	"context"
	"image/color"
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
	cli := pexels.NewClient("abc-123")
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

func getFont() string {
	fonts := []string{"Herr_Von_Muellerhoff/HerrVonMuellerhoff-Regular.ttf", "Homemade_Apple/HomemadeApple-Regular.ttf", "Inspiration/Inspiration-Regular.ttf", "Pacifico/Pacifico-regular.ttf"}
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	font := fonts[rand.Intn(max-min)+min]
	return font
}

func getFontDimensions(property string) float64 {
	rand.Seed(time.Now().UnixNano())
	if property == "x" {
		min := 50
		max := 400
		dimension := rand.Intn(max-min)+min
		return float64(dimension)
	} else if property == "y" {
		min := 125
		max := 175
		dimension := rand.Intn(max-min)+min
		return float64(dimension)
	} else if property == "width" {
		min := 100
		max := 150
		dimension := rand.Intn(max-min)+min
		return float64(dimension)
	} else if property == "spacing" {
		min := 4
		max := 6
		dimension := rand.Intn(max-min)+min
		return float64(dimension)
	}
	return 0
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
	dc.SetColor(color.CMYK{0, 59, 29, 0})
	if err := dc.LoadFontFace(getFont(), 20); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(getCaption(), getFontDimensions("x"), getFontDimensions("y"), 0.5, 0.5, getFontDimensions("width"), getFontDimensions("spacing"), gg.AlignCenter)
	dc.SavePNG("out.png")
}

func main() {
	getPhoto()
	captionPhoto()

}
