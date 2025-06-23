package main

//? When using the go tool, a package’s import path indicates not only where to find it in
//? the local workspace, but where to find it on the Internet so that go get can retrieve and update it

// The go get command can download a single package or an entire subtree or
// repository using the ... notation, as in the previous section.

// ! В Go 1.16+
//* go get 	 — только добавляет зависимость в go.mod, но не ставит бинарник.
//* go install — ставит бинарник + сохраняет код в кеш.

//* go get -u ... - generally retrieves the latest version of each package
//? May be inappropriate for deployed projects, where precise control of dependencies is critical for release hygiene.
//
// The usual solution to this problem is to vendor the code, that is, to make a persistent local copy
// of all the necessary dependencies, and to update this copy carefully and deliberately

//*``` go run 1.Tutorial/1.5_FetchingURL/main.go http://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import
//*``` <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
