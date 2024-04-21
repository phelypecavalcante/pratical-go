package main

import "fmt"

func main() {
	// fmt.Println(div(5, 0))
	fmt.Println(safeDiv(5, 0))
	fmt.Println(safeDiv(5, 2))
}

// named return values
func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv
	// (just like a & b)
	defer func() {
		// e's type is any (or interface{}) *not* error
		if e := recover(); e != nil {
			fmt.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
	/* Miki don't like this kind of programming
	q = a / b
	return
	*/
}

func div(a, b int) int {
	return a / b
}
