package main

// go tool is used for downloading querying, formatting, building, testing, and installing packages of Go code.

//? The only configuration most users ever need is the GOPATH environment variable,
//? which specifies the root of the workspace
// When switching to a different workspace, users update the value of GOPATH

// * GOPATH has three subdirectories. *//
//
//? The src subdirectory holds source code. Используется для
// - Совместимости со старыми проектами, которые ещё не перешли на модули (GO111MODULE=off).
// - Ручного управления зависимостями (если кому-то нужно).
// - Хранения личных/локальных пакетов, которые не хочется публиковать в виде модулей.

//! Сейчас (с Go Modules, Go 1.11+):
//? $GOPATH/src больше не обязателен для новых проектов.
//? Вместо этого используется модульный кеш ($GOPATH/pkg/mod), а проекты можно размещать в любом месте на диске.

// ! Go ищет этот пакет в таком порядке:
// 	*-	В вашем модуле (если ./vendor/github.com/example/mypkg существует).
// 	*-	В кеше модулей ($GOPATH/pkg/mod/github.com/example/mypkg@vX.Y.Z).
// 	*-	В $GOPATH/src (только если модули отключены). !GO111MODULE=on → всегда использует Go Modules, $GOPATH/src игнорируется.
// 	*-	В GOROOT (если это пакет из stdlib, например fmt).

//? The pkg subdirectory is where the build tools store compiled packages,
// and the bin subdirectory holds executable programs like helloworld

// * GOROOT, specifies the root directory of the Go distribution,
// * which provides all the packages of the standard library
//
// Users never need to set GOROOT since, by default,
// the go tool will use the location where it was installed.
