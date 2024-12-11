package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", gifHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
}

func gifHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm Error: %v\n", err)
	}
	cycle, err := strconv.Atoi(request.FormValue("cycle"))
	if err != nil {
		fmt.Fprintf(writer, "cycle Error: %v\n", err)
		return
	}
	if cycle == 0 {
		cycle = 5
	}
	lissajous(writer, float64(cycle))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL Path:=%q\n", request.URL.Path)
	fmt.Fprintf(writer, "URL Query:=%q\n", request.URL.RawQuery)
	fmt.Fprintf(writer, "RemoteAddr:=%q\n", request.RemoteAddr)
	fmt.Fprintf(writer, "Method:=%q\n", request.Method)
	for key, value := range request.Header {
		fmt.Fprintf(writer, "Header[%q]:=%q\n", key, value)
	}
	fmt.Fprintf(writer, "Body:=%q\n", request.Body)
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm Error: %v\n", err)
	}
	fmt.Fprintf(writer, "Form:=%q\n", request.Form)
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
