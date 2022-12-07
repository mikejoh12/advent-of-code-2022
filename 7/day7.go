package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name      string
	files     map[string]File
	subdirs   map[string]*Dir
	parentDir *Dir
}

func (d *Dir) size() (totalSize int) {
	for _, file := range d.files {
		totalSize += file.size
	}
	return
}

func (d *Dir) getTotalSize() int {
	size := d.size()
	for _, dir := range d.subdirs {
		d = dir
		size += d.getTotalSize()
	}
	return size
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	root := Dir{
		name:    "root",
		subdirs: make(map[string]*Dir),
		files:   make(map[string]File),
	}
	position := &root

	for scanner.Scan() {
		t := scanner.Text()
		switch {
		case t == "$ cd /":
			position = &root
		case t[0:3] == "dir":
			if _, ok := position.subdirs[t[4:]]; !ok {
				newDir := Dir{
					name:      t[4:],
					subdirs:   make(map[string]*Dir),
					files:     make(map[string]File),
					parentDir: position,
				}
				position.subdirs[t[4:]] = &newDir
			}
		case t == "$ cd ..":
			position = position.parentDir
		case len(t) >= 5 && t[0:5] == "$ cd " && !strings.Contains(t, "/"):
			data := strings.Split(t, " ")
			position = position.subdirs[data[len(data)-1]]
		case t == "$ ls":
		default:
			var fileSize int
			var fileName string
			fmt.Sscanf(t, "%d %s", &fileSize, &fileName)
			if _, ok := position.files[fileName]; !ok {
				file := File{name: fileName, size: fileSize}
				position.files[fileName] = file
			}
		}
	}

	var totSize int
	var recur func(d *Dir)

	unusedSpace, deleteSpace := 70000000 - root.getTotalSize(), root.getTotalSize()

	recur = func(d *Dir) {
		subDirSize := d.getTotalSize()
		if subDirSize < 100000 {
			totSize += subDirSize
		}
		if unusedSpace+subDirSize >= 30000000 && subDirSize < deleteSpace {
			deleteSpace = subDirSize
		}
		for _, dir := range d.subdirs {
			recur(dir)
		}
	}
	recur(&root)

	fmt.Println("part 1:", totSize)
	fmt.Println("part 2:", deleteSpace)
}
