package main

//! A  string  is  an  immutable  sequence  of  bytes.
// Strings  may  contain  arbitrary  data, including  bytes  with  value  0,
// but  usually  they  contain  human-readable  text.
//! Text strings are conventionally interpreted as UTF-8-encoded sequences of Unicode code points (runes)

//? Про кодировки
//? https://habr.com/ru/articles/478636/

// Кодировки используются, в зависимости от символов, которые используются в предоставленном тексте.
// Сначала появилась аски кодировка, она табличку с символами, вместимостью 256 символов, которые кодируем 8битами(1байт). Первые 7 бит отводились под латинские символы, плюс цифры и всякие знаки, а далее, остаток по комбинациям отводилась под символы страны. Так у каждой страны была частично своя кодировка своих символов.
// Позднее решили что надо создать общую кодировку на всех, так появился юникод. емкость символов юникода составляет от 0 до 10FFFF, но из-за всеобъемлющего количества символов, он платит за это самым большим объемом по памяти.
// Для решения этой проблеммы сделали UTF кодировки
//
// UTF-8 является юникод-кодировкой переменной длинны, с помощью которой можно представить любой символ юникода.
// Первым делом надо сказать, что структурной (атомарной) единицей этой кодировки является байт. То что кодировка переменной длинны, значит, что один символ может быть закодирован разным количеством структурных единиц кодировки, то есть разным количеством байтов. Так например латиница кодируется одним байтом, а кириллица двумя байтами.
// Давайте возьмем символ «o»(англ.) из примера про ASCII выше. Помним что в таблице ASCII символов он находится на 111 позиции, в битовом виде это будет 01101111. В таблице юникода этот символ — U+006F что в битовом виде тоже будет 01101111. И теперь так, как UTF — это кодировка переменной длины, то в ней этот символ будет закодирован одним байтом. То есть представление данного символа в обеих кодировках будет одинаково. И так для всего диапазона символов от 0 до 128
//
// То есть если ваш документ состоит из английского текста то вы не заметите разницы если откроете его и в кодировке UTF-8 и UTF-16 и ASCII

//? UTF 8 кодировка Работает она следующим образом.
//? Первый бит каждого байта кодирующего символ отвечает не за сам символ, а за определение байта.
//?
//? То есть например если ведущий (первый) бит нулевой, то это значит что для кодирования символа используется всего один байт. Что и обеспечивает совместимость с ASCII. Если внимательно посмотрите на таблицу символов ASCII то увидите что первые 128 символов (английский алфавит, управляющие символы и знаки препинания) если их привести к двоичному виду, все начинаются с нулевого бита (будьте внимательны, если будете переводить символы в двоичную систему с помощью например онлайн конвертера, то первый нулевой ведущий бит может быть отброшен, что может сбить с толку).
//? 01001000 — первый бит ноль, значит 1 байт кодирует 1 символ -> «H»
//? 01100101 — первый бит ноль, значит 1 байт кодирует 1 символ -> «e»
//?
//? Для двухбайтовых символов первые три бита должны быть такие — 110
//? 11010000 10111100 — в начале 110, значит 2 байта кодируют 1 символ. Второй байт в таком случае всегда начинается с 10. Итого отбрасываем управляющие биты (начальные, которые выделены красным и зеленым) и берем все оставшиеся (10000111100), переводим их в шестнадцатиричный вид (043С) -> U+043C в юникоде равно символ «м».
//?
//? для трех-байтовых символов в первом байте ведущие биты — 1110
//? 11101000 10000111 101010101  — суммируем все кроме управляющих битов и получаем что в 16-ричной равно 103В5, U+103D5 — древнеперситдская цифра сто (10000001111010101)
//?
//? для четырех-байтовых символов в первом байте ведущие биты — 11110
//? 11110100 10001111 10111111 10111111 — U+10FFFF это последний допустимый символ в таблице юникода (100001111111111111111)

