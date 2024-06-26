package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
	// mapDemo()

	file, err = os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	ws, errN := mostCommonN(file, 10)
	if errN != nil {
		log.Fatalf("error: %s", errN)
	}
	fmt.Println(ws)
}

// Q: What is the most common word (ignore case) in sherlock.text?
// Word frequency

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

/*
Will run before main

	func init() {
		//...
	}
*/
func mapDemo() {
	var stocks map[string]float64 // symbol -> price
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}
	/*
		stocks = make(map[string]float64)
		stocks[sym] = 136.73
	*/

	stocks = map[string]float64{
		sym:    137.73,
		"AAPL": 172.35,
	}

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k := range stocks { // keys
		fmt.Println(k)
	}

	for k, v := range stocks { // key & value
		fmt.Println(k, "->", v)
	}

	for _, v := range stocks { // values
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "AAPL") // no panic
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

type WordFrequency struct {
	Word      string
	Frequency int
}

type ByFrequency []WordFrequency

func (f ByFrequency) Len() int {
	return len(f)
}
func (f ByFrequency) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f ByFrequency) Less(i, j int) bool {
	return f[i].Frequency > f[j].Frequency
}

func mostCommonN(r io.Reader, n int) ([]string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return nil, err
	}

	wfs := make([]WordFrequency, 0, len(freqs))

	for k, v := range freqs {
		wfs = append(wfs, WordFrequency{
			Word:      k,
			Frequency: v,
		})
	}

	sort.Sort(ByFrequency(wfs))

	return maxWords(wfs, n)
}

func maxWords(wfs []WordFrequency, n int) ([]string, error) {
	if len(wfs) == 0 {
		return nil, fmt.Errorf("empty map")
	}

	words := make([]string, 0, n)
	for _, v := range wfs[:n] {
		words = append(words, v.Word)
	}

	return words, nil
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}

	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}
