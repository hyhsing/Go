package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.POST("/upload", func(c *gin.Context) {

		file, _, _ := c.Request.FormFile("file")
		cimg, _ := jpeg.Decode(file)
 

		// 圖片2
		desc := image.NewRGBA(image.Rect(0, 0, cimg.Bounds().Dx(), 210.0))
		gc := draw2dimg.NewGraphicContext(desc)

		// 字型的指定
		//draw2d.SetFontFolder("fonts/") // 字型放置目錄
		draw2d.SetFontNamer(func(fontData draw2d.FontData) string {
			return "msjh.ttf"
		})

		gc.SetFontSize(14)                                                  // 文字大小
		gc.FillStringAt("一個用Go語言寫的圖片上傳並標註文字的範例程式碼", 8, 52) // 內容
		gc.Close()
		gc.FillStroke()

		//starting position of the second image (bottom left)
		sp2 := image.Point{0, cimg.Bounds().Dy()}
		r2 := image.Rectangle{sp2, sp2.Add(desc.Bounds().Size())}

		r := image.Rectangle{image.Point{0, 0}, r2.Max}
		rgba := image.NewRGBA(r)

		// 將圖片繪製至rgba上
		draw.Draw(rgba, cimg.Bounds(), cimg, image.Point{0, 0}, draw.Src)
		draw.Draw(rgba, r2, desc, image.Point{0, 0}, draw.Src)
		//----
		var buff bytes.Buffer
		var opt jpeg.Options
		opt.Quality = 80
		jpeg.Encode(&buff, rgba, &opt)
		// encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
		// htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"

		c.Data(http.StatusOK, "image/jpeg", buff.Bytes())
	})

	r.Run(":3000")
}
