// Within a Go program, every package is identified by a unique string called its import
// path.  These  are  the  strings  that  appear  in  an  import  declaration  like
// "gopl.io/ch2/conv".  The  language  specification  doesn’t  define  where
// these strings come from or what they mean;  it’s up to the tools to interpret them.
// When using the go tool (Chapter 10), an import path denotes a directory containing
// one or more Go source files that together make up the package.
// In addition to its import path, each package has a package name, which is the short
// (and  not  necessarily  unique)  name  that  appears  in  its  package  declaration.  By
// convention, a package’s name matches the last segment of its import path, making it
// easy  to  predict  that  the  package  name  of  gopl.io/ch2/conv  is conv

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/conv"
)

// По умолчанию, Go ищет пакеты в двух местах:
// Внутри модуля: когда включены модули Go, путь к пакетам должен быть определён
// в go.mod, а зависимости загружаются автоматически.
// В $GOPATH: если вы не используете модули, Go будет искать пакеты только
// в директории, определённой переменной окружения GOPATH. В вашем случае это C:\Users\Ismail\go.
func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conv.Fahrenheit(t)
		c := conv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, conv.FToC(f), c, conv.CToF(c))
	}
	if len(os.Args) < 2 {
		fmt.Println("Enter conversion type (FToc, CToF, KToC, FToM, PToKg)")
		var convs string
		fmt.Scanf("%d\n", &convs)
		// var usableFunc func(k any) interface{}
		// switch convs {
		// case "FToc":
		// 	usableFunc = conv.FToC
		// case "CTof":
		// 	usableFunc = conv.CToF
		// case "KToC":
		// 	usableFunc = conv.KToC
		// case "FToM":
		// 	usableFunc = conv.FToM
		// case "PToKg":
		// 	usableFunc = conv.PToKg
		// }
		var num int
		fmt.Print("Enter nums count: ")
		fmt.Scanf("%d\n", &num)
		nums := make([]float64, num)
		for i := 0; i < num; i++ {
			// fmt.Scanln(&nums[i])
			// fmt.Scanln() — он считывает строку до первого символа новой строки,
			// оставляя символ новой строки (\n) в буфере ввода.
			// В результате, если вводите число, а затем нажимаете Enter,
			// символ новой строки остаётся в потоке ввода, что влияет на последующие вызовы Scanln().
			var num float64
			fmt.Printf("Enter number %d:", i+1)
			fmt.Scanf("%f\n", &num)
			nums[i] = num
		}

		for _, i := range nums {
			f := conv.FToM(conv.Feet(i))
			fmt.Println(f)
			// fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			// 	f, conv.FToC(f), c, conv.CToF(c), k, conv.KToC(k))
		}
		// fmt.Println(nums)

	}

}
