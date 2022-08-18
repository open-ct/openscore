package util

import (
	"bufio"
	"flag"
	"fmt"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi      = flag.Float64("dpi", 200, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./font/simhei.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 20, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	width    = 1024
)

/**
生成图片
*/

func UploadPic(name string, text string) (src string) {

	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	rgba := image.NewRGBA(image.Rect(0, 0, width, width))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	// Draw the guidelines.
	for i := 0; i < width; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(width-10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
		rgba.Set(10+i, width-10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(20, 20+int(c.PointToFixed(*size)>>6))

	opts := truetype.Options{}
	opts.Size = *size
	opts.DPI = *dpi
	face := truetype.NewFace(f, &opts)

	for _, x := range []rune(text) {
		w, _ := face.GlyphAdvance(x)
		if pt.X.Round()+w.Round() > width-10 {

			pt.X = fixed.Int26_6(5) << 6
			pt.Y += c.PointToFixed(*size * *spacing)
		}
		pt, err = c.DrawString(string(x), pt)
	}

	// Save that RGBA image to disk.
	name += ".png"
	newPath := "./tmp/" + name

	outFile, err := os.Create(newPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fileBytes, err := ioutil.ReadFile(newPath)
	if err != nil {
		panic(err)
	}

	fileUrl := UploadFileToStorage(name, fileBytes)

	os.Remove(newPath)
	fmt.Println("Wrote out.png OK.")
	return fileUrl
}

func UploadFileToStorage(name string, fileBytes []byte) string {
	fullFilePath := fmt.Sprintf("openscore/img/%s", name)
	fileUrl, _, err := auth.UploadResource("admin", "", "", fullFilePath, fileBytes)
	if err != nil {
		panic(err)
	}

	return fileUrl
}
