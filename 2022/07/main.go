package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

var example = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

const (
	threshold = 100000

	totalSize    = 70000000
	requiredSize = 30000000
)

func main() {
	example1()
	star1()
	example2()
	star2()
}

func example1() {
	res := newFileDir(example)
	res.Print()

	var total int
	addSize(res, threshold, &total)
	fmt.Printf("example total: %d\n", total)

	toTrim := totalSize - requiredSize
	fmt.Printf("to trim: %d bytes\n", toTrim)
	minTrim := res.MinDelete(requiredSize, totalSize)
	fmt.Printf("[example] min trim: %d\n", minTrim)
}

func star1() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	res := newFileDir(inputs)
	var total int
	addSize(res, threshold, &total)
	fmt.Printf("[*] total: %d\n", total)

	toTrim := totalSize - requiredSize
	fmt.Printf("to trim: %d bytes\n", toTrim)
	minTrim := res.MinDelete(toTrim, totalSize)

	fmt.Printf("[**] min trim: %d\n", minTrim)
}
func example2() {
	res := newFileDir(example)
	// res.Print()

	totalSize := res.Size()
	toTrim := totalSize - requiredSize
	fmt.Printf("to trim: %d bytes\n", toTrim)
	minTrim := res.MinDelete(toTrim, totalSize)
	fmt.Printf("[example] min trim: %d\n", minTrim)
}

func star2() {
	inputs, err := util.ReadLines("input.txt")
	util.Must(err)

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	res := newFileDir(inputs)

	sizeUsed := res.Size()
	sizeAvailable := totalSize - sizeUsed
	toTrim := requiredSize - sizeAvailable
	fmt.Printf("size used is: %d, to trim: %d bytes\n", sizeUsed, toTrim)
	minTrim := res.MinDelete(toTrim, totalSize)
	fmt.Printf("[**] min trim: %d\n", minTrim)
}

func addSize(fd *fileDir, threshold int, total *int) {
	size := fd.Size()
	if size <= threshold {
		*total = *total + size
	}
	for _, f := range fd.subDirs {
		addSize(f, threshold, total)
	}
}

type fileDir struct {
	name    string
	parent  *fileDir
	files   map[string]*file
	subDirs map[string]*fileDir
}

func (fd *fileDir) Print() {
	if fd == nil {
		return
	}

	fmt.Printf("name: %s [total size: %d bytes]\n", fd.name, fd.Size())
	if len(fd.files) > 0 {
		fmt.Println("files:")
		for _, file := range fd.files {
			fmt.Printf("- %s\n", file)
		}
	}

	if len(fd.subDirs) > 0 {
		fmt.Println("subdirs:")
		for _, fd := range fd.subDirs {
			fd.Print()
		}
	}
}

func (fd *fileDir) Size() int {
	var total int
	if fd == nil {
		return total
	}

	for _, f := range fd.files {
		total += f.size
	}

	for _, f := range fd.subDirs {
		total += f.Size()
	}

	return total
}

func (fd *fileDir) MinDelete(toTrim, currentMin int) int {
	for _, f := range fd.subDirs {
		size := f.Size()
		if size >= toTrim && size < currentMin {
			// this is eligible to delete
			currentMin = size

			// find another one that could be eligible
			newVal := f.MinDelete(toTrim, currentMin)
			if newVal < currentMin {
				currentMin = newVal
			}
		}
	}

	return currentMin
}

func newDir(name string, parent *fileDir) *fileDir {
	return &fileDir{
		name:    name,
		parent:  parent,
		files:   make(map[string]*file),
		subDirs: make(map[string]*fileDir),
	}
}

type file struct {
	name string
	size int
}

func (f *file) String() string { return fmt.Sprintf("%s [%d bytes]", f.name, f.size) }

func newFile(name string, size int) *file { return &file{name: name, size: size} }

func newFileDir(inputs []string) *fileDir {

	var (
		res       = newDir("/", nil) // root
		activeDir = res
	)

mainLoop:
	for _, input := range inputs {
		switch {
		case strings.HasPrefix(input, "$"):
			// a command
			cmds := strings.Split(strings.TrimLeft(input, "$ "), " ")
			switch cmds[0] {
			case "cd":
				// move dir
				loc := cmds[1]
				switch loc {
				case "..":
					// go up 1 dir
					activeDir = activeDir.parent
				case "/":
					// back to root
					activeDir = res
				default:
					activeDir = activeDir.subDirs[loc]
				}
				continue mainLoop
			case "ls":
				// get list of files
				continue mainLoop
			}
		default:
			// this is an output of a command
			// atm it's only the output of ls
			fileStr := strings.Split(input, " ")
			arg1, arg2 := fileStr[0], fileStr[1]
			switch arg1 {
			case "dir":
				nameDir := arg2
				if _, ok := activeDir.subDirs[nameDir]; !ok {
					activeDir.subDirs[nameDir] = newDir(nameDir, activeDir)
				}
			default:
				size, err := strconv.Atoi(arg1)
				util.Must(err)
				activeDir.files[arg2] = newFile(arg2, size)
			}
		}
	}

	return res
}
