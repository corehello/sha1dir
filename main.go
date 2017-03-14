package main

import (
	"flag"
	"fmt"
	"github.com/corehello/sha1dir/sha1dir"
	"os"
	"strings"
)

func main() {
	rootpath := flag.String("root", ".", "the root directory which to be walked")
	filter := flag.String("filter", "", "the blacklist of the directory, support regexp")
	output := flag.String("output", "sha1result", "the output file name or path")
	if len(os.Args) == 1 {
		fmt.Println("Not enough given arguments")
		flag.Usage()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "help":
		flag.Usage()
		os.Exit(0)
	default:
		flag.Parse()
	}
	if *filter != "" {
		err := sha1dir.Run(*rootpath, strings.Split(*filter, ","), *output)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	} else {
		err := sha1dir.Run(*rootpath, []string{}, *output)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	fmt.Println("Sha1 info for path " + *rootpath + " is stored in the file: " + *output)
}
