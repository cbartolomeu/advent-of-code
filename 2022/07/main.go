package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	files  []int
	dirs   map[string]*Directory
	parent *Directory
}

func toInt(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func getDirectory(filename string) Directory {
	dir := Directory{dirs: make(map[string]*Directory)}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	currDir := &dir
	for fileScanner.Scan() {
		line := fileScanner.Text()

		split := strings.Split(line, " ")

		if split[0] == "$" && split[1] == "cd" {
			to := split[2]
			if to == "/" {
				currDir = &dir
			} else if to == ".." {
				currDir = currDir.parent
			} else {
				currDir = currDir.dirs[to]
			}
		} else if split[0] != "$" {
			if split[0] == "dir" {
				newDir := Directory{parent: currDir, dirs: make(map[string]*Directory)}
				currDir.dirs[split[1]] = &newDir
			} else {
				currDir.files = append(currDir.files, toInt(split[0]))
			}
		}

	}

	return dir
}

func sum(arr []int) int {
	res := 0
	for _, item := range arr {
		res += item
	}
	return res
}

func sumOfDirsWithLessThan(dir Directory, max int) (int, []int) {
	total := sum(dir.files)
	dirs := []int{}
	for _, v := range dir.dirs {
		dirSize, newDirs := sumOfDirsWithLessThan(*v, max)

		dirs = append(dirs, newDirs...)
		total += dirSize
	}

	if total <= max || max == -1 {
		dirs = append(dirs, total)
	}

	return total, dirs
}

func maxBeforeN(dirs []int, n int) int {
	res := -1
	sort.Ints(dirs)
	for _, value := range dirs {
		if value >= n && (value < res || res == -1) {
			res = value
		}
	}

	return res
}

func main() {
	input := os.Args[1]

	dir := getDirectory(input)

	n, dirs := sumOfDirsWithLessThan(dir, 100000)
	fmt.Println(sum(dirs))

	n, dirs = sumOfDirsWithLessThan(dir, -1)
	freeSpace := 70000000 - n
	fmt.Println(maxBeforeN(dirs, 30000000-freeSpace))
}
