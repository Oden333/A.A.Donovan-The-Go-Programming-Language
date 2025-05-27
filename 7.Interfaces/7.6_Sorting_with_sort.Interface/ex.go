package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
)

func colTitle(x, y *Track) bool  { return x.Title < y.Title }
func colArtist(x, y *Track) bool { return x.Artist < y.Artist }
func colAlbum(x, y *Track) bool  { return x.Album < y.Album }
func colYear(x, y *Track) bool   { return x.Year < y.Year }
func colLength(x, y *Track) bool { return x.Length < y.Length }

type byColumns struct {
	tracks  []*Track
	columns []func(x, y *Track) bool
}

func sortByColumns(t []*Track, f ...func(x, y *Track) bool) *byColumns {
	return &byColumns{
		tracks:  t,
		columns: f,
	}
}

func (x byColumns) Len() int      { return len(x.tracks) }
func (x byColumns) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }
func (x byColumns) Less(i, j int) bool {
	a, b := x.tracks[i], x.tracks[j]
	var k int
	// compare columns one by one except the last
	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k]
		switch {
		case f(a, b):
			return true
		case f(b, a):
			return false
		}
	}
	// all equal, use last column as final judgement
	return x.columns[k](a, b)
}

func getTracks() []*Track {
	t := make([]*Track, len(tracks))
	copy(t, tracks)
	return t
}

func useSortByColumns() []*Track {
	t := getTracks()
	sort.Sort(sortByColumns(t, colTitle, colArtist))
	return t
}

func useSortStable() []*Track {
	t := getTracks()
	sort.Stable(byArtist(t))
	sort.Stable(byTitle(t))
	return t
}

func ex7_8() {
	fmt.Println("By Title, Artist")
	printTracks(useSortByColumns())

	fmt.Println("\nUse sort.Stable. By Title, Artist")
	printTracks(useSortStable())
}

var tpl *template.Template

func maian() {
	var err error
	tpl, err = template.ParseFiles("./tmpl.tpl")
	if err != nil {
		log.Fatal("template parsing error", err)
	}
	// err = tpl.Execute(os.Stdout, tracks)
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var err error
	by := r.URL.Query().Get("by")
	switch by {
	default:
	case "title":
		sort.Stable(byTitle(tracks))
	case "artist":
		sort.Stable(byArtist(tracks))
	case "album":
		sort.Sort(sortByColumns(tracks, colAlbum))
	case "year":
		sort.Stable(byYear(tracks))
	case "length":
		sort.Sort(sortByColumns(tracks, colLength))
	}
	err = tpl.Execute(w, tracks)
	if err != nil {
		log.Fatal("Error executing template", err)
	}
}
