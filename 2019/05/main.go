package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

const inp = `3,225,1,225,6,6,1100,1,238,225,104,0,1102,72,20,224,1001,224,-1440,224,4,224,102,8,223,223,1001,224,5,224,1,224,223,223,1002,147,33,224,101,-3036,224,224,4,224,102,8,223,223,1001,224,5,224,1,224,223,223,1102,32,90,225,101,65,87,224,101,-85,224,224,4,224,1002,223,8,223,101,4,224,224,1,223,224,223,1102,33,92,225,1102,20,52,225,1101,76,89,225,1,117,122,224,101,-78,224,224,4,224,102,8,223,223,101,1,224,224,1,223,224,223,1102,54,22,225,1102,5,24,225,102,50,84,224,101,-4600,224,224,4,224,1002,223,8,223,101,3,224,224,1,223,224,223,1102,92,64,225,1101,42,83,224,101,-125,224,224,4,224,102,8,223,223,101,5,224,224,1,224,223,223,2,58,195,224,1001,224,-6840,224,4,224,102,8,223,223,101,1,224,224,1,223,224,223,1101,76,48,225,1001,92,65,224,1001,224,-154,224,4,224,1002,223,8,223,101,5,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1107,677,226,224,1002,223,2,223,1005,224,329,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,344,1001,223,1,223,1107,226,226,224,1002,223,2,223,1006,224,359,1001,223,1,223,8,226,226,224,1002,223,2,223,1006,224,374,101,1,223,223,108,226,226,224,102,2,223,223,1005,224,389,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,404,101,1,223,223,1107,226,677,224,1002,223,2,223,1006,224,419,101,1,223,223,1008,226,677,224,1002,223,2,223,1006,224,434,101,1,223,223,108,677,677,224,1002,223,2,223,1006,224,449,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,479,101,1,223,223,7,226,677,224,1002,223,2,223,1006,224,494,1001,223,1,223,7,677,677,224,102,2,223,223,1006,224,509,101,1,223,223,107,226,677,224,1002,223,2,223,1006,224,524,1001,223,1,223,1007,226,226,224,102,2,223,223,1006,224,539,1001,223,1,223,108,677,226,224,102,2,223,223,1005,224,554,101,1,223,223,1007,677,677,224,102,2,223,223,1006,224,569,101,1,223,223,8,677,226,224,102,2,223,223,1006,224,584,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,599,1001,223,1,223,1007,677,226,224,1002,223,2,223,1005,224,614,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,629,101,1,223,223,1108,677,677,224,1002,223,2,223,1005,224,644,1001,223,1,223,8,226,677,224,1002,223,2,223,1006,224,659,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,674,101,1,223,223,4,223,99,226`
const input = 5

var output int

var (
	errInvalidArgsNum = errors.New("invalid number of arguments")
)

type (
	mode   int
	opcode int
)

const (
	positionMode mode = iota
	immediateMode

	ocOne   opcode = 1
	ocTwo   opcode = 2
	ocThree opcode = 3
	ocFour  opcode = 4
	ocFive  opcode = 5
	ocSix   opcode = 6
	ocSeven opcode = 7
	ocEight opcode = 8
	oc99    opcode = 99
)

type opcodeFunc func([]int, *int, ...int) (bool, error)

var opcodes = map[opcode]opcodeFunc{
	ocOne: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 3 {
			return false, errInvalidArgsNum
		}

		arr[args[2]] = arr[args[0]] + arr[args[1]]
		return false, nil
	},
	ocTwo: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 3 {
			return false, errInvalidArgsNum
		}

		arr[args[2]] = arr[args[0]] * arr[args[1]]
		return false, nil
	},
	ocThree: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 1 {
			return false, errInvalidArgsNum
		}

		arr[args[0]] = input
		return false, nil
	},
	ocFour: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 1 {
			return false, errInvalidArgsNum
		}
		output = arr[args[0]]
		return false, nil
	},
	ocFive: func(arr []int, idx *int, args ...int) (bool, error) {
		if len(args) != 2 {
			return false, errInvalidArgsNum
		}

		if arr[args[0]] == 0 {
			return false, nil
		}

		*idx = arr[args[1]]
		return true, nil
	},
	ocSix: func(arr []int, idx *int, args ...int) (bool, error) {
		if len(args) != 2 {
			return false, errInvalidArgsNum
		}

		if arr[args[0]] != 0 {
			return false, nil
		}

		*idx = arr[args[1]]
		return true, nil
	},
	ocSeven: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 3 {
			return false, errInvalidArgsNum
		}

		var val int
		if arr[args[0]] < arr[args[1]] {
			val = 1
		}

		arr[args[2]] = val
		return false, nil
	},
	ocEight: func(arr []int, _ *int, args ...int) (bool, error) {
		if len(args) != 3 {
			return false, errInvalidArgsNum
		}

		var val int
		if arr[args[0]] == arr[args[1]] {
			val = 1
		}

		arr[args[2]] = val
		return false, nil
	},
	oc99: func(arr []int, _ *int, args ...int) (bool, error) {
		return false, nil
	},
}

