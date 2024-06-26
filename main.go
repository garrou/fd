package main

import (
	"flag"

	"github.com/garrou/fd/lib"
)

func main() {
	hidden := flag.Bool("i", false, "search hidden entries (default: false)")
	recurse := flag.Bool("r", false, "find recursively (default: false)")
	count := flag.Bool("c", false, "count the number of results (default: false)")
	pattern := flag.Bool("p", false, "search pattern (default: false)")
	extension := flag.Bool("e", false, "search files with this extension (default: false)")
	time := flag.Bool("t", false, "get the execution time (default: false)")
	exclude := flag.String("x", "", "exclude matching folders (space separated)")
	flag.Parse()

	walker := lib.NewWalker(
		flag.Args(),
		*recurse,
		*hidden,
		*count,
		*pattern,
		*extension,
		*time,
		*exclude,
	)
	walker.Search()
}
