package sha1dir

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sha1file(path string, info os.FileInfo, err error) error {
	if !match(blacklist, path) {
		switch mode := info.Mode(); {
		case mode.IsRegular():
			wg.Add(1)
			go func() {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
				}
				sum := sha1.Sum(content)
				abspath, _ := filepath.Abs(path)
				o := fmt.Sprintf("%s, %x, %d", abspath, sum, info.Size())
				output <- o
				//fix the data lost for buffered channel
				time.Sleep(1 * time.Millisecond)
				defer wg.Done()
			}()
		}
	}
	return nil
}

func dirwalk(path string) {
	err := filepath.Walk(path, sha1file)
	if err != nil {
		fmt.Println("Walk failed")
	}
}

func match(s []string, e string) bool {
	if len(s) == 0 {
		return false
	}
	for _, a := range s {
		if regexp.MustCompile(a).MatchString(e) {
			return true
		}
	}
	return false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var blacklist []string
var output = make(chan string, 31)

func Run(rootpath string, filter []string, outputfile string) {
	_, err := PathExists(rootpath)
	if err != nil {
		fmt.Println("The given path: " + rootpath + " is not existed on disk!!")
		return
	}
	for _, fil := range filter {
		blacklist = append(blacklist, fil)
	}
	runtime.GOMAXPROCS(4)
	f, _ := os.Create(outputfile)
	defer f.Close()
	go func() {
		fmt.Println("Writing to the output file: " + outputfile)
		for t := range output {
			fmt.Fprintln(f, t)
		}
	}()
	dirwalk(rootpath)
	wg.Wait()
	close(output)
}
