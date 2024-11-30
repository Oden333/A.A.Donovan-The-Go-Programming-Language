package types

type Comic struct {
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}
