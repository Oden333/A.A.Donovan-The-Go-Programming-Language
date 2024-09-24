// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

// Server2 is a minimal "echo" and counter server.
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/liss", handler3)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var cycles = 5 // количество полных оборотов осциллятора по оси X.

func handler3(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	cycles, _ = strconv.Atoi(r.Form.Get("zalupa"))
	lissajous(w)
}

func server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the HTTP request.
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL,
		r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	// used fmt.Fprintf to write to an http.ResponseWriter representing the web browser
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// handler1 echoes the Path component of the requested URL.
func handler1(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// A handler pattern  that  ends  with  a  slash  matches  any  URL  that  has  the  pattern  as  a  prefix.
// Behind the scenes, the server runs the handler for each incoming request in a separate
// goroutine  so  that  it  can  serve  multiple  requests  simultaneously.  However,  if  two
// concurrent  requests  try  to  update  count  at  the  same  time,  it  might  not  be
// incremented  consistently;  the  program  would  have  a  serious  bug  called  a  race
// condition (§9.1). To avoid this problem, we must ensure that at most one goroutine
// accesses  the  variable  at  a  time,  which  is  the  purpose  of  the  mu.Lock()  and
// mu.Unlock() calls that bracket each access of count. We’ll look more closely at
// concurrency with shared variables in Chapter 9.

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func server0() {
	http.HandleFunc("/", handler0) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the requested URL.
func handler0(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{21, 160, 87, 1}, //Exercise 1.5
	color.RGBA{154, 154, 154, 0},
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2
	grayIndex  = 3
)

func lissajous(out io.Writer) error {
	const (
		res       = 0.001 // разрешение угла (чем меньше значение, тем больше точек на кривой).
		size      = 200   // размер изображения (в пикселях), площадь изображения — от -size до +size.
		nframes   = 100   // количество кадров в анимации.
		delay     = 4     // задержка между кадрами в единицах времени (10 мс).
		pointSize = 2     // Размер "точки" (радиус в пикселях)
	)
	freq := 9.00                        // относительная частота осциллятора по оси Y.
	anim := gif.GIF{LoopCount: nframes} // структура для создания GIF, состоящего из nframes кадров.
	phase := 0.0                        // разница фаз между осцилляторами по осям X и Y.

	//Генерация кадров
	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // создаёт прямоугольную область для кадра.
		img := image.NewPaletted(rect, palette)      // создаёт палеточное изображение с заданной палитрой.

		// Устанавливаем цвет фона (например, grayIndex)
		for y := 0; y < 2*size+1; y++ {
			for x := 0; x < 2*size+1; x++ {
				img.SetColorIndex(x, y, whiteIndex) // Закрашиваем фон серым цветом
			}
		}

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res { // проходит по углу от 0 до 2π за один цикл и рисует точку кривой Лиссажу.
			// координаты точки на кривой.
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// //  рисует точку в изображении с заданным цветом
			// img.SetColorIndex(size+int(x*size+0.5),
			// 	size+int(y*size+0.5),
			// 	greenIndex)
			// Центральная координата точки
			px := size + int(x*size+0.5)
			py := size + int(y*size+0.5)

			// // Увеличиваем размер точки закрашиванием области вокруг (квадрат)
			// for dx := -pointSize; dx <= pointSize; dx++ {
			// 	for dy := -pointSize; dy <= pointSize; dy++ {
			// 		// Проверяем, что координаты не выходят за границы изображения
			// 		if px+dx >= 0 && px+dx < 2*size+1 && py+dy >= 0 && py+dy < 2*size+1 {
			// 			img.SetColorIndex(px+dx, py+dy, greenIndex)
			// 		}
			// 	}
			// }

			// Рисуем увеличенные точки (круг)
			for dx := -pointSize; dx <= pointSize; dx++ {
				for dy := -pointSize; dy <= pointSize; dy++ {
					if px+dx >= 0 && px+dx < 2*size+1 && py+dy >= 0 && py+dy < 2*size+1 {
						if dx*dx+dy*dy <= pointSize*pointSize { // условие для круглых точек
							img.SetColorIndex(px+dx, py+dy, greenIndex)
						}
					}
				}
			}
		}

		phase += 0.1 // изменяет фазу для каждого кадра, чтобы кривые менялись
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return fmt.Errorf("failed to encode GIF: %v", err)
	}
	return nil
}
