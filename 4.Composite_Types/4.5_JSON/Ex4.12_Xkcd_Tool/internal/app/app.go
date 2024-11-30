package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"xkcd/types"
)

const getURL = "https://xkcd.com/%d/info.0.json"

var comics = make(map[int]types.Comic)

func Process() {
Start:
	fmt.Println("Which operation do you want to perform:\n1)Get by Comic Number\n2)Get by a search term")
Mode:
	var mode int
	fmt.Scanf("%d\n", &mode)
	switch mode {
	default:
		fmt.Println("Choose beetwen 1 and 2")
		goto Mode
	case 1:
		fmt.Println("Type in the Comic Number")
		var num int
		fmt.Scanf("%d\n", &num)
		comic, ok := comics[num]
		if ok {
			fmt.Println(fmt.Sprintf("Cached comic number %d:", num), "\nComic Transcript:\n", comic.Transcript)
			goto Start
		}
		comic, err := getComicByUrl(num)
		comics[num] = comic
		if err != nil {
			fmt.Fprintln(os.Stdout, "Unable to get comic by Id", err)
		}
		fmt.Println(fmt.Sprintf(getURL, num), "\nComic Transcript:\n", comic.Transcript)
		goto Start
	case 2:
		fmt.Println("Type in string to search in saved comics")
		var str string
		fmt.Scanf("%s\n", &str)
		var found bool
		for num, comic := range comics {
			if strings.Contains(comic.Transcript, str) || strings.Contains(comic.Title, str) {
				found = true
				fmt.Println("Found comic:")
				fmt.Println(fmt.Sprintf(getURL, num), "\nComic Transcript:\n", fmt.Sprintf("%+v", comic.Transcript))
			}
		}
		if !found {
			fmt.Println("Nothing found by your request")
		}
		goto Start
	}
}

func getComicByUrl(n int) (types.Comic, error) {
	var comic types.Comic
	url := fmt.Sprintf(getURL, n)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}
