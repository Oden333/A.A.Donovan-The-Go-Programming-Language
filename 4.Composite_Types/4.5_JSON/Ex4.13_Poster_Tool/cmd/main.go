package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"poster/types"
	"strconv"
	"strings"
)

// Exercise 4.13: The JSON-based web service of the Open Movie Database lets you
// search https://omdbapi.com/ for a movie by name and download its poster
// image.

// Write a tool poster that downloads the poster image for the movie named on the command line.

const (
	reqURL   = "https://omdbapi.com/"
	basePath = `E:\GoLang\A.A.Donovan-The-Go-Programming-Language\4.Composite_Types\4.5_JSON\Ex4.13_Poster_Tool\`
)

var apiKey *string

func init() {
	apiKey = flag.String("api_key", "", "Your API Key for omdbapi.com site (Check https://www.omdbapi.com/apikey.aspx for obtaining one)")
	flag.Parse()
}

func main() {
	if *apiKey == "" {
		fmt.Println("You need to pass api Key param to run app")
		return
	}

Start:
	fmt.Println("Type in Movie Title or IMDb Id to search for")
	var name string
	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ := strconv.Atoi(name)

	queryParams := make(url.Values, 3)
	queryParams.Add("apikey", *apiKey)
	if id != 0 {
		queryParams.Add("i", name)
	} else {
		queryParams.Add("t", name)
	}

	params := queryParams.Encode()
	resp, err := http.Get(reqURL + "?" + params)
	if err != nil {
		fmt.Println("error while making request", err, resp)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp, _ := io.ReadAll(resp.Body)
		fmt.Printf("Unable to get movie, \n%s", resp)
	}

	var movie types.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		fmt.Println("error decoding movie response", err)
		return
	}
	if movie.Response == "False" {
		fmt.Println("Unable to find the movie, try another one")
		goto Start
	}

	prettyPrint, _ := json.MarshalIndent(movie, " ", "  ")
	fmt.Printf("Found movie: %s\n", prettyPrint)

	if movie.Poster == "" {
		return
	}
	resp, err = http.Get(movie.Poster)
	if err != nil {
		fmt.Println("error while making request", err, resp)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp, _ := io.ReadAll(resp.Body)
		fmt.Printf("Unable to get movie poster, \n%s", resp)
		return
	}

	fmt.Printf("The %s movie poster will be saved in Posters dir\n", movie.Title)
	var postersPath = path.Join(basePath, "Posters")

	if err = os.Mkdir(postersPath, fs.ModePerm); err != nil {
		fmt.Println("")
	}
	fileName := strings.Join([]string{movie.Title, filepath.Ext(movie.Poster)}, "")
	file, err := os.Create(path.Join(postersPath, fileName))
	if err != nil {
		fmt.Println("error creating poster file\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("error writing poster\n", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	fmt.Printf("Success on creating poster file for %s movie\n", movie.Title)

}
