package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", login))
	if err != nil {
		return "", 0, fmt.Errorf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("unexpected status code %s", resp.Status)

	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, 2, nil
}

type Reply struct {
	Name        string `json:"name,omitempty"`
	PublicRepos string `json:"public_repos,omitempty"`
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
