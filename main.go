package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"strconv"
	"sort"
)

type file struct {
	path string
	size int64
}

type SizeSorterAsc []file

func (a SizeSorterAsc) Len() int           { return len(a) }
func (a SizeSorterAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SizeSorterAsc) Less(i, j int) bool { return a[i].size < a[j].size }

func addFile(path string, size int64, t []file) {
	sort.Sort(SizeSorterAsc(t))
	if (t[0].size < size){
		t[0] = file{path, size}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected directory to search.")
		os.Exit(1)
	}

	var numberOfFiles int
	var input []string

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		numberOfFiles = 10
		input = os.Args[1:]
	}else{
		numberOfFiles = num
		input = os.Args[2:]
	}

	top := make([]file, numberOfFiles)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	for i, src := range input {
		if strings.Contains(src, "~") {
			input[i] = filepath.Join(usr.HomeDir, src[1:])
		}
	}

	for _, source := range input {
		filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if path == "." || path == ".." || err != nil {
				return nil
			}

			addFile(path, info.Size(), top)

			return nil
		})
	}

	for i, _ := range top {
		fmt.Printf("%d, %s", top[len(top)-i-1].size, top[len(top)-i-1].path)
		fmt.Printf("\n")
	}
}