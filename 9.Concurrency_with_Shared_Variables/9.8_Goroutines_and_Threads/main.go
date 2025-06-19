package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main0() {
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.GOMAXPROCS(0)) // 0 args == return cpu num

	fmt.Println(runtime.GOMAXPROCS(2)) // first Call return cpu num and rewrites it
	fmt.Println(runtime.GOMAXPROCS(2)) // second call return set before number
}

//* GODEBUG=schedtrace=1000 ./main
// SCHED 0ms: gomaxprocs=12 idleprocs=9 threads=5 spinningthreads=1 needspinning=0 idlethreads=0 runqueue=0 [1 0 0 0 0 0 0 0 0 0 0 0]
// 12
// 2

//? gomaxprocs=12 - Сначала выводится эта инфа, информация по коду не успела обновиться
//? idleprocs=9 - idle (P)
//? threads=5 - оптимизация runtime, 5 тредов занимает го шедулер, используя это количество и расширяясь при необходимости
//? idlethreads=0 - idle threads
//? runqueue=0 - GRQ
//? [1 0 0 0 0 0 0 0 0 0 0 0] - количество горутин в LRQ каждого из P

func main1() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumProc:", runtime.GOMAXPROCS(2))       // second call return set before number
	fmt.Println("NumGoroutine:", runtime.NumGoroutine()) // second call return set before number

	time.Sleep(1 * time.Minute)
}

//* GODEBUG=schedtrace=1000 ./main
// SCHED 0ms: gomaxprocs=12 idleprocs=9 threads=5 spinningthreads=1 needspinning=0 idlethreads=0 runqueue=0 [1 0 0 0 0 0 0 0 0 0 0 0]
// NumCPU: 12
// NumProc: 12
// NumGoroutine: 2
// SCHED 1009ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0]
// SCHED 2011ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0]
// SCHED 3014ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0]
// ...

func main2() { // G0
	// P0 -g-> P1
	runtime.GOMAXPROCS(2)
	go func() {
		for {
			fmt.Println(runtime.NumGoroutine())
			time.Sleep(time.Second) // 258 = main + 256 + Sysmon
		}
	}()

	for range 256 {
		go func() {
			for {
				for i := range 10000000 {
					_ = i * i
				}
			}
		}()
	}

	time.Sleep(1 * time.Minute)
}

//* GODEBUG=schedtrace=1000 ./main
// SCHED 0ms: gomaxprocs=12 idleprocs=11 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
// SCHED 1009ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=53 [162 39]
// SCHED 2016ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=84 [113 57]
// SCHED 3022ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=180 [65 9]
// SCHED 4025ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=0 needspinning=1 idlethreads=2 runqueue=177 [16 61]

func main() { // handoff
	// syscall -> блокируется не только горутина, но и сам тред -> после определённого времени эта горутина передаётся другому P - handoff
	runtime.GOMAXPROCS(5)

	for range 10 {
		go func() {
			// var buf [1]byte
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println(scanner.Scan()) // Вечный syscall. При этом планировщик не может поручить этот вызов NP. G блокируется, T блокируется -> создаётся handoff и место в потоке отдаётся другой горутине
		}()
	}
	time.Sleep(time.Minute)
}

//* GODEBUG=schedtrace=1000 ./main
// SCHED 0ms: gomaxprocs=12 idleprocs=9 threads=5 spinningthreads=1 needspinning=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
// SCHED 1002ms: gomaxprocs=5 idleprocs=5 threads=6 spinningthreads=0 needspinning=0 idlethreads=4 runqueue=0 [0 0 0 0 0]
// SCHED 2003ms: gomaxprocs=5 idleprocs=5 threads=6 spinningthreads=0 needspinning=0 idlethreads=4 runqueue=0 [0 0 0 0 0]

//? threads=6
//? 1 M для sysmon.
//? 1 M для main() (sleep).
//? 3 M заблокированы на scanner.Scan().
//? 1 новый M создан для обработки других задач (но в вашем коде их нет, поэтому он idle).
