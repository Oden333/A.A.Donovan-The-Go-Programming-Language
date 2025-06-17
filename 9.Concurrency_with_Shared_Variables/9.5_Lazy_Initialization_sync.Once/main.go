package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

// This version of Icon uses lazy initialization:
func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}
func loadIcon(_ string) image.Image { return nil }

// * In the absence of explicit synchronization, the compiler and CPU are free to reorder accesses
// * to memory in any number of ways, so long as the behavior of each goroutine is sequentially consistent

// The simplest correct way to ensure that all goroutines observe the effects of loadIcons is to synchronize them using a mutex
var mu sync.Mutex // guards icons
var iconss map[string]image.Image

// Concurrency-safe.
func Icons(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()

	//* However, the cost of enforcing mutually exclusive access to icons is that two
	//* goroutines cannot access the variable concurrently, even once the variable has been
	//* safely initialized and will never be modified again

	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

//? There are now two critical sections
//* The goroutine first acquires a reader lock,consults the map, then releases the lock.
//* If no entry was found, the goroutine acquires a writer lock.

var loadIconsOnce sync.Once
var iconsss map[string]image.Image

// Concurrency-safe.
func Iconss(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

//? Each call to Do(loadIcons) locks the mutex and checks the boolean variable.
// In the first call, in which the variable is false, Do calls loadIcons and sets the
// variable to true. Subsequent calls do nothing, but the mutex synchronization ensures
// that the effects of loadIcons on memory (specifically, icons) become visible to
// all goroutines.
