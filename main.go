package main

import (
  "flag"
  "github.com/corehello/sha1dir/sha1dir"
)

func main(){
  rootpath := flag.String("root", ".", "the root directory which to be walked")
  filter := flag.String("filter", "", "the blacklist of the directory")
  output := flag.String("output", "sha1result", "the output file name or path")
  switch os.Args[1] {
  case "help":
    flag.usage()
    os.Exit(0)
  default:
    flag.Parse()
  }
  sha1dir.Run(rootpath, filter.Split(","), output)
}
