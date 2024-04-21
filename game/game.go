package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	var i1 Item
	fmt.Println(i1)
	i1.Show("i1")

	i2 := Item{1, 2}
	i2.Show("i2")

	i3 := Item{
		Y: 10,
		// X: 20,
	}
	i3.Show("i3")
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	i3.Show("i3")

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}

	p1.Show("p1")
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	p1.Move(400, 600)
	p1.Show("p1")

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}

	moveAll(ms, 150, 100)

	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)

	// time.Time import json.Marshaler interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(k)
	p1.FoundKey(k)
	fmt.Println(p1.FoundKey(Key(8)))
	p1.Show("p1")

	players := []Player{
		{
			Name: "Parzival",
			Item: Item{500, 300},
		},
		{
			Name: "Nice",
			Item: Item{700, 500},
		},
		{
			Name: "Luchetti",
			Item: Item{600, 200},
		},
		{
			Name: "Phelype",
			Item: Item{10, 30},
		},
		{
			Name: "Pê",
			Item: Item{550, 500},
		},
	}

	sortByDistance(players, 100, 200)
}

func sortByDistance(players []Player, x, y int) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\nUnsorted")
	for _, p := range players {
		fmt.Printf("%s - %d\n", p.Name, p.Euclidean(x, y))
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Euclidean(x, y) > players[j].Euclidean(x, y)
	})

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\nEuclidean")
	for _, p := range players {
		fmt.Printf("%s - %d\n", p.Name, p.Euclidean(x, y))
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Manhattan(x, y) > players[j].Manhattan(x, y)
	})

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Manhattan")
	for _, p := range players {
		fmt.Printf("%s - %d\n", p.Name, p.Manhattan(x, y))
	}
}

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

func (k Key) Valid() bool {
	if k < Jade || k >= invalidKey {
		return false
	}
	return true
}

/* Exercise
- Add a "Keys" field to Player which is a slice of Key
- Add a "FoundKey(k Key) error" method to player which will add k to Key if it's no there
	- Err if k is not one of the known keys
*/

// Go's version of "enum"
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey
)

type Key byte

// Rule of thumb: Accept interfaces, return types

type mover interface {
	Move(x, y int)
	// Move(int, int)
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Item is an item in the game
type Item struct {
	X int
	Y int
}

func (i Item) Show(name string) {
	fmt.Printf("%s: %#v\n", name, i)
}

func (i Item) xSize(x int) float64 {
	return axisSize(i.X, x)
}

func (i Item) ySize(y int) float64 {
	return axisSize(i.Y, y)
}

func axisSize(pos1, pos2 int) float64 {
	return math.Abs(float64(pos1 - pos2))
}

func (i Item) Manhattan(x, y int) int {
	return int(i.xSize(x) + i.ySize(y))
}

func (i Item) Euclidean(x, y int) int {
	return int(math.Sqrt(math.Pow(i.xSize(x), 2) + math.Pow(i.ySize(y), 2)))
}

type Player struct {
	Name string
	Keys []Key
	Item // Embed Item
}

func (p Player) Show(name string) {
	fmt.Printf("%s: name: %s, item: %v, keys: %v\n", name, p.Name, p.Item, p.Keys)
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

func (p *Player) FoundKey(k Key) error {

	if !k.Valid() {
		return fmt.Errorf("unknown %s", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	// if !slices.Contains(p.Keys, k) {
	// 	p.Keys = append(p.Keys, k)
	// }

	return nil
}

// i is called "the receiver"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item {}
// func NewItem(x, y int) *Item {}
// func NewItem(x, y int) (Item, error) {}
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bound %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// The Go compiler does "escape analysis" and will allocate i on the heap
	return &i, nil
}

// zero vs missing value

const (
	maxX = 1000
	maxY = 600
)
