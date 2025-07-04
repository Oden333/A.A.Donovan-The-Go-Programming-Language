package main

import (
	"fmt"
	"sync/atomic"
)

// 9.8.1 Growable Stacks

//! Each OS thread has a fixed-size block of memory (often as large as 2MB) for its stack
// the work area where it saves the local variables of function calls that are in
// progress or temporarily suspended while another function is called.

// Changing the fixed size can improve space efficiency and allow more threads to be created,
// or it can enable more deeply recursive functions, but it cannot do both.
//! In contrast, a goroutine starts life with a small stack, typically 2KB.

//? A goroutine’s stack, like the stack of an OS thread, holds the local variables of active and
//? suspended function calls, but unlike an OS thread, a goroutine’s stack is not fixed;
//? it grows and shrinks as needed.
//! The size limit for a goroutine stack may be as much as 1GB,

//! Планирование горутин происходит на уровне user space, то есть ими управляет планировщик Go (если точнее — Go Runtime),
//! а планирование тредов на уровне kernel space, то есть ими управляет ОС.

//? Основной ресурс для выполнения программ — ядра процессора
//
//? Ядер мало, а работы у них много, поэтому вводится концепция тредов: треды выполняются на ядрах.
//
//? Планировщик ОС управляет тредами и оптимизирует работу с ними таким образом, чтобы ядра не простаивали без работы.
//
//? Треды ОС большие и страшные — не потому что их плохо спроектировали, а потому что на уровне ОС есть свои ограничения и особенности работы, поэтому тредов нам доступно мало.
//
//? Тредов мало, а работы внутри нашей программы очень много, поэтому вводится концепция горутин: горутины выполняются на тредах
//
//? Планировщик Go управляет горутинами и оптимизирует работу с ними таким образом, чтобы создать максимально эффективно использовать треды — создавать их как можно меньше и не позволять им простаивать

func counter(out chan<- int) {
	for x := 0; x < 10000; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(out chan struct{}, in <-chan int) {
	for v := range in {
		fmt.Printf("%d ", v)
		if v%40 == 0 {
			fmt.Printf("\n")
		}
		out <- struct{}{}
	}
	fmt.Println()
	close(out)
}

func simple_pipeline() {
	naturals := make(chan int)
	squares := make(chan int)
	donner := make(chan struct{})
	go counter(naturals)
	go squarer(squares, naturals)
	go printer(donner, squares)
	for range donner {
	}
}

func dont_run_me() {
	var c atomic.Int64
	for {
		channel := make(chan struct{})
		c.Add(1)
		go func(ch chan struct{}) {
			<-ch
		}(channel)
		fmt.Printf("\r%d", c.Load())
	}
}

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}
