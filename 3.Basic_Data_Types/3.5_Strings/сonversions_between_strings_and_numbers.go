package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func Ch3_5_5() {
	x := 123
	// Itoa (“integer to ASCII”)
	fmt.Println(
		strconv.Itoa(x),
		fmt.Sprintf("%d", x),
	)
	// Приведение к нужной базе (двоичное представление числа)
	fmt.Println(
		strconv.FormatInt(int64(x), 2), "\n",
		fmt.Sprintf("\b%[1]b, %[1]d, %[1]x", x),
	) //	"1111011"

	// Парсинг инта с строки
	y, _ := strconv.ParseInt(strconv.FormatInt(int64(x), 2), 2, 64)
	//? The third argument of ParseInt gives the size of the integer type that the result must fit into;
	//? for example, 16 implies int16, and the special value of 0 implies int.
	//? In any case, the type of the result y is always int64, which you can then convert to a smaller type.
	fmt.Println(y)

	number := 0
	buf := bytes.Buffer{}
	buf.WriteString("1111011")
	_, err := fmt.Fscanf(&buf, "%b", &number)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Число:", number) // Выводит: Число: 123
	}
}
