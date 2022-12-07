package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type File struct {
	name		string
	size		int
}

type Dir struct {
	name		string
	files		map[string]File
	subdirs		map[string]*Dir
	parentDir	*Dir
}

func (d *Dir) size() (totalSize int) {
	for _, file := range d.files {
		totalSize += file.size
	}
	return
}

func (d *Dir) getSubDirSize() int {
	size := d.size()
	for _, dir := range d.subdirs {
		d = dir
		size += d.getSubDirSize()
	}
	return size
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	root := Dir{
		name:  "root",
		subdirs: make(map[string]*Dir),
		files: make(map[string]File),
	}
	position := &root

	for scanner.Scan() {
		t := scanner.Text()
		switch {
		case t == "$ cd /":
			position = &root
		case t[0:3] == "dir":
			newDir := Dir{
				name: t[4:],
				subdirs: make(map[string]*Dir),
				files: make(map[string]File),
				parentDir: position,
			}
			position.subdirs[t[4:]] = &newDir
		case t == "$ cd ..":
			position = position.parentDir
		case len(t) >= 5 && t[0:5] == "$ cd " && !strings.Contains(t, "/"):
			data := strings.Split(t, " ")
			name := data[len(data)-1]
			position = position.subdirs[name]
		case t == "$ ls":
			// No action
		default:
			var fileSize int
			var fileName string
			fmt.Sscanf(t, "%d %s", &fileSize, &fileName)
			file := File{name: fileName, size: fileSize}
			position.files[fileName] = file
		}
	}

	var totSize int
 	var recur func(d *Dir)

	unusedSpace := 70000000 - root.getSubDirSize()
	dirToDeleteSpace := root.getSubDirSize()

	recur = func(d *Dir) {
		subDirSize := d.getSubDirSize()

		if subDirSize < 100000 {
			totSize += subDirSize
		}

		if unusedSpace + subDirSize >= 30000000 {
			dirToDeleteSpace = subDirSize
		}

		for _, dir := range d.subdirs {
			recur(dir)
		}
	}
	recur(&root)
	fmt.Println("recur size:", totSize)
	fmt.Println("dir to delete space:", dirToDeleteSpace)
}