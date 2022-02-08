package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func main() {

	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")

	flag.Parse()

	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
}
