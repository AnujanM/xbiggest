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

type SizeSorterDes []file

func (a SizeSorterDes) Len() int           { return len(a) }
func (a SizeSorterDes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SizeSorterDes) Less(i, j int) bool { return a[i].size > a[j].size }

func parse(folders string) []string{
	data := strings.Split(folders, ",")
	for i, d := range data {
		data[i] = strings.Trim(d, " ")
	}
	return data
}

// func getSmallest(t []file) int{
// 	smallestSize := t[0].size
// 	var index int
// 	for i, _ := range t{
// 		if (t[i].size < smallestSize){
// 			index = i
// 		}
// 	}
// 	return index
// }

func addFile(path string, size int64, t []file) {
	//small := getSmallest(t)
	sort.Sort(SizeSorterAsc(t))
	if (t[0].size < size){
		t[0] = file{path, size}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected query.")
		os.Exit(1)
	}
	input := os.Args[1]

	if len(os.Args) == 2 {
		numberOfFiles := 10
	}else{
		numberOfFiles, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Number of ")
			log.Fatal(err)
		}
	}
	
	top := make([]file, numberOfFiles)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	q := parse(input)
	for i, src := range q {
		if strings.Contains(src, "~") {
			q[i] = filepath.Join(usr.HomeDir, src[1:])
		}
	}

	for _, source := range q {
		filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if path == "." || path == ".." || err != nil {
				return nil
			}

			addFile(path, info.Size(), top)

			// fmt.Printf("%s, %d", path, info.Size())
			// fmt.Printf("\n")

			return nil
		})
	}

	//fmt.Println(getSmallest(top))
	sort.Sort(SizeSorterDes(top))
	for _, thing := range top {
		fmt.Printf("%d, %s", thing.size, thing.path)
		fmt.Printf("\n")
	}
}
