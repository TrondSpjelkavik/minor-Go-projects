package main

import (
	"flag"
	"fmt"
	"strings"
)


var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	x := 1
	p := &x

	fmt.Println(*p)

	*p = 2

	fmt.Println(x)

	fmt.Println(f() == f())

	q := 1
	incr(&q)

	fmt.Println(incr(&q))

	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

var s = f()

func f() *int {
	v := 1
	return &v
}

func incr(b *int) int {
	*b++
	return *b
}

