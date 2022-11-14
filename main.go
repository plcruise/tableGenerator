package main

import (
	"flag"
	"fmt"
)

func main() {
	var outfile string
	flag.StringVar(&outfile, "n", "", "the name of the file you wish to store the latex code in. If empty it will be posted to the command line")
	flag.Parse()
	fmt.Println(outfile)
}
