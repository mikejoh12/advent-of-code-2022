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
			fmt.Println("Adding dir:", t[4:])
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
			fmt.Println("cd to dir:", name)
			position = position.subdirs[name]
		case t == "$ ls":
			// No action
		default:
			var fileSize int
			var fileName string
			fmt.Sscanf(t, "%d %s", &fileSize, &fileName)
			fmt.Println("adding file to dir", position.name ,fileName, fileSize)
			file := File{name: fileName, size: fileSize}
			position.files[fileName] = file
		}
		fmt.Println("dir", position.name, "dir size", position.size())
	}

	var totSize int
 	var recur func(d *Dir)

	recur = func(d *Dir) {
		if d.getSubDirSize() < 100000 {
			totSize += d.getSubDirSize()
		}
		for _, dir := range d.subdirs {
			recur(dir)
		}
	}
	recur(&root)
	fmt.Println("recur size", totSize)
}