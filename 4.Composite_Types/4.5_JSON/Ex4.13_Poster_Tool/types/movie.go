package types

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime,omitempty"`
	Genre    string `json:"Genre,omitempty"`
	Director string `json:"Director,omitempty"`
	Writer   string `json:"Writer,omitempty"`
	Actors   string `json:"Actors,omitempty"`
	Plot     string `json:"Plot,omitempty"`
	Language string `json:"Language,omitempty"`
	Country  string `json:"Country,omitempty"`
	Awards   string `json:"Awards,omitempty"`
	Poster   string `json:"Poster,omitempty"`
	Ratings  []struct {
		Source string `json:"Source,omitempty"`
		Value  string `json:"Value,omitempty"`
	} `json:"Ratings,omitempty"`
	Metascore  string `json:"Metascore,omitempty"`
	ImdbRating string `json:"imdbRating,omitempty"`
	ImdbVotes  string `json:"imdbVotes,omitempty"`
	ImdbID     string `json:"imdbID,omitempty"`
	Type       string `json:"Type,omitempty"`
	DVD        string `json:"DVD,omitempty"`
	BoxOffice  string `json:"BoxOffice,omitempty"`
	Production string `json:"Production,omitempty"`
	Website    string `json:"Website,omitempty"`
	Response   string `json:"Response,omitempty"`
}
