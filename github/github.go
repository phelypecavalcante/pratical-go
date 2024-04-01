package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	name, repos, err := githubInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.Printf("error: %s", err)
			os.Exit(1)
		*/
	}
	fmt.Printf("name: %#v\npublic_repos: %#v\n", name, repos)
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)

	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	var r struct {
		Name        string `json:"name,omitempty"`
		PublicRepos int    `json:"public_repos,omitempty"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.PublicRepos, nil
}

/* JSON <-> Go
true/false <-> true/false
string  <-> string
null <-> nil
number <-> float64, float32, int8, nt16, int32, int64, int, uint8, ...
array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding.json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
