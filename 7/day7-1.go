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

func (d *Dir) totalSizeOfDirsLessThan100k() int {
	size := d.size()
	fmt.Println("Recursive in dir", d.name, "size", d.size())

	for _, dir := range d.subdirs {
		d = dir
		size += d.totalSizeOfDirsLessThan100k()
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
		case t == "$ cd ..":
			position = position.parentDir
		case strings.Contains(t, "cd") && !strings.Contains(t, "/"):
			data := strings.Split(t, " ")
			name := data[len(data)-1]
			fmt.Println("cd to dir:", name)
			position = position.subdirs[name]
		case t == "$ ls":
			// No action
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

	fmt.Println(root.totalSizeOfDirsLessThan100k())
}