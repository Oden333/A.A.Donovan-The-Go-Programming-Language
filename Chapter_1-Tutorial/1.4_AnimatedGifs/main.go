package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
)

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

func main() {
	// Открываем файл для записи, os.Create всегда создаст новый файл или очистит существующий
	f, err := os.Create("out.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	// вызываешь os.Create("out.gif"),
	// это автоматически перезаписывает файл, если он существует, и очищает его содержимое.
	// Поэтому проверять длину файла и очищать его вручную не нужно.
	// Функция os.Create уже делает это

	defer f.Close()
	lissajous(f)

	// НЕ ПОЛУЧАЕТСЯ СОЗДАНИЕ ГИФ ЧЕРЕЗ аргумент cmd,
	// поэтому сделал через явное создание файла внутри проги
	//
	// Обернём os.Stdout в буферизованный writer
	// writer := bufio.NewWriter(os.Stdout)

	// err := lissajous(writer)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error during GIF generation: %v\n", err)
	// 	os.Exit(1)
	// }

	// // Принудительно сбрасываем буфер после записи
	// err = writer.Flush()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error flushing buffer: %v\n", err)
	// 	os.Exit(1)
	// }
}

func lissajous(out io.Writer) error {
	const (
		cycles    = 5     // количество полных оборотов осциллятора по оси X.
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

		for t := 0.0; t < cycles*2*math.Pi; t += res { // проходит по углу от 0 до 2π за один цикл и рисует точку кривой Лиссажу.
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
