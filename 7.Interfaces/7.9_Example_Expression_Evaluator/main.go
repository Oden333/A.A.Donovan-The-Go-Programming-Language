//* Expression language consists of
//* floating-point literals;
//* the binary operators +, -, *, and /;
//* the unary operators -x and +x;
//* function calls pow(x,y), sin(x), and sqrt(x);
//* variables  such  as  x  and  pi;  and  of  course
//* parentheses  and  standard operator  precedence.

// * All  values  are  of  type  float64
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

func server() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", plot)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func ex7_15() {
	// var input = new(bytes.Buffer)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your expression:")
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	expr, err := parseAndCheck(string(line))
	if err != nil {
		log.Fatalf("bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Enter variable values in formate ' x = 12 ' ending with line skip")

	var (
		k   string
		v   float64
		vs  []string
		env = make(map[Var]float64)
	)
	for {
		line, _, err = reader.ReadLine()
		if len(line) == 0 {
			break
		}
		// log.Println(string(line))
		if err != nil {
			log.Output(1, err.Error())
			break
		}

		vs = strings.Split(strings.TrimSpace(string(line)), "=")
		if len(vs) == 2 {
			k = vs[0]
			v, err = strconv.ParseFloat(strings.TrimSpace(vs[1]), 64)
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("invalid format, type in your vars like x=13")
		}

		env[Var(k)] = v
		// log.Println(k, v)
	}
	fmt.Fprintf(os.Stdout, "Your expr: %s with Vars %v\n The calculations results with: \b%v\n", expr.String(), env, expr.Eval(env))
}

// ex7.16
func calcServ() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", home)
	// mux.HandleFunc("/plot", plot)
	mux.HandleFunc("/calc", calc)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

var homePage []byte

func parseVars(s string) (map[Var]float64, error) {
	vars := make(map[Var]float64)
	if s == "" {
		return nil, errors.New("empty variables string")
	}

	for arg := range strings.SplitSeq(strings.TrimSpace(s), ",") {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}

		kv := strings.SplitN(arg, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid variable format: %q (expected key=value)", arg)
		}

		key := strings.TrimSpace(kv[0])
		if key == "" {
			return nil, fmt.Errorf("empty variable name in %q", arg)
		}

		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %q: %v", key, err)
		}

		vars[Var(key)] = val
	}

	if len(vars) == 0 {
		return nil, errors.New("no valid variables provided")
	}
	return vars, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	if len(homePage) == 0 {
		http.Error(w, "home page not loaded", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	bufio.NewReader(bytes.NewReader(homePage)).WriteTo(w)
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "invalid expression: "+err.Error(), http.StatusBadRequest)
		return
	}

	vars, err := parseVars(r.Form.Get("args"))
	if err != nil {
		http.Error(w, "invalid varibles array: "+err.Error(), http.StatusBadRequest)
		return
	}

	result := expr.Eval(vars)

	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "result encoding error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		return
	}
}

func main() {
	//homePage, err = os.ReadFile("./templates/home.html") - one-step read into []byte
	if err := loadHomePage(); err != nil {
		log.Fatal(err)
	}
	calcServ()
}

func loadHomePage() error {
	file, err := os.Open("./templates/home.html")
	if err != nil {
		return err
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	writer.Flush()
	homePage = buf.Bytes()
	return nil
}

func load_invalid() {
	h, err := os.Open("./templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	homePage = make([]byte, 0)

	reader := bufio.NewReader(h)
	writer := bufio.NewWriter(bytes.NewBuffer(homePage))
	// bytes.NewBuffer(homePage) создаёт новый буфер, но:
	// Не расширяет homePage, а просто использует его как начальное значение.
	// После записи через writer данные попадают в новый буфер, но homePage остаётся пустым (len=0).

	// bufio.Writer пишет во внутренний буфер, но:
	// Вы не извлекаете данные из него обратно в homePage.
	// Flush() сбрасывает данные в bytes.Buffer, но не в homePage.

	_, err = reader.WriteTo(writer)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}