func (oc opcode) numArgs() int {
	return map[opcode]int{
		ocOne:   3,
		ocTwo:   3,
		ocThree: 1,
		ocFour:  1,
		ocFive:  2,
		ocSix:   2,
		ocSeven: 3,
		ocEight: 3,
	}[oc]
}

func modeOne(arr []int, ocInput int, idx *int, args ...int) (bool, error) {

	ocstring := fmt.Sprintf("%05d", ocInput)
	ocint, err := strconv.Atoi(ocstring[3:5])
	if err != nil {
		return false, err
	}

	oc := opcode(ocint)

	getMode := func(idx int) int {

		idx = 2 - idx

		m, err := strconv.Atoi(string(ocstring[idx]))
		if err != nil {
			return 0
		}
		return m
	}

	if oc.numArgs() != len(args) {
		return false, errInvalidArgsNum
	}

	// modify all the arguments
	for i := range args {

		mode := mode(getMode(i))

		if mode == immediateMode || i == 2 {
			continue
		}
		args[i] = arr[args[i]]
	}

	var increment bool

	switch oc {
	case ocOne:
		arr[args[2]] = args[0] + args[1]
	case ocTwo:
		arr[args[2]] = args[0] * args[1]
	case ocThree:
		arr[args[0]] = input
	case ocFour:
		output = args[0]
	case ocFive:
		if args[0] == 0 {
			break
		}
		*idx = args[1]
		increment = true
	case ocSix:
		if args[0] != 0 {
			break
		}
		*idx = args[1]
		increment = true
	case ocSeven:
		var val int
		if args[0] < args[1] {
			val = 1
		}
		arr[args[2]] = val
	case ocEight:
		var val int
		if args[0] == args[1] {
			val = 1
		}
		arr[args[2]] = val
	default:
		return false, fmt.Errorf("invalid situation %v %v", oc, args)
	}

	return increment, nil

}

func getOC(ocInput int) (opcode, error) {
	ocstring := fmt.Sprintf("%05d", ocInput)
	ocint, err := strconv.Atoi(ocstring[3:5])
	if err != nil {
		return 0, err
	}

	return opcode(ocint), nil
}

func main() {

	t := time.Now()
	defer func(t time.Time) {
		fmt.Printf("runtime: %v\n", time.Since(t))
	}(t)

	arrString := strings.Split(inp, ",")
	arr := make([]int, len(arrString))
	for i := range arrString {
		num, err := strconv.Atoi(arrString[i])
		util.Must(err)
		arr[i] = num
	}

	var i int
	for {

		oc := opcode(arr[i])
		if oc == oc99 {
			break
		}
		na := oc.numArgs()
		if na == 0 {
			newOC, err := getOC(arr[i])
			util.Must(err)
			na = newOC.numArgs()
		}
		args := make([]int, na)

		for j := 0; j < na; j++ {
			i++
			args[j] = arr[i]
		}

		ocf, ok := opcodes[oc]
		if !ok {
			ocf = func(arr []int, idx *int, args ...int) (bool, error) {
				return modeOne(arr, int(oc), idx, args...)
			}
		}

		increment, err := ocf(arr, &i, args...)
		if err != nil {
			panic(err)
		}

		if increment {
			continue
		}

		i++
	}
	fmt.Printf("part one output: %v\n", output)

}