// UTF-16 также является кодировкой переменной длинны. Главное ее отличие от UTF-8 состоит в том
// что структурной единицей в ней является не один а два байта. То есть в кодировке UTF-16 любой символ юникода
// может быть закодирован либо двумя, либо четырьмя байтами, то есть либо одной кодовой парой, либо двумя.
func main() {

	// s := "hello, world"
	// fmt.Println(len(s))     // "12"
	// fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')

	//! The i-th byte of a string is not necessarily the i-th character of a string, because the
	//! UTF-8  encoding  of  a  non-ASCII  code  point  requires  two  or  more  bytes.  Working
	//! with characters is discussed shortly

	// 	Strings may be compared with comparison operators like == and <; the comparison is
	// done byte by byte, so the result is the natural lexicographic ordering

	//! s[0] = 'L' // compile error: cannot assign to s[0]
	//* Immutability  means  that  it  is  safe  for  two  copies  of  a  string  to  share  the  same
	//* underlying memory, making it cheap to copy strings of any length. Similarly, a string
	//* s  and  a  substring  like  s[7:]  may  safely  share  the  same  data,  so  the  substring
	//* operation  is  also  cheap.  No  new  memory  is  allocated  in  either  case.

	// Go  source  files  are  always  encoded  in  UTF-8  and  Go  text  strings  are
	// conventionally interpreted as UTF-8, we can include Unicode code points in string
	// literals.

	//* One set of escapes handles
	//* ASCII control codes like newline, carriage return, and tab:
	//* \a  “alert” or bell
	// fmt.Println(rune('\a'))
	//* \b backspace
	// fmt.Println(rune('\b'))
	//* \f form feed
	// fmt.Println(rune('\f'))
	//* \n newline
	// fmt.Println(rune('\n'))
	//* \r carriage return
	// fmt.Println(rune('\r'))
	//* \t tab
	// fmt.Println(rune('\t'))
	//* \v vertical tab
	// fmt.Println(rune('\v'))
	//* \' single quote (only in the rune literal '\'')
	// fmt.Println(rune('\''))
	//* \" double quote (only within "..." literals)
	// fmt.Println(string("\""))
	//* \\ backslash
	// fmt.Println(rune('\\'))

	//& There are two forms,
	//& \uhhhh for a 16-bit value and
	//& \Uhhhhhhhh for a 32-bit value,
	//& where each h is a hexadecimal digit
	// fmt.Println(
	// 	"\xe4\xb8\x96\xe7\x95\x8c",
	// 	"\u4e16\u754c",
	// 	"\U00004e16\U0000754c",
	// 	string('世'),
	// 	string('\u4e16'),
	// 	string('\U00004e16'),
	// )

	//? A rune whose value is less than 256 may be written with a single hexadecimal escape,
	//? such as '\x41' for 'A', but for higher values, a \u or \U escape must be used.
	//? Consequently, '\xe4\xb8\x96' is not a legal rune literal, even though those three
	//? bytes are a valid UTF-8 encoding of a single code point.

	//* The  string  contains  13  bytes,
	//* but  interpreted  as  UTF-8,  it  encodes
	//* only  nine  code points or runes:
	// s := "Hello, 世界"
	// fmt.Println(len(s))                    // "13"
	// fmt.Println(utf8.RuneCountInString(s)) // "9"
	// for i := 0; i < len(s); {
	// 	runE, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\t%d bytes\n", i, runE, size)
	// 	i += size
	// }

	//! Go’s range  loop,  when  applied  to  a  string, performs UTF-8 decoding implicitly.
	// for i, r := range "Hello, 世界" {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	// }

	//? What happens if we range over a string containing arbitrary binary data or, for that matter, UTF-8 data containing errors?
	//! Each  time  a  UTF-8  decoder,  whether  explicit  in  a  call  to
	//! utf8.DecodeRuneInString  or  implicit  in  a  range  loop,  consumes  an
	//! unexpected  input  byte,  it  generates  a  special  Unicode  replacement  character,
	//! '\uFFFD' - �,  which  is  usually  printed  as  a  white  question  mark  inside  a  black
	//! hexagonal or diamond-like shape  . When a program encounters this rune value, it’s
	//! often a sign that some upstream part of the system that generated the string data has
	//! been careless in its treatment of text encodings.
	//! fmt.Printf("%c", '\uFFFD')

	//? A []rune  conversion  applied  to  a  UTF-8-encoded  string  returns  the  sequence  of Unicode code points that the string encodes:
	// s := "\xe3\x83\x97\xe3\x83\xad\xe3\x82\xb0\xe3\x83\xa9\xe3\x83\xa0"
	// fmt.Printf("%s", s)
	// s := "プログラム"
	// fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	// r := []rune(s)
	// fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	// for _, ru := range []byte(s) {
	// 	fmt.Printf("% [1]b", ru)
	// }
	// fmt.Println()
	// for _, ru := range r {
	// 	fmt.Printf("% [1]c - %[1]b(2)- (%[1]d)10 - (%[1]x)10\n", ru)
	// }
	// fmt.Println(string(0x30d7))

	Ch3_5_5()
}
